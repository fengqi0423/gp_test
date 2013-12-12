package gp

type CovFunc func(map[int64]float64, map[int64]float64)float64

func CovMatrix(X []*Vector, cov_func CovFunc) (*Matrix) {
    l := int64(len(X))
    ret := NewMatrix()
    for i := int64(0); i < l; i++ {
        for j := i; j < l; j++ {
            c := cov_func(X[i].data, X[j].data)
            ret.SetValue(i, j, c)
            ret.SetValue(j, i, c)
        }
    }
    return ret
}

func CovVector(X []*Vector, y *Vector, cov_func CovFunc) (*Vector) {
    l := int64(len(X))
    ret := NewVector()
    for i := int64(0); i < l; i++ {
        ret.SetValue(i, cov_func(X[i].data, y.data))
    }
    return ret
}