package yoart_msg

import (
	"encoding/binary"
	"unsafe"
)

type YoartMsgTypes uint8

const (
	NoneMsgType YoartMsgTypes = iota
	AckMsgType
	NackMsgType
	ReqMsgType
    AnsMsgType
    BkstMsgType
    FileMsgType
)

type YoartMsgHeader struct {
    Type YoartMsgTypes
    Seq uint8
    Len uint16
    Id uint32
}

func (header *YoartMsgHeader) Size() uint32 {
    results := unsafe.Sizeof(header.Type)
    results += unsafe.Sizeof(header.Seq)
    results += unsafe.Sizeof(header.Len)
    results += unsafe.Sizeof(header.Id)
    return uint32(results)
}

func (header *YoartMsgHeader) ToBytes() []byte {
    bs := make([]byte, 0)

    bs = append(bs, uint8(header.Type))
    bs = append(bs, header.Seq)
    bs = binary.LittleEndian.AppendUint16(bs, header.Len)
    bs = binary.LittleEndian.AppendUint32(bs, header.Id)
    return bs
}

func (header *YoartMsgHeader) calculateChecksum() uint32 {
    checksum := uint32(0)
    checksum += uint32(header.Type)
    checksum += uint32(header.Seq)
    checksum += uint32((header.Len & 0xFF) + ((header.Len >> 8) & 0xFF))
    checksum += uint32((header.Id & 0xFF) + ((header.Id >> 8) & 0xFF) + ((header.Id >> 16) & 0xFF) + ((header.Id >> 24) & 0xFF))
    return checksum
}

func newYoartMsgHeader(msgType YoartMsgTypes, seq uint8, length uint16, id uint32) *YoartMsgHeader {
    header := &YoartMsgHeader {
        Type: msgType,
        Seq: seq,
        Len: length,
        Id: id,
    }
    return header
}

type YoartMsgFooter struct {
    Checksum uint32
}

func (footer *YoartMsgFooter) Size() uint32 {
    result := unsafe.Sizeof(footer.Checksum)
    return uint32(result)
}

func (footer *YoartMsgFooter) ToBytes() []byte {
    bs := make([]byte, 0)
    bs = binary.LittleEndian.AppendUint32(bs, footer.Checksum)
    return bs
}

func newYoartMsgFooter() *YoartMsgFooter {
    footer := &YoartMsgFooter {
        Checksum: 0,
    }
    return footer
}

type YoartMsg struct {
    Header YoartMsgHeader
    Data []uint8
    Footer YoartMsgFooter
}

func (msg *YoartMsg) Size() uint32 {
    results := msg.Header.Size()
    results += uint32(len(msg.Data))
    results += msg.Footer.Size()
    return results
}

func (msg *YoartMsg) ToBytes() []byte {
    bs := msg.Header.ToBytes()
    bs = append(bs, msg.Data...)
    bs = append(bs, msg.Footer.ToBytes()...)
    return bs
}

func (msg *YoartMsg) calculateChecksum() uint32 {
    checksum := uint32(0)
    checksum += msg.Header.calculateChecksum()
    for _, val := range msg.Data {
        checksum += uint32(val)
    }
    return checksum
}

func NewYoartMsg(msgType YoartMsgTypes, seq uint8, id uint32, data []uint8) *YoartMsg {
    msg := &YoartMsg {
        Header: *newYoartMsgHeader(msgType, seq, uint16(len(data)), id),   
        Data: data,
    }
    msg.Footer.Checksum = msg.calculateChecksum()
    return msg
}

func NewYoartRequestMsg(id uint32, data []uint8) *YoartMsg {
    return NewYoartMsg(ReqMsgType, 1, id, data)
}

func NewYoartAnswerMsg(id uint32, data []uint8) *YoartMsg {
    return NewYoartMsg(AnsMsgType, 1, id, data)
}

func NewYoartBroadcastMsg(id uint32, data []uint8) *YoartMsg {
    return NewYoartMsg(BkstMsgType, 1, id, data)
}

func NewYoartFileMsg(seq uint8, data []uint8) *YoartMsg {
    return NewYoartMsg(FileMsgType, seq, 0, data)
}

