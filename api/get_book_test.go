package api

import (
	"net/http"
	"testing"
)

func TestGetBookIntegration(t *testing.T) {
	table := map[string]struct {
		BookID     string
		ExpCode    int
		ExpContent interface{}
	}{
		"correct": {
			BookID:  testBookID1,
			ExpCode: http.StatusOK,
			ExpContent: map[string]string{
				"book_id": testBookID1,
				"name":    "book name 1",
			},
		},
		//TODO: check wrong BookID
	}

	for name, in := range table {
		t.Run(name, func(t *testing.T) {
			e, done := createTestEnvExpect(t)
			defer done()

			p := e.GET("/book/{book_id}", in.BookID).Expect()

			p.Status(in.ExpCode)
			if in.ExpCode == http.StatusOK {
				p.JSON().Equal(in.ExpContent)
			}
		})
	}
}

func BenchmarkGetBookIntegration(b *testing.B) {
	e, done := createTestEnvExpect(b)
	defer done()

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		e.GET("/book/{book_id}", testBookID1).Expect()
	}
}
