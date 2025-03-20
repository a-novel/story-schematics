package api

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"

	authapi "github.com/a-novel/authentication/api"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/story-schematics/api/codegen"
	"github.com/a-novel/story-schematics/internal/dao"
	"github.com/a-novel/story-schematics/internal/services"
	"github.com/a-novel/story-schematics/models"
)

type GenerateBeatsSheetService interface {
	GenerateBeatsSheet(ctx context.Context, request services.GenerateBeatsSheetRequest) ([]models.Beat, error)
}

func (api *API) GenerateBeatsSheet(
	ctx context.Context, req *codegen.GenerateBeatsSheetForm,
) (codegen.GenerateBeatsSheetRes, error) {
	userID, err := authapi.RequireUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user ID: %w", err)
	}

	beatsSheet, err := api.GenerateBeatsSheetService.GenerateBeatsSheet(ctx, services.GenerateBeatsSheetRequest{
		LoglineID:   uuid.UUID(req.GetLoglineID()),
		StoryPlanID: uuid.UUID(req.GetStoryPlanID()),
		UserID:      userID,
	})

	switch {
	case errors.Is(err, dao.ErrLoglineNotFound), errors.Is(err, dao.ErrStoryPlanNotFound):
		return &codegen.NotFoundError{Error: err.Error()}, nil
	case err != nil:
		return nil, fmt.Errorf("generate beats sheet: %w", err)
	}

	return &codegen.BeatsSheet{
		Content: lo.Map(beatsSheet, func(item models.Beat, _ int) codegen.Beat {
			return codegen.Beat{
				Key:     item.Key,
				Title:   item.Title,
				Content: item.Content,
			}
		}),
	}, nil
}
