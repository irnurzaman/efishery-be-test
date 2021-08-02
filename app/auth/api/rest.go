package api

import (
	"efishery-be-test/app/auth/interfaces"
	"efishery-be-test/app/auth/models"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var (
	BadContentType = map[string]string{"reason": "Content type must be application/json"}
	BadRequestBody = map[string]string{"reason": "Invalid parse body request to JSON"}
)

type RESTAPI struct {
	server  *echo.Echo
	host    string
	port    int
	service interfaces.Service
}

func (r *RESTAPI) register(c echo.Context) (err error) {
	var request models.ReqRegisterUser
	resp := map[string]interface{}{}
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusBadRequest, BadContentType)
	}
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequestBody)
	}
	pwd, err := r.service.RegisterUser(request)
	if err != nil {
		resp["remark"] = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}
	resp["password"] = pwd
	return c.JSON(http.StatusOK, resp)
}

func (r *RESTAPI) login(c echo.Context) (err error) {
	var request models.ReqLoginUser
	resp := map[string]interface{}{}
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusBadRequest, BadContentType)
	}
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequestBody)
	}
	token, err := r.service.VerifyUser(request)
	if err != nil {
		resp["remark"] = err.Error()
		return c.JSON(http.StatusUnauthorized, resp)
	}
	resp["token"] = token
	return c.JSON(http.StatusOK, resp)
}

func (r *RESTAPI) verify(c echo.Context) (err error) {
	var request models.ReqVerifyToken
	var claims models.RespVerifyToken
	resp := map[string]interface{}{}
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusBadRequest, BadContentType)
	}
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequestBody)
	}
	payload, err := r.service.VerifyToken(request)
	if err != nil {
		resp["remark"] = err.Error()
		return c.JSON(http.StatusUnauthorized, resp)
	}
	copier.Copy(&claims, payload)
	resp["claims"] = claims
	return c.JSON(http.StatusOK, resp)
}

func (r *RESTAPI) Run() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method} uri=${uri} status=${status} error=${error} ip=${remote_ip} \n",
	}))
	g := r.server.Group("/auth")
	g.POST("/register", r.register)
	g.POST("/login", r.login)
	g.POST("/verify", r.verify)

	addr := fmt.Sprintf("%s:%d", r.host, r.port)
	r.server.Start(addr)
}

func NewRESTAPI(host string, port int, service interfaces.Service) *RESTAPI {
	return &RESTAPI{
		server:  echo.New(),
		host:    host,
		port:    port,
		service: service,
	}
}
