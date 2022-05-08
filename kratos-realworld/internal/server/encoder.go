package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-realworld/internal/errors"
	nethttp "net/http"
)

func errorEncoder(writer nethttp.ResponseWriter, request *nethttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(request, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Set("Content-Type", "application/"+codec.Name())
	writer.WriteHeader(se.Code)
	writer.Write(body)
}
