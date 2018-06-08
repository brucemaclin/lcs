package lcs

import (

	"testing"
	"math/rand"
	"encoding/hex"
	"fmt"
)
var randomStrSlice []string
const randomLength = 1000
const randomStrLen = 10
func init() {
	for i:=0;i<randomLength;i++ {
		var b [randomStrLen/2]byte
		rand.Read(b[:])
		randomStrSlice = append(randomStrSlice,hex.EncodeToString(b[:]))
	}
	fmt.Println("random slice length:", len(randomStrSlice))
}
func getRandomStr(index int) (string,string) {
	index = index%randomLength -1
	if index<0 {
		index = 0
	}
	return randomStrSlice[index],randomStrSlice[index+1]
}
func TestLCS(t *testing.T) {

	ret, str := Lcs("hellolcs", "lhcs")
	if ret != 3 || str != "lcs" {
		t.Error(ret, str)
		return
	}
}

func BenchmarkLcs(b *testing.B) {


	for i := 0; i < b.N; i++ {
		raw,target := getRandomStr(i)
		Lcs(raw,target)
	}
}