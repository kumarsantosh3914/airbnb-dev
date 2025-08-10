import { Request, Response } from "express";
import {
  confirmBookingService,
  createBookingService,
  getBookingsById,
} from "../services/booking.service";

export const createBookingHanlder = async (req: Request, res: Response, next: Function) => {
  try {
    const booking = await createBookingService(req.body);
    res.status(201).json({
      bookingId: booking.bookingId,
      idempotencyKey: booking.idempotencyKey,
    });
  } catch (err) {
    next(err);
  }
};

export const confirmBookingHandler = async (req: Request, res: Response, next: Function) => {
  try {
    const booking = await confirmBookingService(req.params.idempotencyKey);
    res.status(200).json({
      bookingId: booking.id,
      status: booking.status,
    });
  } catch (err) {
    next(err);
  }
};

export const getBookingsByHandler = async (req: Request, res: Response, next: Function) => {
  try {
    const booking = await getBookingsById(Number(req.params.id));
    res.status(200).json({
      status: "success",
      message: "Booking retrieved successfully",
      data: booking,
    });
  } catch (err) {
    next(err);
  }
};
