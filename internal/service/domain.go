package service

import (
	"github.com/LastDarkNes/go-dns/internal/model"
	"github.com/LastDarkNes/go-dns/internal/repository"
)

type DomainService interface {
	Get(string) (*model.Domain, error)
	Create(model.Domain) (*model.Domain, error)
	GetPage(int) []model.Domain
}

type domainService struct {
	repo repository.DomainRepository
}

func NewDomainService() DomainService {
	new_repo := repository.NewDomainRepository()
	return &domainService{
		repo: new_repo,
	}

}

func (d *domainService) Get(name string) (*model.Domain, error) {
	return d.repo.GetByName(name)
}

func (d *domainService) Create(domain model.Domain) (*model.Domain, error) {
	return d.repo.Save(domain)
}

func (d *domainService) GetPage(page int) []model.Domain {
	return d.repo.GetPaginated(10, page*10)
}
