package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// GetXid generates a new XID.
func GetXid(c *gin.Context) {
	id := xid.New()
	c.JSON(http.StatusOK, gin.H{
		"id": id.String(),
	})
}
