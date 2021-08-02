package api

import (
	_ "efishery-be-test/app/auth/docs"
	"efishery-be-test/app/auth/interfaces"
	"efishery-be-test/app/auth/models"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	BadContentType = map[string]string{"remark": "Content type must be application/json"}
	BadRequestBody = map[string]string{"remark": "Invalid parse body request to JSON"}
	MissingAPIKey = map[string]string{"remark": "Missing API key"}
)

type RESTAPI struct {
	server  *echo.Echo
	host    string
	port    int
	service interfaces.Service
}

// RegisterUser godoc
// @Summary Register new user
// @Id RegisterUser
// @Tags Auth
// @Success 200 {object} models.RespRegisterUser "{"password": "{password}"}"
// @Failure 400 {object} models.RespError "{"remark": "Content type must be application/json"}"
// @Failure 400 {object} models.RespError "{"remark": "Invalid parse body request to JSON"}"
// @Failure 422 {object} models.RespError "{"remark": "Phone number has been registered"}"
// @Router /auth/register [post]
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

// LoginUser godoc
// @Summary Login user
// @Id LoginUser
// @Tags Auth
// @Success 200 {object} models.RespLoginUser "{"token": "{token}"}"
// @Failure 400 {object} models.RespError "{"remark": "Content type must be application/json"}"
// @Failure 400 {object} models.RespError "{"remark": "Invalid parse body request to JSON"}"
// @Failure 401 {object} models.RespError "{"remark": "Invalid authentication for phone {phone}"}"
// @Router /auth/login [post]
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

// VerifyToken godoc
// @Summary Verify and extract JWT
// @Id VerifyToken
// @Tags Auth
// @Security ApiKeyAuth
// @Success 200 {object} models.RespVerifyToken "{"claims": Model}"
// @Failure 400 {object} models.RespError "{"remark": "Content type must be application/json"}"
// @Failure 400 {object} models.RespError "{"remark": "Invalid parse body request to JSON"}"
// @Failure 401 {object} models.RespError "{"remark": "Invalid token verification"}"
// @Router /auth/verify [post]
func (r *RESTAPI) verify(c echo.Context) (err error) {
	var claims models.RespVerifyToken
	resp := map[string]interface{}{}
	apikey := c.Request().Header.Get("Authorization")
	if apikey == "" {
		return c.JSON(http.StatusBadRequest, MissingAPIKey)
	}
	payload, err := r.service.VerifyToken(apikey)
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
	r.server.GET("/swagger/*", echoSwagger.WrapHandler)
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
