package main

import (
	"fmt"
)

/**
*
* 型アサーション
* Important!!
*
* interfaceに詰められた値から型を保持して取り出す時に使う
* t := i.(T)
  interface i が具体的な型 T を保持し、基になる T の値を変数　t に代入することをやっている
  ※ i が Tを保持していない場合は、panic

  panicにさせないためには
  t, ok := i.(T)
  okにはboolが入る
*/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic example
	//f := i.(float64)
	//fmt.Println(f)
}
