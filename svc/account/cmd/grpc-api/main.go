package main

import (
	"flag"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/reyhanfahlevi/pkg/go/log"
	"github.com/reyhanfahlevi/poc-test/pkg/config"
	"github.com/reyhanfahlevi/poc-test/svc/account/app/grpc"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/repo"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/usecase"
)

var (
	errLogPath   string
	infoLogPath  string
	debugLogPath string
)

func main() {
	flag.StringVar(&errLogPath, "error_log", "", "error log")
	flag.StringVar(&infoLogPath, "info_log", "", "info log")
	flag.StringVar(&debugLogPath, "debug_log", "", "debug log")
	flag.Parse()

	setLog(log.ErrorLevel, errLogPath)
	setLog(log.InfoLevel, infoLogPath)
	setLog(log.DebugLevel, debugLogPath)

	err := config.Init(config.WithConfigFile("files/etc/account/config.json", "/etc/account/config.json"))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Get()

	db, err := sqlx.Open("postgres", cfg.DB.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	r := repo.New(db)
	uc := usecase.New(cfg.App.Salt, r)

	srv := grpc.New("", uc)
	err = srv.Run(cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func setLog(level log.Level, filePath string) {
	lgr, err := log.NewLogger(&log.Config{
		Level:   level,
		LogFile: filePath,
		Caller:  true,
		AppName: "Account Service",
		UseJSON: os.Getenv("APP_ENV") == "production",
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = log.SetLogger(level, lgr)
	if err != nil {
		log.Fatalln(err)
	}
}
