package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err == nil
}

type UserRepo interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur      UserRepo
	pr      ProfileRepo
	jwtConf *conf.JWT
	log     *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtConf *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtConf: jwtConf, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtConf.GetToken(), username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	user := User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	err := uc.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	user, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(user.PasswordHash, password) {
		return nil, errors.Unauthorized("user", "login failed")
	}
	return &UserLogin{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    uc.generateToken(user.Username),
	}, nil
}
