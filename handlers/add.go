package handlers

import (
	"bytes"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func AddHandler(c *gin.Context) {
	var req AddRequest
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	result := req.A + req.B
	jsonBuf := append([]byte(`{"result":`), strconv.FormatFloat(result, 'f', -1, 64)...)
	jsonBuf = append(jsonBuf, '}')
	c.Data(http.StatusOK, "application/json", jsonBuf)
}
