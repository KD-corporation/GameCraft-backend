/*
  Warnings:

  - Added the required column `Otp` to the `OTP` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "public"."OTP" ADD COLUMN     "Otp" TEXT NOT NULL;
