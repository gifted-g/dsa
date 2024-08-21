package main

type StringBuilder struct {
	strings []string
	size    int
	cap     int
}

func NewStringBuilder() *StringBuilder {
	intialCap := 10
	return &StringBuilder{
		strings: make([]string, intialCap),
		size:    0,
		cap:     intialCap,
	}
}

func (sb *StringBuilder) Append(str string) {
	if sb.size == sb.cap {
		sb.resize()
	}

	sb.strings[sb.size] = str
	sb.size++
}

func (sb *StringBuilder) resize() {
	newCap := sb.cap * 2
	newStrings := make([]string, newCap)
	copy(newStrings, sb.strings)
	sb.strings = newStrings
	sb.cap = newCap
}

func (sb *StringBuilder) ToString() string {
	var result string
	for _, str := range sb.strings[0:sb.size] {
		result += str
	}
	return result
}
