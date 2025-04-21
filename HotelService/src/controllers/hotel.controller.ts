import { NextFunction, Request, Response } from "express";
import { createHotelService, deleteHotelService, getAllHotelsService, getHotelByIdService, updateHotelService } from "../services/hotel.service";

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

export async function getAllHotelsHandler(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await getAllHotelsService();

    res.status(200).json({
        message: "Hotels fetched successfully",
        data: hotelResponse,
        success: true,
    })
}

export async function deleteHotelHandler(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await deleteHotelService(Number(req.params.id));

    res.status(200).json({
        message: "Hotel deleted successfully",
        data: hotelResponse,
        success: true,
    })
}

export async function updateHotelHandler(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await updateHotelService(Number(req.params.id), req.body);

    res.status(200).json({
        message: "Hotel updated successfully",
        data: hotelResponse,
        success: true,
    })
}