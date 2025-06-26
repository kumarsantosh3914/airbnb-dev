import { CreateRoomCategoryDTO } from "../dto/roomCategory.dto";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";

const roomCategoryRepository = new RoomCategoryRepository();
const hotelRepository = new RoomCategoryRepository();

export async function createRoomCategoryService(
  createRoomCategoryDTO: CreateRoomCategoryDTO
) {
  const roomCategory = await roomCategoryRepository.create(
    createRoomCategoryDTO
  );
  return roomCategory;
}

export async function getRoomCategoryByIdService(id: number) {
  const roomCategory = await roomCategoryRepository.findById(id);
  return roomCategory;
}

export async function getAllRoomCategoriesByHotelIdService(hotelId: number) {
  // Check if the hotel exists
  const hotel = await hotelRepository.findById(hotelId);

  if (!hotel) {
    throw new NotFoundError(`Hotel with ID ${hotelId} does not exist.`);
  }

  // find all room categories by hotel ID
  const roomCategories = await roomCategoryRepository.findAllByHotelId(hotelId);
  return roomCategories;
}

export async function deleteRoomCategoryService(id: number) {
  const roomCategory = await roomCategoryRepository.findById(id);

  if (!roomCategory) {
    throw new NotFoundError(`Room category with ID ${id} does not exist.`);
  }

  await roomCategoryRepository.delete({ id });
  return true;
}
