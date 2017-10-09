package calculator

import (
	"fmt"
)

type Trend int
const  (
	TrendSame Trend = iota
	TrendUp
	TrendDown
)

type CryptoType string
const (
	BTC CryptoType = "BTC"
	ETH = "ETH"
	LTC = "LTC"
	LSK = "LSK"
	GAME = "GAME"
	BCC = "BCC"
	DASH = "DASH"
)

const (
	RedColor = "\033[00;31m"
	GreenColor = "\033[00;32m"
	YellowColor = "\033[00;33m"
	OriginalColor = "\033[00;39m"
)

type info struct {
	idx   int
	value float64
	trend Trend
}

type Calculator struct {
	values map[CryptoType]info
}


func New() *Calculator {
	fmt.Print("\033[2J")
	screenGoto(1,1)
	c := &Calculator{}
	c.values = make(map[CryptoType]info)
	c.values[BTC] = info{0, 0.0, TrendSame}
	c.values[ETH] = info{1, 0.0, TrendSame}
	c.values[LTC] = info{2, 0.0, TrendSame}
	c.values[LSK] = info{3, 0.0, TrendSame}
	c.values[GAME] = info{4, 0.0, TrendSame}
	c.values[BCC] = info{5, 0.0, TrendSame}
	c.values[DASH] = info{6, 0.0, TrendSame}
	return c
}

func (c *Calculator) SetValue(crypto CryptoType, value float64) {
	info, _ := c.values[crypto]
	current := info.value
	// info.trend = TrendSame

	if value > current {
		info.trend = TrendUp
	} else if value < current {
		info.trend = TrendDown
	}
	info.value = value
	c.values[crypto] = info
	info.display(crypto)
}


func (i *info) display(crypto CryptoType) {
	screenGoto(i.idx + 1, 1)
	color := OriginalColor
	if i.trend == TrendUp {
		color = GreenColor
	} else if i.trend == TrendDown {
		color = RedColor
	}
	fmt.Printf("%4s: %s%8.2f%s", crypto, color, i.value, OriginalColor)
}

func screenGoto(row int, col int) {
	ctrl := fmt.Sprintf("\033[%d;%dH", row, col)
	fmt.Print(ctrl)
}