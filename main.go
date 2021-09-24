package main

import (
	"fmt"
	"strconv"
	"time"
)

var PBStartPiece string
var PBEndPiece string
var PBPiece string
var PBBackground string
var DoneText string
var WorkingText string

func main() {
	initFakeLoaders()
	progress(2000, "Integer Progress")
	fmt.Printf("\n")
	progressBar(2000, "Progress Bar")
	fmt.Printf("\n")
	progressSpinner(2000, "Progress Spinner")
	fmt.Printf("\n")
}

func initFakeLoaders() {
	PBStartPiece = DarkGray + "[" + "\033[0m"
	PBEndPiece = DarkGray + "]" + "\033[0m"
	PBPiece = White + "█" + "\033[0m"
	PBBackground = DarkGray + "██████████████████████████████████████████████████" + "\033[0m"
	DoneText = Green + "DONE   " + "\033[0m"
	WorkingText = Yellow + "WORKING" + "\033[0m"

}

func progress(speed int, caption string) {
	fmt.Printf("\n%s %s 0%%\n", WorkingText, caption)

	for i := 0; i < 101; i++ {
		fmt.Printf("\033[1A%s %s %s%%\n", WorkingText, caption, strconv.Itoa(i))

		time.Sleep(time.Millisecond * time.Duration(speed/100))
	}

	fmt.Printf("\033[1A%s %s\n", DoneText, caption)
}

func progressBar(speed int, caption string) {
	fmt.Printf("%s %s%s%s %s\033[59D", caption, PBStartPiece, PBBackground, PBEndPiece, WorkingText)

	for i := 0; i < 50; i++ {
		fmt.Print(PBPiece)

		time.Sleep(time.Millisecond * time.Duration(speed/50))
	}

	fmt.Printf("%s %s\n", PBEndPiece, DoneText)
}

func progressSpinner(duration int, caption string) {
	fmt.Printf("- %s\n", caption)

	for i := 0; i < (duration/100)/4; i++ {
		fmt.Printf("\033[1A\\ %s\n", caption)
		time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Printf("\033[1A| %s\n", caption)
		time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Printf("\033[1A/ %s\n", caption)
		time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Printf("\033[1A- %s\n", caption)
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}

const (
	Green        = "\033[1;32m"
	Yellow       = "\033[1;33m"
	White        = "\033[1;37m"
	DarkGray     = "\033[1;90m"
	DefaultColor = "\033[0m"
	Blink        = "\033[1;5m"
)
