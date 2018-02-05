package api

import (
	"net/http"
	"testing"
)

func TestGetBook(t *testing.T) {
	table := map[string]struct {
		BookID     string
		ExpCode    int
		ExpContent map[string]string
	}{
		"correct": {
			BookID:  "913d5f4e-5759-455d-83fe-72939b3ddf3a",
			ExpCode: http.StatusOK,
			ExpContent: map[string]string{
				"book_id": "913d5f4e-5759-455d-83fe-72939b3ddf3a",
				"name":    "book name 1",
			},
		},
		//TODO: check empty BookID
		//TODO: check wrong BookID
	}

	for name, in := range table {
		t.Run(name, func(t *testing.T) {
			e := createTestEnvExpect(t)
			p := e.GET("/book/{book_id}", in.BookID).Expect()

			p.Status(in.ExpCode)
			if in.ExpCode == http.StatusOK {
				p.JSON().Equal(in.ExpContent)
			}
		})
	}
}
