package main

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

func printMessage(col int, row int, bg termbox.Attribute, fg termbox.Attribute, message string) {
    for _, char := range message {
        termbox.SetCell(col, row, char, fg, bg)
        col += runewidth.RuneWidth(char)
    }
}

func runEditor() {
    err := termbox.Init()
    if err != nil {fmt.Println(err); os.Exit(1)}
    for {
        printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "A texticular text editor")
        termbox.Flush()
        event := termbox.PollEvent()
        if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
            termbox.Close()
            break
        }
    }
}

func main() {
    runEditor()    
}
