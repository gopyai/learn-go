package mdw

import (
	"context"
	"net/http"

	"devx/chainer"

	"github.com/julienschmidt/httprouter"
)

type (
	HandlerId struct {
		Token string
		Email string
	}
)

func Login(h chainer.Handler) chainer.Handler {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Check cookie
		c, err := r.Cookie("token")
		if err != nil {
			// No cookie
			// Signup to create a new one
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		// Validate token
		email, err := verifyAndReadToken(c.Value)
		if err != nil {
			// Invalid token
			// Signup to create a new one
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		// Pass the id via context for handler to use
		ctx := context.WithValue(r.Context(), "id", HandlerId{c.Value, email})

		// Run the handler
		h(w, r.WithContext(ctx), p)
	}
}

func ContextGetId(r *http.Request) (hId HandlerId) {
	if value := r.Context().Value("id"); value != nil {
		hId = value.(HandlerId)
	}
	return
}
