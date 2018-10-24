package leetcode

/**
 * @title TinyURL 的加密与解密
 *
 * TinyURL是一种URL简化服务， 比如：当你输入一个URL https://leetcode.com/problems/design-tinyurl 时，它将返回一个简化的URL http://tinyurl.com/4e9iAk.

 * 要求：设计一个 TinyURL 的加密 encode 和解密 decode 的方法。
 * 你的加密和解密算法如何设计和运作是没有限制的，你只需要保证一个URL可以被加密成一个TinyURL，并且这个TinyURL可以用解密方法恢复成原本的URL。
 *
 * @see https://leetcode-cn.com/problems/encode-and-decode-tinyurl/description/
 *
 * 基本思路：
 * 数组存储longUrl，数组下标本身是自增id，id转62进制生成短链
 */

import (
	"strconv"
	"sync"
)

type Codec struct {
	Urls []string //longUrl列表
	Size int64    //当前URL数量，主要用于越界判断
}

var mu sync.Mutex

func (this *Codec) Encode(longUrl string) string {
	mu.Lock()
	defer mu.Unlock()
	this.Urls = append(this.Urls, longUrl)
	this.Size++
	//生成短链接
	return this.int2str(this.Size)
}

func (this *Codec) Decode(shortUrl string) string {
	//解出id
	i := this.str2int(shortUrl)
	if i > this.Size {
		return ""
	}
	return this.Urls[i-1]
}

func (this *Codec) int2str(n int64) (str string) {
	if n < 1 {
		return ""
	}
	return this.int2str(n/62) + func(i int64) (s string) {
		var res string
		if i >= 0 && i <= 9 {
			res = strconv.Itoa(int(i))
		} else if i >= 10 && i <= 35 {
			res = string(i - 10 + 'A')
		} else {
			res = string(i - 36 + 'a')
		}
		return res
	}(n%62)
}

func (this *Codec) str2int(s string) (res int64) {
	for _, v := range s {
		res = res*62 + func(c rune) int64 {
			var res int32
			if c >= '0' && c <= '9' {
				res = c - '0'
			} else if c >= 'A' && c <= 'Z' {
				res = c - 'A' + 10
			} else {
				res = c - 'a' + 36
			}
			return int64(res)
		}(v)
	}
	return
}
