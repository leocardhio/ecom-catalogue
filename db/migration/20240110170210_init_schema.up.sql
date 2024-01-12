CREATE TYPE "product_condition" AS ENUM (
  'Brand New',
  'Like New',
  'Lightly Used',
  'Well Used',
  'Heavily Used'
);

CREATE TABLE "products" (
  "ulid" varchar(26) PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL,
  "price" integer NOT NULL,
  "description" text NOT NULL,
  "condition" product_condition NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "users" (
  "ulid" varchar(26) PRIMARY KEY NOT NULL,
  "username" varchar(16) UNIQUE NOT NULL,
  "city" varchar NOT NULL,
  "rating" real,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reviews" (
  "from_user_ulid" varchar(26) NOT NULL,
  "to_user_ulid" varchar(26) NOT NULL,
  "review" text,
  "rating" smallint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  PRIMARY KEY ("from_user_ulid", "to_user_ulid")
);

CREATE TABLE "users_to_products" (
  "user_ulid" varchar(26) NOT NULL,
  "product_ulid" varchar(26) NOT NULL,
  PRIMARY KEY ("user_ulid", "product_ulid")
);

CREATE TABLE "product_images" (
  "product_ulid" varchar(26) NOT NULL,
  "image_url" varchar NOT NULL,
  PRIMARY KEY ("product_ulid", "image_url")
);

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("condition");

CREATE INDEX ON "products" ("price");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("city");

CREATE INDEX ON "users" ("rating");

CREATE INDEX ON "reviews" ("from_user_ulid");

CREATE INDEX ON "reviews" ("to_user_ulid");

CREATE INDEX ON "users_to_products" ("user_ulid");

CREATE INDEX ON "users_to_products" ("product_ulid");

CREATE INDEX ON "product_images" ("product_ulid");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("user_ulid") REFERENCES "users" ("ulid");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("product_ulid") REFERENCES "products" ("ulid");

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_ulid") REFERENCES "products" ("ulid");

ALTER TABLE "reviews" ADD FOREIGN KEY ("from_user_ulid") REFERENCES "users" ("ulid");

ALTER TABLE "reviews" ADD FOREIGN KEY ("to_user_ulid") REFERENCES "users" ("ulid");
