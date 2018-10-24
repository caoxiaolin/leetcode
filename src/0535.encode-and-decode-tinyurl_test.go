package leetcode

import (
	"testing"
)

func TestTinyURL(t *testing.T) {
	obj := new(Codec)
	for i := 0; i < 1000; i++ {
		obj.Encode("http://www.baidu.com")
	}
	expect := "http://www.google.com"
	shortUrl := obj.Encode(expect)
	result := obj.Decode(shortUrl)
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
