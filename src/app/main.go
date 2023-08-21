package main

import (
	"alqinsidev/jsa-mini-project/aduan/config"
	"alqinsidev/jsa-mini-project/aduan/modules/aduan/delivery/http"
	"alqinsidev/jsa-mini-project/aduan/modules/aduan/repository/postgres"
	"alqinsidev/jsa-mini-project/aduan/modules/aduan/usecase"
	"alqinsidev/jsa-mini-project/aduan/utils"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	cfg := config.EnvConfig()
	db := utils.NewDBConnection(cfg)

	app := fiber.New()

	SetupAduanHandler(app, db.PGsql)

	appPort := fmt.Sprintf("0.0.0.0:%s", viper.GetString("APP_PORT"))
	err := app.Listen(appPort)
	if err != nil {
		log.Error().Err(err).Msg("Cannot start App")
	}

}

func SetupAduanHandler(r *fiber.App, db *sql.DB) {
	aduanRepository := postgres.NewAduanRepository(db)
	aduanUsecase := usecase.NewAduanUsecase(aduanRepository)
	http.NewAduanHandler(r, aduanUsecase)
}
