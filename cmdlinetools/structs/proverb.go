package main

type proverb struct {
	line  string
	chars map[rune]int
}

//attach countChars to proverb
//(p *proverb) is a receiver and makes this function
// become a method (function that has a receiver)
// with * you can modify values on memory
// without * you are using a copy
func (p *proverb) countChars() {
	if p.chars != nil {
		return
	}

	m := make(map[rune]int, 0)
	for _, c := range p.line {
		m[c] = m[c] + 1
	}
	p.chars = m
}

func newProverb(line string) *proverb {
	p := new(proverb)
	p.line = line
	p.countChars()
	return p
}
