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
	"github.com/a-novel/story-schematics/internal/daoai"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

func TestExpandBeat(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type expandBeatData struct {
		resp *models.Beat
		err  error
	}

	testCases := []struct {
		name string

		form *codegen.ExpandBeatForm

		expandBeatData *expandBeatData

		expect    codegen.ExpandBeatRes
		expectErr error
	}{
		{
			name: "Success",

			form: &codegen.ExpandBeatForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TargetKey:    "beat-1",
			},

			expandBeatData: &expandBeatData{
				resp: &models.Beat{
					Key:     "beat-1",
					Title:   "Beat 1 expanded",
					Content: "Beat 1 content expanded",
				},
			},

			expect: &codegen.Beat{
				Key:     "beat-1",
				Title:   "Beat 1 expanded",
				Content: "Beat 1 content expanded",
			},
		},
		{
			name: "BeatsSheetNotFound",

			form: &codegen.ExpandBeatForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TargetKey:    "beat-1",
			},

			expandBeatData: &expandBeatData{
				err: dao.ErrBeatsSheetNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrBeatsSheetNotFound.Error()},
		},
		{
			name: "StoryPlanNotFound",

			form: &codegen.ExpandBeatForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TargetKey:    "beat-1",
			},

			expandBeatData: &expandBeatData{
				err: dao.ErrStoryPlanNotFound,
			},

			expect: &codegen.NotFoundError{Error: dao.ErrStoryPlanNotFound.Error()},
		},
		{
			name: "UnknownTargetKey",

			form: &codegen.ExpandBeatForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TargetKey:    "beat-1",
			},

			expandBeatData: &expandBeatData{
				err: daoai.ErrUnknownTargetKey,
			},

			expect: &codegen.UnprocessableEntityError{Error: daoai.ErrUnknownTargetKey.Error()},
		},
		{
			name: "Error",

			form: &codegen.ExpandBeatForm{
				BeatsSheetID: codegen.BeatsSheetID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TargetKey:    "beat-1",
			},

			expandBeatData: &expandBeatData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockExpandBeatService(t)

			ctx := context.WithValue(t.Context(), authapi.ClaimsAPIKey{}, &authmodels.AccessTokenClaims{
				UserID: lo.ToPtr(uuid.MustParse("00000000-1000-0000-0000-000000000001")),
			})

			if testCase.expandBeatData != nil {
				source.EXPECT().
					ExpandBeat(ctx, services.ExpandBeatRequest{
						BeatsSheetID: uuid.UUID(testCase.form.GetBeatsSheetID()),
						TargetKey:    testCase.form.GetTargetKey(),
						UserID:       uuid.MustParse("00000000-1000-0000-0000-000000000001"),
					}).
					Return(testCase.expandBeatData.resp, testCase.expandBeatData.err)
			}

			handler := api.API{ExpandBeatService: source}

			res, err := handler.ExpandBeat(ctx, testCase.form)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
