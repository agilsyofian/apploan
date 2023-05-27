package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/agilsyofian/kreditplus/models"
	"github.com/agilsyofian/kreditplus/utilitize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) register(ctx *gin.Context) {
	var req models.Konsumen
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	newUUID, _ := uuid.NewRandom()
	password, _ := util.HashPassword(req.Password)
	var konsumen models.Konsumen = req
	konsumen.ID = newUUID
	konsumen.Password = password

	base64 := utilitize.NewBase64(konsumen.FotoKTP)
	valid, mime := base64.CheckMimeType()
	if !valid {
		err := fmt.Errorf("invalid format")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pathKtp := server.config.Assets + "/ktp/" + time.Now().Format("2006-01-02")
	fotoKTP, err := base64.StoreBase64ToImage(mime, pathKtp, newUUID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pathSelfie := server.config.Assets + "/selfie/" + time.Now().Format("2006-01-02")
	fotoSelfie, err := base64.StoreBase64ToImage(mime, pathSelfie, newUUID.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	konsumen.FotoKTP = fotoKTP
	konsumen.FotoSelfie = fotoSelfie

	buildLimit := utilitize.NewFactoryLimit(konsumen)
	limit := buildLimit.BuildLimit()

	response, err := server.store.Register(konsumen, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) profile(ctx *gin.Context) {

	var response models.Register
	authPayload := ctx.MustGet(authorizationPayloadKey).(ClientData)
	id, err := uuid.Parse(authPayload.ClientPayload.UserID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	konsumen, err := server.store.GetKonsumen(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response.Konsumen = *konsumen

	limit, err := server.store.LimitGetByKonsumen(konsumen.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response.Limit = limit

	ctx.JSON(http.StatusOK, response)
}
