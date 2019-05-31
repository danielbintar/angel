package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Adapter func(httprouter.Handle) httprouter.Handle

func Adapt(h httprouter.Handle, adapters ...Adapter) httprouter.Handle {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func MustHaveBody() Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			if len(r.Body) == 0 {
				http.Error(w, "body is required", http.StatusUnprocessableEntity)
				return
			}

			h(w, r, params)
	    }
    }
}
