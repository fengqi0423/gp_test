package gp

import(
    "math"
    "math/rand"
    "time"
)

func GenNoisySin(min float64, max float64, period float64, amp float64, noise float64, count int) (*[]float64, *[]float64) {
    rand.Seed(time.Now().Unix())
    interval := (max - min) / float64(count)
    X := make([]float64, count)
    Y := make([]float64, count)
    for i := 0; i < count; i++ {
        x := min + interval * float64(i) + 0.5*interval
        X[i] = x
        Y[i] = math.Sin((x-min)*2*math.Pi/period) * amp + rand.NormFloat64()*noise
    }
    return &X, &Y
}