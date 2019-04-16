package service

import (
	model "DemoGO/pkg/models"
	"DemoGO/pkg/repository"
)

type OfferService struct {
	repo *repository.OffersRepository
}

//NewService create new service
func NewOfferService(r *repository.OffersRepository) *OfferService {
	return &OfferService{
		repo: r,
	}
}

//Store an bookmark
func (s *OfferService) Store(b *model.Offer) (*model.Offer, error) {
	return s.repo.Store(b)
}

//Find a bookmark
func (s *OfferService) Find(id model.ID) (*model.Offer, error) {
	return s.repo.Find(id)
}

/*//Search bookmarks
func (s *Service) Search(query string) ([]*model.Bid, error) {
	return s.repo.Search(strings.ToLower(query))
}
*/
//FindAll bookmarks
func (s *OfferService) FindAll() ([]model.Offer, error) {
	return s.repo.FindAll()
}

//Delete a bookmark
func (s *OfferService) Delete(id model.ID) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
