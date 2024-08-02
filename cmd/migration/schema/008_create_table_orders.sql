-- +goose Up
CREATE TYPE order_status AS ENUM ('picked_up', 'on_delivery', 'delivered', 'payment_collected');

CREATE TABLE "orders" (
    "id" BIGSERIAL PRIMARY KEY,
    "courier_id" BIGINT,
    "reciever_id" BIGINT NOT NULL,
    "location_x" DECIMAL(9, 6),  
    "location_y" DECIMAL(9, 6),  
    "address" TEXT,
    "status" order_status NOT NULL,
    "ordered_at" TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY ("courier_id") REFERENCES "users"("id") ON DELETE CASCADE,
    FOREIGN KEY ("reciever_id") REFERENCES "users"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "orders";

DROP TYPE "order_status";