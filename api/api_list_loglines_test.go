package api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	authapi "github.com/a-novel/authentication/api"
	authmodels "github.com/a-novel/authentication/models"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/story-schematics/api"
	"github.com/a-novel/story-schematics/api/codegen"
	apimocks "github.com/a-novel/story-schematics/api/mocks"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

func TestListLoglines(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type listLoglinesData struct {
		resp []*models.LoglinePreview
		err  error
	}

	testCases := []struct {
		name string

		params codegen.GetLoglinesParams

		listLoglinesData *listLoglinesData

		expect    codegen.GetLoglinesRes
		expectErr error
	}{
		{
			name: "Success",

			params: codegen.GetLoglinesParams{
				Limit:  codegen.OptInt{Value: 10, Set: true},
				Offset: codegen.OptInt{Value: 2, Set: true},
			},

			listLoglinesData: &listLoglinesData{
				resp: []*models.LoglinePreview{
					{
						Slug:      "slug-1",
						Name:      "Logline 1",
						Content:   "Logline 1 content",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						Slug:      "slug-2",
						Name:      "Logline 2",
						Content:   "Logline 2 content",
						CreatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
					},
				},
			},

			expect: &codegen.GetLoglinesOKApplicationJSON{
				{
					Slug:      "slug-1",
					Name:      "Logline 1",
					Content:   "Logline 1 content",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					Slug:      "slug-2",
					Name:      "Logline 2",
					Content:   "Logline 2 content",
					CreatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Error",

			params: codegen.GetLoglinesParams{
				Limit:  codegen.OptInt{Value: 10, Set: true},
				Offset: codegen.OptInt{Value: 2, Set: true},
			},

			listLoglinesData: &listLoglinesData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockListLoglinesService(t)

			ctx := context.WithValue(t.Context(), authapi.ClaimsAPIKey{}, &authmodels.AccessTokenClaims{
				UserID: lo.ToPtr(uuid.MustParse("00000000-1000-0000-0000-000000000001")),
			})

			if testCase.listLoglinesData != nil {
				source.EXPECT().
					ListLoglines(ctx, services.ListLoglinesRequest{
						UserID: uuid.MustParse("00000000-1000-0000-0000-000000000001"),
						Limit:  testCase.params.Limit.Value,
						Offset: testCase.params.Offset.Value,
					}).
					Return(testCase.listLoglinesData.resp, testCase.listLoglinesData.err)
			}

			handler := api.API{ListLoglinesService: source}

			res, err := handler.GetLoglines(ctx, testCase.params)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
