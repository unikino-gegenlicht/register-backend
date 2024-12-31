-- name: get-ticket-types
SELECT * FROM register.ticket_types;

-- name: get-articles
SELECT * FROM register.articles;

-- name: get-article
SELECT * FROM register.articles WHERE id = $1::uuid;

-- name: insert-article
INSERT INTO register.articles(name, price_members, price_guests, color) VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: update-article
UPDATE register.articles
SET name = $2, enabled = $3, price_members = $4, price_guests =$5, color = $6
WHERE id = $1
RETURNING *;

-- name: delete-article
DELETE FROM register.articles
WHERE id = $1::uuid;