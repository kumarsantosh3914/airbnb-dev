// This file contains all the basic configuration for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number;
    MONGODB_URI: string;
}

type DBConfig = {
    DB_HOST: string;
    DB_USER: string;
    DB_PASSWORD: string;
    DB_NAME: string;
}

type RedisConfig = {
    REDIS_PORT?: number;
    REDIS_HOST?: string;
}

function loadEnv() {
    dotenv.config();
    console.log('Environment variables loaded from .env file');
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3000,
    MONGODB_URI: String(process.env.MONGODB_URI),
}

export const dbConfig: DBConfig = {
    DB_HOST: String(process.env.DB_HOST),
    DB_USER: String(process.env.DB_USER),
    DB_PASSWORD: String(process.env.DB_PASSWORD),
    DB_NAME: String(process.env.DB_NAME),
}

export const redisConfig: RedisConfig = {
    REDIS_PORT: process.env.REDIS_PORT ? Number(process.env.REDIS_PORT) : undefined,
    REDIS_HOST: process.env.REDIS_HOST || 'localhost',
}