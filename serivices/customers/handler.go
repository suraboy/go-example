package customers

/**
 * Created by boy.sirichai on 23/8/2022 AD
 */
import (
	"github.com/gofiber/fiber/v2"
)

// Handler type
type Handler struct {
	service Servicer
}

// NewHandler init route
func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) GetCustomers(c *fiber.Ctx) error {
	result, err := h.service.GetSymbols()
	if err != nil {
		return err
	}
	return c.JSON(result)
}