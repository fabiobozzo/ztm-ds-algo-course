package recursion

import (
	"ds_algo_course/pkg/data_structures/stacks"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"time"
)

type TowerOfHanoi struct {
	pins          map[string]stacks.Stack
	numberOfDisks int
}

func NewTowerOfHanoi(numberOfDisks int) *TowerOfHanoi {
	t := TowerOfHanoi{
		pins: map[string]stacks.Stack{
			"A": stacks.NewLinkedListStack(),
			"B": stacks.NewLinkedListStack(),
			"C": stacks.NewLinkedListStack(),
		},
		numberOfDisks: numberOfDisks,
	}

	for i := numberOfDisks; i > 0; i-- {
		t.pins["A"].Push(strconv.Itoa(i))
	}

	return &t
}

const sleepTime = time.Millisecond * 1000

func (t *TowerOfHanoi) Solve() {
	t.Draw()
	time.Sleep(1)

	t.moveDisk(t.numberOfDisks, "A", "C", "B")
}

func (t *TowerOfHanoi) moveDisk(disk int, fromPin, toPin, auxPin string) {
	if disk == 1 {
		//fmt.Printf("move disk 1 from %s to %s\n", fromPin, toPin)

		movingDisk, err := t.pins[fromPin].Pop()
		if err != nil {
			log.Fatal(err)
		}

		t.pins[toPin].Push(movingDisk)

		t.Draw()
		time.Sleep(sleepTime)

		return
	}

	t.moveDisk(disk-1, fromPin, auxPin, toPin)

	//fmt.Printf("move disk %d from %s to %s\n", disk, fromPin, toPin)

	movingDisk, err := t.pins[fromPin].Pop()
	if err != nil {
		log.Fatal(err)
	}

	t.pins[toPin].Push(movingDisk)

	t.Draw()
	time.Sleep(sleepTime)

	t.moveDisk(disk-1, auxPin, toPin, fromPin)
}

func (t *TowerOfHanoi) Draw() {
	//fmt.Printf("%v\n", t.pins)

	a := stackToArray(copyStack(t.pins["A"]))
	b := stackToArray(copyStack(t.pins["B"]))
	c := stackToArray(copyStack(t.pins["C"]))

	rowCount := len(a)

	if len(b) > rowCount {
		rowCount = len(b)
	}

	if len(c) > rowCount {
		rowCount = len(c)
	}

	var tableData [][]string
	for i := 0; i < rowCount; i++ {
		var row []string
		row = []string{}

		addElementToTableRow(i, rowCount, a, &row)
		addElementToTableRow(i, rowCount, b, &row)
		addElementToTableRow(i, rowCount, c, &row)

		tableData = append(tableData, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"A", "B", "C"})

	for _, v := range tableData {
		table.Append(v)
	}

	fmt.Print("\033[H\033[2J")
	table.Render()
}

func addElementToTableRow(index int, rowCount int, pin []string, row *[]string) {
	padding := rowCount - len(pin)
	index = index - padding

	if index >= 0 && index < len(pin) {
		aValue, err := strconv.Atoi(pin[index])
		if err != nil {
			log.Fatal(err)
		}

		*row = append(*row, diskToString(aValue))
	} else {
		*row = append(*row, "")
	}
}

func diskToString(disk int) string {
	diskString := ""
	for i := 0; i < disk; i++ {
		diskString += "#"
	}

	return diskString
}

func stackToArray(s stacks.Stack) []string {
	array := make([]string, s.Length())

	for i := s.Length() - 1; i >= 0; i-- {
		el, err := s.Pop()
		if err != nil {
			log.Fatal(err)
		}

		array[i] = el
	}

	return array
}

func copyStack(source stacks.Stack) stacks.Stack {
	var tmp []string
	var newStack stacks.Stack
	newStack = stacks.NewLinkedListStack()

	for !source.IsEmpty() {
		el, err := source.Pop()
		if err != nil {
			log.Fatal(err)
		}

		tmp = append(tmp, el)
		newStack.Push(el)
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		source.Push(tmp[i])
	}

	return newStack
}
