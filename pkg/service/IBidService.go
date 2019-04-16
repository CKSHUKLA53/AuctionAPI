package service

import (
	model "DemoGO/pkg/models"
	"DemoGO/pkg/repository"
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
func (s *BidService) Find(id model.ID) (*model.Bid, error) {
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
func (s *BidService) Delete(id model.ID) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
