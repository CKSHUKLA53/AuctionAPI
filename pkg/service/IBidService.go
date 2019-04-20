package service

import (
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/repository"
)

type BidService struct {
	repo *repository.BidRepository
}

//NewService create new service
func NewBidService(r *repository.BidRepository) *BidService {
	return &BidService{
		repo: r,
	}
}

func (s *BidService) Store(b *model.Bid) (*model.Bid, error) {
	return s.repo.Store(b)
}

func (s *BidService) Find(id int) (*model.Bid, error) {
	return s.repo.Find(id)
}

func (s *BidService) FindAll() ([]model.Bid, error) {
	return s.repo.FindAll()
}

func (s *BidService) Delete(id int) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *BidService) Update(id int, key string, value interface{}) (*model.Bid, error) {
	return s.repo.Update(id, key, value)
}
