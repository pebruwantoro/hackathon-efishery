package organization_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	mockRepository "github.com/pebruwantoro/hackathon-efishery/internal/app/repository/mocks"
	"github.com/pebruwantoro/hackathon-efishery/internal/app/usecase/organization"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		ShouldError            bool
		IsAppError             bool
		ExpectedAppErrorStatus int
		CreateOrganizationResp entity.Organization
		CreateOrganizationErr  error
	}{
		"ShouldErrorWhenCreateOrganizationError": {
			ShouldError:            true,
			IsAppError:             true,
			ExpectedAppErrorStatus: 422,
			CreateOrganizationErr:  errors.New("create organization error"),
		},
		"ShouldSuccess": {
			ShouldError:            false,
			CreateOrganizationResp: entity.Organization{ID: 100},
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			organizationRepo := new(mockRepository.Organization)

			organizationRepo.On("Create", mock.Anything, mock.Anything).Return(test.CreateOrganizationErr).Once().Run(func(args mock.Arguments) {
				arg := args.Get(1).(*entity.Organization)
				arg.ID = test.CreateOrganizationResp.ID
			})

			usecase := organization.NewUsecase().SetOrganizationRepository(organizationRepo).Validate()

			resp, err := usecase.Create(context.Background(), organization.CreateUpdateOrganizationRequest{})

			if test.ShouldError {
				assert.NotNil(t, err)

				if test.IsAppError {
					var appErr *apperror.ApplicationError
					assert.ErrorAs(t, err, &appErr)

					appErr = err.(*apperror.ApplicationError)
					assert.Equal(t, test.ExpectedAppErrorStatus, appErr.Status)
				}
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
