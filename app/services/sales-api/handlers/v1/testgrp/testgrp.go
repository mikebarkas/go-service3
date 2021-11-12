// Package testgrp maintains the group of handlers for health checking.
package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/mikebarkas/service3/business/sys/validate"
	"github.com/mikebarkas/service3/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of check enpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler for development.
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		// return 500 and hide error message
		// return errors.New("untrusted data")

		// return 400 and display error message
		return validate.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)

		// k8s to restart system
		// return web.NewShutdownError("restarting service")

		// Test panic recover() middleware
		// panic("testing panic")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
