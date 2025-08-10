import { NextFunction, Request, Response } from "express";
import { AppError } from "../utils/errors/app.error";

export const appErrorHandler = (err: AppError, req: Request, res: Response, next: NextFunction) => {
    const status = typeof err.statusCode === 'number' ? err.statusCode : 500;
    res.status(status).json({
        success: false,
        message: err.message,
    });
}

export const genericErrorHandler = (err: AppError, req: Request, res: Response, next: NextFunction) => {
    const status = typeof err.statusCode === 'number' ? err.statusCode : 500;
    res.status(status).json({
        success: false,
        message: err.message,
    });
}