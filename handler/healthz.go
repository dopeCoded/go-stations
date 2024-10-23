// handler/healthz.go
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HealthzResponse のインスタンスを作成し、Message に "OK" を設定
	response := &model.HealthzResponse{
		Message: "OK",
	}

	// レスポンスヘッダーに Content-Type を設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// JSON エンコードしてレスポンスに書き込む
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		// エンコードに失敗した場合、エラーログを出力し、500 エラーを返す
		log.Println("JSON エンコードエラー:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
