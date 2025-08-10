export type GetAvailableRoomsDTO = {
    roomCategoryId: number;
    checkInDate: Date;
    checkOutDate: Date;
}

export type UpdateBookingIdToRoomsDTO = {
    bookingId: number;
    roomIds: number[];
}