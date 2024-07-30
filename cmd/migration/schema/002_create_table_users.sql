-- +goose Up
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "full_name" VARCHAR(64) NOT NULL,
    "email" VARCHAR(64) NOT NULL,
    "password" TEXT NOT NULL,
    "role_id" BIGINT NOT NULL,
    FOREIGN KEY ("role_id") REFERENCES "roles"("id")
);

CREATE UNIQUE INDEX unique_email ON "users" ("email");

INSERT INTO "users" (
    "full_name",
    "email",
    "password",
    "role_id"
) VALUES (
    'Diyorbek Hasanov',
    'khdiyz.12@gmail.com',
    '38393071776572747975696f704153444647484a60a83648193440eb0b6cb9210f2bff80e2f85462',
    (SELECT "id" FROM "roles" WHERE "name" = 'ADMIN')
);

-- +goose Down
DROP TABLE "users";