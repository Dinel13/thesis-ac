CREATE TABLE "krss" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "teacher" INT NOT NULL,
  "capacity" INT NOT NULL,
  "remain" INT NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "krss" ("id");