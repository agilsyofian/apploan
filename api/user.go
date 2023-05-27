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

type Profile struct {
	Detail ProfileDetail     `json:"detail"`
	Limit  []utilitize.Limit `json:"limit"`
}

type ProfileDetail struct {
	NIK         int64     `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	TempatLahir string    `json:"tempat_lahir"`
	TglLahir    string    `json:"tgl_lahir"`
	Gaji        float64   `json:"gaji"`
	FotoKTP     string    `json:"foto_ktp"`
	FotoSelfie  string    `json:"foto_selfie"`
	CreatedAt   time.Time `json:"created_at"`
}

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

	insertKonsumen, err := server.store.CreateKonsumen(konsumen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := server.store.BuildProfile(*insertKonsumen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) profile(ctx *gin.Context) {
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

	response, err := server.store.BuildProfile(*konsumen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)

	ctx.JSON(http.StatusOK, response)
}
