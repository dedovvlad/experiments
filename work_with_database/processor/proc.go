package processor

import (
	"context"
	"log"
	"math"
	"sync"

	"lerning/work_with_database/models"
)

type (
	service interface {
		GetNumber(ctx context.Context, id string) (models.Passport, error)
		SetNumbersOne(ctx context.Context, series, number string) (bool, error)
		SetNumbersMany(ctx context.Context, series, number string) (bool, error)
		SetNumbersChunk(ctx context.Context, series, number string) (bool, error)
	}
	Proc struct {
		service service
	}
)

type Iter struct {
	Start int
	End   int
}

func NewProc(s service) *Proc {
	return &Proc{
		service: s,
	}
}

func message(count int) {
	log.Printf("total passport numbers in file: %d", count)
	log.Printf("Set passport numbers done")
}

func (p *Proc) AddPassportsOne(filePath string) {
	ctx := context.Background()

	data, count, _ := Parser(filePath, 0, 0)

	for _, v := range data {
		_, err := p.service.SetNumbersOne(ctx, v.Series, v.Number)
		if err != nil {
		}
	}
	message(count)
}

func (p *Proc) AddPassportsMany(filePath string) {
	ctx := context.Background()

	data, count, _ := Parser(filePath, 0, 0)

	for _, v := range data {
		_, err := p.service.SetNumbersMany(ctx, v.Series, v.Number)
		if err != nil {
		}
	}
	message(count)
}

func (p *Proc) AddPassportsChunk(filePath string, chunk int) {

	ctx := context.Background()
	defer ctx.Done()

	data, count, err := Parser(filePath, 0, 0)
	if err != nil {
		log.Print(err)
	}

	var wg sync.WaitGroup

	for _, i := range makeIterationList(chunk, count) {
		wg.Add(1)
		go func(iter *Iter) {
			for _, v := range data[iter.Start:iter.End] {
				_, err = p.service.SetNumbersMany(ctx, v.Series, v.Number)
			}
			wg.Done()
			log.Printf("Set data till: %d untill: %d", iter.Start, iter.End)
		}(i)
	}
	wg.Wait()
	message(count)

}

func makeIterationList(chunk, count int) []*Iter {

	iterArr := make([]*Iter, int(math.Ceil(float64(count)/float64(chunk))))

	start := 0
	for i := 0; i < len(iterArr); i++ {
		iterArr[i] = &Iter{
			Start: start,
			End:   start + chunk,
		}

		if start += chunk; start > count {
			iterArr[i].End = count
		}
	}

	log.Printf("Create %d chunk", len(iterArr))
	return iterArr
}
