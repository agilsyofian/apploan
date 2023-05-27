package api

import (
	"net/http"

	"github.com/agilsyofian/kreditplus/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) kontrakCreate(ctx *gin.Context) {
	var req models.Kontrak
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// newUUID, _ := uuid.NewRandom()
	// password, _ := util.HashPassword(req.Password)
	// var konsumen models.Konsumen = req
	// konsumen.ID = newUUID

	ctx.JSON(http.StatusOK, gin.H{
		"test": "not implemented",
	})
}
