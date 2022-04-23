# **brainfuck-interpreter-Golang**


Brainfuck is an esoteric programming language created in 1993 by Urban Müller.[1]

Notable for its extreme minimalism, the language consists of only eight simple commands, a data pointer and an instruction pointer. While it is fully Turing complete, it is not intended for practical use, but to challenge and amuse programmers. Brainfuck simply requires one to break commands into microscopic steps.

### **Language design**

The language consists of eight commands, listed below. A brainfuck program is a sequence of these commands, possibly interspersed with other characters (which are ignored). The commands are executed sequentially, with some exceptions: an instruction pointer begins at the first command, and each command it points to is executed, after which it normally moves forward to the next command. The program terminates when the instruction pointer moves past the last command.

The brainfuck language uses a simple machine model consisting of the program and instruction pointer, as well as a one-dimensional array of at least 30,000 byte cells initialized to zero; a movable data pointer (initialized to point to the leftmost byte of the array); and two streams of bytes for input and output (most often connected to a keyboard and a monitor respectively, and using the ASCII character encoding).

# BASICS

The idea behind brainfuck is memory manipulation. Basically you are given an array of 30,000 1byte memory blocks. The array size is actually dependent upon the implementation used in the compiler or interpretor, but standard brainfuck states 30,000. Within this array, you can increase the memory pointer, increase the value at the memory pointer, etc. Let me first present to you the 8 operators available to us.

```> = increases memory pointer, or moves the pointer to the right 1 block.
< = decreases memory pointer, or moves the pointer to the left 1 block.
+ = increases value stored at the block pointed to by the memory pointer
- = decreases value stored at the block pointed to by the memory pointer
[ = like c while(cur_block_value != 0) loop.
] = if block currently pointed to's value is not zero, jump back to [
, = like c getchar(). input 1 character.
```

# Usage

`go get github.com/Overover1400/brainfuck`

```go
inp := `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.`

	bf := instructor.NewInstructor()
	if err := bf.CompileBf(inp); err != nil {
		log.Fatalln(err)
	}

	if err := bf.ExecuteBf(); err != nil {
		log.Fatalln(err)
	}

```
========================================================================
