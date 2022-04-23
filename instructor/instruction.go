package instructor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	OpInc      = iota //Operation increment "+"
	OpDec             //Operation decrease  "-"
	OpLeft            //Operation left      "<"
	OpRight           //Operation right     ">"
	OpOut             //Operation output    "."
	OpInp             //Operation input     ","
	OpForward         //Operation forward   "["
	OpBackward        //Operation backward  "]"
)
const DataSize = 65535

type Instructor struct {
	Instruct []Instruction
	Cells    []uint16
}

type Instruction struct {
	Operate uint16
	Operand uint16
}

func NewInstructor() *Instructor {
	return &Instructor{}
}

func (ins *Instructor) CompileBf(input string) error {
	if len(input) <= 0 {
		return errors.New("input is empty")
	}

	var (
		cell, backWPos uint16
		forWPosS       = make([]uint16, 0)
	)
	for _, inp := range []byte(input) {
		switch inp {
		case '+':
			appendIns(ins, OpInc)
		case '-':
			appendIns(ins, OpDec)
		case '>':
			appendIns(ins, OpRight)
		case '<':
			appendIns(ins, OpLeft)
		case '.':
			appendIns(ins, OpOut)
		case ',':
			appendIns(ins, OpInp)
		case '[':
			appendIns(ins, OpForward)
			forWPosS = append(forWPosS, cell)
		case ']':
			appendIns(ins, OpBackward)
			if len(forWPosS) == 0 {
				return errors.New("missing '['")
			}
			backWPos = forWPosS[len(forWPosS)-1]
			forWPosS = forWPosS[:len(forWPosS)-1]
			ins.Instruct[cell].Operand = backWPos
			ins.Instruct[backWPos].Operand = cell
		default:
			cell--
		}
		cell++
	}
	if len(forWPosS) != 0 {
		return errors.New("missing ']'")
	}
	return nil
}

func (ins *Instructor) ExecuteBf() error {

	var (
		memorySlot = make([]uint16, DataSize)
		memoryPos  = 0
		reader     = bufio.NewReader(os.Stdin)
	)

	for i := 0; i < len(ins.Instruct); i++ {

		switch ins.Instruct[i].Operate {
		case OpInc:
			memorySlot[memoryPos]++
		case OpDec:
			memorySlot[memoryPos]--
		case OpLeft:
			memoryPos--
		case OpRight:
			memoryPos++
		case OpOut:
			fmt.Printf("%c", memorySlot[memoryPos])
			ins.Cells = append(ins.Cells, memorySlot[memoryPos])
		case OpInp:
			read, err := reader.ReadByte()
			if err != nil {
				return errors.New("reader has error : " + err.Error())
			}
			memorySlot[memoryPos] = uint16(read)
		case OpForward:
			if memorySlot[memoryPos] == 0 {
				i = int(ins.Instruct[i].Operand)
			}
		case OpBackward:
			if memorySlot[memoryPos] > 0 {
				i = int(ins.Instruct[i].Operand)
			}
		default:
			return errors.New("unknown operator")
		}
	}
	return nil
}

func appendIns(ins *Instructor, op uint16) {
	ins.Instruct = append(ins.Instruct, Instruction{Operate: op})
}
