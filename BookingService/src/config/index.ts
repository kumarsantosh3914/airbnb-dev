// This file contains all the basic configuration for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number;
    REDIS_SERVER_URL: string;
    LOCK_TTL: number;
    HOTEL_SERVICE_URL: string;
}

function loadEnv() {
    dotenv.config();
    console.log('Environment variables loaded from .env file');
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3000,
    REDIS_SERVER_URL: process.env.REDIS_SERVER_URL || 'redis://localhost:6379',
    LOCK_TTL: Number(process.env.LOCK_TTL) || 1000,
    HOTEL_SERVICE_URL: process.env.HOTEL_SERVICE_URL || 'http://localhost:3001/api/v1',
}