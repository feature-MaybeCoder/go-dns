package repository

import (
	"github.com/LastDarkNes/go-dns/internal/model"
	"github.com/LastDarkNes/go-dns/pkg/custom_errors"
)

type DomainRepository interface {
	GetByName(string) (*model.Domain, error)
	Save(model.Domain) (*model.Domain, error)
	GetPaginated(limit int, offset int) []model.Domain
}

type domainRepository struct {
	domains []model.Domain
}

func NewDomainRepository() DomainRepository {
	return &domainRepository{}
}

func (d *domainRepository) GetByName(name string) (*model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == name {
			return &domain, nil
		}
	}

	return nil, custom_errors.NotFoundError{ObjName: "domain"}

}

func (d *domainRepository) Save(new_domain model.Domain) (*model.Domain, error) {
	for _, domain := range d.domains {
		if domain.Name == new_domain.Name {
			return nil, custom_errors.UnexpectedBehaviourError{Message: "Domain with such name already exists"}
		}
	}
	d.domains = append(d.domains, new_domain)

	return &new_domain, nil

}

func (d *domainRepository) GetPaginated(limit int, offset int) []model.Domain {
	n := len(d.domains)

	size := limit

	if offset >= n {
		return nil
	}

	if offset+limit > n {
		size = n - offset
	}

	foundDomains := make([]model.Domain, size)

	copy(foundDomains, d.domains[offset:offset+size])

	return foundDomains
}
