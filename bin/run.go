package main

import (
    "gp"
    "fmt"
)

func main() {
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