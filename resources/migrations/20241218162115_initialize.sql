-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA register;

CREATE TABLE register.articles(
    id uuid default gen_random_uuid() primary key,
    name text not null unique,
    enabled boolean not null default true,
    price_members money,
    price_guests money not null,
    color text not null default '#00a2ff',
);

CREATE TABLE register.menus(
    id uuid default gen_random_uuid() primary key,
    name text not null unique,
    enabled boolean not null default true,
    price_members money,
    price_guests money not null,
    color text not null default '#dc00ff',
    icon text not null default 'star'
);

CREATE TABLE register.transactions(
    id uuid default gen_random_uuid() primary key,
    "timestamp" timestamptz not null default now(),
    amount money not null,
    standard_tickets int not null default 0,
    combined_tickets int not null default 0,
    free_tickets int not null default 0,
    menus uuid[],
    articles uuid[]
)
-- +goose StatementEnd