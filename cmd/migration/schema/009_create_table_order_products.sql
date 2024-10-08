-- +goose Up

CREATE TABLE "order_products" (
    "id" BIGSERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "product_id" BIGINT NOT NULL,
    "product_attributes" JSON NOT NULL,
    "quantity" INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY ("order_id") REFERENCES "orders"("id") ON DELETE CASCADE,
    FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "order_products";