package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/docker/distribution/uuid"
)

type UserService interface {
	Create(req *forms.User) (*models.User, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
	Get(userId string) (*models.User, *se.ServiceError)
	GetAll() ([]*models.User, *se.ServiceError)
	Edit(userId string, req *forms.EditUser) (*models.User, *se.ServiceError)
	Delete(userId string) *se.ServiceError
	DeleteAuth(userId, accessToken string) *se.ServiceError
	ClearAuth(userId, accessToken string) *se.ServiceError
}

type userSrv struct {
	userRepo repo.UserRepo
	authRepo repo.AuthRepo
	authSrv  AuthService
	emailSrv EmailService
}

func (u *userSrv) Create(req *forms.User) (*models.User, *se.ServiceError) {
	err := Validator.validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	if ok, err := u.userRepo.EmailExists(req.Email); ok {
		return nil, se.ConflictOrInternal(err, "user already exists")
	}

	if ok, err := u.userRepo.PhoneExists(req.PhoneNumber); ok {
		return nil, se.ConflictOrInternal(err, "phone already taken")
	}

	password, err := Crypto.hashPassword(req.Password)
	if err != nil {
		return nil, se.Internal(err, "could not hash password")
	}

	var user models.User

	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Avatar = req.Avatar
	user.Password = password
	user.Role = req.Role

	usr, err := u.userRepo.Add(&user)
	if err != nil {
		return nil, se.Internal(err)
	}

	return usr, nil
}

func (u *userSrv) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	err := Validator.validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user does not exist")
	}

	if ok := Crypto.comparePassword(user.Password, req.Password); !ok {
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

	ath.User = user
	return ath, nil
}

func (u *userSrv) Get(userId string) (*models.User, *se.ServiceError) {
	if _, err := uuid.Parse(userId); err != nil {
		return nil, se.Internal(err, "invalid user id")
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

func (u *userSrv) Edit(userId string, req *forms.EditUser) (*models.User, *se.ServiceError) {
	err := Validator.validate(req)
	if err != nil {
		return nil, se.Validating(err)
	}

	_, err = uuid.Parse(userId)
	if err != nil {
		return nil, se.Internal(err, "invalid user id")
	}

	user, err := u.userRepo.GetById(userId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	user, er := u.getEdit(req, user)
	if er != nil {
		return nil, er
	}

	user, err = u.userRepo.Edit(userId, user)
	if err != nil {
		return nil, se.Internal(err, "error when editing user")
	}

	return user, nil
}

func (u *userSrv) Delete(userId string) *se.ServiceError {
	_, err := uuid.Parse(userId)
	if err != nil {
		return se.Internal(err, "invalid user id")
	}

	err = u.userRepo.Delete(userId)
	if err != nil {
		return se.NotFoundOrInternal(err, "user not found")
	}

	return nil
}

func (u *userSrv) DeleteAuth(userId, accessToken string) *se.ServiceError {
	if _, err := uuid.Parse(userId); err != nil {
		return se.Internal(err, "invalid user id")
	}

	err := u.authRepo.Delete(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func (u *userSrv) ClearAuth(userId, accessToken string) *se.ServiceError {
	if _, er := uuid.Parse(userId); er != nil {
		return se.Internal(er, "invalide user id")
	}

	err := u.authRepo.Clear(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func NewUserService(repo repo.UserRepo, authRepo repo.AuthRepo, authSrv AuthService, emailSrv EmailService) UserService {
	return &userSrv{userRepo: repo, authRepo: authRepo, authSrv: authSrv, emailSrv: emailSrv}
}

// Auxillary Functions
func (u *userSrv) getEdit(req *forms.EditUser, user *models.User) (*models.User, *se.ServiceError) {
	if req.FirstName != user.FirstName && req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != user.LastName && req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Address != user.Address && req.Address != "" {
		user.Address = req.Address
	}

	if req.Avatar != user.Avatar && req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if req.Email != user.Email && req.Email != "" {
		ok, err := u.userRepo.EmailExists(req.Email)
		if ok {
			return nil, se.Conflict("email already in use")
		}

		if err != nil {
			return nil, se.Internal(err)
		}

		user.Email = req.Email
	}

	if req.PhoneNumber != user.PhoneNumber && req.PhoneNumber != "" {
		ok, err := u.userRepo.PhoneExists(req.PhoneNumber)
		if ok {
			return nil, se.Conflict("phone already in use")
		}

		if err != nil {
			return nil, se.Internal(err)
		}

		user.PhoneNumber = req.PhoneNumber
	}

	return user, nil
}
