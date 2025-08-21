import express from "express";
import request from "supertest";
import * as hotelService from "../../src/services/hotel.service";
import router from "../../src/routers/v1/index.router";

jest.mock("../../src/services/hotel.service");

describe("Hotel Controller", () => {
  const app = express().use(express.json()).use("/api/v1", router);

  it("POST /api/v1/hotels return 201 with created hotel", async () => {
    (hotelService.createHotelService as jest.Mock).mockResolvedValue({
      id: 1,
      name: "Hotel A",
    });

    const res = await request(app).post("/api/v1/hotels").send({
      name: "Hotel A",
      address: "123 Main St",
      location: "Delhi",
      rating: 4.5,
      ratingCount: 10,
    });

    expect(res.status).toBe(201);
    expect(res.body?.data).toMatchObject({ id: 1, name: "Hotel A" });
    expect(hotelService.createHotelService).toHaveBeenCalledWith(
      expect.objectContaining({ name: "Hotel A" })
    );
  });

  it("handles service errors -> 400/500", async () => {
    (hotelService.createHotelService as jest.Mock).mockRejectedValue(
      new Error("Bad")
    );

    const res = await request(app).post("/api/v1/hotels").send({ name: "X" });

    // assert your error middleware mapping
    expect(res.status).toBeGreaterThanOrEqual(400);
  });

  describe("GET /api/v1/hotels/:id", () => {
    it("should return 200 with hotel data", async () => {
      const mockHotel = { id: 1, name: "Hotel A", location: "New Delhi" };
      (hotelService.getHotelByIdService as jest.Mock).mockResolvedValue(
        mockHotel
      );

      const res = await request(app).get("/api/v1/hotels/1");

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: "Hotel fetched successfully",
        data: mockHotel,
        success: true,
      });
      expect(hotelService.getHotelByIdService).toHaveBeenCalledWith(1);
    });
  });

  describe("GET /api/v1/hotels", () => {
    it("should return 200 with hotel data", async () => {
      const mockHotel = { id: 1, name: "Hotel A", location: "New Delhi" };
      (hotelService.getHotelByIdService as jest.Mock).mockResolvedValue(
        mockHotel
      );

      const res = await request(app).get("/api/v1/hotels/1");

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: "Hotel fetched successfully",
        data: mockHotel,
        success: true,
      });
      expect(hotelService.getHotelByIdService).toHaveBeenCalledWith(1);
    });

    it("should handle invalid hotel ID -> 400/500", async () => {
      (hotelService.getHotelByIdService as jest.Mock).mockRejectedValue(
        new Error("Hotel not found")
      );

      const res = await request(app).get("/api/v1/hotels/9999999");

      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(hotelService.getHotelByIdService).toHaveBeenLastCalledWith(
        9999999
      );
    });

    it("shuld convert string ID to number", async () => {
      const mockHotel = { id: 323, name: "Test Hotel" };
      (hotelService.getHotelByIdService as jest.Mock).mockResolvedValue(
        mockHotel
      );

      await request(app).get("/api/v1/hotels/323");

      expect(hotelService.getHotelByIdService).toHaveBeenCalledWith(323);
    });
  });

  describe("GET /api/v1/hotels", () => {
    it("should return 200 with all hotels", async () => {
      const mockHotels = [
        { id: 1, name: "Hotel A" },
        { id: 2, name: "Hotel B" },
      ];
      (hotelService.getAllHotelsService as jest.Mock).mockResolvedValue(
        mockHotels
      );

      const res = await request(app).get("/api/v1/hotels");

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: "Hotels fetched successfully",
        data: mockHotels,
        success: true,
      });
      expect(hotelService.getAllHotelsService).toHaveBeenCalledWith();
    });

    it("should return empty array no hotels exists", async () => {
      (hotelService.getAllHotelsService as jest.Mock).mockResolvedValue([]);

      const res = await request(app).get("/api/v1/hotels");

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: "Hotels fetched successfully",
        data: [],
        success: true,
      });
    });
  });

  describe('DELETE /api/v1/hotels/:id', () => {
    it('should return 200 when hotel is deleted successfully', async () => {
      const mockDeletedHotel = { id: 1, name: 'Deleted Hotel' };
      (hotelService.deleteHotelService as jest.Mock).mockResolvedValue(mockDeletedHotel);

      const res = await request(app)
        .delete('/api/v1/hotels/1');

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: 'Hotel deleted successfully',
        data: mockDeletedHotel,
        success: true,
      });
      expect(hotelService.deleteHotelService).toHaveBeenCalledWith(1);
    });

    it('should handle non-existent hotel deletion', async () => {
      (hotelService.deleteHotelService as jest.Mock).mockRejectedValue(
        new Error('Hotel not found')
      );

      const res = await request(app)
        .delete('/api/v1/hotels/999');
      
      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(hotelService.deleteHotelService).toHaveBeenCalledWith(999);
    });

    it('should convert string ID to number', async () => {
      const mockDeletedHotel = { id: 456, name: 'Hotel A' };
      (hotelService.deleteHotelService as jest.Mock).mockResolvedValue(mockDeletedHotel);

      await request(app)
        .delete('/api/v1/hotels/456');
      
      expect(hotelService.deleteHotelService).toHaveBeenCalledWith(456);
    });
  });

  describe('PUT /api/v1/hotels/:id', () => {
    it('should return 200 with updated hotel data', async () => {
      const updatedData = {
        name: 'Updated Hotel',
        address: '456 Updated St',
        location: 'NYC',
        rating: 4.8,
        ratingCount: 20
      };
      const mockUpdatedHotel = { id: 1, ...updatedData };
      (hotelService.updateHotelService as jest.Mock).mockResolvedValue(mockUpdatedHotel);

      const res = await request(app)
        .put('/api/v1/hotels/1')
        .send(updatedData);

      expect(res.status).toBe(200);
      expect(res.body).toEqual({
        message: 'Hotel updated successfully',
        data: mockUpdatedHotel,
        success: true,
      });
      expect(hotelService.updateHotelService).toHaveBeenCalledWith(1, updatedData);
    });

    it('should handle update of non-existent hotel', async () => {
      (hotelService.updateHotelService as jest.Mock).mockRejectedValue(
        new Error('Hotel not found')
      );

      const res = await request(app)
        .put('/api/v1/hotels/999')
        .send({
          name: 'Non-existent Hotel',
          address: '123 Fake St',
          location: 'Nowhere',
          rating: 4.0,
          ratingCount: 5
        });

      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(hotelService.updateHotelService).toHaveBeenCalledWith(999, {
        name: 'Non-existent Hotel',
        address: '123 Fake St',
        location: 'Nowhere',
        rating: 4.0,
        ratingCount: 5
      });
    });

    it('should convert string ID to number', async () => {
      const updatedData = {
        name: 'Test update',
        address: '789 Test Ave',
        location: 'Test City',
        rating: 4.2,
        ratingCount: 15
      };
      const mockUpdatedHotel = { id: 789, ...updatedData };
      (hotelService.updateHotelService as jest.Mock).mockResolvedValue(mockUpdatedHotel);

      await request(app)
        .put('/api/v1/hotels/789')
        .send(updatedData);

      expect(hotelService.updateHotelService).toHaveBeenCalledWith(789, updatedData);
    });

    it('should handle validation errors for invalid data', async () => {
      // Clear previous mock calls for this specific test
      jest.clearAllMocks();

      const res = await request(app)
        .put('/api/v1/hotels/1')
        .send({ name: '' }); // Invalid data - missing required fields

      expect(res.status).toBeGreaterThanOrEqual(400);
      // The service shouldn't be called if validation fails
      expect(hotelService.updateHotelService).not.toHaveBeenCalled();
    });

    it('should handle service validation errors', async () => {
      (hotelService.updateHotelService as jest.Mock).mockRejectedValue(
        new Error('Invalid update data')
      );

      const validData = {
        name: 'Valid Hotel',
        address: '123 Valid St',
        location: 'Valid City',
        rating: 4.0,
        ratingCount: 10
      };

      const res = await request(app)
        .put('/api/v1/hotels/1')
        .send(validData);

      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(hotelService.updateHotelService).toHaveBeenCalledWith(1, validData);
    });
  });
});
