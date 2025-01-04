-- name: ListCommonConditions :many
select *
from common_conditions;

-- name: GetCommonConditionByCode :one
select *
from common_conditions
where code = $1
limit 1;

-- name: GetCommonConditionById :one
select *
from common_conditions
where id = $1
limit 1;

-- name: ListCommonCulpritsForCondition :many
select cf.id as food_id, cf.ph_value, ccc.reasoning
from common_conditions cc
         join common_conditions_culprits ccc on cc.id = ccc.condition_id
         join common_foods cf on ccc.food_id = cf.id
where cc.id = $1;

-- name: ListCommonSymptomsForCondition :many
select cs.id as symptom_id, cs.name, cs.code, cs.description
from common_conditions cc
         join public.common_conditions_symptoms ccs on cc.id = ccs.condition_id
         join public.common_symptoms cs on cs.id = ccs.symptom_id
where cc.id = $1;