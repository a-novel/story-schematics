package services

import (
	"errors"

	"github.com/google/uuid"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/story-schematics/internal/dao"
	"github.com/a-novel/story-schematics/models"
)

var ErrSelectBeatsSheetService = errors.New("SelectBeatsSheetService.SelectBeatsSheet")

func NewErrSelectBeatsSheetService(err error) error {
	return errors.Join(err, ErrSelectBeatsSheetService)
}

type SelectBeatsSheetSource interface {
	SelectBeatsSheet(ctx context.Context, data uuid.UUID) (*dao.BeatsSheetEntity, error)
	SelectLogline(ctx context.Context, data dao.SelectLoglineData) (*dao.LoglineEntity, error)
}

type SelectBeatsSheetRequest struct {
	BeatsSheetID uuid.UUID
	UserID       uuid.UUID
}

type SelectBeatsSheetService struct {
	source SelectBeatsSheetSource
}

func (service *SelectBeatsSheetService) SelectBeatsSheet(
	ctx context.Context, request SelectBeatsSheetRequest,
) (*models.BeatsSheet, error) {
	data, err := service.source.SelectBeatsSheet(ctx, request.BeatsSheetID)
	if err != nil {
		return nil, NewErrSelectBeatsSheetService(err)
	}

	// Make sure the selected beats sheet is linked to a logline that belongs to the user.
	_, err = service.source.SelectLogline(ctx, dao.SelectLoglineData{
		ID:     data.LoglineID,
		UserID: request.UserID,
	})
	if err != nil {
		return nil, NewErrSelectBeatsSheetService(err)
	}

	return &models.BeatsSheet{
		ID:          data.ID,
		LoglineID:   data.LoglineID,
		StoryPlanID: data.StoryPlanID,
		Content:     data.Content,
		CreatedAt:   data.CreatedAt,
	}, nil
}

func NewSelectBeatsSheetServiceSource(
	selectBeatsSheetDAO *dao.SelectBeatsSheetRepository,
	selectLoglineDAO *dao.SelectLoglineRepository,
) SelectBeatsSheetSource {
	return &struct {
		*dao.SelectBeatsSheetRepository
		*dao.SelectLoglineRepository
	}{
		SelectBeatsSheetRepository: selectBeatsSheetDAO,
		SelectLoglineRepository:    selectLoglineDAO,
	}
}

func NewSelectBeatsSheetService(source SelectBeatsSheetSource) *SelectBeatsSheetService {
	return &SelectBeatsSheetService{source: source}
}
