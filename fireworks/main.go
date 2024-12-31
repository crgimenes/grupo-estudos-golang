package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"golang.org/x/term"
)

const (
	ANSI_RESET          = "\033[0m"
	ANSI_CLEAR          = "\033[2J"
	ANSI_HIDE_CURSOR    = "\033[?25l"
	ANSI_SHOW_CURSOR    = "\033[?25h"
	ANSI_BLACK          = "\033[30m"
	ANSI_RED            = "\033[31m"
	ANSI_GREEN          = "\033[32m"
	ANSI_YELLOW         = "\033[33m"
	ANSI_BLUE           = "\033[34m"
	ANSI_MAGENTA        = "\033[35m"
	ANSI_CYAN           = "\033[36m"
	ANSI_WHITE          = "\033[37m"
	ANSI_SAVE_CURSOR    = "\033[s"
	ANSI_RESTORE_CURSOR = "\033[u"

	BANNER = `Feliz Ano Novo! Muito sucesso, felicidade! E claro, muitos códigos! Grupos de estudos de Go! `
)

var (
	rows, cols int
	mx         sync.Mutex
	bannerBuff []rune
)

func updateTerminalSize() {
	c, r, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("error getting size: %s\r\n", err)
	}
	mx.Lock()
	rows, cols = r, c

	bannerBuff = make([]rune, 0, cols)
	runes := []rune(BANNER)

	for i := 0; i < cols; i++ {
		bannerBuff = append(bannerBuff, runes[i%len(runes)])
	}

	os.Stdout.WriteString(ANSI_CLEAR)
	mx.Unlock()
}

func randomColor() string {
	colors := []string{
		ANSI_RED,
		ANSI_GREEN,
		ANSI_YELLOW,
		ANSI_BLUE,
		ANSI_MAGENTA,
		ANSI_CYAN,
		ANSI_WHITE,
	}
	return colors[rand.Intn(len(colors))]
}

func drawExplosion(x, y int, maxRadius int) {
	color := randomColor()
	for radius := 1; radius <= maxRadius; radius++ {
		for angle := 0; angle < 360; angle += 10 {
			theta := float64(angle) * (math.Pi / 180)
			px := x + int(float64(radius)*2.0*math.Cos(theta))
			py := y + int(float64(radius)*math.Sin(theta))
			sy := strconv.Itoa(py)
			sx := strconv.Itoa(px)
			os.Stdout.WriteString("\033[" + sy + ";" + sx + "H" + color + "*")
		}
		time.Sleep(100 * time.Millisecond)
	}
	for radius := 1; radius <= maxRadius; radius++ {
		for angle := 0; angle < 360; angle += 10 {
			theta := float64(angle) * (math.Pi / 180)
			px := x + int(float64(radius)*2.0*math.Cos(theta))
			py := y + int(float64(radius)*math.Sin(theta))
			sy := strconv.Itoa(py)
			sx := strconv.Itoa(px)
			os.Stdout.WriteString("\033[" + sy + ";" + sx + "H" + ANSI_RESET + " ")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func drawBanner() {
	np := 0
	for {
		np++

		x := 1
		mx.Lock()
		y := rows / 2
		mx.Unlock()

		sy := strconv.Itoa(y)
		sx := strconv.Itoa(x)

		os.Stdout.WriteString(ANSI_SAVE_CURSOR +
			"\033[" + sy + ";" + sx + "H" +
			ANSI_WHITE + string(bannerBuff) +
			ANSI_RESTORE_CURSOR)

		if np%4 == 0 {
			copy(bannerBuff, bannerBuff[1:])
			bannerBuff[len(bannerBuff)-1] = bannerBuff[0]
		}

		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	updateTerminalSize()

	os.Stdout.WriteString(ANSI_CLEAR + ANSI_HIDE_CURSOR)

	go func() {
		for {
			mx.Lock()
			x := rand.Intn(cols-10) + 5
			y := rand.Intn(rows-10) + 5
			mx.Unlock()
			maxRadius := rand.Intn(5) + 3
			go drawExplosion(x, y, maxRadius)
			time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH, os.Interrupt, syscall.SIGTERM)
	go func() {
		for caux := range ch {
			switch caux {
			case syscall.SIGWINCH:
				updateTerminalSize()
			case os.Interrupt, syscall.SIGTERM:
				os.Stdout.WriteString(
					ANSI_SHOW_CURSOR +
						ANSI_RESET +
						ANSI_CLEAR +
						"\033[1;1H\r\n",
				)

				os.Exit(0)
			}
		}
	}()
	ch <- syscall.SIGWINCH

	os.Stdout.WriteString(ANSI_CLEAR + ANSI_HIDE_CURSOR)

	go drawBanner()

	c := make(chan struct{})
	<-c
}
