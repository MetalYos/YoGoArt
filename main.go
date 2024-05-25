package main

import (
	"log"
    settings "yogoart/utils/settings"
    ser "yogoart/yoserial"
    // "yogoart/tui"
    vid "yogoart/utils/video_utils"
)

func main() {
    appSettings := settings.LoadSettings("settings.json")

    yoSerial := ser.NewYoSerial(appSettings.BaudRate)
    err := yoSerial.Open(appSettings.SerialPort)
    if err != nil {
        log.Fatal(err)
    }

    fb := vid.ReadFrameAsJpeg("./sample_data/in1.mp4", 5, 128, 32)

    // tui.RunTui() 

    n, err := yoSerial.SendBkstMessage(2000, fb)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("Sent message, number of bytes sent:", n)
    }
}
