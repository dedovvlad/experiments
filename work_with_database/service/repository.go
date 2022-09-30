package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

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

func (s *Service) CreateNumbersOne(ctx context.Context, series, number string) error {
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

func (s *Service) CreateNumbersPrepare(ctx context.Context, stmt *sql.Stmt, series, number string) error {
	id := uuid.New().String()
	_, err := stmt.ExecContext(ctx, id, series, number)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateNumbersChunk(ctx context.Context, params []*models.Passport) error {
	tx, err := s.db.BeginTx(ctx, nil)

	paramsString := make([]string, 0, len(params))
	for _, param := range params {
		paramsString = append(paramsString, fmt.Sprintf("('%s', '%s', '%s')", param.ID, param.Series, param.Number))
	}

	insert := fmt.Sprintf(`INSERT INTO passports (id, series, number) VALUES %s`, strings.Join(paramsString, ","))

	_, err = tx.QueryContext(ctx, insert)
	if err != nil {
		log.Print(err)
	}
	return tx.Commit()
}

func (s *Service) Number(ctx context.Context, id string) (models.Passport, error) {
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
