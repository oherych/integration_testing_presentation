package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestService_GetLogoHandlerIntegration(t *testing.T) {
	e, done := createTestEnvExpect(t)
	defer done()

	p := e.GET("/logo").Expect()

	p.Status(http.StatusOK)
	p.Header("Content-Length").Equal("16567")
	p.Header("Content-Disposition").Equal("attachment; filename=logo.png")
	p.ContentType("image/png")

	hash := md5.Sum([]byte(p.Body().Raw()))
	assert.Equal(t, "716a9a93c3a28db4dead8e36d3046035", hex.EncodeToString(hash[:]))
}
