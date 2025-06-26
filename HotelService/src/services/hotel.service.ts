import { createHotelDto } from "../dto/hotel.dto";
import { HotelRepository } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();

export async function createHotelService(hotelDto: createHotelDto) {
    const hotel = await hotelRepository.create(hotelDto);
    return hotel;
}

export async function getHotelByIdService(id: number) {
    const hotel = await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotels = await hotelRepository.findAll();
    return hotels;
}

export async function deleteHotelService(id: number) {
    const hotel = await hotelRepository.softDelete(id);
    return hotel;
}

export async function updateHotelService(id: number, hotelDto: createHotelDto) {
    const hotel = await hotelRepository.update(id, hotelDto);
    return hotel;
}
