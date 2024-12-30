-- +goose Up
-- +goose StatementBegin
ALTER TABLE register.articles ALTER price_members TYPE numeric(12,2);
ALTER TABLE register.articles ALTER price_guests TYPE numeric(12,2);
ALTER TABLE register.menus ALTER price_members TYPE numeric(12,2);
ALTER TABLE register.menus ALTER price_guests TYPE numeric(12,2);
ALTER TABLE register.ticket_types ALTER price TYPE numeric(12,2);
-- +goose StatementEnd
