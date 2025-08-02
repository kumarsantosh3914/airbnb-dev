import express from 'express';
import { validateRequestBody } from '../../validators';
import { RoomGenerationRequestSchema } from '../../dto/roomGeneration.dto';
import { generateRoomHandler } from '../../controllers/roomGeneration.controller';

const roomGenerationRouter = express.Router();

roomGenerationRouter.post(
    '/',
    validateRequestBody(RoomGenerationRequestSchema),
    generateRoomHandler
)

export default roomGenerationRouter;