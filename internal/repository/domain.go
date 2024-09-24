package repository

import (
	"github.com/LastDarkNes/go-dns/internal/model"
	"github.com/LastDarkNes/go-dns/pkg/custom_errors"
)

type DomainRepository interface {
	GetByName(string) (model.Domain, error)
	Save(model.Domain) (model.Domain, error)
	GetPaginated(int, int) []model.Domain
}

type domainRepository struct {
	domains []model.Domain
}

func NewDomainRepository() domainRepository {
	return domainRepository{}
}

func (d *domainRepository) GetByName(name string) (model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == name {
			return domain, nil
		}
	}

	return model.Domain{}, custom_errors.NotFoundError{ObjName: "domain"}

}

func (d *domainRepository) Save(new_domain model.Domain) (model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == new_domain.Name {
			return model.Domain{}, custom_errors.UnexpectedBehaviourError{Message: "Domain with such name already exists"}
		}
	}
	d.domains = append(d.domains, new_domain)

	return new_domain, nil

}

func (d *domainRepository) GetPaginated(limit int, offset int) []model.Domain {
	var foundDomains []model.Domain
	total_domains := len(d.domains)
	for index, domain := range d.domains {
		if index >= offset || index < total_domains {
			foundDomains = append(foundDomains, domain)
		}
	}
	return foundDomains

}
