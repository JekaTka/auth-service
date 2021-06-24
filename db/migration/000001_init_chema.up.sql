CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "nickname" varchar UNIQUE NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "created_at" timestamp DEFAULT (now()),
                         "updated_at" timestamp DEFAULT (now()),
                         "deleted_at" timestamp
);

CREATE INDEX ON "users" ("nickname");

CREATE INDEX ON "users" ("email");
