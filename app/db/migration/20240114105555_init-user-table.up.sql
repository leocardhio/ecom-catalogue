CREATE TABLE "users" (
  "id" varchar(26) PRIMARY KEY NOT NULL,
  "username" varchar(16) UNIQUE NOT NULL,
  "city" varchar NOT NULL,
  "rating" real,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("city");

CREATE INDEX ON "users" ("rating");