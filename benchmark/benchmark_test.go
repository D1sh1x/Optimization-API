package benchmark

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"optimization/handlers"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkAddHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/add", handlers.AddHandler)

	body, _ := json.Marshal(map[string]float64{
		"a": 12.3,
		"b": 67.1,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/add", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
}
