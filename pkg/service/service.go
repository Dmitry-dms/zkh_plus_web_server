package service

import "github.com/dmitry-dms/rest-gin/pkg/repository"

type Authorization interface {
}
type UserList interface {
}

type Service struct {
	Authorization
	UserList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
