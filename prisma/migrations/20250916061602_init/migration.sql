/*
  Warnings:

  - A unique constraint covering the columns `[Username]` on the table `OTP` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[Email]` on the table `OTP` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[Username]` on the table `User` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "OTP_Username_key" ON "public"."OTP"("Username");

-- CreateIndex
CREATE UNIQUE INDEX "OTP_Email_key" ON "public"."OTP"("Email");

-- CreateIndex
CREATE UNIQUE INDEX "User_Username_key" ON "public"."User"("Username");
