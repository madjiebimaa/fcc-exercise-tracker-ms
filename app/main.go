package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/app/config"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/app/mongo"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/handlers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/middlewares"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/repositories"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/services"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
)

func main() {
	config.GetEnv()

	ctx := context.Background()
	mongoConfig := mongo.NewConfigDB(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"))
	cl := mongoConfig.Init(ctx)
	defer func() {

		if err := cl.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cl.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	db := cl.Database(os.Getenv("DB_NAME"))
	userColl := db.Collection(os.Getenv("USER_COLLECTION"))

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userRepo := repositories.NewMongoUserRepository(userColl)
	userService := services.NewUserService(userRepo, timeoutContext)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.New()
	r.Use(middlewares.CORS())

	r.POST("/api/users", userHandler.Register)

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal("can't connect to server")
	}
}
