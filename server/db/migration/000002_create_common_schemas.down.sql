drop table if exists common_conditions_symptoms;
drop table if exists common_conditions;
drop table if exists common_symptoms;



create unique index if not exists "common_symptoms_name_uidx" on common_symptoms (name);
create unique index if not exists "common_symptoms_code_uidx" on common_symptoms (code);

create table if not exists common_conditions_symptoms
(
    condition_id bigserial not null,
    symptom_id   bigserial not null,
    constraint pk_common_conditions_symptoms primary key (condition_id, symptom_id),
    constraint fk_common_conditions foreign key(condition_id) references common_conditions(id),
    constraint fk_common_symptoms foreign key(symptom_id) references common_symptoms(id)
)
