package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danielgatis/go-vte/ansi"
)

func main() {
	var seqs []ansi.Sequence
	reader := bufio.NewReader(os.Stdin)
	decoder := ansi.NewDecoder(reader)
	if err := decoder.Decode(&seqs); err != nil {
		panic(err)
	}

	for _, seq := range seqs {
		fmt.Println(seq.Info())
	}
}
