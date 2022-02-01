package main

import (
	"testing"
	"time"
)

func TestSingleInstance(t *testing.T) {
	SingleInstance()
	time.Sleep(1*time.Second)
	t.Log("ok")
}