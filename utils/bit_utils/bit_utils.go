package bit_utils

func BoolToUint8(val bool) uint8 {
    if val {
        return 1
    }
    return 0
}

func Uint32ToUint8Array(val uint32) []uint8 {
    res := make([]uint8, 4)
    res[0] = uint8(val & 0xFF)
    res[1] = uint8((val >> 8) & 0xFF)
    res[2] = uint8((val >> 16) & 0xFF)
    res[3] = uint8((val >> 24) & 0xFF)
    return res
}
