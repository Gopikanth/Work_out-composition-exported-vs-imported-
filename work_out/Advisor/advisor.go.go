//here we created a simple advicing project to buy a product next time or not using -"composition"

package main

import "fmt"

type shop struct {
	what_item string
	price     int
	profit    int
	loss      int
}
type advice struct {
	advicing string
	shop     shop
}

func (s shop) show() {
	fmt.Println("showing details like what item", s.what_item)
	fmt.Println("what is the actual price", s.price)
	fmt.Println("profit", s.profit)
	fmt.Println("loss", s.loss)

}
func (a advice) showing() {
	fmt.Println("my advice is ", a.advicing)
	a.shop.show()
}
func main() {
	vegitable := shop{
		what_item: "tomato",
		price:     60,
		profit:    5,
		loss:      0,
	}
	if vegitable.loss < vegitable.profit {
		fmt.Println("you can buy this item again :-", vegitable.what_item)

	} else {
		fmt.Println("Try to avoid")
	}
	snack := shop{
		what_item: "potato",
		price:     60,
		profit:    0,
		loss:      4,
	}
	if snack.loss < snack.profit {
		fmt.Println("you can buy this item again :-", snack.what_item)

	} else {
		fmt.Println("Try to avoid :-", snack.what_item)
	}

}
