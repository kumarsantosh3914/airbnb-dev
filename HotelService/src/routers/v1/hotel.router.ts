import express from 'express';
import { validateRequestBody } from '../../validators';
import { hotelScheama } from '../../validators/hotel.validator'
import { createHotelHanlder, deleteHotelHandler, getAllHotelsHandler, getHotelHandler, updateHotelHandler } from '../../controllers/hotel.controller';

const hotelRouter = express.Router();

hotelRouter.post('/', validateRequestBody(hotelScheama), createHotelHanlder);
hotelRouter.get('/:id', getHotelHandler);
hotelRouter.get('/', getAllHotelsHandler);
hotelRouter.delete('/:id', deleteHotelHandler);
hotelRouter.put('/:id', validateRequestBody(hotelScheama), updateHotelHandler);


export default hotelRouter;