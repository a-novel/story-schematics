package api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	"github.com/a-novel/story-schematics/api"
	"github.com/a-novel/story-schematics/api/codegen"
	apimocks "github.com/a-novel/story-schematics/api/mocks"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

func TestCreateStoryPlan(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type createStoryPlanData struct {
		resp *models.StoryPlan
		err  error
	}

	testCases := []struct {
		name string

		form *codegen.CreateStoryPlanForm

		createStoryPlanData *createStoryPlanData

		expect    codegen.CreateStoryPlanRes
		expectErr error
	}{
		{
			name: "Success",

			form: &codegen.CreateStoryPlanForm{
				Slug:        "slug",
				Name:        "name",
				Description: "description",
				Beats: []codegen.BeatDefinition{
					{
						Name:      "Beat 1",
						Key:       "beat-1",
						KeyPoints: []string{"beat 1 - key point 1", "beat 1 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 2,
					},
					{
						Name:      "Beat 2",
						Key:       "beat-2",
						KeyPoints: []string{"beat 2 - key point 1", "beat 2 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 1,
					},
				},
			},

			createStoryPlanData: &createStoryPlanData{
				resp: &models.StoryPlan{
					ID:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Slug:        "slug",
					Name:        "name",
					Description: "description",
					Beats: []models.BeatDefinition{
						{
							Name:      "Beat 1",
							Key:       "beat-1",
							KeyPoints: []string{"beat 1 - key point 1", "beat 1 - key point 2"},
							Purpose:   "purpose",
							MinScenes: 1,
							MaxScenes: 2,
						},
						{
							Name:      "Beat 2",
							Key:       "beat-2",
							KeyPoints: []string{"beat 2 - key point 1", "beat 2 - key point 2"},
							Purpose:   "purpose",
							MinScenes: 1,
							MaxScenes: 1,
						},
					},
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},

			expect: &codegen.StoryPlan{
				ID:          codegen.StoryPlanID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				Slug:        "slug",
				Name:        "name",
				Description: "description",
				Beats: []codegen.BeatDefinition{
					{
						Name:      "Beat 1",
						Key:       "beat-1",
						KeyPoints: []string{"beat 1 - key point 1", "beat 1 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 2,
					},
					{
						Name:      "Beat 2",
						Key:       "beat-2",
						KeyPoints: []string{"beat 2 - key point 1", "beat 2 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 1,
					},
				},
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Error",

			form: &codegen.CreateStoryPlanForm{
				Slug:        "slug",
				Name:        "name",
				Description: "description",
				Beats: []codegen.BeatDefinition{
					{
						Name:      "Beat 1",
						Key:       "beat-1",
						KeyPoints: []string{"beat 1 - key point 1", "beat 1 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 2,
					},
					{
						Name:      "Beat 2",
						Key:       "beat-2",
						KeyPoints: []string{"beat 2 - key point 1", "beat 2 - key point 2"},
						Purpose:   "purpose",
						MinScenes: 1,
						MaxScenes: 1,
					},
				},
			},

			createStoryPlanData: &createStoryPlanData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockCreateStoryPlanService(t)

			if testCase.createStoryPlanData != nil {
				source.EXPECT().
					CreateStoryPlan(t.Context(), services.CreateStoryPlanRequest{
						Slug:        models.Slug(testCase.form.GetSlug()),
						Name:        testCase.form.GetName(),
						Description: testCase.form.GetDescription(),
						Beats: lo.Map(
							testCase.form.GetBeats(),
							func(item codegen.BeatDefinition, _ int) models.BeatDefinition {
								return models.BeatDefinition{
									Name:      item.GetName(),
									Key:       item.GetKey(),
									KeyPoints: item.GetKeyPoints(),
									Purpose:   item.GetPurpose(),
									MinScenes: item.GetMinScenes(),
									MaxScenes: item.GetMaxScenes(),
								}
							},
						),
					}).
					Return(testCase.createStoryPlanData.resp, testCase.createStoryPlanData.err)
			}

			handler := api.API{CreateStoryPlanService: source}

			res, err := handler.CreateStoryPlan(t.Context(), testCase.form)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
