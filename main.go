package main

import (
	"flag"
	"github.com/Overover1400/brainfuck/instructor"
	"io/ioutil"
	"log"
	"os"
)

// Flags
var (
	input = flag.String("i", "", "optional you can put input your brainfuck command for example : ++++.")
	//cellSize = flag.Bool("c", true, "optional cell size if you want to")
	//verbose    = flag.Bool("v", false, "verbose debugging")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	var inp string

	//input from arg :
	if len(os.Args) == 2 {
		inp = os.Args[1]
		// input from flag -i :
	} else if *input != "" {
		inp = *input
		// input from file :
	} else if file, err := os.OpenFile("brainf.bf", os.O_RDWR, 0755); err == nil {
		defer file.Close()
		if inpByte, err := ioutil.ReadAll(file); err == nil {
			inp = string(inpByte)
		}
		// default input :
	} else {
		inp = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.[][].++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.[][]."
	}

	bf := instructor.NewInstructor()
	if err := bf.CompileBf(inp); err != nil {
		log.Fatalln(err)
	}

	if err := bf.ExecuteBf(); err != nil {
		log.Fatalln(err)
	}

}
