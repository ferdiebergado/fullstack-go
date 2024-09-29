package stdout

const (
	Escape = "\x1b["
	Reset  = Escape + "0m"

	Red     = Escape + "31m"
	Green   = Escape + "32m"
	Yellow  = Escape + "33m"
	Blue    = Escape + "34m"
	Magenta = Escape + "35m"
	Cyan    = Escape + "36m"
)
