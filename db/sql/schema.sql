-- Enable UUID extension 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Create pokemon table
CREATE TABLE "pokemon" (
    "id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    PRIMARY KEY ("id")
);
-- Create vote table with foreign key references to pokemon and default random ID
CREATE TABLE "vote" (
    "id" TEXT NOT NULL DEFAULT (uuid_generate_v4()),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "voted_for_id" INTEGER NOT NULL,
    "voted_against_id" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "vote_voted_for_id_fkey" FOREIGN KEY ("voted_for_id") REFERENCES "pokemon"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "vote_voted_against_id_fkey" FOREIGN KEY ("voted_against_id") REFERENCES "pokemon"("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
-- Create indexes for better query performance
CREATE INDEX "vote_voted_for_id_idx" ON "vote"("voted_for_id");
CREATE INDEX "vote_voted_against_id_idx" ON "vote"("voted_against_id");