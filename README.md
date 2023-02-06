# Go BT Lock
Locks the screen in Windows using the distance value (RSSI) of your Bluetooth device.

# Build
1. Download the ZIP file using git clone or the [link](https://github.com/sercanarga/go-bt-lock/archive/refs/heads/main.zip).
2. Customize the values in the `main.go` file.
```go
deviceMAC       = ""                       // Bluetooth Device MAC ID
minRSSI   int16 = -60                      // Minimum RSSI tolerance value of the bluetooth device
timeout         = 20 * time.Second         // Timeout Time
adapter         = bluetooth.DefaultAdapter // Bluetooth Receiver
```
3. Take the windowless build.
```shell
go build -ldflags -H=windowsgui .
```

# Add Startup
To add it to startup, use `WIN+R` to open the run window and enter the command `shell:startup`. Can copy the compiled `bt-lock.exe` file to the startup folder. It will then run at startup.

<img src="https://i.imgur.com/pkHr8UM.png" width="200">

# Performance
In the tests I did, I found that it uses about 2.7mb of RAM and 0.2% to 1% CPU usage.

<img src="https://i.imgur.com/kfZYEev.png" width="350">
