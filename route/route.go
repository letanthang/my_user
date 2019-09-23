package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/letanthang/my_framework/config"
	"github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/my_framework/handlers"
)

func All(e *echo.Echo) {
	Public(e)
	Staff(e)
	Private(e)
}

func Public(e *echo.Echo) {
	publicRoute := e.Group("/api/v1/public")
	publicRoute.POST("/user/register", handlers.RegisterAccount)
	publicRoute.PATCH("/user/login", handlers.LoginAccount)
	publicRoute.GET("/health", handlers.CheckHealth)
}

func Staff(e *echo.Echo) {
	staffRoute := e.Group("/api/v1/staff")
	staffRoute.DELETE("/user", handlers.DeleteAccount)
}

func Private(e *echo.Echo) {

	jwtConfig := middleware.JWTConfig{
		Claims:     &types.MyClaims{},
		SigningKey: []byte(config.Config.Encryption.JWTSecret),
	}
	privateRoute := e.Group("/api/v1/private")
	privateRoute.Use(middleware.JWTWithConfig(jwtConfig))
	privateRoute.PUT("/user", handlers.UpdateAccount)

}
