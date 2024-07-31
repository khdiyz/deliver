-- +goose Up
CREATE TABLE "categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL
);

-- +goose Down
DROP TABLE "categories";