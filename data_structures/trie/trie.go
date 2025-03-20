package trie

type Trie struct {
	m   map[string]*Trie
	end bool
}

func Constructor() Trie {
	return Trie{
		m: make(map[string]*Trie),
	}
}

func (t *Trie) Insert(word string) {
	tr := t
	for _, w := range word {
		existed, ok := tr.m[string(w)]
		if !ok {
			next := Constructor()
			tr.m[string(w)] = &next
			tr = &next
			continue
		}
		tr = existed
	}

	tr.end = true
}

func (t *Trie) Search(word string) bool {
	tr := t
	var ok bool
	for _, w := range word {
		tr, ok = tr.m[string(w)]
		if !ok {
			return false
		}
	}

	return tr.end
}

func (t *Trie) StartsWith(prefix string) bool {
	tr := t
	var ok bool
	for _, w := range prefix {
		tr, ok = tr.m[string(w)]
		if !ok {
			return false
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
