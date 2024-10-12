package service

import (
	"github.com/LastDarkNes/go-dns/internal/model"
	"github.com/LastDarkNes/go-dns/internal/repository"
	"github.com/LastDarkNes/go-dns/pkg/config"
)

type DomainRepository interface {
	GetByName(string) (*model.Domain, error)
	Save(model.Domain) (*model.Domain, error)
	GetPaginated(limit int, offset int) []model.Domain
}

type domainService struct {
	repo DomainRepository
}

func NewDomainService() *domainService {
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
	return d.repo.GetPaginated(
		config.CommonConfig.ItemsPerPaginationPage,
		page*config.CommonConfig.ItemsPerPaginationPage,
	)
}
