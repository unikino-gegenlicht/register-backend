-- +goose Up
-- +goose StatementBegin

BEGIN;
CREATE TABLE IF NOT EXISTS register.screenings (
    id bigserial primary key ,
    title text not null,
    start timestamptz not null,
    allow_reservations boolean not null default true,
    max_occupancy int not null default 80,
    wpid int not null
);

CREATE TABLE IF NOT EXISTS register.tickets (
    id uuid primary key default gen_random_uuid (),
    issued_at timestamptz default now(),
    type uuid not null references register.ticket_types (id) MATCH FULL ON UPDATE CASCADE ON DELETE RESTRICT,
    screening int not null references register.screenings(id) MATCH FULL ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS register.reservations (
    id uuid primary key default gen_random_uuid(),
    first_name text not null,
    last_name text not null,
    email text not null,
    seats int not null,
    tickets_issued boolean not null default false,
    screening int not null REFERENCES register.screenings(id) MATCH FULL ON UPDATE CASCADE ON DELETE RESTRICT
);

ALTER TABLE register.ticket_types
ADD COLUMN enabled boolean not null default true;
COMMIT;
-- +goose StatementEnd