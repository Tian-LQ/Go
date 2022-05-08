package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// User PO持久化对象
type User struct {
	gorm.Model
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user biz.User) error {
	userPO := User{
		Email:        user.Email,
		Username:     user.Username,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: user.PasswordHash,
	}
	ret := ur.data.db.Create(&userPO)
	return ret.Error
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	userPO := User{}
	ret := ur.data.db.Where(&User{Email: email}).First(&userPO)
	if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not found by email")
	}
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &biz.User{
		Email:        userPO.Email,
		Username:     userPO.Username,
		Bio:          userPO.Bio,
		Image:        userPO.Image,
		PasswordHash: userPO.PasswordHash,
	}, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
