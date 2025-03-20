package services_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/a-novel/story-schematics/internal/daoai"
	"github.com/a-novel/story-schematics/internal/services"
	servicesmocks "github.com/a-novel/story-schematics/internal/services/mocks"
	"github.com/a-novel/story-schematics/models"
)

func TestGenerateLoglines(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type generateLoglinesData struct {
		resp []models.LoglineIdea
		err  error
	}

	testCases := []struct {
		name string

		request services.GenerateLoglinesRequest

		generateLoglinesData *generateLoglinesData

		expect    []models.LoglineIdea
		expectErr error
	}{
		{
			name: "Success",

			request: services.GenerateLoglinesRequest{
				Count:  5,
				Theme:  "test-theme",
				UserID: uuid.MustParse("00000000-0000-0000-1000-000000000001"),
			},

			generateLoglinesData: &generateLoglinesData{
				resp: []models.LoglineIdea{
					{
						Name:    "Logline 1",
						Content: "Content 1",
					},
					{
						Name:    "Logline 2",
						Content: "Content 2",
					},
				},
			},

			expect: []models.LoglineIdea{
				{
					Name:    "Logline 1",
					Content: "Content 1",
				},
				{
					Name:    "Logline 2",
					Content: "Content 2",
				},
			},
		},
		{
			name: "Error",

			request: services.GenerateLoglinesRequest{
				Count:  5,
				Theme:  "test-theme",
				UserID: uuid.MustParse("00000000-0000-0000-1000-000000000001"),
			},

			generateLoglinesData: &generateLoglinesData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			ctx := t.Context()

			source := servicesmocks.NewMockGenerateLoglinesSource(t)

			if testCase.generateLoglinesData != nil {
				source.EXPECT().
					GenerateLoglines(ctx, daoai.GenerateLoglinesRequest{
						Count:  testCase.request.Count,
						Theme:  testCase.request.Theme,
						UserID: testCase.request.UserID.String(),
					}).
					Return(testCase.generateLoglinesData.resp, testCase.generateLoglinesData.err)
			}

			service := services.NewGenerateLoglinesService(source)

			resp, err := service.GenerateLoglines(ctx, testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, resp)

			source.AssertExpectations(t)
		})
	}
}
