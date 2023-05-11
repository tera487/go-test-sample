package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// Sleeper はテストと実装で共通化するため
type Sleeper interface {
	Sleep()
}

// DefaultSleeper は実装
type DefaultSleeper struct{}

// Sleep はSleeper interfaceのための実装
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper はカウントダウンの設定
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Countdown はカウントダウンで数値を表示する関数
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; 0 < i; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// Sleep は
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
