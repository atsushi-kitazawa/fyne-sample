package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	doMain()
}

func doMain() {
	a := app.New()
	w := a.NewWindow("timer")
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.Show()

	w2 := a.NewWindow("hello world")
	w2.SetContent(container.NewVBox(makeUI()))
	w2.Resize(fyne.NewSize(300, 100))
	w2.Show()

	a.Run()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	out := widget.NewLabel("Hello world")
	in := widget.NewEntry()
	btn := widget.NewButton("clear", func() {
		in.SetText("")
	})

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + " !")
	}
	return out, in, btn
}
