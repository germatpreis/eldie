package controllers

import (
	"context"
	db "github.com/germatpreis/eldie/server/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommonConditionController struct {
	db  *db.Queries
	ctx context.Context
}

func (c *CommonConditionController) ListConditions(ctx *gin.Context) {
	result, err := c.db.ListCommonConditions(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve conditions", "error": err.Error()})
		return
	}

	if result == nil {
		result = []db.CommonCondition{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved conditions", "result": result})
}

func (c *CommonConditionController) ListCulpritsForConditions(ctx *gin.Context) {
	conditionId, err := c.ensureConditionExists(ctx)
	if err == nil {
		return
	}

	culprits, err := c.db.ListCommonCulpritsForCondition(ctx, conditionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve culprits for the condition with this ID"})
		return
	}

	if culprits == nil {
		culprits = []db.ListCommonCulpritsForConditionRow{}
	}

	ctx.JSON(http.StatusOK, culprits)
}

func (c *CommonConditionController) ListSymptomsForConditions(ctx *gin.Context) {
	conditionId, err := c.ensureConditionExists(ctx)
	if err != nil {
		return
	}

	symptoms, err := c.db.ListCommonSymptomsForCondition(ctx, conditionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve symptoms for the condition with this ID"})
		return
	}

	if symptoms == nil {
		symptoms = []db.ListCommonSymptomsForConditionRow{}
	}

	ctx.JSON(http.StatusOK, symptoms)
}

func (c *CommonConditionController) ensureConditionExists(ctx *gin.Context) (conditionId int64, err error) {
	conditionId, err = StringToInt64(ctx.Param("conditionId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed to parse conditionId", "error": err.Error()})
		return conditionId, err
	}

	_, err = c.db.GetCommonConditionById(ctx, conditionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve condition with this ID"})
		return conditionId, err
	}
	return conditionId, nil
}

func NewCommonConditionController(db *db.Queries, ctx context.Context) *CommonConditionController {
	return &CommonConditionController{
		db,
		ctx,
	}
}

func StringToInt32(value string) (parsedValue int32, err error) {
	parsedValue64, err := StringToInt64(value)
	if err != nil {
		return 0, err
	}
	return int32(parsedValue64), nil
}

func StringToInt64(value string) (parsedValue int64, err error) {
	parsedValue, err = strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return parsedValue, nil
}
