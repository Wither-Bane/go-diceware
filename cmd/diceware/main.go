package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sethvargo/go-diceware/diceware"

	//DIPO-CODE
	"crypto/rand"
	"math/big"
	//DIPO-CODE
)

var (
	flagWords = flag.Int("words", 6,
		"number of words to generate")
	flagSeparator = flag.String("separator", "-",
		"character to use between words")

	stdout, stderr = os.Stdout, os.Stderr
)

//DIPO-CODE
func gen_rand_num_prefix() int {
	var min int64 = 10
	var max int64 = 99

	max_iterations := 100

	var num big.Int

	for i := 0; i < max_iterations; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(max-min))
		if err != nil {
			panic(err)
		}

		if num.Int64() >= min {
			break
		}
	}
	return int(num.Int64())
}

//DIPO-CODE

func main() {
	flag.Parse()

	list, err := diceware.Generate(*flagWords)
	if err != nil {
		fmt.Fprintf(stderr, "error: %s\n", err)
		os.Exit(2)
	}

	//DIPO-CODE

	num1 := gen_rand_num_prefix()
	num2 := gen_rand_num_prefix()

	fmt.Fprint(stdout, num1)
	fmt.Fprint(stdout, *flagSeparator)
	//DIPO-CODE
	for i, w := range list {
		fmt.Fprint(stdout, w)
		if i < len(list)-1 {
			fmt.Fprint(stdout, *flagSeparator)
		}
	}
	//DIPO-CODE
	fmt.Fprint(stdout, *flagSeparator)
	fmt.Fprint(stdout, num2)
	//DIPO-CODE

	if fi, _ := stdout.Stat(); fi == nil || (fi.Mode()&os.ModeCharDevice) != 0 {
		fmt.Fprintln(stdout)
	}
}
