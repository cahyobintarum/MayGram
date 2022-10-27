package v1

import (
	engine "github.com/cahyobintarum/MayGram/config/gin"
	"github.com/cahyobintarum/MayGram/pkg/domain/comment"
	"github.com/cahyobintarum/MayGram/pkg/server/http/middleware"
	"github.com/cahyobintarum/MayGram/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type CommentRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	commentHandler comment.CommentHandler
	authMiddleware middleware.AuthMiddleware
}

func (p *CommentRouterImpl) get() {
	p.routerGroup.GET("", p.authMiddleware.CheckJWTAuth, p.commentHandler.GetCommentsHdl)
}

func (p *CommentRouterImpl) post() {
	p.routerGroup.POST("", p.authMiddleware.CheckJWTAuth, p.commentHandler.CreateCommentHdl)
}

func (p *CommentRouterImpl) put() {
	p.routerGroup.PUT("/:commentId", p.authMiddleware.CheckJWTAuth, p.commentHandler.UpdateCommentHdl)
}

func (p *CommentRouterImpl) delete() {
	p.routerGroup.DELETE("/:commentId", p.authMiddleware.CheckJWTAuth, p.commentHandler.DeleteCommentHdl)
}

func (p *CommentRouterImpl) Routers() {
	p.get()
	p.post()
	p.put()
	p.delete()
}

func NewCommentRouter(ginEngine engine.HttpServer, commentHandler comment.CommentHandler, authMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/comments")
	return &CommentRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, commentHandler: commentHandler, authMiddleware: authMiddleware}
}
