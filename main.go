package main

import (
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadUI() fyne.CanvasObject {
    list := container.NewVBox(
        widget.NewLabel("label 1"),
        widget.NewLabel("label 2"),
    )

    bar := widget.NewToolbar(
        widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
        widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
    )

    content := widget.NewMultiLineEntry()

    left := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)

    ui := container.NewHSplit(left, content)
    ui.Offset = 0.25
    return ui
}

func main() {
    app := app.New()
    window := app.NewWindow("Notes")

    window.Resize(fyne.NewSize(1280, 720))
    window.SetContent(loadUI())
    window.ShowAndRun()
}
