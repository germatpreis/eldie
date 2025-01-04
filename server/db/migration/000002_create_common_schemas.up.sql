create table if not exists common_conditions
(
    id                       bigserial    not null primary key,
    code                     char(6)      not null,
    name                     varchar(128) not null,
    description              text         not null,
    duration_healing_phase   int          not null,
    duration_challenge_phase int          not null,
    days_between_phases      int          not null
);

create unique index if not exists "common_conditions_name_uidx" on common_conditions (name);

create table if not exists common_symptoms
(
    id          bigserial    not null primary key,
    name        varchar(128) not null,
    code        char(6)      not null,
    description text         not null
);

create unique index if not exists common_symptoms_name_uidx on common_symptoms (name);
create unique index if not exists common_symptoms_code_uidx on common_symptoms (code);

create table if not exists common_conditions_symptoms
(
    condition_id bigserial not null,
    symptom_id   bigserial not null,
    constraint pk_common_conditions_symptoms primary key (condition_id, symptom_id),
    constraint fk_common_conditions foreign key (condition_id) references common_conditions (id),
    constraint fk_common_symptoms foreign key (symptom_id) references common_symptoms (id)
);

create table if not exists common_foods
(
    id    bigserial    not null primary key,
    name  varchar(255) not null,
    food_group varchar(255),
    ph_value float
);

create unique index if not exists common_foods_name_uidx on common_foods (name);

create table if not exists common_conditions_culprits
(
    condition_id bigserial not null,
    food_id bigserial not null,
    reasoning text,
    constraint pk_common_conditions_culprits primary key (condition_id, food_id),
    constraint fk_common_conditions_culprits_conditions foreign key (condition_id) references common_conditions (id),
    constraint fk_common_conditions_culprits_foods foreign key (food_id) references common_foods (id)
);
