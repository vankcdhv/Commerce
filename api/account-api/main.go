package account_api

import "net/http"

type accountHandler struct {
	
}

func (h *accountHandler) handler() (response http.ResponseWriter, request *http.Request) {
	response.Write(http.StatusOK,)
	return
}
