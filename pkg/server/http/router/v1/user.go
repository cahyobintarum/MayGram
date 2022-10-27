package v1

import (
	engine "github.com/cahyobintarum/MayGram/config/gin"
	"github.com/cahyobintarum/MayGram/pkg/domain/user"
	"github.com/cahyobintarum/MayGram/pkg/server/http/middleware"
	"github.com/cahyobintarum/MayGram/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	userHandler    user.UserHandler
	authMiddleware middleware.AuthMiddleware
}

func (u *UserRouterImpl) post() {
	u.routerGroup.POST("/register", u.userHandler.RegisterUserHdl)
}

func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/:user_id", u.userHandler.GetUserByIdHdl)
}

func (u *UserRouterImpl) put() {
	u.routerGroup.PUT("", u.authMiddleware.CheckJWTAuth, u.userHandler.UpdateUserHdl)
}

func (u *UserRouterImpl) delete() {
	u.routerGroup.DELETE("", u.authMiddleware.CheckJWTAuth, u.userHandler.DeleteUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.get()
	u.post()
	u.put()
	u.delete()
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler, authMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/users")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler, authMiddleware: authMiddleware}
}
