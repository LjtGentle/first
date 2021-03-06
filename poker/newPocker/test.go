package main

import "fmt"

type ValueCounter uint64


// cards := []byte{0x02, 0x13, 0x04, 0x05, 0x06}
func (this *ValueCounter) Set(cards []byte) {
	//*this = 0
	l := uint8(len(cards))
	for i := uint8(0); i < l; i++ {
		card := (cards[i] & 0xF) * 3   //面值*3  等价于 x= x+(x<<1)
		fmt.Println("card==",card)
		count := ((*this >> card) & 0x07) + 1 //
		fmt.Println("count==",count)
		*this &= (^(0x07 << card))
		*this |= (count << card)

	}
}

func (this *ValueCounter) Get(card byte) uint8 {
	return uint8((*this >> (card & 0xF * 3 )) & 0x07)
}



func main() {
	// var CARDS = []byte{
	// 	//方块
	// 	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
	// 	//梅花
	// 	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
	// 	//红桃
	// 	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E,
	// 	//黑桃
	// 	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E,
	// }
	//fmt.Println("string(CARDS)",string(CARDS))

	var v ValueCounter
	cards := []byte{0x02, 0x13, 0x04, 0x05, 0x06}
	v.Set(cards)
	// fmt.Println("v=",v)
	//
	// fmt.Println("=====>",0x07>>0x02)
	//
	// fmt.Println("=====>",((0>>00001001)&0x07)+1)
	//
	// fmt.Println("---->",0x07 << (cards[0] & 0xF) * 3)
	//fmt.Println("int(0xf)=",int(0x0f))

	fmt.Println("0x07<<0x06=",0x07<<0x06)
}

