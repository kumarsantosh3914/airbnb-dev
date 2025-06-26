import Room from "../db/models/room";
import BaseRepository from "./base.repository";

class RoomRespository extends BaseRepository<Room> {
    constructor() {
        super(Room);
    }

}

export default RoomRespository;