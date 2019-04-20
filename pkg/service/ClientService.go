package service

import (
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/repository"
)

type ClientService struct {
	repo *repository.ClientRepository
}

//NewService create new service
func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{
		repo: r,
	}
}

func (s *ClientService) Find(b int) (*model.Client, error) {
	return s.repo.Find(b)
}

//FindByUsername
func (s *ClientService) FindByUsername(clientname string) ([]*model.Client, error) {
	return s.repo.FindByKey("client_name", clientname)
}

func (s *ClientService) Store(b *model.Client) (int, error) {
	return s.repo.Store(b)
}
