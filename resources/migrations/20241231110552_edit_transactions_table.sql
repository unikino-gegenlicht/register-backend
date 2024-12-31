-- +goose Up
-- +goose StatementBegin
ALTER TABLE register.transactions DROP COLUMN amount;
ALTER TABLE register.transactions DROP COLUMN menus;
ALTER TABLE register.transactions DROP COLUMN tickets;
ALTER TABLE register.transactions DROP COLUMN articles;
ALTER TABLE register.transactions ADD COLUMN data text not null;
-- +goose StatementEnd