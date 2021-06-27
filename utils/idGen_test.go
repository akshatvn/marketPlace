package utils

import (
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	allIds := make(map[string]bool)
	for i :=0;i<100000;i++ {
		s := GenerateID()
		if _, ok := allIds[s]; ok {
			t.Errorf("id collided")
		}
		allIds[s] = true
		time.Sleep(time.Microsecond) // granularity of ID
	}
}
