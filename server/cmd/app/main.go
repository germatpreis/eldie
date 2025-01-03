package main

import (
	"context"
	dbCon "github.com/Geoff89/sqlccrud/db/sqlc"
	"github.com/germatpreis/eldie/server/controllers"
	"github.com/germatpreis/eldie/server/routes"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.ContactController
	ContactRouters    routes.ContactRoutes
)

func main() {
}
