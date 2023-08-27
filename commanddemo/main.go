package main

import c "designpattern/commanddemo/command"

func main() {
	tv := &c.Tivi{}
	onCommand := &c.OnCommand{
		Device: tv,
	}
	offCommand := &c.OffCommand{
		Device: tv,
	}

	onButton := &c.Button{
		Command: onCommand,
	}
	onButton.Press()
	offButton := &c.Button{
		Command: offCommand,
	}
	offButton.Press()
}
