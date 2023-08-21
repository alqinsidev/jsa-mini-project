package http

import (
	"alqinsidev/jsa-mini-project/aduan/domain"
	"alqinsidev/jsa-mini-project/aduan/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AduanHandler struct {
	AUsecase domain.AduanUsecase
}

func NewAduanHandler(r *fiber.App, AUsecase domain.AduanUsecase) {
	handler := &AduanHandler{AUsecase: AUsecase}

	r.Get("/aduan", handler.Fetch)
	r.Get("/aduan/summary", handler.FetchSummary)
	r.Get("/aduan/:id", handler.GetById)
	r.Put("/aduan", handler.UpdateStatus)
}

func (h *AduanHandler) Fetch(c *fiber.Ctx) error {
	reqQuery := helpers.GetRequestQuery(c)
	result, totalData, err := h.AUsecase.Fetch(reqQuery)
	if err != nil {
		log.Error().Err(err).Msg("err")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    helpers.PaginateResponse(reqQuery, result, totalData),
		"code":    http.StatusOK,
		"message": "success",
	})
}

func (h *AduanHandler) GetById(c *fiber.Ctx) error {
	aduanId := c.Params("id")
	parsedUUID, err := uuid.Parse(aduanId)
	if err != nil {
		log.Error().Err(err).Msg("error parsing uuid in aduan Get By ID")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "message": "fail", "code": http.StatusBadRequest})
	}

	result, err := h.AUsecase.FindById(parsedUUID)

	if err != nil {
		log.Error().Err(err)

		switch err {
		case domain.ErrAduanNotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error(), "message": "fail", "code": http.StatusNotFound})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data":    helpers.MapRawAduanDetail(result),
	})
}
func (h *AduanHandler) FetchSummary(c *fiber.Ctx) error {
	result, err := h.AUsecase.FetchSummary()
	if err != nil {
		log.Error().Err(err).Msg("error handler summary")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data":    result,
	})
}
func (h *AduanHandler) UpdateStatus(c *fiber.Ctx) error {
	update := new(domain.UpdateStatusPayload)
	if err := c.BodyParser(update); err != nil {
		log.Error().Err(err).Msg("error parsing update status")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	result, err := h.AUsecase.UpdateStatus(update)
	if err != nil {
		switch err {
		case domain.ErrAduanNotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error(), "message": "fail", "code": http.StatusNotFound})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data":    result,
	})
}
