import express from 'express';
import  {sendNotificationHandler}  from '../../controllers/notification.controller';

const notificationRouter = express.Router();

notificationRouter.post('/send', sendNotificationHandler)

export default notificationRouter;