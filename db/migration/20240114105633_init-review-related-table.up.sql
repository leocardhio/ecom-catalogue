CREATE TABLE "reviews" (
  "from_user_ulid" varchar(26) NOT NULL,
  "to_user_ulid" varchar(26) NOT NULL,
  "review" text,
  "rating" smallint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  PRIMARY KEY ("from_user_ulid", "to_user_ulid")
);

CREATE INDEX ON "reviews" ("from_user_ulid");

CREATE INDEX ON "reviews" ("to_user_ulid");

ALTER TABLE "reviews" ADD FOREIGN KEY ("from_user_ulid") REFERENCES "users" ("ulid");

ALTER TABLE "reviews" ADD FOREIGN KEY ("to_user_ulid") REFERENCES "users" ("ulid");