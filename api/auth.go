package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/agilsyofian/kreditplus/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) auth(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.AuthKonsumen(req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid username")))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid password")))
		return
	}

	tokenPayload := ClientPayload{
		TypeToken: "token",
		UserID:    user.ID,
	}

	duration := server.config.AccessTokenDuration
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(tokenPayload, duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshTokenPayload := ClientPayload{
		TypeToken: "refresh token",
		UserID:    user.ID,
	}

	refreshDuration := server.config.AccessTokenRefreshDuration
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(refreshTokenPayload, refreshDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var sesData models.Session = models.Session{
		ID:           refreshPayload.ID,
		KonsumenID:   user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		ExpiredAt:    refreshPayload.ExpiredAt,
		IsBlocked:    false,
	}

	_, err = server.store.SessionCreate(sesData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := struct {
		AccessToken           string    `json:"access_token"`
		RefreshToken          string    `json:"refresh_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	}{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.SessionGet(refreshPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("session not found")))
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	clientPayload := convertClientPayload(refreshPayload.Payload)

	if session.KonsumenID != clientPayload.UserID {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	tokenPayload := ClientPayload{
		TypeToken: "token",
		UserID:    session.KonsumenID,
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		tokenPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
