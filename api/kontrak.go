package api

import (
	"errors"
	"net/http"

	"github.com/agilsyofian/kreditplus/models"
	"github.com/agilsyofian/kreditplus/utilitize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) kontrakCreate(ctx *gin.Context) {
	var req models.Kontrak
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(ClientData)

	kontrakFact := utilitize.NewKontrak(req.Otr, 0.06, 0.05, req.Tenor)
	kontrakGen := kontrakFact.BuildKontrak()

	noKontrak, _ := uuid.NewRandom()
	payload := models.Kontrak{
		No:         noKontrak,
		KonsumenID: authPayload.ClientPayload.UserID,
		Otr:        req.Otr,
		AdminFee:   kontrakGen.AdminFee,
		JmlCicilan: kontrakGen.JmlCicilan,
		JmlBunga:   kontrakGen.JmlBunga,
		NamaAsset:  req.NamaAsset,
		Tenor:      req.Tenor,
		Status:     "inpg",
	}

	konsumen, err := server.store.GetKonsumen(authPayload.ClientPayload.UserID)
	if err != nil {
		err := errors.New("invalid credential")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	profileKonsumen, err := server.store.BuildProfile(*konsumen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	limitTenor := float64(0)
	for _, limit := range profileKonsumen.Limit {
		if limit.Tenor == req.Tenor {
			limitTenor = limit.Limit
		}
	}

	if req.Otr > limitTenor {
		err := errors.New("user limit reached")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	kontrak, err := server.store.KontrakCreate(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, kontrak)
}
