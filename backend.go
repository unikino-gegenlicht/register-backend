package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"register-backend/internal/configuration"
	_ "register-backend/internal/database"
	"register-backend/internal/middleware"
	"register-backend/routes/articles"
	"register-backend/routes/tickets"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/gin-contrib/requestid"
	"github.com/wisdom-oss/common-go/v3/types"
)

func main() {

	config := configuration.Config.Sub("http")

	router := gin.New()
	router.Use(requestid.New())
	router.Use(middleware.ErrorHandler)
	router.Use(gin.CustomRecovery(middleware.RecoveryHandler))
	router.HandleMethodNotAllowed = true
	router.RedirectFixedPath = true
	router.UseH2C = true
	router.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, types.ServiceError{
			Type:   "https://www.rfc-editor.org/rfc/rfc9110.html#section-15.5.6",
			Status: http.StatusMethodNotAllowed,
			Title:  "Method Not Allowed",
			Detail: "The used HTTP method is not allowed on this route. Please check the documentation and your request",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, types.ServiceError{
			Type:   "https://www.rfc-editor.org/rfc/rfc9110.html#section-15.5.5",
			Status: http.StatusNotFound,
			Title:  "Route Not Found",
			Detail: "The requested path does not exist in this microservice. Please check the documentation and your request",
		})
	})

	// fixing not getting the json field name in validation errors as it's a
	// buggy behavior in go-playground/validator
	//
	// ref: https://github.com/go-playground/validator/issues/935
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})
	}

	/* Route Configuration */
	article := router.Group("/articles")
	{
		article.GET("/", articles.GetAll)
		article.PUT("/", articles.New)
		article.GET("/:articleID", articles.GetSingle)
		article.PATCH("/:articleID", articles.Edit)
		article.DELETE("/:articleID", articles.Delete)
	}
	ticketing := router.Group("/tickets")
	{
		ticketing.GET("/", tickets.GetTypes)
		ticketing.POST("/convert-reservation/:reservationID", tickets.ConvertReservation)
		ticketing.POST("/", tickets.IssueTicket)
	}

	reservations := router.Group("/reservation")
	{
		reservations.POST("/new")
		reservations.GET("/:reservationID/")
		reservations.GET("/:reservationID/tickets", tickets.ConvertReservation)
	}

	server := &http.Server{
		Addr:              net.JoinHostPort(config.GetString("host"), config.GetString("port")),
		Handler:           router.Handler(),
		ReadHeaderTimeout: 30 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("an error occurred while running the http server", "error", err.Error())
		}
	}()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("ready to serve requests")
	<-shutdownSignal

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("an error occurred shutting down the backend", "error", err.Error())
	}

}
