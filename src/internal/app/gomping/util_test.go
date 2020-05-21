package gomping

import (
	"github.com/gdamore/tcell"
	"testing"
	"time"
)

type caseTestLossColor []struct {
	desc   string
	loss   float64
	result tcell.Color
}

func getTestLossColor() caseTestLossColor {
	n := float64(1)
	return caseTestLossColor{
		{
			desc:   "Loss    0%",
			loss:   float64(0),
			result: tcell.ColorWhite,
		}, {
			desc:   "Loss    1%",
			loss:   n / float64(100),
			result: tcell.ColorRed,
		}, {
			desc:   "Loss   50%",
			loss:   n / float64(2),
			result: tcell.ColorRed,
		}, {
			desc:   "Loss  0.1%",
			loss:   n / float64(1000),
			result: tcell.ColorRed,
		}, {
			desc:   "Loss 0.01%",
			loss:   n / float64(10000),
			result: tcell.ColorRed,
		},
	}
}

func TestLossColor(t *testing.T) {
	t.Helper()
	testCase := getTestLossColor()

	for _, v := range testCase {
		t.Run(v.desc, func(st *testing.T) {
			c := lossColor(v.loss)
			if c != v.result {
				st.Errorf("Expect: %v, Result: %v, Loss: %v%%)", v.result, c, v.loss)
			}
		})
	}
}

type caseTestRttColor []struct {
	desc   string
	rtt    time.Duration
	result tcell.Color
}

func getTestRttColor() caseTestRttColor{
	return caseTestRttColor{
		{
			desc: "  1ms",
			rtt: time.Millisecond,
			result: tcell.ColorWhite,
		},{
			desc: " 49ms",
			rtt: time.Millisecond * 50,
			result: tcell.ColorWhite,
		},{
			desc: " 50ms",
			rtt: time.Millisecond * 51,
			result: tcell.ColorYellow,
		},{
			desc: "149ms",
			rtt: time.Millisecond * 150,
			result: tcell.ColorYellow,
		},{
			desc: "150ms",
			rtt: time.Millisecond * 151,
			result: tcell.ColorRed,
		},

	}
}

func TestRttColor(t *testing.T) {
	t.Helper()
	testCase := getTestRttColor()

	redLine = time.Millisecond * 150
	yellowLine = time.Millisecond *50

	for _, v := range testCase {
		t.Run(v.desc, func(st *testing.T) {
			c := rttColor(v.rtt)
			if c != v.result {
				st.Errorf("Expect: %v, Result: %v, Rtt: %v)", v.result, c, v.rtt)
			}
		})
	}
}

type caseTestRttFormat []struct{
	desc string
	rtt time.Duration
	result string
}

func getTestRttFormat() caseTestRttFormat{
	return caseTestRttFormat{
		{
			desc: "1ms",
			rtt: time.Millisecond,
			result: "1.00ms",
		},{
			desc: "10ms",
			rtt: time.Millisecond * 10,
			result: "10.00ms",
		},{
			desc: "100ms",
			rtt: time.Millisecond * 100,
			result: "100.00ms",
		},{
			desc: "100.1ms",
			rtt: time.Microsecond * 100100,
			result: "100.10ms",
		},{
			desc: "100.11ms",
			rtt: time.Microsecond * 100110,
			result: "100.11ms",
		},{
			desc: "0.1ms",
			rtt: time.Millisecond / 10,
			result: "100.00µs",
		},{
			desc: "100.111ms",
			rtt: time.Microsecond * 100111,
			result: "100.11ms",
		},{
			desc: "100.888ms",
			rtt: time.Microsecond * 100888,
			result: "100.89ms",
		},{
			desc: "90µs",
			rtt: time.Microsecond * 90,
			result: "90.00µs",
		},{
			desc: "50ns",
			rtt: time.Nanosecond * 50,
			result: "0.05µs",
		},{
			desc: "1s",
			rtt: time.Second,
			result: "1.00s",
		},
	}
}

func TestRttFormat(t *testing.T) {
	t.Helper()
	testCase := getTestRttFormat()

	for _, v := range testCase {
		t.Run(v.desc, func(st *testing.T) {
			s := rttFormat(v.rtt)
			if s != v.result {
				st.Errorf("aa %v", s)
			}
		})
	}
}
