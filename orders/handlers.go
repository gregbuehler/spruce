package orders

import (
	"fmt"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func CreationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func ResendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}
