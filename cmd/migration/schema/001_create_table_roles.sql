-- +goose Up
CREATE TABLE "roles" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "description" TEXT
);

CREATE UNIQUE INDEX unique_role_name ON "roles" ("name");

INSERT INTO "roles" ("name") VALUES 
('ADMIN'),
('CUSTOMER'),
('COURIER');

-- +goose Down
DROP TABLE "roles";