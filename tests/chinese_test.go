package tests

import (
	"fmt"
	"github.com/qbhy/chinese"
	"testing"
)

func TestFindWords(t *testing.T) {
	fmt.Println(chinese.RelatedWords("我说呢"))
}
