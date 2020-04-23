package fetching

import (
	"errors"
	"testing"

	"github.com/ManuRua/golab/codelytv_golang_course/refactor_to_cobra/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestFetchByID(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid beer":     {input: 127, want: 127, err: nil},
		"not found beer": {input: 99999, err: errors.New("error")},
	}
	repo := memory.NewRepository()

	service := NewService(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := service.FetchByID(tc.input)
			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.ProductID)
		})
	}
}
