package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"moneda/evaluation/flight"
	"moneda/evaluation/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func main() {
	flightRepository, err := flight.NewRepository(flight.NewConfig())

	if err != nil {
		panic(err)
	}

	flightMapper := flight.NewMapper()
	flightService := flight.NewService(flightRepository, flightMapper)
	securityService := security.NewService()

	httpPort := os.Getenv("PORT")
	e := echo.New()

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			return securityService.ApiKeyExists(c.Request().Context(), key)
		},
	}))

	e.POST("/getFlightData", func(c echo.Context) error {
		// implement your code how you wish here
		log.Println("getFlightData, request: ", c.Request().Body)
		req := &flight.GetFlightDataRequest{}
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		resp, err := flightService.GetFlightData(c.Request().Context(), req)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, resp)
	})

	e.POST("/searchFlightInfo", func(c echo.Context) error {
		// implement your code how you wish here
		req := &flight.SearchFlightInfoRequest{}

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		resp, err := flightService.SearchFlightInfo(c.Request().Context(), req)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, resp)
	})

	server := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	go func() {
		if err := e.StartH2CServer(fmt.Sprintf(":%s", httpPort), server); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(fmt.Sprintf("Shutting down the server: %s", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
