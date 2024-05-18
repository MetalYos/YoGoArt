package settings

import (
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
    SerialPort string `json:"port_name"`
    BaudRate uint32 `json:"baud_rate"`
}

func LoadSettings(filename string) *Settings {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Printf("Can't open file! %s", err)
        return nil
    }

    settings := &Settings{}
    err = json.Unmarshal(data, settings)
    if err != nil {
        log.Println(err)
        return nil
    }

    if settings.SerialPort == "" {
        log.Fatalln("Serial Port is a mandatory field in the settings JSON file!")
        return nil
    }

    if settings.BaudRate == 0 {
        log.Println("BaudRate was not in the settings JSON file! Setting it to 115200")
        settings.BaudRate = 115200
    }

    log.Println("Loaded settings successfully!", settings)
    return settings
}
