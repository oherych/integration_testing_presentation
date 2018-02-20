package api

import (
	"net/http"
	"sync"
	"testing"
)

func TestCreateBookIntegration(t *testing.T) {
	table := map[string]struct {
		BookName   string
		ExpCode    int
		ExpContent map[string]string
	}{
		"correct": {
			BookName: "new book",
			ExpCode:  http.StatusOK,
			ExpContent: map[string]string{
				"status": "ok",
			},
		},
	}

	for name, in := range table {
		t.Run(name, func(t *testing.T) {
			e, done := createTestEnvExpect(t)
			defer done()

			p := e.POST("/book").
				WithFormField("name", in.BookName).
				Expect()

			p.Status(in.ExpCode)
			if in.ExpCode == http.StatusOK {
				p.JSON().Equal(in.ExpContent)
			}
		})
	}
}

func TestCreateBookLoading(t *testing.T) {
	const requests = 10
	var wg sync.WaitGroup

	e, done := createTestEnvExpect(t)
	defer done()

	wg.Add(requests)
	for j := 0; j < requests; j++ {
		go func() {
			defer wg.Done()

			e.POST("/book").
				WithFormField("name", "test book").
				Expect()
		}()
	}

	wg.Wait()
}
