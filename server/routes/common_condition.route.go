package routes

import (
	"github.com/germatpreis/eldie/server/controllers"
	"github.com/gin-gonic/gin"
)

type CommonConditionRoutes struct {
	controller controllers.CommonConditionController
}

func NewCommonConditionRoute(controller controllers.CommonConditionController) CommonConditionRoutes {
	return CommonConditionRoutes{controller}
}

func (c *CommonConditionRoutes) Router(rg *gin.RouterGroup) {
	router := rg.Group("conditions")
	router.GET("/", c.controller.ListConditions)
	router.GET("/:conditionId/culprits", c.controller.ListCulpritsForConditions)
	router.GET("/:conditionId/symptoms", c.controller.ListSymptomsForConditions)
}
