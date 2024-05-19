package main

import "fmt"

func main() {
	l := Constructor()
	fmt.Println(l.ShouldPrintMessage(1, "foo"))
	fmt.Println(l.ShouldPrintMessage(2, "bar"))
	fmt.Println(l.ShouldPrintMessage(3, "foo"))
	fmt.Println(l.ShouldPrintMessage(8, "bar"))
	fmt.Println(l.ShouldPrintMessage(10, "foo"))
	fmt.Println(l.ShouldPrintMessage(11, "foo"))
}

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{
		m: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	v, ok := l.m[message]
	if !ok {
		l.m[message] = timestamp
		return true
	}

	if timestamp-v >= 10 {
		l.m[message] = timestamp
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
