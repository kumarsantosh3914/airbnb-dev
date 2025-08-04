import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer";
import { Request, Response } from "express";

export async function generateRoomHandler(req: Request, res: Response) {
    await addRoomGenerationJobToQueue(req.body);

    res.status(201).json({
        message: "Room generation job added to queue successfully",
        success: true,
        data: {},
    });
}