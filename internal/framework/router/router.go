package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	PrefixPath() string
	RegisterAPI(engine *gin.Engine)
}
