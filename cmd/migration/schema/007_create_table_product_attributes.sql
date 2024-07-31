-- +goose Up
CREATE TABLE "product_attributes" (
    "id" BIGSERIAL PRIMARY KEY,
    "product_id" BIGINT NOT NULL,
    "attribute_id" BIGINT NOT NULL,
    FOREIGN KEY ("product_id") REFERENCES "products"("id"),
    FOREIGN KEY ("attribute_id") REFERENCES "attributes"("id")
);

-- +goose Down
DROP TABLE "product_attributes";