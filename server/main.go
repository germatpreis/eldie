package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/germatpreis/eldie/server/controllers"
	dbCon "github.com/germatpreis/eldie/server/db/sqlc"
	"github.com/germatpreis/eldie/server/routes"
	"github.com/germatpreis/eldie/server/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.ContactController
	ContactRoutes     routes.ContactRoutes

	CommonConditionController controllers.CommonConditionController
	CommonConditionRoutes     routes.CommonConditionRoutes
)

func init() {
	ctx = context.TODO()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	db = dbCon.New(conn)

	fmt.Println("connected to postgres")

	ContactController = *controllers.NewContactController(db, ctx)
	ContactRoutes = routes.NewRouteContact(ContactController)

	CommonConditionController := *controllers.NewCommonConditionController(db, ctx)
	CommonConditionRoutes = routes.NewCommonConditionRoute(CommonConditionController)

	server = gin.Default()
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	router := server.Group("/api")
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	ContactRoutes.Router(router)
	CommonConditionRoutes.Router(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "route does not exist"})
	})

	err = server.Run(":" + config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
