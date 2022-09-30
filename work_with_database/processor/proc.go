package processor

import (
	"context"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/google/uuid"

	"lerning/work_with_database/models"
)

type (
	service interface {
		Number(ctx context.Context, id string) (models.Passport, error)
		CreateNumbersOne(ctx context.Context, series, number string) error
		CallPrepare(ctx context.Context) (*sql.Stmt, error)
		CreateNumbersPrepare(ctx context.Context, stmt *sql.Stmt, series, number string) error
		CreateNumbersChunk(ctx context.Context, params []*models.Passport) error
	}
	Proc struct {
		service service
	}
)

const (
	series = iota
	number
)

func NewProc(s service) *Proc {
	return &Proc{
		service: s,
	}
}

func (p *Proc) AddPassportsOne(filePath string) {

	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	rider := csv.NewReader(file)
	log.Printf("Set passport numbers start")
	var count int
	for {
		rows, err := rider.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
		}

		err = p.service.CreateNumbersOne(ctx, rows[series], rows[number])
		if err != nil {
			log.Print(err)
		}
		count++
	}

	log.Printf("Total: %d", count)
	log.Print("Set passport numbers done")
}

func (p *Proc) AddPassportsPrepare(filePath string) {

	ctx := context.Background()

	stmt, err := p.service.CallPrepare(ctx)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	rider := csv.NewReader(file)
	log.Printf("Set passport numbers start")
	var count int
	for {
		rows, err := rider.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
		}

		err = p.service.CreateNumbersPrepare(ctx, stmt, rows[series], rows[number])
		if err != nil {
			log.Print(err)
		}
		count++
	}

	log.Printf("Total: %d", count)
	log.Print("Set passport numbers done")
}

func (p *Proc) AddPassportsChunk(filePath string, volume int) {
	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	rider := csv.NewReader(file)
	buffer := make([]*models.Passport, 0, volume)

	var count, total int

	insertToDB := func(rows []*models.Passport) {
		err := p.service.CreateNumbersChunk(ctx, rows)
		if err != nil {
			log.Print(err)
		}
		log.Printf("Added %d passports number", total)
	}

	log.Printf("Set passport numbers start")

	for {
		rows, err := rider.Read()
		if err == io.EOF {
			insertToDB(buffer)
			break
		}
		if err != nil {
			log.Print(err)
		}
		if count == volume {
			insertToDB(buffer)
			buffer = buffer[:0]
			count = 0
		}
		buffer = append(buffer, &models.Passport{
			ID:     uuid.New().String(),
			Series: rows[series],
			Number: rows[number],
		})
		count++
		total++
	}

	log.Print("Set passport numbers done")
}
