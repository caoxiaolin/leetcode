package leetcode

import (
	"testing"
)

func TestTinyURL(t *testing.T) {
	obj := new(Codec)
	for i := 0; i < 10000; i++ {
		obj.Encode("http://www.baidu.com")
	}
	expect := "2bJ"
	url := "http://www.google.com"
	shortUrl := obj.Encode(url)
	if shortUrl != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, shortUrl)
	}
	expect = url
	result := obj.Decode(shortUrl)
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}

	shortUrl = "A48s"
	result = obj.Decode(shortUrl)
	expect = ""
	if result != expect {
		t.Errorf("\n expect: %v\nbut got: %v", expect, result)
	}
}
