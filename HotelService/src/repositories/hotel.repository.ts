import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel> {
    constructor() {
        super(Hotel);
    }

    async findAll() {
        const hotels = await this.model.findAll({
            where: {
                deletedAt: null // Assuming soft delete is implemented
            }
        });

        if(!hotels) {
            logger.error("No hotels found");
            throw new NotFoundError("No hotels found");
        }

        logger.info(`Hotels found: ${hotels.length}`);
        return hotels;
    }

    async softDelete(id: number) {
        const hotel = await this.model.findByPk(id);

        if(!hotel) {
            logger.error(`Hotel not found for deletion: ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }

        hotel.deletedAt = new Date(); // Assuming deletedAt is a Date field for soft delete
        await hotel.save();
        logger.info(`Hotel soft deleted: ${hotel.id}`);
        return true;
    }
}