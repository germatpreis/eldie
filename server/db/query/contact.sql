-- name: CreateContact :one
insert into contacts(first_name,
                     last_name,
                     phone_number,
                     street,
                     created_at,
                     updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetContactById :one
select * from contacts
where contact_id = $1 limit 1;

-- name: ListContacts :many
select * from contacts
order by contact_id
limit $1
offset $2;

-- name: UpdateContact :one
update contacts
set
    first_name = coalesce(sqlc.narg('first_name'), first_name),
    last_name = coalesce(sqlc.narg('last_name'), last_name),
    phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
    street = coalesce(sqlc.narg('street'), street),
    last_name = coalesce(sqlc.narg('last_name'), last_name),
    updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
where contact_id = sqlc.arg('contact_id')
returning *;

-- name: DeleteContact :exec
delete from contacts
where contact_id = $1;
