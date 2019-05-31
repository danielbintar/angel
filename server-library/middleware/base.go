package middleware

import (
	"context"
	"encoding/json"
	"io"
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

func MustHaveForm(form interface{}) Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			err := json.NewDecoder(r.Body).Decode(&form)
			switch {
				case err == io.EOF:
					http.Error(w, "form must be filled", http.StatusUnprocessableEntity)
					return
				case err != nil:
					http.Error(w, err.Error(), http.StatusUnprocessableEntity)
					return
			}

			ctx := context.WithValue(r.Context(), "form", form)
			h(w, r.WithContext(ctx), params)
	    }
    }
}
