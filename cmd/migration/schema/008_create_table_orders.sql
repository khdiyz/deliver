-- +goose Up
CREATE TYPE order_status AS ENUM ('picked_up', 'on_delivery', 'delivered', 'payment_collected');

CREATE TABLE "orders" (
    "id" BIGSERIAL PRIMARY KEY,
    "courier_id" BIGINT NOT NULL,
    "reciever_id" BIGINT NOT NULL,
    "status" order_status NOT NULL,
    "ordered_at" TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY ("courier_id") REFERENCES "users"("id"),
    FOREIGN KEY ("reciever_id") REFERENCES "users"("id")
);

-- +goose Down
DROP TABLE "orders";