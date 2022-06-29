package randomnumber

import (
	"fmt"
	"net/http"
)

type handler struct {
	impl Interface
}

func NewHandler(impl Interface) http.Handler {
	return &handler{impl: impl}
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	resp, err := h.impl.GetRandomNumber()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "%d", resp)
}
