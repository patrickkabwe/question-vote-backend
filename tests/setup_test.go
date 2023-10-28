package tests

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before tests")
	// setup

	m.Run()
	fmt.Println("After tests")
	// teardown
}
