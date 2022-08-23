package customers

/**
 * Created by boy.sirichai on 23/8/2022 AD
 */
import (
	"net/http"
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


func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}