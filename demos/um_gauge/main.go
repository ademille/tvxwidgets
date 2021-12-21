// Demo code for the bar chart primitive.
package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	gauge := tvxwidgets.NewUtilModeGauge()
	gauge.SetLabel("cpu usage:")
	gauge.SetLabelColor(tcell.ColorLightSkyBlue)
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetWarnPercentage(65)
	gauge.SetCritPercentage(80)
	gauge.SetBorder(true)

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				rangeLower := 0
				rangeUpper := 100
				randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				gauge.SetValue(float64(randomNum))
				app.Draw()
			}
		}
	}
	go update()

	if err := app.SetRoot(gauge, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
