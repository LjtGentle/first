package fire

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Poker 存放 文件中比较数据的结构体
type poker struct {
	Alice  string `json:"alice"`
	Bob    string `json:"bob"`
	Result int    `json:"result"`
}

// Match 用于 存放读取文件的json格式数据
type Match struct {
	Matches []poker `json:"matches"`
}

//  CardCom用于同类型比较的函数传递参数
type cardCom struct {
	cardSizeMap1 map[byte]int
	cardSizeMap2 map[byte]int
	max1, max2   byte
}

// SizeTranByte 对面值转译整对应的byte值  -- 方便大小比较
func SizeTranByte(card byte) (res byte) {

	switch card {
	case 50:
		// 2
		res = 0x02
	case 51:
		res = 0x03
	case 52:
		res = 0x04
	case 53:
		res = 0x05
	case 54:
		res = 0x06
	case 55:
		res = 0x07
	case 56:
		res = 0x08
	case 57:
		res = 0x09
	case 84:
		res = 0x0A
	case 74:
		res = 0x0B
	case 81:
		res = 0x0C
	case 75:
		res = 0x0D
	case 65:
		res = 0x0E
	case 88:
		res = 0x10

	}
	return
}

// ReadFile 把数据从文件中读取出来 分别放在切片中返回
func ReadFile(filename string) (alices, bobs []string, results []int) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var matches Match
	err = json.Unmarshal(buf, &matches)
	if err != nil {
		panic(err)
	}

	alices = make([]string, len(matches.Matches))
	bobs = make([]string, len(matches.Matches))
	results = make([]int, len(matches.Matches))

	for k, v := range matches.Matches {
		alices[k] = v.Alice
		bobs[k] = v.Bob
		results[k] = v.Result
	}
	return
}

// JudgmentGroupNew 判断牌的类型
func JudgmentGroupNew(card []byte) (judeCardType uint8, cardSizeMap map[byte]int, resMax byte) {
	cardColorMap := make(map[byte]int, 5)
	cardSizeMap = make(map[byte]int, 5)
	// 扫描牌 分别放好大小，花色   --key放的是花色或是面值，--value放的是出现的次数
	for i, v := range card {
		if i%2 == 0 {
			// 存放大小
			if _, ok := cardSizeMap[v]; ok {
				cardSizeMap[v] ++
			} else {
				cardSizeMap[v] = 1
			}
			// 存放颜色
		} else {
			if _, ok := cardColorMap[v]; ok {
				cardColorMap[v] ++
			} else {
				cardColorMap[v] = 1
			}
		}
	}
	// 获取map的长度
	sizeLen := len(cardSizeMap)
	colorLen := len(cardColorMap)
	// 同花的时候，5个颜色一样，所以 colorLen = 1
	if colorLen > 1 {
		// 非同花
		switch sizeLen {
		case 4:
			// 一对
			judeCardType = 9
			return
		case 2: // 3带2  或是 4带1
			// 遍历map value
			for _, v := range cardSizeMap {
				if v == 4 {
					judeCardType = 3
					return
				}
			}
			judeCardType = 4
			return
		case 3:
			// 3条 或是 两对
			for _, v := range cardSizeMap {
				if v == 3 {
					judeCardType = 7
					return
				}
			}
			judeCardType = 8
			return
		case 5:
			// 单牌或是顺子
			isShun, max := IsShunZiNew(card)
			if isShun {
				resMax = max
				judeCardType = 6
				return
			}
			judeCardType = 10
			return

		}

	} else {
		// 同花 或是 同花顺
		isShun, max := IsShunZiNew(card)
		if isShun {
			resMax = max
			judeCardType = 1
		} else {
			judeCardType = 5
		}

	}

	return

}

// IsShunZiNew 判断是否是顺子 返回顺子的最大值和是否是顺子
func IsShunZiNew(card []byte) (shunZi bool, max byte) {
	shunZi = false
	saves := make([]byte, 14)
	// 把扑克牌放如slice中
	for i, v := range card {
		if i%2 == 0 {
			switch v {
			case 50:
				saves[1] = v
			case 51:
				saves[2] = v
			case 52:
				saves[3] = v
			case 53:
				saves[4] = v
			case 54:
				saves[5] = v
			case 55:
				saves[6] = v
			case 56:
				saves[7] = v
			case 57:
				saves[8] = v
			case 84:
				saves[9] = v
			case 74:
				saves[10] = v
			case 81:
				saves[11] = v
			case 75:
				saves[12] = v
			case 65:
				saves[13] = v
				saves[0] = v
			default:
				fmt.Println("无法解析的扑克牌", "card --v=", v)
			}
		}

	}
	// 判断数组是否连续 倒序遍历
	sum := 0
	for i := len(saves) - 1; i >= 0; i-- {
		// slice有值
		if saves[i] != 0x00 {
			sum++
		} else {
			sum = 0
		}
		// 5个连续
		if sum >= 5 {
			shunZi = true
			max = saves[i+4] // 返回顺子的最大值
			return
		}
	}
	return
}

// QuickSortByte 快排 对字节 逆序
func QuickSortByte(bs []byte) []byte {
	if len(bs) <= 1 {
		return bs
	}
	splitdata := bs[0]           // 第一个数据
	low := make([]byte, 0, 0)    // 比我小的数据
	hight := make([]byte, 0, 0)  // 比我大的数据
	mid := make([]byte, 0, 0)    // 与我一样大的数据
	mid = append(mid, splitdata) // 加入一个
	for i := 1; i < len(bs); i++ {
		if bs[i] > splitdata {
			low = append(low, bs[i])
		} else if bs[i] < splitdata {
			hight = append(hight, bs[i])
		} else {
			mid = append(mid, bs[i])
		}
	}
	low, hight = QuickSortByte(low), QuickSortByte(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

// SingleCardCompareSizeNew 同类型单牌比较 返回值是比较结果 0是平局 1是前面赢 2是后面赢
func (com *cardCom) SingleCardCompareSizeNew() (result int) {

	cardSizeSlice1 := make([]byte, len(com.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(com.cardSizeMap1))

	// 遍历map，把面值放到slice中
	i := 0
	for k := range com.cardSizeMap1 {
		cardSizeSlice1[i] = SizeTranByte(k)
		i++
	}
	i = 0
	for k := range com.cardSizeMap2 {
		cardSizeSlice2[i] = SizeTranByte(k)
		i++
	}

	// 比较5张牌的面值
	result = SingleCardSizeCom(5, cardSizeSlice1, cardSizeSlice2)

	return
}

// SingleCardSizeCom 对比单牌 大小0是平局 1是前面赢 2是后面赢
func SingleCardSizeCom(comLen int, cardSizeSlice1, cardSizeSlice2 []byte) (result int) {
	// 对传进来的slice逆序排序
	cardSizeSlice1 = QuickSortByte(cardSizeSlice1)
	cardSizeSlice2 = QuickSortByte(cardSizeSlice2)

	// 一个个对比
	for i := 0; i < comLen; i++ {
		if cardSizeSlice1[i] > cardSizeSlice2[i] {
			return 1
		} else if cardSizeSlice1[i] < cardSizeSlice2[i] {
			return 2
		}
	}
	return 0
}

// aPairComNew 同类型一对比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) aPairComNew() (result int) {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]byte, len(com.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(com.cardSizeMap1))
	// 用于存放对子的面值
	var pair1 byte
	var pair2 byte
	i := 0
	for k, v := range com.cardSizeMap1 {
		k = SizeTranByte(k) // 对牌子转译，才可以比较大小
		if v == 2 {
			pair1 = k
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
		if v == 2 {
			pair2 = SizeTranByte(k)

		} else {
			cardSizeSlice2[i] = SizeTranByte(k)
			i++
		}
	}
	// 先比较对子的大小
	if pair1 > pair2 {
		return 1
	} else if pair1 < pair2 {
		return 2
	} else {
		// 再单牌大小
		result = SingleCardSizeCom(3, cardSizeSlice1, cardSizeSlice2)
		return
	}

}

// twoPairComNew 同类型的两对比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) twoPairComNew() (result int) {
	// 用于存放两对的牌子
	pairs1 := make([]byte, 2)
	pairs2 := make([]byte, 2)
	// 用于存放单牌
	var val1 byte
	var val2 byte

	i := 0
	for k, v := range com.cardSizeMap1 {
		k = SizeTranByte(k) // 转译面值成可以比较的
		if v == 2 {
			pairs1[i] = k
			i++
		} else {
			val1 = k
		}
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
		k = SizeTranByte(k)
		if v == 2 {
			pairs2[i] = k
			i++
		} else {
			val2 = k
		}

	}
	// 比较对子的大小
	result = SingleCardSizeCom(2, pairs1, pairs2)
	if result != 0 {
		return
	}

	// 再比较单牌的大小
	if val1 > val2 {
		return 1
	} else if val1 < val2 {
		return 2
	} else {
		return 0
	}

}

// onlyThreeComNew 同类型的三条比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlyThreeComNew() (result int) {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]byte, len(com.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(com.cardSizeMap1))
	// 用于存放三条的面值
	var three1 byte
	var three2 byte
	i := 0
	for k, v := range com.cardSizeMap1 {
		k = SizeTranByte(k)
		cardSizeSlice1[i] = k
		if v == 3 {
			three1 = k
		} else {
			i++
		}
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
		k = SizeTranByte(k)
		cardSizeSlice2[i] = k
		if v == 3 {
			three2 = k
		} else {
			i++
		}
	}
	// 先比较三条的面值
	if three1 > three2 {
		return 1
	} else if three1 < three2 {
		return 2
	} else {
		// 再比较单牌的
		result = SingleCardSizeCom(2, cardSizeSlice1, cardSizeSlice2)
		return
	}
}

// onlyShunZiNew 同类型顺子的比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlyShunZiNew() (result int) {
	// max 是顺子的最大的牌，只要比较这张牌就行了
	if com.max1 > com.max2 {
		return 1
	} else if com.max1 < com.max2 {
		return 2
	}
	return 0
}

// onlySameFlowerNew 是同类型同花的比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlySameFlowerNew() (result int) {
	// 同类型同花 只要比较牌面值最大的，可以看着是单牌比较面值大小
	result = com.SingleCardCompareSizeNew()
	return
}

// straightFlushNew 同类型同花顺比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) straightFlushNew() (result int) {
	// 同类型同花顺比较，可以看作顺子之间比较
	return com.onlyShunZiNew()
}

// fourComNew 同类型4条比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) fourComNew() (result int) {
	// 存放四条的面值
	var four1 byte
	var four2 byte
	// 存放单牌的面值
	var val1 byte
	var val2 byte

	for k, v := range com.cardSizeMap1 {
		k = SizeTranByte(k) // 对面值转译成可以比较的
		if v == 4 {
			four1 = k
		} else {
			val1 = k
		}
	}
	for k, v := range com.cardSizeMap2 {
		k = SizeTranByte(k) // 对面值转译成可以比较的
		if v == 4 {
			four2 = k
		} else {
			val2 = k
		}
	}
	// 先比较4条大小
	if four1 > four2 {
		return 1
	} else if four1 < four2 {
		return 2
	} else {
		// 再比较单牌的大小
		if val1 > val2 {
			return 1
		} else if val1 < val2 {
			return 2
		} else {
			return 0
		}
	}
}

// threeAndTwoNew 同类型3带2比较  0是平局 1是前面赢 2是后面赢
func (com *cardCom) threeAndTwoNew() (result int) {
	// 存放3条的面值
	var three1 byte
	var three2 byte
	// 存放对子的面值
	var two1 byte
	var two2 byte
	for k, v := range com.cardSizeMap1 {
		if v == 3 {
			three1 = SizeTranByte(k)
		} else {
			two1 = SizeTranByte(k)
		}
	}
	for k, v := range com.cardSizeMap2 {
		if v == 3 {
			three2 = SizeTranByte(k)
		} else {
			two2 = SizeTranByte(k)
		}
	}
	// 先对比3条的面值
	if three1 > three2 {
		return 1
	} else if three1 < three2 {
		return 2
	} else {
		// 再对比对子的面值
		if two1 > two2 {
			return 1
		} else if two1 < two2 {
			return 2
		} else {
			return 0
		}
	}

}

// PokerMan 5张遍历判断 文件扑克牌的函数
func PokerMan() {
	file := "../resources/match_result.json"
	alices, bobs, results := ReadFile(file)
	t1 := time.Now()
	k := 0
	// 遍历全部对比
	for i := 0; i < len(alices); i++ {
		result := -1
		// 分牌型
		val1, cardSizesMap1, max1 := JudgmentGroupNew([]byte(alices[i]))
		val2, cardSizesMap2, max2 := JudgmentGroupNew([]byte(bobs[i]))
		if val1 < val2 {
			result = 1
		} else if val1 > val2 {
			result = 2
		} else {
			// 牌型相同的处理情况
			// ...
			cardCom := cardCom{
				cardSizeMap1: cardSizesMap1,
				cardSizeMap2: cardSizesMap2,
				max1:         max1,
				max2:         max2,
			}
			switch val1 {
			case 10:
				// 同类型下的单张大牌比较
				result = cardCom.SingleCardCompareSizeNew()
			case 9:
				// 同类型的一对
				result = cardCom.aPairComNew()
			case 8:
				// 同类型两对
				result = cardCom.twoPairComNew()
			case 7:
				// 同类型三条
				result = cardCom.onlyThreeComNew()
			case 6:
				// 同类型顺子
				result = cardCom.onlyShunZiNew()
			case 5:
				// 同类型同花
				result = cardCom.onlySameFlowerNew()
			case 4:
				// 同类型3带2
				result = cardCom.threeAndTwoNew()
			case 3:
				// 同类型四条
				result = cardCom.fourComNew()
			case 1: // 同类型同花顺
				result = cardCom.straightFlushNew()
			}

			// 最后比较结果
		}
		// 打印判断出错的信息
		if result != results[i] {
			k++
			fmt.Printf("[%#v]5张判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n", k, alices[i], bobs[i], results[i], result)
		}
	}
	t2 := time.Now()
	fmt.Println("time--->", t2.Sub(t1))

}
