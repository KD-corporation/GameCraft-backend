/*
  Warnings:

  - You are about to drop the column `expiresAt` on the `OTP` table. All the data in the column will be lost.
  - Added the required column `ExpiresAt` to the `OTP` table without a default value. This is not possible if the table is not empty.
  - Added the required column `FirstName` to the `OTP` table without a default value. This is not possible if the table is not empty.
  - Added the required column `LastName` to the `OTP` table without a default value. This is not possible if the table is not empty.
  - Added the required column `Password` to the `OTP` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "public"."OTP" DROP COLUMN "expiresAt",
ADD COLUMN     "ExpiresAt" TIMESTAMP(3) NOT NULL,
ADD COLUMN     "FirstName" TEXT NOT NULL,
ADD COLUMN     "LastName" TEXT NOT NULL,
ADD COLUMN     "Password" TEXT NOT NULL;
