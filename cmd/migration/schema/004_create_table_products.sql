-- +goose Up
CREATE TABLE "products" (
    "id" BIGSERIAL PRIMARY KEY,
    "category_id" BIGINT NOT NULL,
    "name" VARCHAR(64) NOT NULL,
    "description" TEXT,
    "photo" VARCHAR(64),
    "price" INTEGER,
    FOREIGN KEY ("category_id") REFERENCES "categories"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "products";