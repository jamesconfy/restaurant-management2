package handlers

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Logout(c *gin.Context)
	ClearAuth(c *gin.Context)
}

type userHandler struct {
	userSrv service.UserService
}

var defaultCookieName = "Authorization"

// Register User godoc
// @Summary	Register route
// @Description	Register route
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Create	true "Signup Details"
// @Success	200  {object}  response.SuccessMessage{data=models.User}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/auth/register [post]
func (u *userHandler) Create(c *gin.Context) {
	var req forms.Create

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	req.Role = ""

	if strings.Contains(c.Request.RequestURI, "admin") {
		req.Role = "ADMIN"
	}

	user, err := u.userSrv.Create(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "user created successfully", user)
}

// Login User godoc
// @Summary	Login route
// @Description	Login route
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Login	true "Login Details"
// @Success	200  {object}  response.SuccessMessage{data=models.User}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/auth/login [post]
func (u *userHandler) Login(c *gin.Context) {
	var req forms.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	auth, err := u.userSrv.Login(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	u.setCookie(c, auth.AccessToken, 0)
	response.Success(c, "user logged in successfully", auth)
}

// Get User godoc
// @Summary	Get user route
// @Description	Get user by id
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	userId	path	string	true	"User id"
// @Success	200  {object}  response.SuccessMessage{data=models.User}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/users/{userId} [get]
func (u *userHandler) Get(c *gin.Context) {
	user, err := u.userSrv.Get(c.Param("userId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "user gotten successfully", user)
}

// Get All User godoc
// @Summary	Get all user route
// @Description	Get all users in the system
// @Tags	Users
// @Accept	json
// @Produce	json
// @Success	200  {object}  response.SuccessMessage{data=[]models.User}
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/users [get]
func (u *userHandler) GetAll(c *gin.Context) {
	users, err := u.userSrv.GetAll()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "user gotten successfully", users, len(users))
}

// Logout User godoc
// @Summary	Logout user route
// @Description	Logout user
// @Tags	User
// @Produce	json
// @Success	200  {string}	string	"Logged out successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/logout [post]
// @Security ApiKeyAuth
func (u *userHandler) Logout(c *gin.Context) {
	err := u.userSrv.DeleteAuth(c.GetString("userId"), u.getAuth(c))
	if err != nil {
		response.Error(c, *err)
		return
	}

	u.setCookie(c, "", -1)
	response.Success201(c, "Logged out successfully", nil)
}

// Clear Login Auth godoc
// @Summary	Clear Login Auth route
// @Description	Clear user auth
// @Tags	User
// @Produce	json
// @Success	200  {string}	string	"Logged out from all other device successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/profile/clear [post]
// @Security ApiKeyAuth
func (u *userHandler) ClearAuth(c *gin.Context) {
	err := u.userSrv.ClearAuth(c.GetString("userId"), u.getAuth(c))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success201(c, "Logged out from all other device successfully", nil)
}

func NewUserHandler(userSrv service.UserService) UserHandler {
	return &userHandler{userSrv: userSrv}
}

// Auxillary function
func (u *userHandler) setCookie(c *gin.Context, value string, max_age int) {
	c.SetCookie(defaultCookieName, value, 0, "/", "", false, true)
}

func (u *userHandler) getAuth(c *gin.Context) (auth string) {
	auth, _ = c.Cookie(defaultCookieName)

	if auth != "" {
		return
	}

	auth = c.GetHeader(defaultCookieName)
	return
}
