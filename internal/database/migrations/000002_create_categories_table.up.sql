CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE UNIQUE INDEX ON "categories" ("user_id", "name");

ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");