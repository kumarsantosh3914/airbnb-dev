// This file contains all the basic configuration for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number;
    MONGODB_URI: string;
    REDIS_PORT?: number;
    REDIS_HOST?: string;
    MAIL_USER?: string;
    MAIL_PASSWORD?: string;
}

function loadEnv() {
    dotenv.config();
    console.log('Environment variables loaded from .env file');
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3000,
    MONGODB_URI: String(process.env.MONGODB_URI),
    REDIS_PORT: process.env.REDIS_PORT ? Number(process.env.REDIS_PORT) : 6379,
    REDIS_HOST: process.env.REDIS_HOST ? String(process.env.REDIS_HOST) : 'localhost',
    MAIL_USER: process.env.MAIL_USER,
    MAIL_PASSWORD: process.env.MAIL_PASSWORD,
};