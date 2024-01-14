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

CREATE INDEX ON "users_to_products" ("user_ulid");

CREATE INDEX ON "users_to_products" ("product_ulid");

CREATE INDEX ON "product_images" ("product_ulid");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("user_ulid") REFERENCES "users" ("ulid");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("product_ulid") REFERENCES "products" ("ulid");

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_ulid") REFERENCES "products" ("ulid");