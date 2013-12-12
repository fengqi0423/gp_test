#! /bin/sh

go run run.go
gnuplot -e "set datafile sep ','; plot 'real.csv' u 1:2 w l, 'real.csv' u 1:3 w l, 'real.csv' u 1:4 w l, 'sample.csv' u 1:2; pause -1"
