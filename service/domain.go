package service

import (
	"github.com/LastDarkNes/go-dns/custom_errors"
	"github.com/LastDarkNes/go-dns/model"
)

type DomainServiceInterface interface {
	Get(string) (model.Domain, error)
	Create(model.Domain) model.Domain
	GetMultiple(int, int) []model.Domain
}

type DomainService struct {
	domains []model.Domain
}

func (d *DomainService) Get(name string) (model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == name {
			return domain, nil
		}
	}

	return model.Domain{}, custom_errors.NotFoundError{ObjName: "domain"}

}

func (d *DomainService) Create(new_domain model.Domain) (model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == new_domain.Name {
			return model.Domain{}, custom_errors.UnexpectedBehaviourError{Message: "Domain with such name already exists"}
		}
	}
	d.domains = append(d.domains, new_domain)

	return new_domain, nil

}

func (d *DomainService) GetMultiple(limit int, offset int) []model.Domain {
	var foundDomains []model.Domain
	total_domains := len(d.domains)
	for index, domain := range d.domains {
		if index >= offset || index < total_domains {
			foundDomains = append(foundDomains, domain)
		}
	}
	return foundDomains

}
