import { generateRooms } from "../services/roomGeneration.service";
import { Request, Response } from "express";

export async function generateRoomHandler(req: Request, res: Response) {
    const response = await generateRooms(req.body);

    res.status(201).json({
        message: "Rooms generated successfully",
        success: true,
        data: response,
    });
}