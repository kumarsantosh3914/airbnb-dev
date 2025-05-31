import { NotificationDto } from "../dto/notification.dto";
import { mailerQueue } from "../queues/email.queue";

export const MAILER_PAYLOAD = "payload:mail";

export const addEmailToQueue = (payload: NotificationDto) => {
    mailerQueue.add(MAILER_PAYLOAD, payload);
    console.log(`Email added to queue: ${JSON.stringify(payload)}`);
};
