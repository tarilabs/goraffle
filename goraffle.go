package main

import (
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
)

func raffle(textarea string) []string {
	lines := strings.Split(textarea, "\n")

	results := make([]string, len(lines))
	for i, value := range rand.Perm(len(lines)) {
		results[i] = lines[value]
	}

	return results
}

func raffleHTML(textarea string) string {
	results := raffle(textarea)

	var str strings.Builder
	for _, line := range results {
		str.WriteString(fmt.Sprintf("<li>%s</li>", line))
	}

	return str.String()
}

func raffleJS(this js.Value, p []js.Value) interface{} {
	textarea := p[0].String()
	return js.ValueOf(raffleHTML(textarea))
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("raffle", js.FuncOf(raffleJS))

	<-c
}
