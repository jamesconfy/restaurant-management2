package service

import (
	"fmt"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/docker/distribution/uuid"
)

type UserService interface {
	Create(req *forms.Create) (*models.User, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
	Get(userId string) (*models.User, *se.ServiceError)
	GetAll() ([]*models.User, *se.ServiceError)
	DeleteAuth(userId, accessToken string) *se.ServiceError
	ClearAuth(userId, accessToken string) *se.ServiceError
}

type userSrv struct {
	userRepo     repo.UserRepo
	authRepo     repo.AuthRepo
	validatorSrv ValidationService
	cryptoSrv    CryptoService
	authSrv      AuthService
	emailSrv     EmailService
	cashbin      *casbin.Enforcer
}

func (u *userSrv) Create(req *forms.Create) (*models.User, *se.ServiceError) {
	err := u.validatorSrv.Validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	if ok, err := u.userRepo.EmailExists(req.Email); ok {
		return nil, se.ConflictOrInternal(err, "user already exists")
	}

	if ok, err := u.userRepo.PhoneExists(req.PhoneNumber); ok {
		return nil, se.ConflictOrInternal(err, "phone already taken")
	}

	password, err := u.cryptoSrv.HashPassword(req.Password)
	if err != nil {
		return nil, se.Internal(err, "could not hash password")
	}

	var user models.User

	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Password = password
	user.Role = req.Role

	usr, err := u.userRepo.Add(&user)
	if err != nil {
		return nil, se.Internal(err)
	}

	defer func() {
		obj := fmt.Sprintf("/api/v1/users/%v", usr.Id)
		u.cashbin.AddPolicy(usr.Id, obj, "(GET)|(POST)|(DELETE)|(PATCH)|(PUT)", "allow")
		u.cashbin.AddGroupingPolicy(usr.Id, fmt.Sprintf("role::%v", strings.ToLower(usr.Role)))

		u.cashbin.SavePolicy()
	}()

	return usr, nil
}

func (u *userSrv) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	err := u.validatorSrv.Validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user does not exist")
	}

	ok := u.cryptoSrv.ComparePassword(user.Password, req.Password)
	if !ok {
		return nil, se.BadRequest("password does not match")
	}

	auth := new(models.Auth)
	auth.UserId = user.Id

	auth.AccessToken, auth.RefreshToken, err = u.authSrv.Create(user)
	if err != nil {
		return nil, se.Internal(err, "Error when creating token")
	}

	// Create auth row
	ath, err := u.authRepo.Add(auth)
	if err != nil {
		return nil, se.Internal(err, "Error when adding/updating user token")
	}

	return ath, nil
}

func (u *userSrv) Get(userId string) (*models.User, *se.ServiceError) {
	_, err := uuid.Parse(userId)
	if err != nil {
		return nil, se.NotFound("user not found")
	}

	user, err := u.userRepo.GetById(userId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	return user, nil
}

func (u *userSrv) GetAll() ([]*models.User, *se.ServiceError) {
	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return users, nil
}

func (u *userSrv) DeleteAuth(userId, accessToken string) *se.ServiceError {
	if _, er := uuid.Parse(userId); er != nil {
		return se.NotFound("user not found")
	}

	err := u.authRepo.Delete(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func (u *userSrv) ClearAuth(userId, accessToken string) *se.ServiceError {
	if _, er := uuid.Parse(userId); er != nil {
		return se.NotFound("user not found")
	}

	err := u.authRepo.Clear(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func NewUserService(repo repo.UserRepo, authRepo repo.AuthRepo, validator ValidationService, cryptoSrv CryptoService, authSrv AuthService, emailSrv EmailService, cashbin *casbin.Enforcer) UserService {
	return &userSrv{userRepo: repo, authRepo: authRepo, validatorSrv: validator, cryptoSrv: cryptoSrv, authSrv: authSrv, emailSrv: emailSrv, cashbin: cashbin}
}
