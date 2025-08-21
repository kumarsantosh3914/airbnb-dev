import express from "express";
import request from "supertest";
import { StatusCodes } from "http-status-codes";
import * as roomService from "../../src/services/room.service";
import router from "../../src/routers/v1/index.router";

jest.mock("../../src/services/room.service");

describe("Room Controller", () => {
  const app = express().use(express.json()).use("/api/v1", router);

  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe("GET /api/v1/rooms/available", () => {
    it("should return 200 with available rooms", async () => {
      const mockRooms = [
        { id: 1, roomNumber: "101", roomCategoryId: 1, isAvailable: true },
        { id: 2, roomNumber: "102", roomCategoryId: 1, isAvailable: true },
        { id: 3, roomNumber: "201", roomCategoryId: 2, isAvailable: true },
      ];

      (roomService.getAvailableRoomsService as jest.Mock).mockResolvedValue(mockRooms);

      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "1",
          checkInDate: "2025-08-25",
          checkOutDate: "2025-08-30",
        });

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        success: true,
        message: "Available rooms fetched successfully",
        data: mockRooms,
      });
      expect(roomService.getAvailableRoomsService).toHaveBeenCalledWith({
        roomCategoryId: 1,
        checkInDate: new Date("2025-08-25"),
        checkOutDate: new Date("2025-08-30"),
      });
    });

    it("should return empty array when no rooms are available", async () => {
      (roomService.getAvailableRoomsService as jest.Mock).mockResolvedValue([]);

      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "2",
          checkInDate: "2025-12-24",
          checkOutDate: "2025-12-26",
        });

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        success: true,
        message: "Available rooms fetched successfully",
        data: [],
      });
    });

    it("should convert string parameters to correct types", async () => {
      const mockRooms = [
        { id: 5, roomNumber: "505", roomCategoryId: 3, isAvailable: true },
      ];

      (roomService.getAvailableRoomsService as jest.Mock).mockResolvedValue(mockRooms);

      await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "3",
          checkInDate: "2025-09-01",
          checkOutDate: "2025-09-05",
        });

      expect(roomService.getAvailableRoomsService).toHaveBeenCalledWith({
        roomCategoryId: 3, // Should be converted to number
        checkInDate: new Date("2025-09-01"), // Should be converted to Date
        checkOutDate: new Date("2025-09-05"), // Should be converted to Date
      });
    });

    it("should handle service errors", async () => {
      (roomService.getAvailableRoomsService as jest.Mock).mockRejectedValue(
        new Error("Database connection failed")
      );

      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "1",
          checkInDate: "2025-08-25",
          checkOutDate: "2025-08-30",
        });

      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(roomService.getAvailableRoomsService).toHaveBeenCalled();
    });

    it("should handle invalid date formats", async () => {
      (roomService.getAvailableRoomsService as jest.Mock).mockRejectedValue(
        new Error("Invalid date format")
      );

      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "1",
          checkInDate: "invalid-date",
          checkOutDate: "2025-08-30",
        });

      expect(res.status).toBeGreaterThanOrEqual(400);
    });

    it("should handle missing query parameters", async () => {
      const res = await request(app)
        .get("/api/v1/rooms/available");

      expect(res.status).toBeGreaterThanOrEqual(400);
    });

    it("should handle non-numeric roomCategoryId", async () => {
      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "abc",
          checkInDate: "2025-08-25",
          checkOutDate: "2025-08-30",
        });

      // This will call the service with NaN, which should be handled
      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(roomService.getAvailableRoomsService).toHaveBeenCalledWith({
        roomCategoryId: NaN,
        checkInDate: new Date("2025-08-25"),
        checkOutDate: new Date("2025-08-30"),
      });
    });

    it("should handle date range edge cases", async () => {
      const mockRooms = [
        { id: 1, roomNumber: "101", roomCategoryId: 1, isAvailable: true },
      ];

      (roomService.getAvailableRoomsService as jest.Mock).mockResolvedValue(mockRooms);

      // Same check-in and check-out date (single day booking)
      const res = await request(app)
        .get("/api/v1/rooms/available")
        .query({
          roomCategoryId: "1",
          checkInDate: "2025-08-25",
          checkOutDate: "2025-08-25",
        });

      expect(res.status).toBe(StatusCodes.OK);
      expect(roomService.getAvailableRoomsService).toHaveBeenCalledWith({
        roomCategoryId: 1,
        checkInDate: new Date("2025-08-25"),
        checkOutDate: new Date("2025-08-25"),
      });
    });
  });

//   describe("POST /api/v1/rooms/update-booking-id", () => {
//     it("should return 200 when booking ID is updated successfully", async () => {
//       const requestBody = {
//         bookingId: "12345", // String as per schema
//         roomIds: [1, 2, 3],
//       };

//       const mockResult = {
//         updatedRooms: [
//           { id: 1, roomNumber: "101", bookingId: 12345 },
//           { id: 2, roomNumber: "102", bookingId: 12345 },
//           { id: 3, roomNumber: "103", bookingId: 12345 },
//         ],
//         totalUpdated: 3,
//       };

//       (roomService.updateBookingIdToRoomsService as jest.Mock).mockResolvedValue(mockResult);

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBe(StatusCodes.OK);
//       expect(res.body).toEqual({
//         success: true,
//         message: "Booking ID updated to rooms successfully",
//         data: mockResult,
//       });
//       expect(roomService.updateBookingIdToRoomsService).toHaveBeenCalledWith(requestBody);
//     });

//     it("should handle single room update", async () => {
//       const requestBody = {
//         bookingId: "98765", // String as per schema
//         roomIds: [1],
//       };

//       const mockResult = {
//         updatedRooms: [
//           { id: 1, roomNumber: "101", bookingId: 98765 },
//         ],
//         totalUpdated: 1,
//       };

//       (roomService.updateBookingIdToRoomsService as jest.Mock).mockResolvedValue(mockResult);

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBe(StatusCodes.OK);
//       expect(res.body.data).toEqual(mockResult);
//       expect(roomService.updateBookingIdToRoomsService).toHaveBeenCalledWith(requestBody);
//     });

//     it("should handle empty room IDs array", async () => {
//       const requestBody = {
//         bookingId: "12345", // String as per schema
//         roomIds: [],
//       };

//       const mockResult = {
//         updatedRooms: [],
//         totalUpdated: 0,
//       };

//       (roomService.updateBookingIdToRoomsService as jest.Mock).mockResolvedValue(mockResult);

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBe(StatusCodes.OK);
//       expect(res.body.data).toEqual(mockResult);
//     });

//     it("should handle service errors", async () => {
//       const requestBody = {
//         bookingId: "12345",
//         roomIds: [1, 2],
//       };

//       (roomService.updateBookingIdToRoomsService as jest.Mock).mockRejectedValue(
//         new Error("Room not found")
//       );

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBeGreaterThanOrEqual(400);
//       expect(roomService.updateBookingIdToRoomsService).toHaveBeenCalledWith(requestBody);
//     });

//     it("should handle invalid booking ID", async () => {
//       const requestBody = {
//         bookingId: null,
//         roomIds: [1, 2],
//       };

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBeGreaterThanOrEqual(400);
//       // Service shouldn't be called if validation fails
//       expect(roomService.updateBookingIdToRoomsService).not.toHaveBeenCalled();
//     });

//     it("should handle invalid room IDs", async () => {
//       const requestBody = {
//         bookingId: "12345",
//         roomIds: ["invalid", "room", "ids"], // Should be numbers
//       };

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBeGreaterThanOrEqual(400);
//       // Service shouldn't be called if validation fails
//       expect(roomService.updateBookingIdToRoomsService).not.toHaveBeenCalled();
//     });

//     it("should handle missing request body", async () => {
//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send({});

//       expect(res.status).toBeGreaterThanOrEqual(400);
//       // Service shouldn't be called if validation fails
//       expect(roomService.updateBookingIdToRoomsService).not.toHaveBeenCalled();
//     });

//     it("should handle partial request body", async () => {
//       const requestBody = {
//         roomIds: [1, 2, 3],
//         // missing bookingId
//       };

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBeGreaterThanOrEqual(400);
//       // Service shouldn't be called if validation fails
//       expect(roomService.updateBookingIdToRoomsService).not.toHaveBeenCalled();
//     });

//     it("should handle database constraint errors", async () => {
//       const requestBody = {
//         bookingId: "12345",
//         roomIds: [999, 1000], // Non-existent room IDs
//       };

//       (roomService.updateBookingIdToRoomsService as jest.Mock).mockRejectedValue(
//         new Error("Rooms not found")
//       );

//       const res = await request(app)
//         .post("/api/v1/rooms/update-booking-id")
//         .send(requestBody);

//       expect(res.status).toBeGreaterThanOrEqual(400);
//     });
//   });
});