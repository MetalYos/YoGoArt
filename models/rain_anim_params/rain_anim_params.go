package rain_anim_params

import (
    bu "yogoart/utils/bit_utils"
)

type RainAnimParams struct {
    NumDropletSources uint32
    RainHeight uint32
    MinVelocity uint32
    MaxVelocity uint32
    DropletMinLen uint32
    DropletMaxLen uint32
    IsSwapAxes bool
}

func (params *RainAnimParams) ToUint8Array() []uint8 {
    res := bu.Uint32ToUint8Array(params.NumDropletSources) 
    res = append(res, bu.Uint32ToUint8Array(params.RainHeight)...)
    res = append(res, bu.Uint32ToUint8Array(params.MinVelocity)...)
    res = append(res, bu.Uint32ToUint8Array(params.MaxVelocity)...)
    res = append(res, bu.Uint32ToUint8Array(params.DropletMinLen)...)
    res = append(res, bu.Uint32ToUint8Array(params.DropletMaxLen)...)
    res = append(res, bu.BoolToUint8(params.IsSwapAxes))

    return res
}

func NewRainAnimParams(numDropletSources uint32, rainHeight uint32, velocity []uint32, dropletLen []uint32, isSwapAxes bool) *RainAnimParams {
    params := &RainAnimParams{
        NumDropletSources: numDropletSources, 
        RainHeight: rainHeight,
        MinVelocity: velocity[0],
        MaxVelocity: velocity[1],
        DropletMinLen: dropletLen[0],
        DropletMaxLen: dropletLen[1],
        IsSwapAxes: isSwapAxes,
    }
    return params
}

func NewRainAnimParamsDefaults() *RainAnimParams {
    params := &RainAnimParams{
        NumDropletSources: 128,
        RainHeight: 32,
        MinVelocity: 1,
        MaxVelocity: 10,
        DropletMinLen: 1,
        DropletMaxLen: 5,
        IsSwapAxes: false,
    }
    return params
}

