package main

import(
	"fmt"
)

func GenDisplaceFn(
	acc, initVel, initDisp float64,
) func(float64) float64 {

	return func (t float64) float64  {
		return 0.5 * acc * t*t + initVel*t + initDisp
	}
}

func main()  {
	var acc,initVel,initDisp,time float64

	fmt.Print(
		"Enter respectively \nacceleration, \ninitial velocity, \ninitial displacement, \ntime\n")
	fmt.Scan(&acc, &initVel, &initDisp, &time)

	fn:= GenDisplaceFn(acc,initVel,initDisp)

	fmt.Println("verdade: ",fn(time))
}