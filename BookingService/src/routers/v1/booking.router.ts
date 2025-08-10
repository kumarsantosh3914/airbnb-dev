import express from 'express';
import { validateRequestBody } from '../../validators';
import { createBookingSchema } from '../../validators/booking.validator';
import { confirmBookingHandler, createBookingHanlder, getBookingsByHandler } from '../../controllers/booking.controller';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(createBookingSchema), createBookingHanlder);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler);
bookingRouter.get('/:id', getBookingsByHandler);

bookingRouter.get('/helth', (req, res) => {
    res.status(200).json('OK');
});

export default bookingRouter;