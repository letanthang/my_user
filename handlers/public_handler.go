package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/my_framework/db"
	types "github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/my_framework/utils"
	"github.com/letanthang/validator"
)

func LoginAccount(c echo.Context) error {
	var req types.LoginReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	account, err := db.GetOneAccount(req.Phone, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}

	token, _ := utils.GetToken(account.ID, "driver", account.FirstName+account.LastName, account.Phone, account.Email)
	response := echo.Map{
		"account": account,
		"token":   token,
	}
	return c.JSON(http.StatusOK, response)
}

func RegisterAccount(c echo.Context) error {
	var req types.AccountAddReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	if err := validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	res, _ := db.InsertAccount(req)
	return c.JSON(http.StatusOK, res)
}

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
