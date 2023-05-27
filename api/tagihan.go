package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/agilsyofian/kreditplus/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) tagihan(ctx *gin.Context) {
	var req struct {
		ID string `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	kontrakNo, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	kontrakDetail, err := server.store.KontrakGetByID(kontrakNo)
	if err != nil {
		err := errors.New("invalid kontrak no")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagihan, err := server.store.TagihanList(kontrakNo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if len(tagihan) == 0 {
		err := errors.New("invalid kontrak no")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(ClientData)
	userID := authPayload.ClientPayload.UserID
	if userID != kontrakDetail.KonsumenID {
		err := errors.New("tagihan doesn't belong to the authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tagihan)
}

func (server *Server) tagihanPay(ctx *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if len(req.IDs) == 0 {
		err := errors.New("invalid ids")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagihanID := req.IDs[0]

	kontrak, err := server.store.KontrakGetByTagihan(tagihanID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	authorizationPayload := ctx.MustGet(authorizationPayloadKey).(ClientData)
	if kontrak.KonsumenID != authorizationPayload.ClientPayload.UserID {
		err := errors.New("tagihan doesn't belong to the authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	tgl_paid := time.Now().Format("2006-01-02")
	payload := models.Tagihan{
		KontrakNo: kontrak.No,
		Status:    "paid",
		TglPaid:   &tgl_paid,
	}
	err = server.store.TagihanUpdateBatch(req.IDs, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
