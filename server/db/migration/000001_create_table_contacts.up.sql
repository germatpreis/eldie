create table if not exists contacts(
    contact_id uuid not null default gen_random_uuid(),
    first_name varchar not null,
    last_name varchar not null,
    phone_number varchar not null,
    street varchar not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null,
    constraint "contacts_pkey" primary key("contact_id")
);

create unique index if not exists "contacts_phone_number_key" on "contacts"("phone_number");