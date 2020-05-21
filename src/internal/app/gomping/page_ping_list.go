package gomping

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type tableHeader struct {
	name string
	color tcell.Color
	selectable bool
	align int
}

var tbopts = map[int]tableHeader{
	cHostname: {
		name: "Host Name",
		color: tcell.ColorYellow,
		selectable: true,
		align: tview.AlignLeft,
	},
	cIp: {
		name: "IP Address",
		color: tcell.ColorDarkCyan,
		selectable: true,
		align: tview.AlignLeft,
	},
	cRtt: {
		name: "RTT",
		color: tcell.ColorYellowGreen,
		selectable: true,
		align: tview.AlignRight,
	},
	cCnt: {
		name: "Cnt",
		color: tcell.ColorWhite,
		selectable: true,
		align: tview.AlignRight,
	},
	cLoss: {
		name: "Loss",
		color: tcell.ColorRed,
		selectable: true,
		align: tview.AlignRight,
	},
	cAvg: {
		name: "Avg",
		color: tcell.ColorWhite,
		selectable: true,
		align: tview.AlignRight,
	},
}

const cHostname, cIp, cRtt, cCnt, cLoss, cAvg = 0, 1, 2, 3, 4, 5

var tables []*tview.Table

func DrawPingList(nextSlide func(), gpp int, idx int) (title string, content tview.Primitive) {
	flex := tview.NewFlex()

	s := len(tg[idx:])
	for i := 0; i < gpp; i++ {
		if i >= s {
			break
		}
		groupFlex := tview.NewFlex()
		table := tview.NewTable().SetFixed(1, 1)

		// set table headers
		for j := 0; j < len(tbopts); j++ {
			table.SetCell(0, j, tview.NewTableCell(tbopts[j].name).SetTextColor(tbopts[j].color).SetSelectable(tbopts[j].selectable).SetAlign(tbopts[j].align))
		}
		// set table cells
		for j := 0; j < len(tg[idx+i].Hosts); j++ {
			table.SetCell(j+1, cHostname, tview.NewTableCell(tg[idx+i].Hosts[j].Hostname).SetAlign(tview.AlignLeft))
			table.SetCell(j+1, cIp, tview.NewTableCell(tg[idx+i].Hosts[j].Ip).SetAlign(tview.AlignLeft))
			table.SetCell(j+1, cRtt, tview.NewTableCell("").SetAlign(tview.AlignRight))
			table.SetCell(j+1, cCnt, tview.NewTableCell("").SetAlign(tview.AlignRight))
			table.SetCell(j+1, cLoss, tview.NewTableCell("0%").SetAlign(tview.AlignRight))
			table.SetCell(j+1, cAvg, tview.NewTableCell("").SetAlign(tview.AlignRight))
		}

		table.SetSelectable(true, false).SetBorder(false)
		tables = append(tables, table)

		groupFlex.AddItem(table, 0, 1, false)
		groupFlex.SetBorder(true).SetTitle(tg[idx+i].GroupName).SetTitleColor(tcell.ColorYellow).SetTitleAlign(tview.AlignLeft)

		flex.AddItem(groupFlex, 0, 1, true)
	}
	return "PingList", flex
}
