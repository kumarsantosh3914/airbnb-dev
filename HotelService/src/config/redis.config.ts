import Redis from "ioredis";
import { redisConfig } from ".";
import logger from "./logger.config";

function connectToRedis() {
    try {
        let connection: Redis;

        const RedisConfig = {
            port: redisConfig.REDIS_PORT,
            host: redisConfig.REDIS_HOST,
            maxRetriesPerRequest: null, // Disable automatic reconnection
        }

        return () => {
            if(!connection) {
                connection = new Redis(RedisConfig);
                return connection;
            }

            return connection;
        }
    } catch (error) {
        logger.error('Error connecting to Redis:', error);
        throw error;
    }
}

export const getRedisConnObject = connectToRedis();