package utils

import (
	"testing"
)

func TestScape(t *testing.T) {

}

func TestScapeSlice(t *testing.T) {
	scape := ScapeSlice([]string{"<a>hello</a>", "< >"})
	t.Log(scape)
}
