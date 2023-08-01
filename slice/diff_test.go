package slice

import (
	"log"
	"testing"
)

func TestMax(t *testing.T) {

}

func TestSliceMax(t *testing.T) {
	mockSlice := []int{1, 5, 8, 9, 111111, 2345}
	max, err := Max(mockSlice)
	if err != nil {
		return
	}
	log.Default().Printf("max res = %v", max)
}
