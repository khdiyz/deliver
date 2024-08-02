-- +goose Up
CREATE TABLE "product_attributes" (
    "id" BIGSERIAL PRIMARY KEY,
    "product_id" BIGINT NOT NULL,
    "attribute_id" BIGINT NOT NULL,
    FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE,
    FOREIGN KEY ("attribute_id") REFERENCES "attributes"("id") ON DELETE CASCADE,
    CONSTRAINT unique_product_attribute UNIQUE ("product_id", "attribute_id")
);

-- +goose Down
DROP TABLE "product_attributes";