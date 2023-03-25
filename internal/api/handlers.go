package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rarifbkn/api-rest-gorm-echo/encryption"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/api/dtos"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/models"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/service"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})

	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	//TODO create JWT
	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error "})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie) // se le pasa al response

	return c.JSON(http.StatusOK, map[string]string{"succes": "true"})

}

func (a *API) AddProduct(c echo.Context) error {
	//	TODO: get auh token from cookie
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})

	}

	// TODO: get the payload from request

	email := claims["email"].(string)

	//	TODO: parse the jwt
	ctx := c.Request().Context()
	params := dtos.AddProduct{}

	err = c.Bind(&params)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	p := models.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	err = a.serv.AddProduct(ctx, p, email)
	if err != nil {
		log.Println(err)
		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Invalid Permissions"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, nil)
}
