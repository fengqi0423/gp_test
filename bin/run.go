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

func output_XY(X *[]float64, Y *[]float64, noise *[]float64, outf string) {
    fo, _ := os.Create(outf)
    for i := 0; i < len(*X); i++ {
        fo.WriteString(fmt.Sprintf("%f,%f,%f,%f\n", (*X)[i], (*Y)[i], (*Y)[i]-(*noise)[i], (*Y)[i]+(*noise)[i]))
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
    min := -5.0
    max := 5.0
    period := 4.0
    amp := 1.0
    noise := 0.1
    count := 100

    radius := 0.2
    camp := 40.0

    cf := gp.CovSEARD{}
    radiuses := make(map[int64]float64)
    radiuses[0] = radius
    cf.Init(radiuses, camp)

    X, Y := gp.GenNoisySin(min, max, period, amp, noise, count)

    t := array2vector(Y)
    Xv := array2vectorlist(X)
    cov := gp.CovMatrix(Xv, cf.Cov)

    C_inv_t := gp.ApproximateInversion(cov, t, 1e-8, int64(count))

    x_pred := make([]float64, 13)
    m_Pred := make([]float64, 13)
    s_Pred := make([]float64, 13)
    for i := -6.0; i <= 6.0; i = i + 1.0 {
        x := gp.NewVector()
        x.SetValue(0, i)

        // Get mean
        k := gp.CovVector(Xv, x, cf.Cov)
        pred := k.Dot(C_inv_t)

        // Get std
        C_inv_k := gp.ApproximateInversion(cov, k, 1e-8, int64(count))
        std := math.Sqrt(cf.Cov(x.GetData(), x.GetData()) - k.Dot(C_inv_k))

        x_pred[int64(i)+6] = i
        m_Pred[int64(i)+6] = pred
        s_Pred[int64(i)+6] = std 
    }

    // Output
    Noise := make([]float64, len(*X))
    for i := 0; i < len(Noise); i++ {
        Noise[i] = noise
    }
    output_XY(X, Y, &Noise, "train.csv")
    output_XY(&x_pred, &m_Pred, &s_Pred, "predict.csv")
}

func main() {
    test_gp()
}