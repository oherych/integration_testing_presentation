package api

import (
	"net/http"
	"sync"
	"testing"
)

func TestCreateBookIntegration(t *testing.T) {
	table := map[string]struct {
		BookID     string
		BookName   string
		ExpCode    int
		ExpContent interface{}
	}{
		"correct": {
			BookID:   newCorrectBookID,
			BookName: "new book",
			ExpCode:  http.StatusOK,
			ExpContent: map[string]string{
				"status": "ok",
			},
		},
		"update": {
			BookID:   testBookID1,
			BookName: "updated name",
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

			p := e.POST("/book/{book_id}", in.BookID).
				WithFormField("name", in.BookName).
				Expect()

			p.Status(in.ExpCode)
			p.JSON().Equal(in.ExpContent)

		})
	}
}

func TestCreateBookLoading(t *testing.T) {
	const requests = 20
	var wg sync.WaitGroup

	e, done := createTestEnvExpect(t)
	defer done()

	wg.Add(requests)
	for j := 0; j < requests; j++ {
		// TODO: run parallel
		func() {
			defer wg.Done()

			p := e.POST("/book/{book_id}", newCorrectBookID).
				WithFormField("name", "test book").
				Expect()

			p.Status(http.StatusOK)
		}()
	}

	wg.Wait()
}
