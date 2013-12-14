#! /bin/sh

go run run.go
gnuplot -e "set datafile sep ','; plot 'train.csv' u 1:2 w l, 'train.csv' u 1:3 w l, 'train.csv' u 1:4 w l, 'predict.csv' u 1:2:3:4 with yerrorbars; pause -1"
