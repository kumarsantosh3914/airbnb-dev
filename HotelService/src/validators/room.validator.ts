import { z } from "zod";

export const getAvailableRoomsSchema = z.object({
    roomCategoryId: z.string({message: "Room category ID is must be present"}),
    checkInDate: z.string({message: "Check-in date must be a valid date"}),
    checkOutDate: z.string({message: "Check-out date must be a valid date"}),
})

export const updateBookingIdToRoomsSchema = z.object({
    bookingId: z.string({message: "Booking ID is required"}),
    roomIds: z.array(z.number(), {message: "Room IDs must be an array of numbers"})
});