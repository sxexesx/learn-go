package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestLeakyFunc(t *testing.T) {
	leakyFunc()
}

func TestLeakyChannels(t *testing.T) {
	leakyChannels()
}
