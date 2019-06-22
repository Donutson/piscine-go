package piscine

import "math"

func CollatzCountdown(start int) int {
	count:=0
	for start!=1{
		if start%2==0{
			start=start/2
			count++	
		}
		start=start*3+1
		count++
		if start> math.MaxInt32{
			return count
		}
	}
	return count
}
