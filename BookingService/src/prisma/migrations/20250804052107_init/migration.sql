/*
  Warnings:

  - Added the required column `checkInDate` to the `booking` table without a default value. This is not possible if the table is not empty.
  - Added the required column `checkOutDate` to the `booking` table without a default value. This is not possible if the table is not empty.
  - Added the required column `roomCategoryId` to the `booking` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `booking` ADD COLUMN `checkInDate` DATETIME(3) NOT NULL,
    ADD COLUMN `checkOutDate` DATETIME(3) NOT NULL,
    ADD COLUMN `roomCategoryId` INTEGER NOT NULL;
