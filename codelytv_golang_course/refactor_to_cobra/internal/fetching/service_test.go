package fetching

import (
	"testing"

	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/storage/memory"
)

func TestFetchByID(t *testing.T) {
	repo := memory.NewRepository()

	service := NewService(repo)

	expected := 127
	b, err := service.FetchByID(expected)
	if err != nil {
		t.Fatalf("expected %d, got an error %v", expected, err)
	}

	if b.ProductID != expected {
		t.Fatalf("expected %d, got: %d", expected, b.BeerID)
	}
}
