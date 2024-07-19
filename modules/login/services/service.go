package services

import (
	"log"
	"users_v1/modules/login/models"
	"users_v1/modules/login/repositories"
)

type IService interface {
	LoginS(login models.LoginRequest) error
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) LoginS(login models.LoginRequest) error {
	err := s.r.LoginR(login)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
