#! /bin/bash

printf "slice"
sliceTime=$(time go run slice.go)
printf "%s" "$sliceTime"
printf "%s\n" "---------------"

printf "slice with references"
sliceTime=$(time go run slice_star.go)
printf "%s" "$sliceTime"
printf "%s\n" "---------------"


printf "map with references"
mapTime=$(time go run map_star.go)
printf "%s" "$mapTime"
printf "%s\n" "---------------"

printf "map without references"
mapTime=$(time go run map_no_star.go)
printf "%s" "$mapTime"
printf "%s\n" "---------------"

printf "map split without references"
mapTime=$(time go run map_no_star.go)
printf "%s" "$mapTime"
printf "%s\n" "---------------"