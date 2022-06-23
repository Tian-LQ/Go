package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(context context.Context, request *v1.LoginRequest) (*v1.UserReply, error) {
	userLogin, err := s.uc.Login(context, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    userLogin.Email,
			Token:    userLogin.Token,
			Username: userLogin.Username,
			Bio:      userLogin.Bio,
			Image:    userLogin.Image,
		},
	}, nil
}

func (s *RealWorldService) Register(context context.Context, request *v1.RegisterRequest) (*v1.UserReply, error) {
	ul, err := s.uc.Register(context, request.User.Username, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    ul.Email,
			Username: ul.Username,
			Token:    ul.Token,
		},
	}, nil
}

func (s *RealWorldService) UpdateUser(context context.Context, request *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) GetProfile(context context.Context, request *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
