import { z } from 'zod';

export const createBookingSchema = z.object({
    userId: z.number({message: "User ID must be a present"}),
    hotelId: z.number({message: "Hotel ID must be a present"}),
    totalGuests: z.number({message: "Total guests must be a present"}).min(1, {message: "Total guests must be at least 1"}),
    bookingAmount: z.number({message: "Booking amount must be a present"}).min(0, {message: "Booking amount must be at least 1"}),
});

