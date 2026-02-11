CREATE TABLE "transactions" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "category_id" int NOT NULL,
  "type" text NOT NULL,
  "amount" numeric(14,2) NOT NULL,
  "description" text,
  "occurred_at" timestamptz NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "transactions" ("user_id", "type");

CREATE INDEX ON "transactions" ("user_id", "occurred_at");

CREATE INDEX ON "transactions" ("category_id", "occurred_at");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");