package api_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	authapi "github.com/a-novel/authentication/api"
	authmodels "github.com/a-novel/authentication/models"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/story-schematics/api"
	"github.com/a-novel/story-schematics/api/codegen"
	apimocks "github.com/a-novel/story-schematics/api/mocks"
	"github.com/a-novel/story-schematics/internal/dao"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

func TestRegenerateBeats(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type regenerateBeatsData struct {
		resp []models.Beat
		err  error
	}

	testCases := []struct {
		name string

		form *codegen.RegenerateBeatsForm

		regenerateBeatsData *regenerateBeatsData

		expect    codegen.RegenerateBeatsRes
		expectErr error
	}{
		{
			name: "Success",

			form: &codegen.RegenerateBeatsForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				RegenerateKeys: []string{
					"beat-1",
					"beat-2",
				},
			},

			regenerateBeatsData: &regenerateBeatsData{
				resp: []models.Beat{
					{
						Key:     "beat-1",
						Title:   "Regenerated Beat 1",
						Content: "Regenerated Content 1",
					},
					{
						Key:     "beat-2",
						Title:   "Regenerated Beat 2",
						Content: "Regenerated Content 2",
					},
				},
			},

			expect: &codegen.BeatsSheet{
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Regenerated Beat 1",
						Content: "Regenerated Content 1",
					},
					{
						Key:     "beat-2",
						Title:   "Regenerated Beat 2",
						Content: "Regenerated Content 2",
					},
				},
			},
		},
		{
			name: "BeatsSheetNotFound",

			form: &codegen.RegenerateBeatsForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				RegenerateKeys: []string{
					"beat-1",
					"beat-2",
				},
			},

			regenerateBeatsData: &regenerateBeatsData{
				err: dao.ErrBeatsSheetNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrBeatsSheetNotFound.Error()},
		},
		{
			name: "LoglineNotFound",

			form: &codegen.RegenerateBeatsForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				RegenerateKeys: []string{
					"beat-1",
					"beat-2",
				},
			},

			regenerateBeatsData: &regenerateBeatsData{
				err: dao.ErrLoglineNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrLoglineNotFound.Error()},
		},
		{
			name: "StoryPlanNotFound",

			form: &codegen.RegenerateBeatsForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				RegenerateKeys: []string{
					"beat-1",
					"beat-2",
				},
			},

			regenerateBeatsData: &regenerateBeatsData{
				err: dao.ErrStoryPlanNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrStoryPlanNotFound.Error()},
		},
		{
			name: "Error",

			form: &codegen.RegenerateBeatsForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				RegenerateKeys: []string{
					"beat-1",
					"beat-2",
				},
			},

			regenerateBeatsData: &regenerateBeatsData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockRegenerateBeatsService(t)

			ctx := context.WithValue(t.Context(), authapi.ClaimsAPIKey{}, &authmodels.AccessTokenClaims{
				UserID: lo.ToPtr(uuid.MustParse("00000000-1000-0000-0000-000000000001")),
			})

			if testCase.regenerateBeatsData != nil {
				source.EXPECT().
					RegenerateBeats(ctx, services.RegenerateBeatsRequest{
						BeatsSheetID:   uuid.UUID(testCase.form.GetBeatsSheetID()),
						UserID:         uuid.MustParse("00000000-1000-0000-0000-000000000001"),
						RegenerateKeys: testCase.form.GetRegenerateKeys(),
					}).
					Return(testCase.regenerateBeatsData.resp, testCase.regenerateBeatsData.err)
			}

			handler := api.API{RegenerateBeatsService: source}

			res, err := handler.RegenerateBeats(ctx, testCase.form)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
