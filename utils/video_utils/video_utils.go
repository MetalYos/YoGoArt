package video_utils

import (
    "bytes"
	"fmt"
	"os"

    ffmpeg "github.com/u2takey/ffmpeg-go"
    "github.com/disintegration/imaging"
)

type RGBA struct {
    R uint8
    G uint8
    B uint8
    A uint8
}

func NewRGBA(r, g, b, a uint8) RGBA {
    return RGBA {
        R: r,
        G: g,
        B: b,
        A: a,
    }
}

func ReadFrameAsJpeg(inFileName string, frameNum int, frameWidth int, frameHeight int) []uint8 {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

    img, err := imaging.Decode(buf)
    if err != nil {
        panic(err)
    }
    thumb := imaging.Thumbnail(img, frameWidth, frameHeight, imaging.Lanczos)
    thumb = imaging.Grayscale(thumb)
    thumb = imaging.AdjustContrast(thumb, 20)

    frameBuffer := make([]uint8, 0)
    currentByte := uint8(0)
    pixIndex := 0
    for k := 0; k < len(thumb.Pix); k += 4 {
        bitIndex := pixIndex % 8
        pixel := NewRGBA(thumb.Pix[k], thumb.Pix[k + 1], thumb.Pix[k + 2], thumb.Pix[k + 3])

        // If a pixel is visible and it's grayscale value is above 50 set its bit to 1
        // Otherwise set it to 0
        if pixel.A > 127  {
            if pixel.R > 50 {
                currentByte |= (1 << bitIndex) 
            }
        }

        if bitIndex == 7 {
            frameBuffer = append(frameBuffer, currentByte)
            currentByte = 0
        }

        pixIndex++
    }
    fmt.Println("Finished decoding frame", frameNum)
    return frameBuffer
}
