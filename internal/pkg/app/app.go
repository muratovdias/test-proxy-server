package app

import (
	"fmt"
	"net/http"

	"github.com/muratovdias/test-proxy-server/internal/app/cache"
	"github.com/muratovdias/test-proxy-server/internal/app/delivery"
	"github.com/muratovdias/test-proxy-server/internal/app/usecase"
)

type App struct {
	cache   *cache.Cache
	usecase *usecase.Usecase
	handler *delivery.Handler
}

func NewApp() *App {
	var app App
	app.cache = cache.NewCache()
	app.usecase = usecase.NewUsecase(app.cache)
	app.handler = delivery.NewHandler(app.usecase)
	return &app
}

func (app *App) Run() error {
	router := app.handler.InitRoutes()
	server := ServerUp(router)
	fmt.Println("Server started at port 8888")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// fmt.Println("Server started at port 8888")
	return nil
}

func ServerUp(handler *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":8888",
		Handler: handler,
	}
}
