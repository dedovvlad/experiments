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

func (s *Service) SetNumbersOne(ctx context.Context, series, number string) error {
	id := uuid.New().String()

	err := s.q.insertPassportData(ctx, insertPassportDataParams{
		ID:     id,
		Series: series,
		Number: number,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CallPrepare(ctx context.Context) (*sql.Stmt, error) {
	p, err := Prepare(ctx, s.db)
	if err != nil {
		return nil, err
	}

	return p.insertPassportDataStmt, nil
}

func (s *Service) SetNumbersPrepare(ctx context.Context, stmt *sql.Stmt, series, number string) error {
	id := uuid.New().String()
	_, err := stmt.ExecContext(ctx, id, series, number)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SetNumbersChunk(ctx context.Context, params []*models.Passport) error {
	tx, err := s.db.BeginTx(ctx, nil)
	q := s.q.WithTx(tx)

	for _, v := range params {
		err = q.insertPassportData(ctx, insertPassportDataParams{
			ID:     v.ID,
			Series: v.Series,
			Number: v.Number,
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
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
