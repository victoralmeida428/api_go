package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	// Dados de exemplo para o teste
	data := map[string]string{"message": "Hello, World!"}
	status := http.StatusOK
	headers := http.Header{
		"X-Custom-Header": []string{"TestValue"},
	}

	// Cria um ResponseRecorder para capturar a resposta
	rec := httptest.NewRecorder()

	// Chama a função WriteJSON
	err := WriteJSON(rec, data, status, headers)
	if err != nil {
		t.Fatalf("WriteJSON returned an error: %v", err)
	}

	// Verifica o código de status
	if rec.Code != status {
		t.Errorf("Expected status %d, got %d", status, rec.Code)
	}

	// Verifica o conteúdo do cabeçalho
	for key, value := range headers {
		if rec.Header().Get(key) != value[0] {
			t.Errorf("Expected header %s to be %s, got %s", key, value[0], rec.Header().Get(key))
		}
	}

	// Verifica o cabeçalho Content-Type
	expectedContentType := "application/json; charset=utf-8"
	if rec.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, rec.Header().Get("Content-Type"))
	}

	// Verifica o corpo da resposta JSON
	var responseBody map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if responseBody["message"] != data["message"] {
		t.Errorf("Expected response body %v, got %v", data, responseBody)
	}
}
