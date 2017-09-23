package handlers

import (
	"net/http"
	"testing"

	"github.com/lyakhov/microssy/pkg/config"
	"github.com/lyakhov/microssy/pkg/logger"
	"github.com/lyakhov/microssy/pkg/logger/standard"
	"github.com/lyakhov/microssy/pkg/router/bitroute"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
