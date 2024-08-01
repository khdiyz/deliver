-- +goose Up

CREATE TABLE "cart" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT UNIQUE NOT NULL REFERENCES "users"("id") ON DELETE CASCADE
);

CREATE TABLE "cart_products" (
    "id" BIGSERIAL PRIMARY KEY,
    "cart_id" BIGINT NOT NULL REFERENCES "cart"("id") ON DELETE CASCADE,
    "product_id" BIGINT NOT NULL REFERENCES "products"("id"),
    "quantity" BIGINT NOT NULL
);

CREATE TABLE "cart_product_attributes" (
    "id" BIGSERIAL PRIMARY KEY,
    "cart_product_id" BIGINT NOT NULL REFERENCES "cart_products"("id") ON DELETE CASCADE,
    "attribute_id" BIGINT NOT NULL REFERENCES "attributes"("id"),
    "option_id" BIGINT NOT NULL REFERENCES "options"("id")
);

-- +goose Down
DROP TABLE "cart_product_attributes";

DROP TABLE "cart_products";

DROP TABLE "cart";