-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA register;

CREATE TABLE register.articles(
    id uuid default gen_random_uuid() primary key,
    name text not null unique,
    enabled boolean not null default true,
    price_members money,
    price_guests money not null,
    color text not null default '#00a2ff'
);

CREATE TABLE register.menus(
    id uuid default gen_random_uuid() primary key,
    name text not null unique,
    enabled boolean not null default true,
    price_members money,
    price_guests money not null,
    color text not null default '#d90080',
    tickets uuid[],
    items uuid[]
);

CREATE TABLE register.ticket_types(
    id uuid default gen_random_uuid() primary key,
    name text not null unique,
    price money not null,
    color text not null default '#b1d300'
);

CREATE TABLE register.transactions(
    id uuid default gen_random_uuid() primary key,
    "timestamp" timestamptz not null default now(),
    amount money not null,
    tickets uuid[],
    menus uuid[],
    articles uuid[]
);
-- +goose StatementEnd