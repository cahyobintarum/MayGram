package v1

import (
	engine "github.com/cahyobintarum/MayGram/config/gin"
	"github.com/cahyobintarum/MayGram/pkg/domain/auth"
	"github.com/cahyobintarum/MayGram/pkg/server/http/middleware"
	"github.com/cahyobintarum/MayGram/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type AuthRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	authHandler    auth.AuthHandler
	authMiddleware middleware.AuthMiddleware
}

func (a *AuthRouterImpl) post() {
	a.routerGroup.POST("/login", a.authHandler.LoginUserHdl)
	a.routerGroup.POST("/refresh", a.authMiddleware.CheckJWTAuth, a.authHandler.GetRefreshTokenHdl)
}

func (a *AuthRouterImpl) Routers() {
	a.post()
}

func NewAuthRouter(ginEngine engine.HttpServer, authHandler auth.AuthHandler, auhtMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/auth")
	return &AuthRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, authHandler: authHandler, authMiddleware: auhtMiddleware}
}
