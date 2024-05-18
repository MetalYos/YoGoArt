package main

import (
	"log"
    settings "yogoart/utils/settings"
    ser "yogoart/yoserial"
    "yogoart/tui"
)

func main() {
    appSettings := settings.LoadSettings("settings.json")

    yoSerial := ser.NewYoSerial(appSettings.BaudRate)
    err := yoSerial.Open(appSettings.SerialPort)
    if err != nil {
        log.Fatal(err)
    }

    tui.RunTui() 

    /*
    n, err := yoSerial.SendBkstMessage(1002, params.ToUint8Array())
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("Sent message, number of bytes sent:", n)
    }
    */
}
