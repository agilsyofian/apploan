package api

import (
	"github.com/agilsyofian/golang/pasetomaker"
	"github.com/agilsyofian/kreditplus/config"
	"github.com/agilsyofian/kreditplus/models"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     config.Config
	store      *models.Database
	tokenMaker pasetomaker.Maker
	router     *gin.Engine
}

func NewServer(config config.Config, store *models.Database, tokenMaker pasetomaker.Maker) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	if server.config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.Recovery())

	router.POST("/register", server.register)

	router.POST("/auth", server.auth)
	router.POST("/auth-renew", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server))
	authRoutes.GET("/profile", server.profile)
	authRoutes.POST("/kontrak", server.kontrakCreate)

	// authRoutes.GET("/konsumen", server.KonsumenGetList)
	// authRoutes.POST("/konsumen", server.konsumenCreate)
	// authRoutes.POST("/konsumen/:id", server.KonsumenUpdate)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// func validJSON(ctx *gin.Context) (bool, error) {
// 	jsonData, err := io.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		fmt.Println("====>")
// 		fmt.Println(err)
// 		return false, errors.New("bad request")
// 	}

// 	if !json.Valid(jsonData) {
// 		fmt.Println(">")
// 		fmt.Println(json.Valid(jsonData))
// 		return false, errors.New("bad request")
// 	}

// 	return true, nil
// }
