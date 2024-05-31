package main

import (
	"errors"
	"image"
	"log"
	settings "yogoart/utils/settings"
	ser "yogoart/yoserial"

	// "yogoart/tui"

    "time"
	vidio "github.com/AlexEidt/Vidio"
)

func SendVideo(yoSerial *ser.YoSerial, filename string, width int, height int, fps int) {
    vid, err := vidio.NewVideo(filename)
    if err != nil {
        panic(err)
    }

    if (vid.Width() != width) || (vid.Height() != height) {
        panic(errors.New("Video file size is different then frame buffer to send size!"))
    }

    img := image.NewRGBA(image.Rect(0, 0, vid.Width(), vid.Height()))
    vid.SetFrameBuffer(img.Pix)


    frameTimeUsec := uint32((1.0 / float32(fps)) * 1000 * 1000)
    log.Println("frameTimeUsec =", frameTimeUsec)

    fbSend := make([]uint8, width * height / 8)
    startTime := time.Now()
    for vid.Read() {
        for {
            elapsedTime := time.Since(startTime).Microseconds()
            if elapsedTime > int64(frameTimeUsec) {
                log.Println("Send frame after elapsedTime =", elapsedTime)
                break
            }
        }

        fbByte := uint8(0)
        fbBit := 0
        fbPixIndex := 0
        for pix := 0; pix < len(img.Pix); pix += 4 {
            R := img.Pix[pix]
            G := img.Pix[pix + 1]
            B := img.Pix[pix + 2]
            if (R < 50) && (G > 200) && (B < 50) {
                fbByte |= (1 << fbBit)
            }
            
            fbBit++
            if fbBit == 8 {
                fbSend[fbPixIndex] = fbByte

                fbBit = 0
                fbByte = 0
                fbPixIndex++
            }
        }

        // Send Frame
        yoSerial.SendBkstMessage(2000, fbSend)

        startTime = time.Now()
    }
}

func main() {
    appSettings := settings.LoadSettings("settings.json")

    yoSerial := ser.NewYoSerial(appSettings.BaudRate)
    err := yoSerial.Open(appSettings.SerialPort)
    if err != nil {
        log.Fatal(err)
    }

    // tui.RunTui() 

    SendVideo(yoSerial, "sample_data/in1.mp4", 128, 32, 10)
}
