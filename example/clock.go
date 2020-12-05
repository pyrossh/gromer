package main

import (
	"time"

	. "github.com/pyros2097/wapp"
)

func Clock(c *RenderContext) UI {
	timeValue, setTime := c.UseState(time.Now())
	running, setRunning := c.UseState(false)
	startTimer := func() {
		setRunning(true)
		go func() {
			for running().(bool) {
				setTime(time.Now())
				time.Sleep(time.Second * 1)
			}
		}()
	}
	stopTimer := func() {
		setRunning(false)
	}
	c.UseEffect(func() func() {
		startTimer()
		return stopTimer
	})

	return Col(
		Row(
			Div(Css("text-6xl"),
				Text("Clock"),
			),
		),
		Row(
			Div(Css("mt-10"),
				Text(timeValue().(time.Time).Format("15:04:05")),
			),
		),
		Row(
			Div(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(startTimer),
				Text("Start"),
			),
			Div(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(stopTimer),
				Text("Stop"),
			),
		),
	)
}
