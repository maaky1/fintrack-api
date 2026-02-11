CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "clerk_user_id" text UNIQUE NOT NULL,
  "fullname" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);