import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDto } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelDto: createHotelDto) {
    const hotel = await Hotel.create({
        name: hotelDto.name,
        address: hotelDto.address,
        location: hotelDto.location,
        rating: hotelDto.rating,
        ratingCount: hotelDto.ratingCount,
    });

    logger.info(`Hotel created: ${hotel.id}`);
    return hotel;
}

export async function getHotelById(id: number) {
    const hotel = await Hotel.findByPk(id);

    if(!hotel) {
        logger.error(`Hotel not found: ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }

    logger.info(`Hotel found: ${hotel.id}`);
    return hotel;
}