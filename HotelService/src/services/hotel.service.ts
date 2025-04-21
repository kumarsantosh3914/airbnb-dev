import { createHotelDto } from "../dto/hotel.dto";
import { createHotel, getHotelById } from "../repositories/hotel.repository";

export async function createHotelService(hotelDto: createHotelDto) {
    const hotel = await createHotel(hotelDto);
    return hotel;
}

export async function getHotelByIdService(id: number) {
    const hotel = await getHotelById(id);
    return hotel;
}