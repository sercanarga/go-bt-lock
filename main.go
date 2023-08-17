package main

import (
	"os/exec"
	"time"
	"tinygo.org/x/bluetooth"
)

var (
	deviceMAC       = ""      		   // Bluetooth Device MAC ID
	minRSSI   int16 = -60                      // Minimum RSSI tolerance value of the bluetooth device
	timeout         = 30 * time.Second         // Timeout Time
	adapter         = bluetooth.DefaultAdapter // Bluetooth Receiver
)

var timer = time.NewTicker(timeout)

func lockWindows() {
	err := exec.Command("cmd", "/C", "rundll32.exe user32.dll,LockWorkStation").Run()
	if err != nil {
		return
	}
}

func startTimer() {
	go func() {
		for range timer.C {
			lockWindows()
		}
	}()
}

func main() {
	startTimer()

	err := adapter.Enable()
	if err != nil {
		panic("Failed to initialize adapter!")
	}

	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if device.Address.String() == deviceMAC {
			if device.RSSI > minRSSI {
				timer.Reset(timeout)
			}
		}
	})
	if err != nil {
		panic("Scan Error!")
	}
}
