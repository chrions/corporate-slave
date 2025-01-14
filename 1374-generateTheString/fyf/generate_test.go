package fyf

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	ret := generateTheString(5)
	log.Println(ret)
}
