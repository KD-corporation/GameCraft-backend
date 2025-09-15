-- CreateTable
CREATE TABLE "public"."OTP" (
    "id" SERIAL NOT NULL,
    "Email" TEXT NOT NULL,
    "expiresAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "OTP_pkey" PRIMARY KEY ("id")
);
