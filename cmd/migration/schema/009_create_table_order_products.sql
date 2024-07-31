-- +goose Up

CREATE TABLE "order_products" (
    "id" BIGSERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "product_id" BIGINT NOT NULL,
    "product_attributes" JSON NOT NULL,
    "count" INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY ("order_id") REFERENCES "orders"("id"),
    FOREIGN KEY ("product_id") REFERENCES "products"("id")
);

-- +goose Down
DROP TABLE "order_products";