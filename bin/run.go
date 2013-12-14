package main

import (
    "os"
    "gp"
    "fmt"
    "math"
)

func test_inversion() {
    m := gp.NewMatrix()
    m.SetValue(0,0,1)
    m.SetValue(0,1,0)
    m.SetValue(1,0,0)
    m.SetValue(1,1,1)
    u := gp.NewVector()
    u.SetValue(0, 2)
    u.SetValue(1, 2)
    v := gp.ApproximateInversion(m, u, 1e-8, 2)
    fmt.Printf("%f %f\n", v.GetValue(0), v.GetValue(1))
}

func output_XY(X *[]float64, Y *[]float64, noise float64, outf string) {
    fo, _ := os.Create(outf)
    for i := 0; i < len(*X); i++ {
        fo.WriteString(fmt.Sprintf("%f,%f,%f,%f\n", (*X)[i], (*Y)[i], (*Y)[i]-noise, (*Y)[i]+noise))
    }
    fo.Close()
}

func array2vector(a *[]float64) (*gp.Vector) {
    ret := gp.NewVector()
    for i := int64(0); i < int64(len((*a))); i++ {
        ret.SetValue(i, (*a)[i])
    }
    return ret
}

func array2vectorlist(a *[]float64) ([]*gp.Vector) {
    ret := make([](*gp.Vector), len(*a))
    for i := int64(0); i < int64(len(*a)); i++ {
        v := gp.NewVector()
        v.SetValue(0, (*a)[i])
        ret[i] = v
    }
    return ret
}

func test_gp() {
    min := -2.0
    max := 2.0
    period := 4.0
    amp := 1.0
    noise := 0.1
    count := 200

    radius := 0.05
    camp := 5.0

    x_pred := 1.0

    cf := gp.CovSEARD{}
    radiuses := make(map[int64]float64)
    radiuses[0] = radius
    cf.Init(radiuses, camp)

    X, Y := gp.GenNoisySin(min, max, period, amp, noise, count)
    for i := 0; i < len(*X); i ++ {
        fmt.Printf("%f -> %f\n", (*X)[i], (*Y)[i])
    }
    t := array2vector(Y)
    Xv := array2vectorlist(X)
    cov := gp.CovMatrix(Xv, cf.Cov)

    C_inv_t := gp.ApproximateInversion(cov, t, 1e-8, int64(count))

    x := gp.NewVector()
    x.SetValue(0, x_pred)

    // Get mean
    k := gp.CovVector(Xv, x, cf.Cov)
    pred := k.Dot(C_inv_t)

    // Get std
    C_inv_k := gp.ApproximateInversion(cov, k, 1e-6, int64(count))
    std := math.Sqrt(cf.Cov(x.GetData(), x.GetData()) - k.Dot(C_inv_k))

    fmt.Printf("%v+-%v\n", pred, std)

    // Output
    output_XY(X, Y, noise, "train.csv")
    tmp1 := make([]float64, 1)
    tmp1[0] = x.GetValue(0)
    tmp2 := make([]float64, 1)
    tmp2[0] = pred
    output_XY(&tmp1, &tmp2, std, "predict.csv")
}

func main() {
    test_gp()
}