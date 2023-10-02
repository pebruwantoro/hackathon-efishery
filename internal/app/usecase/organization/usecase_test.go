package organization_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mockRepository "github.com/pebruwantoro/hackathon-efishery/internal/app/repository/mocks"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/organization"
)

func TestNewUsecase(t *testing.T) {
	organizationRepo := new(mockRepository.Organization)

	t.Run("ShouldPanicWhenOrganizationRepoIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			organization.NewUsecase().Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			organization.NewUsecase().SetOrganizationRepository(organizationRepo).Validate()
		})
	})
}
