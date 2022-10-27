package main

import (
	"net/http"
	"os"
	"time"

	engine "github.com/cahyobintarum/MayGram/config/gin"
	"github.com/cahyobintarum/MayGram/config/postgres"
	authrepo "github.com/cahyobintarum/MayGram/pkg/repository/auth"
	commentrepo "github.com/cahyobintarum/MayGram/pkg/repository/comment"
	photorepo "github.com/cahyobintarum/MayGram/pkg/repository/photo"
	socialmediarepo "github.com/cahyobintarum/MayGram/pkg/repository/socialmedia"
	userrepo "github.com/cahyobintarum/MayGram/pkg/repository/user"
	authhandler "github.com/cahyobintarum/MayGram/pkg/server/http/handler/auth"
	commenthandler "github.com/cahyobintarum/MayGram/pkg/server/http/handler/comment"
	photohandler "github.com/cahyobintarum/MayGram/pkg/server/http/handler/photo"
	socialmediahandler "github.com/cahyobintarum/MayGram/pkg/server/http/handler/socialmedia"
	userhandler "github.com/cahyobintarum/MayGram/pkg/server/http/handler/user"
	"github.com/cahyobintarum/MayGram/pkg/server/http/middleware"
	router "github.com/cahyobintarum/MayGram/pkg/server/http/router/v1"
	authusecase "github.com/cahyobintarum/MayGram/pkg/usecase/auth"
	commentusecase "github.com/cahyobintarum/MayGram/pkg/usecase/comment"
	photousecase "github.com/cahyobintarum/MayGram/pkg/usecase/photo"
	socialmediausecase "github.com/cahyobintarum/MayGram/pkg/usecase/socialmedia"
	userusecase "github.com/cahyobintarum/MayGram/pkg/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	postgresHost := os.Getenv("MY_GRAM_POSTGRES_HOST")
	postgresPort := os.Getenv("MY_GRAM_POSTGRES_PORT")
	postgresDatabase := os.Getenv("MY_GRAM_POSTGRES_DATABASE")
	postgresUsername := os.Getenv("MY_GRAM_POSTGRES_USERNAME")
	postgresPassword := os.Getenv("MY_GRAM_POSTGRES_PASSWORD")

	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         postgresHost,
		Port:         postgresPort,
		User:         postgresUsername,
		Password:     postgresPassword,
		DatabaseName: postgresDatabase,
	})

	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "server up and running",
			"start_time": startTime,
		})
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHandler(userUsecase)

	authRepo := authrepo.NewAuthRepo(postgresCln)
	authUsecase := authusecase.NewAuthUsecase(authRepo, userUsecase)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo, userUsecase)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)

	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo, photoUsecase)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase)

	socialMediaRepo := socialmediarepo.NewSocialMediaRepo(postgresCln)
	socialMediaUsecase := socialmediausecase.NewSocialMediaUsecase(socialMediaRepo)
	socialMediaHandler := socialmediahandler.NewSocialMediaHandler(socialMediaUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	router.NewUserRouter(ginEngine, userHandler, authMiddleware).Routers()
	router.NewAuthRouter(ginEngine, authHandler, authMiddleware).Routers()
	router.NewPhotoRouter(ginEngine, photoHandler, authMiddleware).Routers()
	router.NewCommentRouter(ginEngine, commentHandler, authMiddleware).Routers()
	router.NewSocialMediaRouter(ginEngine, socialMediaHandler, authMiddleware).Routers()

	ginEngine.Serve()
}
