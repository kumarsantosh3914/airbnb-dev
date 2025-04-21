import express from 'express';
import { validateRequestBody } from '../../validators';
import { hotelScheama } from '../../validators/hotel.validator'
import { createHotelHanlder, getHotelHandler } from '../../controllers/hotel.controller';

const hotelRouter = express.Router();

hotelRouter.post('/', validateRequestBody(hotelScheama), createHotelHanlder);
hotelRouter.get('/:id', getHotelHandler);

export default hotelRouter;