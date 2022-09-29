package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"lerning/work_with_database/models"
)

type Service struct {
	db *sql.DB
	q  *Queries
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
		q:  New(db),
	}
}

func (s *Service) SetNumbersOne(ctx context.Context, series, number string) (bool, error) {

	id := uuid.New().String()

	_, err := s.q.insertPassportData(ctx, insertPassportDataParams{
		ID:     id,
		Series: series,
		Number: number,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) SetNumbersMany(ctx context.Context, series, number string) (bool, error) {

	id := uuid.New().String()

	tx, err := s.db.BeginTx(ctx, nil)

	q := s.q.WithTx(tx)

	_, err = q.insertPassportData(ctx, insertPassportDataParams{
		ID:     id,
		Series: series,
		Number: number,
	})
	if err != nil {
		return false, err
	}
	return true, tx.Commit()
}

func (s *Service) SetNumbersChunk(ctx context.Context, series, number string) (bool, error) {

	id := uuid.New().String()

	tx, err := s.db.BeginTx(ctx, nil)

	q := s.q.WithTx(tx)

	_, err = q.insertPassportData(ctx, insertPassportDataParams{
		ID:     id,
		Series: series,
		Number: number,
	})
	if err != nil {
		return false, err
	}
	return true, tx.Commit()
}

func (s *Service) GetNumber(ctx context.Context, id string) (models.Passport, error) {
	dto, err := s.q.getPassportData(ctx, id)
	if err != nil {
		return models.Passport{}, err
	}
	return models.Passport{
		ID:     dto.ID,
		Series: dto.Series,
		Number: dto.Number,
	}, nil
}
