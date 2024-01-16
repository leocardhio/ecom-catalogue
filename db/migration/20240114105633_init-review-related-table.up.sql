CREATE TABLE "reviews" (
  "from_user_id" varchar(26) NOT NULL,
  "to_user_id" varchar(26) NOT NULL,
  "review" text,
  "rating" smallint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  PRIMARY KEY ("from_user_id", "to_user_id")
);

CREATE INDEX ON "reviews" ("from_user_id");

CREATE INDEX ON "reviews" ("to_user_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");