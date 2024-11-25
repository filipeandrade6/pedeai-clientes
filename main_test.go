package main

import (
	"os"
	"testing"
)

// func TestMain(t *testing.T) {
// 	t.Run("test run", func(t *testing.T) {
// 		if err := run(context.Background(), nil, nil, os.Stdin, os.Stdout, os.Stderr, true); err != nil {
// 			t.Errorf("should have not return any error, got: %s", err)
// 		}
// 	})
// }

func TestMain(m *testing.M) {
	if os.Getenv("CI") == "" {
		m.Run()
	}
}
