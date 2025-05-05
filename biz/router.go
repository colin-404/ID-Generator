package biz

import (
	v1 "github.com/colin-404/ID-Generator/biz/handler/v1"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/xid", v1.GetXid)
}
