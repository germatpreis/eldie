package routes

import (
	"github.com/germatpreis/eldie/server/controllers"
	"github.com/gin-gonic/gin"
)

type ContactRoutes struct {
	contactController controllers.ContactController
}

func NewRouteContact(contactController controllers.ContactController) ContactRoutes {
	return ContactRoutes{contactController}
}

func (cr *ContactRoutes) Router(rg *gin.RouterGroup) {
	router := rg.Group("contacts")
	router.POST("/", cr.contactController.CreateContact)
	router.GET("/", cr.contactController.GetAllContacts)
	router.PATCH("/:contactId", cr.contactController.UpdateContact)
	router.GET("/:contactId", cr.contactController.GetContactById)
	router.DELETE("/:contactId", cr.contactController.DeleteContact)
}
