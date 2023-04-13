package server

import (
	"booking-service/infra"
	"booking-service/internal/repos"
	"booking-service/internal/services/pricing"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
}

type ServiceList struct {
	pricing pricing.IPricing
}

func NewServer() (*Server, error) {
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{
		router: router,
	}

	return s, nil
}

func (s *Server) Init() {
	mongodb, err := infra.InitMongoDB()
	if err != nil {
		zap.S().Errorf("Init mongodb error: %v", err)
		panic(err)
	}

	repo := repos.NewMongoDBRepo(mongodb)
	serviceList := s.initServices(repo)

	s.initRouters(serviceList)
}

func (s *Server) Listen() error {
	//address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	//fmt.Sprintf(":%s", os.Getenv("PORT"))
	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    ":3002",
	}

	return s.httpServer.ListenAndServe()
}
