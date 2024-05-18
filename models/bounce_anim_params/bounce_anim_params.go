package bounce_anim_params

import (
    bu "yogoart/utils/bit_utils"
)


/*
struct BounceParams {
    int NumBalls = BOUNCE_ANIM_NUM_BALLS; 
    int MinDiameter = BOUNCE_ANIM_MIN_DIAM;
    int MaxDiameter = BOUNCE_ANIM_MAX_DIAM;
    int MinVelocity = BOUNCE_ANIM_MIN_VEL;
    int MaxVelocity = BOUNCE_ANIM_MAX_VEL;
    bool IsSelfColliding = false;
};
*/


type BounceAnimParams struct {
   NumBalls uint32    
   MinDiameter uint32    
   MaxDiameter uint32    
   MinVelocity uint32    
   MaxVelocity uint32    
   IsSelfColliding bool    
}

func (params *BounceAnimParams) ToUint8Array() []uint8 {
    res := bu.Uint32ToUint8Array(params.NumBalls) 
    res = append(res, bu.Uint32ToUint8Array(params.MinDiameter)...)
    res = append(res, bu.Uint32ToUint8Array(params.MaxDiameter)...)
    res = append(res, bu.Uint32ToUint8Array(params.MinVelocity)...)
    res = append(res, bu.Uint32ToUint8Array(params.MaxVelocity)...)
    res = append(res, bu.BoolToUint8(params.IsSelfColliding))

    return res
}

func NewBounceAnimParams(numBalls uint32, diameter []uint32, velocity []uint32, isSelfColliding bool) *BounceAnimParams {
    params := &BounceAnimParams{
        NumBalls: numBalls, 
        MinDiameter: diameter[0],
        MaxDiameter: diameter[1],
        MinVelocity: velocity[0],
        MaxVelocity: velocity[1],
        IsSelfColliding: isSelfColliding,
    }
    return params
}

func NewBounceAnimParamsDefaults() *BounceAnimParams {
    params := &BounceAnimParams{
        NumBalls: 5,
        MinDiameter: 2,
        MaxDiameter: 20,
        MinVelocity: 2,
        MaxVelocity: 15,
        IsSelfColliding: false,
    }
    return params
}

