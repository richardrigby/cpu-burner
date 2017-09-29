package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// func getCPUSpeed() {
// 	for {
// 		cpuInfo, err := cpu.Info()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, v := range cpuInfo {
// 			fmt.Printf(" [CPU:%d %.0fMhz]\n ", v.CPU, v.Mhz)
// 		}
// 		fmt.Println("")
// 		fmt.Println("")

// 		time.Sleep(10 * time.Second)
// 	}
// }

func main() {
	fmt.Println("MAX CPU - maxs out all CPU cores to raise the temperature of the CPU.")
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())

	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					math.Floor(float64(10) / 2) // Pointless calculation
				}
			}
		}()
	}

	//go getCPUSpeed()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("q to quit")
		fmt.Println("")
		text, _ := reader.ReadString('\n')
		if text == "q\n" {
			os.Exit(0)
		}
	}

}
