package main

import (
	"iter"
)

// Countdown - это наша функция-итератор.
// Она возвращает другую функцию типа iter.Seq[int].
// iter.Seq[int] - это, по сути, func(yield func(int) bool).
func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i >= 0; i-- {
			// "Отдаем" значение в цикл for range.
			// Если yield вернет false, мы должны прекратить итерацию.
			if !yield(i) {
				return
			}
		}
	}
}

type JediList []Jedi

type Jedi struct {
	Name string
	Rank JediRank
}

type JediRank int

const (
	Youngling JediRank = iota
	Padawan
	Master
)

func (jl JediList) Younglings() iter.Seq[Jedi] {
	return func(yield func(Jedi) bool) {
		for _, j := range jl {
			if j.Rank != Youngling {
				continue
			}

			if !yield(j) {
				break
			}
		}
	}
}

func main() {
	jedis := JediList{
		{Name: "Anakin Skywalker", Rank: Padawan},
		{Name: "Ahsoka Tano", Rank: Youngling},
		{Name: "Obi-Wan Kenobi", Rank: Master},
		{Name: "Yoda", Rank: Master},
		{Name: "Luke Skywalker", Rank: Youngling},
		{Name: "Leia Organa", Rank: Youngling},
		{Name: "Rey Skywalker", Rank: Youngling},
		{Name: "Grogu", Rank: Youngling},
		{Name: "Ezra Bridger", Rank: Youngling},
	}

	next, stop := iter.Pull(jedis.Younglings())
	defer stop()

	for {
		if jedi, ok := next(); ok {
			println(jedi.Name)
		} else {
			break
		}
	}

}
