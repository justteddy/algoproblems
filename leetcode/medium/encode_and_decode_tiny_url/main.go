package main

import "fmt"

// https://leetcode.com/problems/encode-and-decode-tinyurl/
func main() {
	codec := Constructor()
	original := "http://some-url"
	short := codec.encode(original)
	long := codec.decode(short)

	fmt.Println(original, short, long)
}

type Codec struct {
	counter rune
	m       map[string]string
}

func Constructor() Codec {
	return Codec{
		counter: '0',
		m:       make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.counter++
	short := string(this.counter)
	this.m[short] = longUrl

	return short
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}
