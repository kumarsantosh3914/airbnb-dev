import { createHotelDto } from "../dto/hotel.dto";
import { createHotel, deleteHotel, getAllHotels, getHotelById, updateHotel } from "../repositories/hotel.repository";

export async function createHotelService(hotelDto: createHotelDto) {
    const hotel = await createHotel(hotelDto);
    return hotel;
}

export async function getHotelByIdService(id: number) {
    const hotel = await getHotelById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotels = await getAllHotels();
    return hotels;
}

export async function deleteHotelService(id: number) {
    const hotel = await deleteHotel(id);
    return hotel;
}

export async function updateHotelService(id: number, hotelDto: createHotelDto) {
    const hotel = await updateHotel(id, hotelDto);
    return hotel;
}
