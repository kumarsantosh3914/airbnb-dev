import axios from "axios";
import { serverConfig } from "../config";

export const getAvailableRooms = async (roomCategoryId: number, checkInDate: string, checkOutDate: string) => {
    const response = await axios.get(`${serverConfig.HOTEL_SERVICE_URL}rooms/available`, {
        params: {
            roomCategoryId, // will be sent as a number
            checkInDate,
            checkOutDate
        },
    });

    return response.data;
}

export const updateBookingIdToRooms = async (bookingId: string, roomIds: number[]) => {
    const response = await axios.post(`${serverConfig.HOTEL_SERVICE_URL}rooms/update-booking-id`, {
        bookingId,
        roomIds
    });

    return response.data;
}