package main
/*
import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle) //类型断言
	if ok {
		whistle.MakeSound()
	}
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Toyco Canary"))

	var x interface{}
	x = "pprof.cn"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
*/