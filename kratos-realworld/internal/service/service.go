package service

import (
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUsecase
}

func NewRealWorldService(uc *biz.UserUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
