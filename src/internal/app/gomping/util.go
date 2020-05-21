package gomping

import (
	"fmt"
	"github.com/gdamore/tcell"
	"time"
)

func lossColor(loss float64) tcell.Color {
	switch {
	case loss > 0:
		return tcell.ColorRed
	default:
		return tcell.ColorWhite
	}
}

func rttColor(rtt time.Duration) tcell.Color {
	switch {
	case rtt > redLine:
		return tcell.ColorRed
	case rtt > yellowLine:
		return tcell.ColorYellow
	default:
		return tcell.ColorWhite
	}
}

func rttFormat(rtt time.Duration) string {
	switch {
	case rtt < time.Millisecond:
		return fmt.Sprintf("%0.2fµs", float64(rtt.Nanoseconds())/float64(time.Microsecond))
	case rtt < time.Second:
		return fmt.Sprintf("%0.2fms", float64(rtt.Nanoseconds())/float64(time.Millisecond))
	default:
		return fmt.Sprintf("%0.2fs", rtt.Seconds())
	}
}
