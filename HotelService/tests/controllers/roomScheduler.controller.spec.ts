import express from "express";
import request from "supertest";
import { StatusCodes } from "http-status-codes";
import router from "../../src/routers/v1/index.router";
import * as roomScheduler from "../../src/schedular/roomScheduler";

jest.mock("../../src/schedular/roomScheduler");

describe("Room Scheduler Controller", () => {
  const app = express().use(express.json()).use("/api/v1", router);

  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe("POST /api/v1/scheduler/start", () => {
    it("should return 200 and start scheduler", async () => {
      (roomScheduler.startScheduler as jest.Mock).mockImplementation(() => {});

      const res = await request(app).post("/api/v1/scheduler/start");

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        message: "Room availability extension scheduler started successfully",
        success: true,
        data: { status: "started" },
      });
      expect(roomScheduler.startScheduler).toHaveBeenCalled();
    });

    it("should handle start errors -> 500", async () => {
      (roomScheduler.startScheduler as jest.Mock).mockImplementation(() => {
        throw new Error("Failed to start");
      });

      const res = await request(app).post("/api/v1/scheduler/start");

      expect(res.status).toBe(StatusCodes.INTERNAL_SERVER_ERROR);
      expect(roomScheduler.startScheduler).toHaveBeenCalled();
    });
  });

  describe("POST /api/v1/scheduler/stop", () => {
    it("should return 200 and stop scheduler", async () => {
      (roomScheduler.stopScheduler as jest.Mock).mockImplementation(() => {});

      const res = await request(app).post("/api/v1/scheduler/stop");

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        message: "Room availability extension scheduler stopped successfully",
        success: true,
        data: { status: "stopped" },
      });
      expect(roomScheduler.stopScheduler).toHaveBeenCalled();
    });

    it("should handle stop errors -> 500", async () => {
      (roomScheduler.stopScheduler as jest.Mock).mockImplementation(() => {
        throw new Error("Failed to stop");
      });

      const res = await request(app).post("/api/v1/scheduler/stop");

      expect(res.status).toBe(StatusCodes.INTERNAL_SERVER_ERROR);
      expect(roomScheduler.stopScheduler).toHaveBeenCalled();
    });
  });

  describe("GET /api/v1/scheduler/status", () => {
    it("should return 200 with scheduler status", async () => {
      (roomScheduler.getSchedulerStatus as jest.Mock).mockReturnValue({ isRunning: true });

      const res = await request(app).get("/api/v1/scheduler/status");

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        message: "Scheduler status retrieved successfully",
        success: true,
        data: { isRunning: true },
      });
      expect(roomScheduler.getSchedulerStatus).toHaveBeenCalled();
    });

    it("should handle status errors -> 500", async () => {
      (roomScheduler.getSchedulerStatus as jest.Mock).mockImplementation(() => {
        throw new Error("Status error");
      });

      const res = await request(app).get("/api/v1/scheduler/status");

      expect(res.status).toBe(StatusCodes.INTERNAL_SERVER_ERROR);
      expect(roomScheduler.getSchedulerStatus).toHaveBeenCalled();
    });
  });

  describe("POST /api/v1/scheduler/extend", () => {
    it("should return 200 when manual extension completes", async () => {
      (roomScheduler.manualExtendAvailability as jest.Mock).mockResolvedValue(undefined);

      const res = await request(app).post("/api/v1/scheduler/extend");

      expect(res.status).toBe(StatusCodes.OK);
      expect(res.body).toEqual({
        message: "Manual room availability extension completed successfully",
        success: true,
        data: { action: "manual_extension_completed" },
      });
      expect(roomScheduler.manualExtendAvailability).toHaveBeenCalled();
    });

    it("should handle manual extension errors -> 500", async () => {
      (roomScheduler.manualExtendAvailability as jest.Mock).mockRejectedValue(new Error("Extend error"));

      const res = await request(app).post("/api/v1/scheduler/extend");

      expect(res.status).toBe(StatusCodes.INTERNAL_SERVER_ERROR);
      expect(roomScheduler.manualExtendAvailability).toHaveBeenCalled();
    });
  });
});
