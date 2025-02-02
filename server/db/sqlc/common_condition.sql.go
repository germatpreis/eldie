// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: common_condition.sql

package db

import (
	"context"
	"database/sql"
)

const getCommonConditionByCode = `-- name: GetCommonConditionByCode :one
select id, code, name, description, duration_healing_phase, duration_challenge_phase, days_between_phases
from common_conditions
where code = $1
limit 1
`

func (q *Queries) GetCommonConditionByCode(ctx context.Context, code string) (CommonCondition, error) {
	row := q.queryRow(ctx, q.getCommonConditionByCodeStmt, getCommonConditionByCode, code)
	var i CommonCondition
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.DurationHealingPhase,
		&i.DurationChallengePhase,
		&i.DaysBetweenPhases,
	)
	return i, err
}

const getCommonConditionById = `-- name: GetCommonConditionById :one
select id, code, name, description, duration_healing_phase, duration_challenge_phase, days_between_phases
from common_conditions
where id = $1
limit 1
`

func (q *Queries) GetCommonConditionById(ctx context.Context, id int64) (CommonCondition, error) {
	row := q.queryRow(ctx, q.getCommonConditionByIdStmt, getCommonConditionById, id)
	var i CommonCondition
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Description,
		&i.DurationHealingPhase,
		&i.DurationChallengePhase,
		&i.DaysBetweenPhases,
	)
	return i, err
}

const listCommonConditions = `-- name: ListCommonConditions :many
select id, code, name, description, duration_healing_phase, duration_challenge_phase, days_between_phases
from common_conditions
`

func (q *Queries) ListCommonConditions(ctx context.Context) ([]CommonCondition, error) {
	rows, err := q.query(ctx, q.listCommonConditionsStmt, listCommonConditions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CommonCondition{}
	for rows.Next() {
		var i CommonCondition
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.Description,
			&i.DurationHealingPhase,
			&i.DurationChallengePhase,
			&i.DaysBetweenPhases,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommonCulpritsForCondition = `-- name: ListCommonCulpritsForCondition :many
select cf.id as food_id, cf.ph_value, ccc.reasoning
from common_conditions cc
         join common_conditions_culprits ccc on cc.id = ccc.condition_id
         join common_foods cf on ccc.food_id = cf.id
where cc.id = $1
`

type ListCommonCulpritsForConditionRow struct {
	FoodID    int64           `json:"food_id"`
	PhValue   sql.NullFloat64 `json:"ph_value"`
	Reasoning sql.NullString  `json:"reasoning"`
}

func (q *Queries) ListCommonCulpritsForCondition(ctx context.Context, id int64) ([]ListCommonCulpritsForConditionRow, error) {
	rows, err := q.query(ctx, q.listCommonCulpritsForConditionStmt, listCommonCulpritsForCondition, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCommonCulpritsForConditionRow{}
	for rows.Next() {
		var i ListCommonCulpritsForConditionRow
		if err := rows.Scan(&i.FoodID, &i.PhValue, &i.Reasoning); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommonSymptomsForCondition = `-- name: ListCommonSymptomsForCondition :many
select cs.id as symptom_id, cs.name, cs.code, cs.description
from common_conditions cc
         join public.common_conditions_symptoms ccs on cc.id = ccs.condition_id
         join public.common_symptoms cs on cs.id = ccs.symptom_id
where cc.id = $1
`

type ListCommonSymptomsForConditionRow struct {
	SymptomID   int64  `json:"symptom_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (q *Queries) ListCommonSymptomsForCondition(ctx context.Context, id int64) ([]ListCommonSymptomsForConditionRow, error) {
	rows, err := q.query(ctx, q.listCommonSymptomsForConditionStmt, listCommonSymptomsForCondition, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCommonSymptomsForConditionRow{}
	for rows.Next() {
		var i ListCommonSymptomsForConditionRow
		if err := rows.Scan(
			&i.SymptomID,
			&i.Name,
			&i.Code,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
