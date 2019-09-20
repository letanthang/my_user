package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
	types "github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/validator"
)

func UpdateAccount(c echo.Context) error {
	var req types.AccountUpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: "Bad parameter"})
	}

	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: err.Error()})
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.MyClaims)
	i, err := db.UpdateAccount(claims, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "bad request", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, i)
}
