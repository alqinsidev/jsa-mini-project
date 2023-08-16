package utils

import (
	"alqinsidev/jsa-mini-project/aduan/config"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
)

type Conn struct {
	PGsql *sql.DB
}

func NewDBConnection(cfg *config.Config) *Conn {
	return &Conn{
		PGsql: InitPGsql(cfg),
	}
}

func InitPGsql(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DB.DSN)

	if err != nil {
		log.Error().Err(err).Msg("Fail to init DB")
	}

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Fail to ping DB")
	}

	log.Info().Msg("Connected to DB")
	return db
}
