package yoserial

import (
	"errors"
	"log"
	ymsg "yogoart/models/yoart_msg"

	"go.bug.st/serial"
)

type YoSerial struct {
    Port serial.Port
    Baudrate uint32
    IsOpen bool
}

func NewYoSerial(baudrate uint32) *YoSerial {
    serialCtrl := &YoSerial {
        Port: nil,
        Baudrate: baudrate,
        IsOpen: false,
    }
    return serialCtrl
}

func GetYoSerialPorts() []string {
    ports, err := serial.GetPortsList()
    if err != nil {
        return make([]string, 0)
    }

    return ports
}

func (serialCtrl *YoSerial) Open(port string) error {
    mode := &serial.Mode{
        BaudRate: int(serialCtrl.Baudrate),
        Parity: serial.NoParity,
        DataBits: 8,
        StopBits: serial.OneStopBit,
    }

    openedPort, err := serial.Open(port, mode)
    if err != nil {
        serialCtrl.IsOpen = false
        return err
    }
    
    serialCtrl.Port = openedPort
    serialCtrl.IsOpen = true
    return nil
}

func (serialCtrl *YoSerial) Close() {
    if !serialCtrl.IsOpen {
        return
    }

    serialCtrl.Port.Close()
}

func (serialCtrl *YoSerial) sendYoartMsg(msg *ymsg.YoartMsg) (int, error) {
    if !serialCtrl.IsOpen {
        return 0, errors.New("Serial port is not open!")
    }

    bs := make([]byte, 1)
    bs[0] = 0x1
    bs = append(bs, msg.ToBytes()...)
    log.Println("sendYoartMsg: bs =", bs)
    return serialCtrl.Port.Write(bs)
}
func (serialCtrl *YoSerial) SendReqMessage(id uint32, data []uint8) (int, error) {
    msg := ymsg.NewYoartRequestMsg(id, data) 
    return serialCtrl.sendYoartMsg(msg)
}

func (serialCtrl *YoSerial) SendAnsMessage(id uint32, data []uint8) (int, error) {
    msg := ymsg.NewYoartAnswerMsg(id, data) 
    return serialCtrl.sendYoartMsg(msg)
}

func (serialCtrl *YoSerial) SendBkstMessage(id uint32, data []uint8) (int, error) {
    msg := ymsg.NewYoartBroadcastMsg(id, data) 
    return serialCtrl.sendYoartMsg(msg)
}

