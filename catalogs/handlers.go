package catalogs

import (
	"fmt"
	"net/http"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}
