package fyf

import (
	"log"
	"testing"
)

//
//["Skiplist","add","add","add","add","search","erase","search","search","search"]
//[[],        [0],[5],[2],[1],[0],[5],[2],[3],[2]]

func TestSkipList(t *testing.T) {
	sk := Constructor()
	sk.Add(0)
	sk.Add(5)
	sk.Add(2)
	sk.Add(1)
	log.Println(sk.Search(0))
	log.Println(sk.Erase(5))
	log.Println(sk.Search(2))
	log.Println(sk.Search(3))
	log.Println(sk.Search(2))
}
