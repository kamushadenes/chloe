package utils

import (
	"github.com/kamushadenes/chloe/utils"
	"testing"
	"time"

	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
)

func TestStartAndWait(t *testing.T) {
	timeout := time.After(1 * time.Second)

	req := structs.CompletionRequest{}

	// Start Channel not nil
	startChannel := make(chan bool)
	req.StartChannel = startChannel
	continueChannel := make(chan bool)
	req.ContinueChannel = continueChannel
	go func() {
		<-startChannel
		continueChannel <- true
	}()

	done := make(chan bool)
	go func() {
		StartAndWait(&req)
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("timeout")
	case <-done:
	}
}

func TestWriteResult(t *testing.T) {
	req := structs.CompletionRequest{}
	result := "some result"

	// Result Channel not nil
	resultChannel := make(chan interface{})
	req.ResultChannel = resultChannel
	go func() {
		WriteResult(&req, result)
	}()
	assert.Equal(t, result, <-resultChannel)
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "abcd", utils.Truncate("abcdefg", 4))
	assert.Equal(t, "abcdefgh", utils.Truncate("abcdefgh", 10))
}
