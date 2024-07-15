package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"src/config"
	"src/internal/controller/http"
	mypostgres "src/internal/repository/postgres"
	"src/internal/service"
	"src/pkg/logging"
	httpserver "src/pkg/server/http"
	"src/pkg/storage/postgres"
)

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	handler := gin.New()
// 	// controller := http.NewRouter(handler)

// }

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		cfg.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	l := logging.New(cfg.Logger.Level, loggerFile)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.Database,
	))

	userRepo := mypostgres.NewUserRepository(db)
	supplierRepo := mypostgres.NewSupplierRepository(db)
	racketRepo := mypostgres.NewRacketRepository(db)

	userService := service.NewUserService(l, userRepo)
	// supplierService := service.NewSupplierService(l, supplierRepo)
	racketService := service.NewRacketService(l, racketRepo, supplierRepo)

	handler := gin.New()
	controller := http.NewRouter(handler)

	controller.SetUserRoute(l, userService)
	controller.SetProductRoute(l, racketService)

	// Create router
	router := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Starting server
	err = router.Start()
	if err != nil {
		log.Fatal(err)
	}

}
