package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type ClientPayload struct {
	TypeToken string    `json:"type_token"`
	UserID    uuid.UUID `json:"user_id"`
}

type ClientData struct {
	UserAgent     string `json:"user_agent"`
	ClientIp      string `json:"client_ip"`
	ClientPayload ClientPayload
}

func authMiddleware(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var client = ClientData{
			UserAgent: ctx.Request.UserAgent(),
			ClientIp:  ctx.ClientIP(),
		}

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := server.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		clientPayload := convertClientPayload(payload.Payload)
		if clientPayload.TypeToken != "token" {
			err := errors.New("invalid authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		client.ClientPayload.UserID = clientPayload.UserID
		ctx.Set(authorizationPayloadKey, client)
		ctx.Next()
	}
}

func convertClientPayload(any interface{}) ClientPayload {
	data := any.(map[string]interface{})

	return ClientPayload{
		UserID:    uuid.MustParse(data["user_id"].(string)),
		TypeToken: data["type_token"].(string),
	}
}
