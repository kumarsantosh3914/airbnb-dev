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

export async function getAllHotels() {
    const hotels = await Hotel.findAll();
    logger.info(`Hotels found: ${hotels.length}`);
    return hotels;
}

export async function deleteHotel(id: number) {
    const deleteCount = await Hotel.destroy({
        where: { id }
    });

    if (deleteCount === 0) {
        logger.error(`Hotel not found: ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }

    logger.info(`Hotel deleted: ${id}`);
    return deleteCount;
}

export async function updateHotel(id: number, hotelDto: createHotelDto) {
    const hotel = await Hotel.findByPk(id);

    if(!hotel) {
        logger.error(`Hotel not found: ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }

    hotel.name = hotelDto.name;
    hotel.address = hotelDto.address;
    hotel.location = hotelDto.location;
    hotel.rating = hotelDto.rating;
    hotel.ratingCount = hotelDto.ratingCount;

    await hotel.save();
    logger.info(`Hotel updated: ${hotel.id}`);
    return hotel;
}