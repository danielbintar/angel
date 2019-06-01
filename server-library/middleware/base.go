package middleware

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// use this to use various middleware
// example
// router.POST("/users", middleware.Adapt(handler, middleware.MustHaveForm(&form)))
func Adapt(h httprouter.Handle, adapters ...Adapter) httprouter.Handle {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// middleware
// validate have json body in request
// body will be unmarshal to context "form"
// example to access
// r.Context().Value("form").(*Form)
func MustHaveForm(form interface{}) Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			if r.Body == nil {
				http.Error(w, "form must be filled", http.StatusUnprocessableEntity)
				return
			}

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

type Adapter func(httprouter.Handle) httprouter.Handle
