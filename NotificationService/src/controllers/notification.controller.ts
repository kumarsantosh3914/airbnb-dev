import { NextFunction, Request, Response } from "express";
import logger from "../config/logger.config";
import { NotificationDto } from "../dto/notification.dto";
import { addEmailToQueue } from "../producers/email.producer";

export const sendNotificationHandler = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  logger.info("Notification request received");

  const notificationData: NotificationDto = req.body;

  logger.info(`Processing notification request for ${notificationData.to}`, {
    templateId: notificationData.templateId,
    subject: notificationData.subject,
  });

  // Add email to queue for processing
  await addEmailToQueue({
    to: notificationData.to,
    subject: notificationData.subject,
    templateId: notificationData.templateId,
    params: notificationData.params || {},
  });

  logger.info(`Email successfully queued for ${notificationData.to}`);

  res.status(200).json({
    success: true,
    message: "Notification sent successfully",
    data: {
        to: notificationData.to,
        subject: notificationData.subject,
        templateId: notificationData.templateId
    }
  });
};
