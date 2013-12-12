package gp

import (
    "math"
)

/*
Covariance functions
Function names are influenced by Rasmussen and Nickisch's gpml
*/

type CovSEARD struct {
    // Squared error covariance function
    // ARD = auto relevance detection, and here indicates there is a scaling/radius factor per dimension
    Radiuses map[int64]float64
    Amp float64
}

func (cov_func *CovSEARD) Init(radiuses map[int64]float64, amp float64) {
    cov_func.Radiuses = radiuses
    cov_func.Amp = amp
}

func (cov_func *CovSEARD) Cov(x1 map[int64]float64, x2 map[int64]float64) float64 {
    ret := 0.0
    tmp := 0.0
    for key, r := range cov_func.Radiuses {
        v1, _ := x1[key]
        v2, _ := x2[key]
        tmp = (v1-v2)/r
        ret += tmp * tmp
    }
    ret = cov_func.Amp * math.Exp(-ret)
    return ret
}