package organization

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/repository"
)

type OrganizationUsecase interface {
	Create(ctx context.Context, req CreateUpdateOrganizationRequest) (resp OrganizationResponse, err error)
}

type usecase struct {
	organizationRepository repository.Organization
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetOrganizationRepository(repo repository.Organization) *usecase {
	s.organizationRepository = repo
	return s
}

func (s *usecase) Validate() OrganizationUsecase {
	if s.organizationRepository == nil {
		panic("organizationRepository is nil")
	}

	return s
}
