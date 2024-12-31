package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"golang.org/x/term"
)

const (
	ANSI_RESET           = "\033[0m"
	ANSI_CLEAR           = "\033[2J"
	ANSI_HIDE_CURSOR     = "\033[?25l"
	ANSI_SHOW_CURSOR     = "\033[?25h"
	ANSI_MOVE_CURSOR_TOP = "\033[H"
)

var (
	rows, cols   int
	mx           sync.RWMutex
	preSinX      []float64
	preCosX      []float64
	preSinY      []float64
	preCosY      []float64
	preSinXY     [][]float64
	preCosXY     [][]float64
	colorStrings [256]string
)

func updateTerminalSize() {
	c, r, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("error getting size: %s\r\n", err)
	}
	mx.Lock()
	rows, cols = r, c
	precalculateValues(rows, cols)
	mx.Unlock()
}

func precalculateValues(r, c int) {
	preSinX = make([]float64, c)
	preCosX = make([]float64, c)
	preSinY = make([]float64, r)
	preCosY = make([]float64, r)
	preSinXY = make([][]float64, r)
	preCosXY = make([][]float64, r)

	for i := 0; i < r; i++ {
		preSinXY[i] = make([]float64, c)
		preCosXY[i] = make([]float64, c)
	}

	for x := 0; x < c; x++ {
		aX := (float64(x) / float64(c)) * math.Pi * 3
		preSinX[x] = math.Sin(aX)
		preCosX[x] = math.Cos(aX)
	}
	for y := 0; y < r; y++ {
		aY := (float64(y) / float64(r)) * math.Pi * 3
		preSinY[y] = math.Sin(aY)
		preCosY[y] = math.Cos(aY)
		for x := 0; x < c; x++ {
			aXY := ((float64(x) + float64(y)) / float64(c)) * math.Pi * 3
			preSinXY[y][x] = math.Sin(aXY)
			preCosXY[y][x] = math.Cos(aXY)
		}
	}
}

func precalculateColorStrings() {
	for i := 0; i < 256; i++ {
		colorStrings[i] = "\x1b[48;5;" + strconv.Itoa(i) + "m\x1b[38;5;0m"
	}
}

func main() {
	fmt.Print(ANSI_HIDE_CURSOR)
	precalculateColorStrings()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH, os.Interrupt)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGWINCH:
				updateTerminalSize()
			case os.Interrupt:
				os.Stdout.WriteString(
					ANSI_RESET +
						ANSI_SHOW_CURSOR +
						ANSI_CLEAR +
						ANSI_MOVE_CURSOR_TOP,
				)
				os.Stdout.WriteString("Bye!\n")
				os.Exit(0)
			}
		}
	}()

	// Força cálculo inicial
	c <- syscall.SIGWINCH

	marquee := "Grupo de estudos de Go! "
	mLen := len(marquee)
	start := time.Now()

	bufCap := 20000
	buf := make([]byte, 0, bufCap)

	for {
		t := time.Since(start).Seconds()
		mx.RLock()
		r, cc := rows, cols
		mx.RUnlock()

		sT := math.Sin(t)
		cT := math.Cos(t)

		buf = buf[:0]
		buf = append(buf, ANSI_MOVE_CURSOR_TOP...)

		for y := 0; y < r; y++ {
			for x := 0; x < cc; x++ {
				sumSin := preSinX[x] + preSinY[y] + preSinXY[y][x]
				sumCos := preCosX[x] + preCosY[y] + preCosXY[y][x]
				val := sumSin*cT + sumCos*sT

				color := int((val + 3) / 6 * 255)
				if color < 0 {
					color = 0
				} else if color > 255 {
					color = 255
				}

				buf = append(buf, colorStrings[color]...)
				buf = append(buf, marquee[x%mLen])
			}
		}

		os.Stdout.Write(buf)
		time.Sleep(60 * time.Millisecond)
	}
}
