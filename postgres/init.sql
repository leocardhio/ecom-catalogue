CREATE TABLE "users" (
  "id" varchar(26) PRIMARY KEY NOT NULL,
  "username" varchar(16) UNIQUE NOT NULL,
  "city" varchar NOT NULL,
  "rating" real,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TYPE "product_condition" AS ENUM (
  'Brand New',
  'Like New',
  'Lightly Used',
  'Well Used',
  'Heavily Used'
);

CREATE TABLE "products" (
  "id" varchar(26) PRIMARY KEY NOT NULL,
  "name" varchar(50) NOT NULL,
  "price" integer NOT NULL,
  "is_sold" boolean NOT NULL DEFAULT (FALSE),
  "description" text NOT NULL,
  "condition" product_condition NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tags" (
  "id" serial PRIMARY KEY NOT NULL,
  "name" varchar UNIQUE NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "reviews" (
  "from_user_id" varchar(26) NOT NULL,
  "to_user_id" varchar(26) NOT NULL,
  "review" text,
  "rating" smallint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  PRIMARY KEY ("from_user_id", "to_user_id")
);

CREATE TABLE "users_to_products" (
  "user_id" varchar(26) NOT NULL,
  "product_id" varchar(26) NOT NULL,
  PRIMARY KEY ("user_id", "product_id")
);

CREATE TABLE "product_images" (
  "product_id" varchar(26) NOT NULL,
  "image_url" varchar NOT NULL,
  PRIMARY KEY ("product_id", "image_url")
);

CREATE TABLE "product_tags" (
  "product_id" varchar NOT NULL,
  "tag_id" smallint NOT NULL,
  PRIMARY KEY ("product_id", "tag_id")
);


CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("city");

CREATE INDEX ON "users" ("rating");

CREATE INDEX ON "reviews" ("to_user_id");

CREATE INDEX ON "products" ("condition");

CREATE INDEX ON "products" ("price");

CREATE INDEX ON "products" ("is_sold");

CREATE INDEX ON "tags" ("name");

CREATE INDEX ON "users_to_products" ("user_id");

CREATE INDEX ON "users_to_products" ("product_id");

CREATE INDEX ON "product_images" ("product_id");

CREATE INDEX ON "product_tags" ("product_id");

CREATE INDEX ON "product_tags" ("tag_id");

CREATE INDEX ON "reviews" ("from_user_id");

CREATE INDEX ON "products" ("name");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_to_products" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "product_tags" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");
