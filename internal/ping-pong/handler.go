package ping_pong

import "github.com/gofiber/fiber/v2"

type PingPongService interface {
	Ping(request PingRequest) (PongResponse, error)
}
type Logger interface {
	Debugf(template string, args ...interface{})
}
type Handler struct {
	logger  Logger
	service PingPongService
}

func NewHandler(logger Logger, service PingPongService) *Handler {
	if logger == nil || service == nil {
		return nil
	}
	return &Handler{logger: logger, service: service}
}
func (h *Handler) Ping(ctx *fiber.Ctx) error {
	h.logger.Debugf("ping request arrived to handler")
	var pingRequest PingRequest
	if err := ctx.BodyParser(&pingRequest); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	pong, err := h.service.Ping(pingRequest)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(pong)
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/ping", h.Ping)
}
