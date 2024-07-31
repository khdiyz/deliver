-- +goose Up
CREATE TABLE "attributes" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL
);

-- +goose Down
DROP TABLE "attributes";