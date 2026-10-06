package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	// 1) CPU
	c, _ := cpu.Percent(time.Second, false)
	if len(c) > 0 {
		fmt.Printf("CPU:    %.2f%%\n", c[0])
	}

	// 2) Memory
	m, _ := mem.VirtualMemory()
	if m != nil {
		fmt.Printf("Memory: %.2f%% (%.2f/%.2f GB)\n",
			m.UsedPercent,
			float64(m.Used)/1024/1024/1024,
			float64(m.Total)/1024/1024/1024)
	}

	// 3) Disk
	d, _ := disk.Usage("/")
	if d != nil {
		fmt.Printf("Disk /: %.2f%%\n", d.UsedPercent)
	}

	// 4) Processes
	procs, _ := process.Processes()

	for _, p := range procs {
		p.CPUPercent()
	}
	time.Sleep(time.Second)

	n := 5
	if len(procs) < n {
		n = len(procs)
	}

	// 5) Top 5 by CPU
	sort.Slice(procs, func(i, j int) bool {
		a, _ := procs[i].CPUPercent()
		b, _ := procs[j].CPUPercent()
		return a > b
	})
	fmt.Println("\n=== Top 5 by CPU ===")
	for _, p := range procs[:n] {
		name, _ := p.Name()
		c, _ := p.CPUPercent()
		fmt.Printf("%-8d %-25s %6.2f%%\n", p.Pid, name, c)
	}

	// 6) Top 5 by Memory
	sort.Slice(procs, func(i, j int) bool {
		a, _ := procs[i].MemoryPercent()
		b, _ := procs[j].MemoryPercent()
		return a > b
	})
	fmt.Println("\n=== Top 5 by Memory ===")
	for _, p := range procs[:n] {
		name, _ := p.Name()
		m, _ := p.MemoryPercent()
		fmt.Printf("%-8d %-25s %6.2f%%\n", p.Pid, name, m)
	}
} 