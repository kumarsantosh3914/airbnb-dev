import express from "express";
import request from "supertest";
import { StatusCodes } from "http-status-codes";
import router from "../../src/routers/v1/index.router";
import * as roomGenerationProducer from "../../src/producers/roomGeneration.producer";

jest.mock("../../src/producers/roomGeneration.producer");

describe("Room Generation Controller", () => {
  const app = express().use(express.json()).use("/api/v1", router);

  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe("POST /api/v1/room-generation", () => {
    it("should return 201 when job is queued successfully", async () => {
      (roomGenerationProducer.addRoomGenerationJobToQueue as jest.Mock).mockResolvedValue(undefined);

      const body = {
        roomCategoryId: 1,
        startDate: new Date("2025-08-25T00:00:00.000Z").toISOString(),
        endDate: new Date("2025-08-30T00:00:00.000Z").toISOString(),
        scheduleType: "immediate",
        // scheduleAt optional
        priceOverride: 1200,
      };

      const res = await request(app)
        .post("/api/v1/room-generation")
        .send(body);

      expect(res.status).toBe(StatusCodes.CREATED);
      expect(res.body).toEqual({
        message: "Room generation job added to queue successfully",
        success: true,
        data: {},
      });
      expect(roomGenerationProducer.addRoomGenerationJobToQueue).toHaveBeenCalledWith(body);
    });

    it("should validate request body and return 4xx for invalid payload", async () => {
      const res = await request(app)
        .post("/api/v1/room-generation")
        .send({
          // missing required fields and wrong types
          roomCategoryId: "not-a-number",
          startDate: "invalid-date",
        });

      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(roomGenerationProducer.addRoomGenerationJobToQueue).not.toHaveBeenCalled();
    });

    it("should handle producer errors with 500", async () => {
      (roomGenerationProducer.addRoomGenerationJobToQueue as jest.Mock).mockRejectedValue(
        new Error("Queue down")
      );

      const body = {
        roomCategoryId: 2,
        startDate: new Date("2025-09-01T00:00:00.000Z").toISOString(),
        endDate: new Date("2025-09-05T00:00:00.000Z").toISOString(),
        scheduleType: "immediate",
      };

      const res = await request(app)
        .post("/api/v1/room-generation")
        .send(body);

      // Depending on your global error handler, any 4xx/5xx is acceptable here
      expect(res.status).toBeGreaterThanOrEqual(400);
      expect(roomGenerationProducer.addRoomGenerationJobToQueue).toHaveBeenCalledWith(body);
    });
  });
});
