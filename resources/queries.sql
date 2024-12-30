-- name: get-ticket-types
SELECT * FROM register.ticket_types;

-- name: get-articles
SELECT * FROM register.articles;

-- name: get-article
SELECT * FROM register.articles WHERE id = $1::uuid OR name = $1::text;