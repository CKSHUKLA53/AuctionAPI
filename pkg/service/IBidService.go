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

//Store an bookmark
func (s *BidService) Store(b *model.Bid) (*model.Bid, error) {
	return s.repo.Store(b)
}

//Find a bookmark
func (s *BidService) Find(id int) (*model.Bid, error) {
	return s.repo.Find(id)
}

/*//Search bookmarks
func (s *Service) Search(query string) ([]*model.Bid, error) {
	return s.repo.Search(strings.ToLower(query))
}
*/
//FindAll bookmarks
func (s *BidService) FindAll() ([]model.Bid, error) {
	return s.repo.FindAll()
}

//Delete a bookmark
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
