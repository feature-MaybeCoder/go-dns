package controller

import (
	"github.com/LastDarkNes/go-dns/model"
	"github.com/LastDarkNes/go-dns/service"
)

type DomainControllerInterface interface {
	Create(string) (model.Domain, error)
	Get(model.Domain) (model.Domain, error)
	GetPaginated(int) []model.Domain
}

type DomainController struct {
	Service service.DomainService
}

func (dc *DomainController) Create(new_domain model.Domain) (model.Domain, error) {
	return dc.Service.Create(new_domain)
}

func (dc *DomainController) Get(new_domain model.Domain) (model.Domain, error) {
	return dc.Service.Create(new_domain)
}

func (dc *DomainController) GetPaginated(page int) []model.Domain {
	return dc.Service.GetMultiple(10, 10*page)
}
