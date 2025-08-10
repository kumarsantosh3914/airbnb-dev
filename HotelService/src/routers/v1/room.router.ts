import express from "express";
import { getAvailableRoomsHandler } from "../../controllers/room.controller";
import { validateQueryParams, validateRequestBody } from "../../validators";
import {
  getAvailableRoomsSchema,
  updateBookingIdToRoomsSchema,
} from "../../validators/room.validator";

const roomRouter = express.Router();

roomRouter.get(
  "/available",
  validateQueryParams(getAvailableRoomsSchema),
  getAvailableRoomsHandler
);
roomRouter.post(
  "/update-booking-id", 
  validateRequestBody(updateBookingIdToRoomsSchema),
  getAvailableRoomsHandler
);

roomRouter.get('/helth-check', (req, res) => {
  res.status(200).json({ message: "Room service is running" });
});

export default roomRouter;
