export type CreateBookingDto = {
    userId: number;
    hotelId: number;
    totalGuests: number;
    bookingAmount: number;
    checkInDate: string;
    checkOutDate: string;
    roomCategoryId: number;
}