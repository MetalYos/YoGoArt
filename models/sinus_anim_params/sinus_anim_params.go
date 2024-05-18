package sinus_anim_params

import (
    bu "yogoart/utils/bit_utils"
)

type SinusAnimParams struct {
    NumSinus uint32
    MinOffset uint32
    MaxOffset uint32
    MinVelocity uint32
    MaxVelocity uint32
}

func (params *SinusAnimParams) ToUint8Array() []uint8 {
    res := bu.Uint32ToUint8Array(params.NumSinus) 
    res = append(res, bu.Uint32ToUint8Array(params.MinOffset)...)
    res = append(res, bu.Uint32ToUint8Array(params.MaxOffset)...)
    res = append(res, bu.Uint32ToUint8Array(params.MinVelocity)...)
    res = append(res, bu.Uint32ToUint8Array(params.MaxVelocity)...)

    return res
}

func NewSinusAnimParams(numSinusWaves uint32, offsetY []uint32, velocity []uint32) *SinusAnimParams {
    params := &SinusAnimParams{
        NumSinus: numSinusWaves, 
        MinOffset: offsetY[0],
        MaxOffset: offsetY[1],
        MinVelocity: velocity[0],
        MaxVelocity: velocity[1],
    }
    return params
}

func NewSinusAnimParamsDefaults() *SinusAnimParams {
    params := &SinusAnimParams{
        NumSinus: 5,
        MinOffset: 50,
        MaxOffset: 90,
        MinVelocity: 3,
        MaxVelocity: 10,
    }
    return params
}

