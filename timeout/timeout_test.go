package timeout

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kamushadenes/chloe/config"
)

func TestWaitTimeout(t *testing.T) {
	timeout := time.Millisecond * 500
	ctx := context.Background()

	t.Run("Timeout", func(t *testing.T) {
		fn := func(ch chan interface{}, errCh chan error) {
			<-time.After(time.Second)
		}

		_, err := WaitTimeout(ctx, timeout, fn)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Function returns value", func(t *testing.T) {
		fn := func(ch chan interface{}, errCh chan error) {
			ch <- "value"
		}

		result, err := WaitTimeout(ctx, timeout, fn)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if result != "value" {
			t.Errorf("Expected result to be 'value', got %v", result)
		}
	})

	t.Run("Function returns error", func(t *testing.T) {
		fn := func(ch chan interface{}, errCh chan error) {
			errCh <- errors.New("test error")
		}

		_, err := WaitTimeout(ctx, timeout, fn)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Slowness warning", func(t *testing.T) {
		warning := config.Timeouts.SlownessWarning
		fn := func(ch chan interface{}, errCh chan error) {
			<-time.After(warning * 2)
		}

		_, err := WaitTimeout(ctx, timeout, fn)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

}
