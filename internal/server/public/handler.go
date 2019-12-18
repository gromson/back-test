package public

import (
	http_ext "back-api/pkg/http"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func AddMessageHandle() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		success := http_ext.NewJsonSuccess(struct{ Hello string }{Hello: "world"})
		success.Respond(w, req)
	}
}
