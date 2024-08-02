-- +goose Up
CREATE TABLE "options" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "attribute_id" BIGINT NOT NULL,
    FOREIGN KEY ("attribute_id") REFERENCES "attributes"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "options";