package seven

import (
	"NewPocker/fire"
	"fmt"
	"time"
)

// Seven 调用同类型比较函数们的参数传入
type Seven struct {
	// 存放面值的map
	cardSizeMap1, cardSizeMap2   map[byte]int
	cardColorMap1, cardColorMap2 map[byte]int
	max1, max2                   byte
	card1, card2                 []byte
}

// IsShunZi 判断是不是顺子，并且返回顺子的最大值  传进来的已经转译好的面值
func IsShunZi(seq []byte) (flag bool, max byte) {
	flag = false
	saves := make([]byte, 14)
	// 遍历序列， 存放入对应的序列
	for _, v := range seq {
		switch v {
		case 0x02:
			saves[1] = v
		case 0x03:
			saves[2] = v
		case 0x04:
			saves[3] = v
		case 0x05:
			saves[4] = v
		case 0x06:
			saves[5] = v
		case 0x07:
			saves[6] = v
		case 0x08:
			saves[7] = v
		case 0x09:
			saves[8] = v
		case 0x0A:
			saves[9] = v
		case 0x0B:
			saves[10] = v
		case 0x0C:
			saves[11] = v
		case 0x0D:
			saves[12] = v
		case 0x0E:
			saves[13] = v
			saves[0] = v
		default:
			fmt.Println("无法解析的扑克牌", "card --v=", v)
		}

	}
	// 判断数组是否连续
	sum := 0
	for i := len(saves) - 1; i >= 0; i-- {
		if saves[i] != 0x00 {
			// slice有值
			sum++
		} else {
			// 没值重置
			sum = 0
		}
		// 判断到有连续5个
		if sum >= 5 {
			flag = true
			max = saves[i+4] // 返回顺子最大值
			return
		}
	}
	return
}

// JudgmentGroup 判断牌型
func JudgmentGroup(card []byte) (cardType uint8, cardSizeMap, cardColorMap map[byte]int, resMax byte) {
	cardColorMap = make(map[byte]int, 7)
	cardSizeMap = make(map[byte]int, 7)
	// 扫描牌 分别放好大小，花色
	for i, v := range card {
		if i%2 == 0 {
			// 大小 判断map是否有值，也就是之前是否出现过  最终 key是面值，value是该面值出现的次数
			if _, ok := cardSizeMap[v]; ok {
				cardSizeMap[v] ++
			} else {
				cardSizeMap[v] = 1
			}
			// 颜色 同上
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
	flag := false
	for _, v := range cardColorMap {
		if v >= 5 { // 5个花色一样的
			flag = true
			break
		}
	}
	if flag {
		// 已经是同花
		// 然后判断是不是顺子
		seq := SameFlowerSeq(cardColorMap, card)
		isShun, max := IsShunZi(seq)
		if isShun {
			// 同花顺
			resMax = max
			cardType = 1
			return
		}
		// 单纯的同花
		cardType = 5
		return
	}
	// 然后再根据存放面值的map长度来判断类型
	switch sizeLen {
	case 7: // 不是顺子 就是7个单牌
		// 判断是不是顺子
		if isShun, max := fire.IsShunZiNew(card); isShun {
			cardType = 6 // 顺子的类型
			resMax = max
			return
		}
		cardType = 10 // 单牌的类型
		return
	case 6: // 1对 或是 顺子
		// 判断是不是顺子
		if isShun, max := fire.IsShunZiNew(card); isShun {
			resMax = max
			cardType = 6
			return
		}
		// 返回一对的类型
		cardType = 9
		return
	case 5: // 可以是顺子 两对 或是 3条
		// 顺子大先判断是不是顺子
		if isShun, max := fire.IsShunZiNew(card); isShun {
			resMax = max
			cardType = 6
			return
		}
		// cardSizeMap 的key是面值 value是该面值出现的次数
		// 然后判断是不是3条
		for _, v := range cardSizeMap {
			if v == 3 {
				cardType = 7 // 3条
				return
			}
		}
		// 2对
		cardType = 8
		return
	case 4: // 可以是 4条 3带2  两对（3个对子）
		for _, v := range cardSizeMap {
			if v == 4 {
				// 4条
				cardType = 3
				return
			} else if v == 3 {
				// 3条
				cardType = 4
				return
			}
		}
		// 剩下两对
		cardType = 8
		return
	case 3: // 4条（4条1对） 3带2（3条和3条， 3条和两对）
		for _, v := range cardSizeMap {
			if v == 4 {
				cardType = 3
				return
			}
		}
		cardType = 4
		return
	case 2: // 4条（4条和3条）
		cardType = 3
		return

	}

	return

}

// SingleCard 同类型单牌比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) SingleCard() (result int) {
	// 用于存放面值大小的slice
	cardSizeSlice1 := make([]byte, len(sevenCom.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenCom.cardSizeMap2))
	i := 0
	for k := range sevenCom.cardSizeMap1 {
		cardSizeSlice1[i] = fire.SizeTranByte(k)
		i++
	}
	i = 0
	for k := range sevenCom.cardSizeMap2 {
		cardSizeSlice2[i] = fire.SizeTranByte(k)
		i++
	}
	result = fire.SingleCardSizeCom(5, cardSizeSlice1, cardSizeSlice2)
	return
}

// APair 同类型一对的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) APair() (result int) {
	// 存放 单牌的面值的slice
	cardSizeSlice1 := make([]byte, len(sevenCom.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenCom.cardSizeMap1))
	// 存放对子的面值
	var val1 byte
	var val2 byte
	i := 0
	for k, v := range sevenCom.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			val1 = k
			continue
		}
		cardSizeSlice1[i] = k
		i++
	}
	i = 0
	for k, v := range sevenCom.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			val2 = k
			continue
		}
		cardSizeSlice2[i] = k
		i++
	}
	// 先对比对子的面值大小
	if val1 > val2 {
		return 1
	} else if val1 < val2 {
		return 2
	}
	// 然后对比各自单牌中最大的三种
	result = fire.SingleCardSizeCom(3, cardSizeSlice1, cardSizeSlice2)
	return
}

// TwoPair 同类型两对的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) TwoPair() (result int) {
	// 存放对子的slice
	pairs1 := make([]byte, 3)
	pairs2 := make([]byte, 3)
	// 存放单牌的slice
	vals1 := make([]byte, 3)
	vals2 := make([]byte, 3)

	j := 0
	i := 0
	for k, v := range sevenCom.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			pairs1[i] = k
			i++
		} else {
			vals1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range sevenCom.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			pairs2[i] = k
			i++
		} else {
			vals2[j] = k
			j++
		}
	}
	// 对对子序列排序  逆序
	pairs1 = fire.QuickSortByte(pairs1)
	pairs2 = fire.QuickSortByte(pairs2)

	// 对子比较 只是比较最大的两个，因为有可能是3对的
	for i := 0; i < 2; i++ {
		if pairs1[i] > pairs2[i] {
			return 1
		} else if pairs1[i] < pairs2[i] {
			return 2
		}
	}
	// 对剩余的单牌排序
	vals1 = fire.QuickSortByte(vals1)
	vals2 = fire.QuickSortByte(vals2)
	// 跟各自的对子slice的第三个比较， 因为可能存在第三个对子
	if vals1[0] < pairs1[2] {
		vals1[0] = pairs1[2]
	}
	if vals2[0] < pairs2[2] {
		vals2[0] = pairs2[2]
	}
	// 对于剩下的单牌比较，只比较最大的一个
	if vals1[0] > vals2[0] {
		return 1
	} else if vals1[0] < vals2[0] {
		return 2
	} else {
		return 0
	}
}

// OnlyThree 同类型3条比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) OnlyThree() (result int) {
	// 存放单牌的slice
	cardSizeSlice1 := make([]byte, len(sevenCom.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenCom.cardSizeMap2))
	// 存放3条的面值
	var three1 byte
	var three2 byte

	i := 0
	for k, v := range sevenCom.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v == 3 {
			three1 = k
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range sevenCom.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 3 {
			three2 = k
		} else {
			cardSizeSlice2[i] = k
			i++
		}
	}
	// 比较3条的面值大小
	if three1 > three2 {
		return 1
	} else if three1 < three2 {
		return 2
	} else {
		// 比较张单牌的大小
		result = fire.SingleCardSizeCom(2, cardSizeSlice1, cardSizeSlice2)
		return
	}

}

// OnlyShunZi 同类型顺子的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) OnlyShunZi() (result int) {
	max1 := fire.SizeTranByte(sevenCom.max1)
	max2 := fire.SizeTranByte(sevenCom.max2)
	if max1 > max2 {
		return 1
	} else if max1 < max2 {
		return 2
	} else {
		return 0
	}
}

// SameFlowerSeq 找到同花对应的面值序列
func SameFlowerSeq(cardColorMap map[byte]int, card []byte) (sizeSlice []byte) {
	// 存放同花的花色
	var color byte

	sliceLen := 0
	for k, v := range cardColorMap {
		if v >= 5 {
			color = k    // 记录颜色
			sliceLen = v // 颜色出现的次数
			break
		}
	}
	sizeSlice = make([]byte, sliceLen) // 大小为颜色出现的次数
	j := 0
	for i := 1; i < len(card); i += 2 {
		if card[i] == color {
			sizeSlice[j] = fire.SizeTranByte(card[i-1]) // 取颜色前一个 也就是面值
			j++
		}
	}
	return

}

// onlySameFlower 同类型同花的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) onlySameFlower() (result int) {
	// 得到同花对应的面值序列
	sizeSlice1 := SameFlowerSeq(sevenCom.cardColorMap1, sevenCom.card1)
	sizeSlice2 := SameFlowerSeq(sevenCom.cardColorMap2, sevenCom.card2)
	// 只对比各自最大的五个
	result = fire.SingleCardSizeCom(5, sizeSlice1, sizeSlice2)
	return
}

// ThreeAndTwo 同类型3带2的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) ThreeAndTwo() (result int) {
	// 存放 3条的面值
	threes1 := make([]byte, 2)
	threes2 := make([]byte, 2)
	// 存放对子的面值
	twos1 := make([]byte, 2)
	twos2 := make([]byte, 2)

	i := 0
	j := 0
	for k, v := range sevenCom.cardSizeMap1 {
		k = fire.SizeTranByte(k) // 对面值进行转译 才可以大小比较
		if v == 3 {
			threes1[i] = k
			i++
		} else if v == 2 {
			twos1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range sevenCom.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 3 {
			threes2[i] = k
			i++
		} else if v == 2 {
			twos2[j] = k
			j++
		}
	}

	// 对3条的序列进行倒序排序 可能出现2个3条的情况
	threes1 = fire.QuickSortByte(threes1)
	threes2 = fire.QuickSortByte(threes2)
	// 对比3条
	if threes1[0] > threes2[0] {
		return 1
	} else if threes1[0] < threes2[0] {
		return 2
	} else {
		// 对对子的序列进行倒序排序
		twos1 = fire.QuickSortByte(twos1)
		twos2 = fire.QuickSortByte(twos2)
		// 可能出现2个3条的情况
		if twos1[0] < threes1[1] {
			twos1[0] = threes1[1]
		}
		if twos2[0] < threes2[1] {
			twos2[0] = threes2[1]
		}
		// 对比对子
		if twos1[0] > twos2[0] {
			return 1
		} else if twos1[0] < twos2[0] {
			return 2
		} else {
			return 0
		}
	}
}

// FourCom 同类型4条的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) FourCom() (result int) {
	// 存放非4条的牌面值
	cardSizeSlice1 := make([]byte, len(sevenCom.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenCom.cardSizeMap2))
	// 存放4条的面值
	var four1 byte
	var four2 byte

	i := 0
	for k, v := range sevenCom.cardSizeMap1 {
		k = fire.SizeTranByte(k) // 面值转换才可以用于比较
		if v == 4 {
			four1 = k
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range sevenCom.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 4 {
			four2 = k
		} else {
			cardSizeSlice2[i] = k
			i++
		}
	}
	// 先比较4条的大小
	if four1 > four2 {
		return 1
	} else if four1 < four2 {
		return 2
	} else {
		// 再比较各自单牌中最大一张
		result = fire.SingleCardSizeCom(1, cardSizeSlice1, cardSizeSlice2)
		return
	}

}

// straightFlush 同类型同花顺的比较 返回值0是平局 1是前面赢 2是后面赢
func (sevenCom *Seven) straightFlush() (result int) {
	// 同类型同花顺比较可以看作是同类型顺子比较
	return sevenCom.OnlyShunZi()
}

// PokerMan 7张的主函数
func PokerMan() {
	file := "../resources/seven_cards_with_ghost.json"
	alices := make([]string, 1024)
	bobs := make([]string, 1024)
	results := make([]int, 1024)
	alices, bobs, results = fire.ReadFile(file)
	t1 := time.Now()

	k := 0
	for i := 0; i < len(alices); i++ {
		result := -1
		val1, cardSizesMap1, cardColorMap1, max1 := JudgmentGroup([]byte(alices[i]))
		val2, cardSizesMap2, cardColorMap2, max2 := JudgmentGroup([]byte(bobs[i]))
		seven := &Seven{
			cardSizeMap1:  cardSizesMap1,
			cardSizeMap2:  cardSizesMap2,
			cardColorMap1: cardColorMap1,
			cardColorMap2: cardColorMap2,
			card1:         []byte(alices[i]),
			card2:         []byte(bobs[i]),
			max1:          max1,
			max2:          max2,
		}
		// 牌型比较
		if val1 < val2 {
			result = 1
		} else if val1 > val2 {
			result = 2
		} else {
			// 同牌型下的比较
			switch val1 {
			case 1:
				// 同花顺
				seven.straightFlush()
			case 3:
				// 四条
				result = seven.FourCom()
			case 4:
				// 3带2
				result = seven.ThreeAndTwo()
			case 5:
				// 同花
				result = seven.onlySameFlower()
			case 6:
				// 顺子
				result = seven.OnlyShunZi()
			case 7:
				// 3条
				result = seven.OnlyThree()
			case 8:
				// 2对
				result = seven.TwoPair()
			case 9:
				// 一对
				result = seven.APair()
			case 10:
				// 单牌
				result = seven.SingleCard()
			}
		}

		if results[i] != result {
			fmt.Printf("[%#v]7张判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n",
				k, alices[i], bobs[i], results[i], result)
			k++
		}
	}

	t2 := time.Now()
	fmt.Println("time----->>>", t2.Sub(t1))
}
