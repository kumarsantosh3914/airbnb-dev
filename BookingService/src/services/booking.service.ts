import { CreateBookingDto } from "../dto/booking.dto";
import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizedIdempotencyKey,
  getBookingById,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import prismaClient from "../prisma/client";
import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";
import { getAvailableRooms, updateBookingIdToRooms } from "../api/hotel.api";

type AvailableRoom = {
  id: number;
  roomCategoryId: number;
  dateOfAvailability: Date;
};

export async function createBookingService(createBookingDto: CreateBookingDto) {
  const ttl = serverConfig.LOCK_TTL;
  const bookingResource = `hotel:${createBookingDto.hotelId}`;

  const availableRooms = await getAvailableRooms(
    createBookingDto.roomCategoryId,
    createBookingDto.checkInDate,
    createBookingDto.checkOutDate
  );

  const checkoutDate = new Date(createBookingDto.checkOutDate);
  const checkinDate = new Date(createBookingDto.checkInDate);

  const totalNights = Math.ceil(
    (checkoutDate.getTime() - checkinDate.getTime()) / (1000 * 60 * 60 * 24)
  );

  if( availableRooms.length === 0 || availableRooms < totalNights) {
    throw new BadRequestError("No available rooms for the selected dates");
  }

  try {
      await redlock.acquire([bookingResource], ttl);
      const booking = await createBooking({
        userId: createBookingDto.userId,
        hotelId: createBookingDto.hotelId,
        totalGuests: createBookingDto.totalGuests,
        bookingAmount: createBookingDto.bookingAmount,
        checkInDate: new Date(createBookingDto.checkInDate),
        checkOutDate: new Date(createBookingDto.checkOutDate),
        roomCategoryId: createBookingDto.roomCategoryId,
      });

      const idempotencyKey = generateIdempotencyKey();

      await createIdempotencyKey(idempotencyKey, booking.id);

      await updateBookingIdToRooms(booking.id.toString(), availableRooms.data.map((room: AvailableRoom) => room.id));

      return {
        bookingId: booking.id,
        idempotencyKey: idempotencyKey,
      };
  } catch (error) {
    throw new InternalServerError(
      "Failed to acquire lock for booking resource"
    );
  }
}

export async function confirmBookingService(idempotencyKey: string) {
  return prismaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey
    );

    if (!idempotencyKeyData || !idempotencyKeyData.bookingId) {
      throw new NotFoundError("Idempotency key not found");
    }

    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("Idempotency key already finalized");
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
    await finalizedIdempotencyKey(tx, idempotencyKey);

    return booking;
  });
}

export async function getBookingsById(bookingId: number) {
  const booking = await getBookingById(bookingId);

  if (!booking) {
    throw new NotFoundError("Booking not found");
  }

  return booking;
}


