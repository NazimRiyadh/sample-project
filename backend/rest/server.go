package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/Product"
	"ecommerce/rest/handlers/Test"
	"ecommerce/rest/handlers/User"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

//InitRoutes(mux, *manager)

type Server struct {
	productHandler *Product.ProductHandler
	userHandler    *User.UserHandler
	testHandler    *Test.TestHandler
	config         *config.Config
}

func NewServer(productHandler *Product.ProductHandler,
	userHandler *User.UserHandler,
	testHandler *Test.TestHandler,
	config *config.Config) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		testHandler:    testHandler,
		config:         config,
	}
}

func (s *Server) Server() {
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.Hudai, middlewares.Arekta, middlewares.PreflightRequest, middlewares.CORS)

	port := ":" + strconv.Itoa(s.config.HttpPort)

	s.productHandler.RegProductRoutes(mux, manager)
	s.userHandler.RegUserRoutes(mux, manager)
	s.testHandler.RegtestHandler(mux, manager)

	//globalRouter := middlewares.GlobalRouter(mux)
	fmt.Println("Server started running on port", port)
	err := http.ListenAndServe(port, manager.With(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
