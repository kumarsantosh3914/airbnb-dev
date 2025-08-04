import { Request, Response } from "express";
import {
  confirmBookingService,
  createBookingService,
  getBookingsById,
} from "../services/booking.service";

export const createBookingHanlder = async (req: Request, res: Response) => {
  const booking = await createBookingService(req.body);

  res.status(201).json({
    bookingId: booking.bookingId,
    idempotencyKey: booking.idempotencyKey,
  });
};

export const confirmBookingHandler = async (req: Request, res: Response) => {
  const booking = await confirmBookingService(req.params.idempotencyKey);

  res.status(200).json({
    bookingId: booking.id,
    status: booking.status,
  });
};

export const getBookingsByHandler = async (req: Request, res: Response) => {

  const booking = await getBookingsById(Number(req.params.id));

  res.status(200).json({
    status: "success",
    message: "Booking retrieved successfully",
    data: booking,
  });
};
