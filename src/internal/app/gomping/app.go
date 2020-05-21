package gomping

import (
	"fmt"
	"github.com/digineo/go-ping"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"math"
	"net"
	"strconv"
	"time"
)

var e error

var redLine, yellowLine time.Duration

var config *StartupConfig
var pinger *ping.Pinger

type Slide func(nextSlide func(), screenSize int, index int) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()
var tg []TargetGroup

const needWidth = 60

func Run(conf string) {
	// load config and config check
	config, e = NewConfig(conf)
	if e != nil {
		log.Fatal(e)
	}

	// create pinger
	pinger, e = newPinger()
	if e != nil {
		log.Fatal(e)
	}

	pages := tview.NewPages()
	pageState := false

	tg = config.TargetGroups

	if config.Settings.RedLine == 0 {
		redLine = time.Millisecond * 150
	} else {
		redLine = time.Millisecond * time.Duration(config.Settings.RedLine)
	}
	if config.Settings.YellowLine == 0 {
		yellowLine = time.Millisecond * 50
	} else {
		yellowLine = time.Millisecond * time.Duration(config.Settings.YellowLine)
	}

	// ページ当たりのグループ数
	var gpp int
	c := len(tg)
	if config.Settings.GroupPerPage != 0 {
		// コンフィグで指定している場合
		gpp = config.Settings.GroupPerPage
	} else {
		// terminal sizeからページ数を算出
		w, _, _ := terminal.GetSize(0)
		// Group per page
		gpp = w / needWidth
	}
	// ページ数
	p := int(math.Ceil(float64(c) / float64(gpp)))

	// 最下部にページとマニュアルのショートカット表示
	footerInfo := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		}).
		SetTextAlign(tview.AlignLeft)

	// ページ移動
	previousSlide := func() {
		slide, _ := strconv.Atoi(footerInfo.GetHighlights()[0])
		slide = (slide - 1 + p) % p
		footerInfo.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(footerInfo.GetHighlights()[0])
		slide = (slide + 1) % p
		footerInfo.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

	// Build PingGroup pages
	for i, j := 0, 0; i < c; i, j = i+gpp, j+1 {
		_, primitive := DrawPingList(nextSlide, gpp, i)
		pages.AddPage(strconv.Itoa(j), primitive, true, j == 0)
		_, _ = fmt.Fprintf(footerInfo, `["%d"][darkcyan]%d[white][""] `, j, j+1)
	}

	pages.AddPage("manualModal", DrawManual(), true, false)

	footerInfo.Highlight("0")

	footerManual := tview.NewTextView().
		SetText("Ctrl+M: Manual ").
		SetTextAlign(tview.AlignRight)

	footer := tview.NewFlex().
		AddItem(footerInfo, 0, 1, false).
		AddItem(footerManual, 0, 1, false)

	// main layout
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(footer, 1, 1, false)

	// キーアクション
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !pageState {
			if event.Key() == tcell.KeyCtrlN || event.Key() == tcell.KeyRight {
				nextSlide()
			} else if event.Key() == tcell.KeyCtrlP || event.Key() == tcell.KeyLeft {
				previousSlide()
			}
		}
		if event.Key() == tcell.KeyCtrlM {
			pageState = true
			footerInfo.Highlight("manualModal").
				ScrollToHighlight()
		} else if event.Key() == tcell.KeyCtrlQ {
			pageState = false
			footerInfo.Highlight("0").
				ScrollToHighlight()
		}

		return event
	})

	PingLoop()

	// Start app
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func PingLoop() {
	for i := 0; i < len(tg); i++ {
		for j := 0; j < len(tg[i].Hosts); j++ {
			ip, e := net.ResolveIPAddr("ip", tg[i].Hosts[j].Ip)
			if e != nil {
				log.Fatal(e)
			}
			go func(i, j int, ip *net.IPAddr) {
				var rtt time.Duration
				var e interface{}
				var loss float64
				var avg string

				for {
					rtt, e = pinger.Ping(ip, opts.resolverTimeout)

					// set count cell
					tables[i].GetCell(j+1, cCnt).SetText(fmt.Sprint(tg[i].Hosts[j].Stats.Count))

					// set rtt cell
					if e != nil {
						tg[i].Hosts[j].Stats.Loss++
						tables[i].GetCell(j+1, cRtt).SetText("-")
					} else {
						tg[i].Hosts[j].Stats.Count++
						tg[i].Hosts[j].Stats.sum = tg[i].Hosts[j].Stats.sum + rtt
						tables[i].GetCell(j+1, cRtt).SetText(rttFormat(rtt)).SetTextColor(rttColor(rtt))
					}

					// set loss cell
					loss = tg[i].Hosts[j].Stats.Loss / (tg[i].Hosts[j].Stats.Count + tg[i].Hosts[j].Stats.Loss) * 100
					tables[i].GetCell(j+1, cLoss).SetText(fmt.Sprintf("%0.1f%%", loss)).SetTextColor(lossColor(loss))

					// set avg cell
					if tg[i].Hosts[j].Stats.Count > 1 {
						avg = rttFormat(time.Duration(tg[i].Hosts[j].Stats.sum.Nanoseconds() / int64(tg[i].Hosts[j].Stats.Count)))
					} else {
						avg = "0"
					}
					tables[i].GetCell(j+1, cAvg).SetText(avg)

					app.Draw()
					time.Sleep(opts.interval)
				}
			}(i, j, ip)
		}
	}
}
