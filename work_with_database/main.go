package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"lerning/work_with_database/processor"
	"lerning/work_with_database/service"
)

func main() {

	cfg := InitConfig()

	db, err := sql.Open("postgres", cfg.DB)
	if err != nil {
		log.Fatalf("init db error: %s", err.Error())
	}

	proc := service.NewService(db)
	action := processor.NewProc(proc)

	start := time.Now()
	action.AddPassportsOne(cfg.FilePath)
	log.Printf("<<< [one] >>> Total time: %f", time.Since(start).Seconds())

	start = time.Now()
	action.AddPassportsPrepare(cfg.FilePath)
	log.Printf("<<< [prepare] >>> Total time: %f", time.Since(start).Seconds())

	start = time.Now()
	action.AddPassportsChunk(cfg.FilePath, 100)
	log.Printf("<<< [chunk] >>> Total time: %f", time.Since(start).Seconds())
}
