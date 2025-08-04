import express from 'express';
import { validateRequestBody } from '../../validators';
import { createBookingSchema } from '../../validators/booking.validator';
import { confirmBookingHandler, createBookingHanlder, getBookingsByHandler } from '../../controllers/booking.controller';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(createBookingSchema), createBookingHanlder);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler);
bookingRouter.get('/:id', getBookingsByHandler);

export default bookingRouter;