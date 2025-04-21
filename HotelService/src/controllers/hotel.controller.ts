import { NextFunction, Request, Response } from "express";
import { createHotelService, getHotelByIdService } from "../services/hotel.service";

export async function createHotelHanlder(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await createHotelService(req.body);

    res.status(201).json({
        message: "Hotel created successfully",
        data: hotelResponse,
        success: true,
    });
}

export async function getHotelHandler(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await getHotelByIdService(Number(req.params.id));

    res.status(200).json({
        message: "Hotel fetched successfully",
        data: hotelResponse,
        success: true,
    })
}