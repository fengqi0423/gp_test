package main

import (
    "os"
    "gp"
    "fmt"
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

func test_data_gen() {
    min := -2.0
    max := 2.0
    period := 4.0
    amp := 10.0 
    noise := 2.0
    count := 100
    sample_count := 20

    X, Y := gp.GenNoisySin(min, max, period, amp, 0.0, count)
    fo, _ := os.Create("real.csv")
    for i := 0; i < count; i++ {
        fo.WriteString(fmt.Sprintf("%f,%f,%f,%f\n", (*X)[i], (*Y)[i], (*Y)[i]-noise, (*Y)[i]+noise))
    }
    fo.Close()

    X, Y = gp.GenNoisySin(min, max, period, amp, noise, sample_count)
    fo, _ = os.Create("sample.csv")
    for i := 0; i < sample_count; i++ {
        fo.WriteString(fmt.Sprintf("%f,%f\n", (*X)[i], (*Y)[i]))
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
    noise := 0.2
    count := 100

    radius := 0.1
    camp := 1.0

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
/*
    for i := int64(0); i < int64(count); i++ {
        for j := int64(0); j < int64(count); j++ {
            fmt.Printf("%3.3e  ", cov.GetValue(i,j))
        }
        fmt.Println()
    }
    */
    C_inv_t := gp.ApproximateInversion(cov, t, 1e-8, 1)

    x := gp.NewVector()
    x.SetValue(0, 1.0)
    k := gp.CovVector(Xv, x, cf.Cov)
    pred := k.Dot(C_inv_t)

    fmt.Printf("%v\n", pred)
}

func main() {
    test_gp()
}