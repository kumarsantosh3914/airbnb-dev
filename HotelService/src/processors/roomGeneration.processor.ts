import { Job, Worker } from "bullmq"
import { RoomGenerationJob } from "../dto/roomGeneration.dto"
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue"
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer"
import logger from "../config/logger.config"
import { generateRooms } from "../services/roomGeneration.service"
import { getRedisConnObject } from "../config/redis.config"

export const setupRoomGenerationWorker = () => {

    const roomGenerationProcessor = new Worker<RoomGenerationJob>(
        ROOM_GENERATION_QUEUE,
        async (job: Job) => {

            if(job.name !== ROOM_GENERATION_PAYLOAD) {
                throw new Error("Invalid job name");
            }

            const payload = job.data;
            logger.info(`Processing room generation job for room ID: ${payload.roomId}`);

            await generateRooms(payload);
            logger.info(`Room generation job completed for room ID: ${payload.roomId}`);
        }, 
        {
            connection: getRedisConnObject(),
        }
    )

    roomGenerationProcessor.on("failed", () => {
        logger.error(`Room generation job failed`);
    });

    roomGenerationProcessor.on("completed", () => {
        logger.info(`Room generation job completed successfully`);
    });
}