-- name: get-ticket-types
SELECT * FROM register.ticket_types;

-- name: get-articles
SELECT * FROM register.articles;

-- name: get-article
SELECT * FROM register.articles WHERE id = $1::uuid OR name = $1::text;

-- name: insert-article
INSERT INTO register.articles(name, price_members, price_guests, color) VALUES
($1, $2, $3, $4)
RETURNING id, name, enabled, price_members, price_guests, color;