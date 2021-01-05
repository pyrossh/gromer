package main

import (
	"time"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
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

	return Col(Css("text-3xl text-gray-700"),
		Header(c),
		Row(
			Div(Css("underline"),
				Text("Clock"),
			),
		),
		Row(
			Div(Css("mt-10"),
				Text(timeValue().(time.Time).Format("15:04:05")),
			),
		),
		Row(
			Button(Css("btn m-20"), OnClick(startTimer),
				Text("Start"),
			),
			Button(Css("btn m-20"), OnClick(stopTimer),
				Text("Stop"),
			),
		),
	)
}
