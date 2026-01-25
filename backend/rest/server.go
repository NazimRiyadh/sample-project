package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func Server(cnf config.Config) {
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.Logger, middlewares.Hudai, middlewares.Arekta, middlewares.PreflightRequest, middlewares.CORS)

	InitRoutes(mux, *manager)

	port := ":" + strconv.Itoa(cnf.HttpPort)

	//globalRouter := middlewares.GlobalRouter(mux)
	err := http.ListenAndServe(port, manager.With(mux))
	fmt.Println("Server started running on port", port)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
