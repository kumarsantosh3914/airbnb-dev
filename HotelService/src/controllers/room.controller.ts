import { NextFunction, Request, Response } from "express";
import { getAvailableRoomsService, updateBookingIdToRoomsService } from "../services/room.service";
import { StatusCodes } from "http-status-codes";

export async function getAvailableRoomsHandler(req: Request, res: Response, next: NextFunction) {

    const rooms = await getAvailableRoomsService({
        roomCategoryId: Number(req.query.roomCategoryId),
        checkInDate: new Date(req.query.checkInDate as string),
        checkOutDate: new Date(req.query.checkOutDate as string),
    });

    res.status(StatusCodes.OK).json({
        success: true,
        message: "Available rooms fetched successfully",
        data: rooms,
    });
}

export async function updateBookingIdToRoomsHandler(req: Request, res: Response, next: NextFunction) {
    const result = await updateBookingIdToRoomsService(req.body);

    res.status(StatusCodes.OK).json({
        success: true,
        message: "Booking ID updated to rooms successfully",
        data: result,
    });
}