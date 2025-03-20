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
	"github.com/a-novel/story-schematics/internal/dao"
	"github.com/a-novel/story-schematics/internal/lib"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

func TestCreateBeatsSheet(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type createBeatsSheetData struct {
		resp *models.BeatsSheet
		err  error
	}

	testCases := []struct {
		name string

		form *codegen.CreateBeatsSheetForm

		createBeatsSheetData *createBeatsSheetData

		expect    codegen.CreateBeatsSheetRes
		expectErr error
	}{
		{
			name: "Success",

			form: &codegen.CreateBeatsSheetForm{
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
			},

			createBeatsSheetData: &createBeatsSheetData{
				resp: &models.BeatsSheet{
					ID:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					LoglineID:   uuid.MustParse("00000000-0000-0000-1000-000000000001"),
					StoryPlanID: uuid.MustParse("00000000-0000-1000-0000-000000000001"),
					Content: []models.Beat{
						{
							Key:     "beat-1",
							Title:   "Beat 1",
							Content: "Beat 1 content",
						},
						{
							Key:     "beat-2",
							Title:   "Beat 2",
							Content: "Beat 2 content",
						},
					},
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},

			expect: &codegen.BeatsSheet{
				ID:          codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Error/LoglineNotFound",

			form: &codegen.CreateBeatsSheetForm{
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
			},

			createBeatsSheetData: &createBeatsSheetData{
				err: dao.ErrLoglineNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrLoglineNotFound.Error()},
		},
		{
			name: "Error/StoryPlanNotFound",

			form: &codegen.CreateBeatsSheetForm{
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
			},

			createBeatsSheetData: &createBeatsSheetData{
				err: dao.ErrStoryPlanNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrStoryPlanNotFound.Error()},
		},
		{
			name: "Error/InvalidStoryPlan",

			form: &codegen.CreateBeatsSheetForm{
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
			},

			createBeatsSheetData: &createBeatsSheetData{
				err: lib.ErrInvalidStoryPlan,
			},

			expect: &codegen.UnprocessableEntityError{Error: lib.ErrInvalidStoryPlan.Error()},
		},
		{
			name: "Error/CreateBeatsSheet",

			form: &codegen.CreateBeatsSheetForm{
				LoglineID:   codegen.LoglineID(uuid.MustParse("00000000-0000-0000-1000-000000000001")),
				StoryPlanID: codegen.StoryPlanID(uuid.MustParse("00000000-0000-1000-0000-000000000001")),
				Content: []codegen.Beat{
					{
						Key:     "beat-1",
						Title:   "Beat 1",
						Content: "Beat 1 content",
					},
					{
						Key:     "beat-2",
						Title:   "Beat 2",
						Content: "Beat 2 content",
					},
				},
			},

			createBeatsSheetData: &createBeatsSheetData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockCreateBeatsSheetService(t)

			ctx := context.WithValue(t.Context(), authapi.ClaimsAPIKey{}, &authmodels.AccessTokenClaims{
				UserID: lo.ToPtr(uuid.MustParse("00000000-1000-0000-0000-000000000001")),
			})

			if testCase.createBeatsSheetData != nil {
				source.EXPECT().
					CreateBeatsSheet(ctx, services.CreateBeatsSheetRequest{
						LoglineID:   uuid.UUID(testCase.form.GetLoglineID()),
						UserID:      uuid.MustParse("00000000-1000-0000-0000-000000000001"),
						StoryPlanID: uuid.UUID(testCase.form.GetStoryPlanID()),
						Content: lo.Map(testCase.form.GetContent(), func(item codegen.Beat, _ int) models.Beat {
							return models.Beat{
								Key:     item.GetKey(),
								Title:   item.GetTitle(),
								Content: item.GetContent(),
							}
						}),
					}).
					Return(testCase.createBeatsSheetData.resp, testCase.createBeatsSheetData.err)
			}

			handler := api.API{CreateBeatsSheetService: source}

			res, err := handler.CreateBeatsSheet(ctx, testCase.form)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
