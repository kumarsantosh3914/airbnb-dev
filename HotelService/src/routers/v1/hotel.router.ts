import express from 'express';
import { validateRequestBody } from '../../validators';
import { hotelSchema } from '../../validators/hotel.validator'
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelHandler, updateHotelHandler } from '../../controllers/hotel.controller';

const hotelRouter = express.Router();

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);
hotelRouter.get('/:id', getHotelHandler);
hotelRouter.get('/', getAllHotelsHandler);
hotelRouter.delete('/:id', deleteHotelHandler);
hotelRouter.put('/:id', validateRequestBody(hotelSchema), updateHotelHandler);

export default hotelRouter;