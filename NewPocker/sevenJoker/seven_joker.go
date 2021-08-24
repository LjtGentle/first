package sevenJoker

import (
	"NewPocker/fire"
	"fmt"
	"time"
)

// comSevenJoker 调用同类型比较函数传入的参数
type comSevenJoker struct {
	cardSizeMap1, cardSizeMap2 map[byte]int
	resCard1, resCard2         byte
	joker1, joker2             int // 记录癞子数量
}

// SameFlowerSeq 找到同花对应的面值序列
func SameFlowerSeq(cardColorMap map[byte]int, card []byte, joker int) (sizeSlice []byte) {
	// 存放同花的花色
	var color byte
	sliceLen := 0
	for k, v := range cardColorMap {
		if v+joker >= 5 { // joker是癞子假设癞子变花色
			color = k
			sliceLen = v + joker
			break
		}
	}
	sizeSlice = make([]byte, sliceLen)
	j := 0
	for i := 1; i < len(card); i += 2 {
		if card[i] == color {
			sizeSlice[j] = fire.SizeTranByte(card[i-1]) // 花色前一个是它对应的面值，转译成可以比较的
			j++
		}
	}
	// 将所有jock都变A ， 因为同花下比较单牌 癞子变最大的单牌
	for joker > 0 && sliceLen != 0 {
		sizeSlice[j] = 0x0E
		j++
		joker--
	}
	return

}

// IsShunZi 判断是不是顺子，返回顺子的一个最大面值 传进来的已经转译好的面值
func IsShunZi(seq []byte, joker int) (shunZi bool, max byte) {
	saves := make([]byte, 15)
	// 遍历seq，把它对应放到saves中，安位入座
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
		case 0x10:
			saves[14] = v
		default:
			fmt.Println("IsShunZi say 无法解析的扑克牌", "card --v=", v)
		}
	}
	sum := 0
	// 判断数组是否连续
	if joker < 1 {
		// 没有癞子的顺子
		for i := len(saves) - 1; i >= 0; i-- {
			if saves[i] != 0x00 {
				sum++
			} else {
				sum = 0
			}
			if sum >= 5 {
				shunZi = true
				max = saves[i+4]
				return
			}
		}
	} else {
		tmp := joker
		sum = 0
		for i := len(saves) - 1; i >= 0; i-- {
			if saves[i] != 0 {
				sum++
			} else if joker > 0 {
				joker--
				sum++
			} else {
				// 这里是回退到一开始的下一个
				i = i + sum - joker
				// 重置 joker 和sum
				joker = tmp
				sum = 0
			}
			if sum >= 5 {
				// 是顺子
				max = saves[i+4]
				if max == 0 {
					// 那么癞子就变成这个下标对应的值
					max = IndexTranByte(i + 4)
				}
				shunZi = true
				return
			}
		}
	}
	return
}

// IndexTranByte 当是顺子时，癞子去补顺子的最大值调用这个函数去转换下标对应的面值
func IndexTranByte(index int) (b byte) {
	switch index {
	// 这个癞子补充的值，不可能是很前面的
	case 4:
		b = 0x05
	case 5:
		b = 0x06
	case 6:
		b = 0x07
	case 7:
		b = 0x08
	case 8:
		b = 0x09
	case 9:
		b = 0x0A
	case 10:
		b = 0x0B
	case 11:
		b = 0x0C
	case 12:
		b = 0x0D
	case 13:
		b = 0x0E
	case 14:
		b = 0x10
	default:
		fmt.Println("IsShunZi say 无法解析的扑克牌", "card --b=", b)
	}
	return
}

// IsShunZiNoTran 判断是不是顺子，并且返回顺子的最大牌面值 传进来的card还没有转译
func IsShunZiNoTran(card []byte, joker int) (shunZi bool, max byte) {
	shunZi = false
	saves := make([]byte, 14)
	// 将面值对号入座
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
			case 88:
				// continue
				// fmt.Println("88")
			default:
				fmt.Println("无法解析的扑克牌", "card --v=", v)
			}
		}

	}
	// 下面判断数组是否连续
	sum := 0

	if joker < 1 {
		// 没有癞子的顺子
		for i := len(saves) - 1; i >= 0; i-- {
			if saves[i] != 0x00 {
				sum++
			} else {
				sum = 0
			}
			if sum >= 5 {
				shunZi = true
				max = saves[i+4]
				return
			}
		}
	} else {
		tmp := joker
		sum = 0
		for i := len(saves) - 1; i >= 0; i-- {
			if saves[i] != 0 {
				sum++
			} else if joker > 0 {
				joker--
				sum++
			} else {
				// 这里是回退到一开始的下一个
				i = i + sum - joker
				joker = tmp
				sum = 0
			}
			if sum >= 5 {
				// 是顺子
				max = saves[i+4]
				if max == 0 {
					// 需要癞子来补充这个最大的面值
					max = IndexFindByte(i + 4)
				}
				shunZi = true
				return
			}
		}
	}

	return
}

// IndexFindByte 下标转成面值的byte
func IndexFindByte(index int) (b byte) {
	switch index {
	case 5:
		b = 54
	case 6:
		b = 55
	case 7:
		b = 56
	case 8:
		b = 57
	case 9:
		b = 84
	case 10:
		b = 74
	case 11:
		b = 81
	case 12:
		b = 75
	case 13:
		b = 65
	case 0:
		b = 65
	case 88:
	case 1:
		b = 50
	case 2:
		b = 51
	case 3:
		b = 52
	case 4:
		b = 53
	default:
		fmt.Println("无法解析的下标", "card --index=", index)

	}
	return
}

//  judgmentGroup 判断牌型  传进来的参数：面值+花色
func judgmentGroup(card [] byte) (cardType uint8, cardSizeMap, cardColorMap map[byte]int, resCard byte, joker int) {
	//  对传进来的牌分拣成  面值 花色
	cardColorMap = make(map[byte]int, 7)
	cardSizeMap = make(map[byte]int, 7)
	// 扫描牌 分别放好大小，花色
	for i, v := range card {
		if i%2 == 0 {
			// 大小
			// 判断是否有癞子
			if v == 88 {
				joker++ // 记录多少个癞子但是不放入map中
				continue
			}
			if _, ok := cardSizeMap[v]; ok {
				cardSizeMap[v]++
			} else {
				cardSizeMap[v] = 1
			}
			// 颜色
		} else {
			if v == 110 { // 癞子的花色，不作处理
				continue
			}
			if _, ok := cardColorMap[v]; ok {
				cardColorMap[v]++
			} else {
				cardColorMap[v] = 1
			}
		}
	}
	sizeLen := len(cardSizeMap)
	flag := false
	for _, v := range cardColorMap {
		if v+joker >= 5 {
			flag = true
			break
		}
	}

	if flag {
		// 同花
		// 判断是不是顺子
		seq := SameFlowerSeq(cardColorMap, card, joker)
		isShun, max := IsShunZi(seq[0:len(seq)-joker], joker)
		if isShun {
			// 同花顺
			resCard = max
			cardType = 1
			return
		}
		if joker == 0 {
			cardType = 5
			return
		}
		// 有癞子的情况下，是同花，也有可能是变成4条 或是 三带2 它们都比同花大
		i := 0
		for _, v := range cardSizeMap {
			if v+joker == 4 {
				// 4 条
				cardType = 3
				return
			} else if v+joker == 3 {
				i++
			}
		}
		if i == 2 {
			// 3带2
			cardType = 4
			return
		}
		cardType = 5
		return

	}
	// 不是同花
	// 根据面值的map长度判断
	switch sizeLen {
	case 7: // 单牌也没有癞子
		// 判断是不是顺子
		if isShun, max := IsShunZiNoTran(card, joker); isShun {
			cardType = 6
			resCard = max
			return
		}
		cardType = 10
		return
	case 6: // 可能 顺子 单牌+癞子 或是 一对
		if isShun, max := IsShunZiNoTran(card, joker); isShun {
			resCard = max
			cardType = 6
			return
		}
		// 就算有癞子也是一对， 因为癞子不放入map，map len为6 如果有癞子的话，其他牌是单牌
		cardType = 9
		return

	case 5: // 顺子 3条 或是 两对
		// 判断是不是顺子
		if isShun, max := IsShunZiNoTran(card, joker); isShun {
			resCard = max
			cardType = 6
			return
		}
		// 有癞子的五张牌一定是3条
		if joker > 0 {
			cardType = 7
			return
		}
		for _, v := range cardSizeMap {
			if v == 3 {
				// 3条
				cardType = 7
				return
			}
		}
		// 2对
		cardType = 8
		return

	case 4: // 因为最多只有一个癞子 就可能出现4条  3带2 或是 2对（3个对子） 有癞子的话也不组顺子，因为可以组4条或是3带2比顺子大
		i := 0
		j := 0
		for _, v := range cardSizeMap {
			if v+joker == 4 {
				// 4条
				cardType = 3
				return
			} else if v+joker == 3 {
				//  有可能是本来有两对再加一个癞子
				i++

			} else if v == 2 {
				j++
			}
		}
		if j < 2 {
			cardType = 4
			return
		}
		// 这里的j=3，也就是三对并且没有癞子的
		cardType = 8
		return
	case 3: // 不是3带2 就是4条了
		for _, v := range cardSizeMap {
			if v+joker == 4 {
				cardType = 3
				return
			}
		}
		cardType = 4
		return
	case 2: // 只能是4条
		cardType = 3
		return

	}

	return
}

// straightFlush 同类型同花顺比较 --1
func (sevenJoker *comSevenJoker) straightFlush() (result int) {
	// 比较顺子最大的牌
	if sevenJoker.resCard1 > sevenJoker.resCard2 {
		return 1
	} else if sevenJoker.resCard1 < sevenJoker.resCard2 {
		return 2
	} else {
		return 0
	}
}

// fourCom 同类型四条比较  --3
func (sevenJoker *comSevenJoker) fourCom() (result int) {
	// 存放单牌的slice
	cardSizeSlice1 := make([]byte, len(sevenJoker.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenJoker.cardSizeMap2))
	// 存放4条的面值
	var four1 byte
	var four2 byte

	i := 0
	for k, v := range sevenJoker.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker1 == 4 {
			if four1 < k {
				four1 = k
			}
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range sevenJoker.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker2 == 4 {
			if four2 < k {
				four2 = k
			}
		} else {
			cardSizeSlice2[i] = k
			i++
		}
	}

	if four1 > four2 {
		return 1
	} else if four1 < four2 {
		return 2
	} else {
		// 比较单牌中最大的一张
		result = fire.SingleCardSizeCom(1, cardSizeSlice1, cardSizeSlice2)
		return
	}
}

// threeAndTwo同类型3带2比较  --4
func (sevenJoker *comSevenJoker) threeAndTwo() (result int) {
	// 存放3条的slice
	threes1 := make([]byte, 3)
	threes2 := make([]byte, 3)
	// 存放对子的slice
	twos1 := make([]byte, 3)
	twos2 := make([]byte, 3)

	i := 0
	j := 0

	for k, v := range sevenJoker.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker1 == 3 {
			threes1[i] = k
			i++
		} else if v == 2 {
			twos1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range sevenJoker.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker2 == 3 {
			threes2[i] = k
			i++
		} else if v == 2 {
			twos2[j] = k
			j++
		}
	}
	// 对3条排序
	threes1 = fire.QuickSortByte(threes1)
	threes2 = fire.QuickSortByte(threes2)

	// 因为有可能出现 2个3条
	if threes1[0] > threes2[0] {
		return 1
	} else if threes1[0] < threes2[0] {
		return 2
	}

	// 对 对子排序
	twos1 = fire.QuickSortByte(twos1)
	twos2 = fire.QuickSortByte(twos2)

	// 当两个3条的时候， 对子就没有值，那么对子的最大应该是3条的最小
	if twos1[0] < threes1[1] {
		twos1[0] = threes1[1]
	}
	if twos2[0] < threes2[1] {
		twos2[0] = threes2[1]
	}
	if twos1[0] > twos2[0] {
		return 1
	} else if twos1[0] < twos2[0] {
		return 2
	} else {
		return 0
	}
}

// onlyFlush 同类型同花比较  --5  1
func (sevenJoker *comSevenJoker) onlyFlush(cardColorMap1, cardColorMap2 map[byte]int, card1, card2 []byte) (result int) {
	// 找到同花对应的面值序列
	sizeSlice1 := SameFlowerSeq(cardColorMap1, card1, sevenJoker.joker1)
	sizeSlice2 := SameFlowerSeq(cardColorMap2, card2, sevenJoker.joker2)
	// 这个序列可以大于5个， 所以只比较各自最大的五个
	result = fire.SingleCardSizeCom(5, sizeSlice1, sizeSlice2)
	return
}

// OnlyShunZi 同类型顺子比较 --6
func (sevenJoker *comSevenJoker) OnlyShunZi() (result int) {
	v1 := fire.SizeTranByte(sevenJoker.resCard1)
	v2 := fire.SizeTranByte(sevenJoker.resCard2)
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return 2
	} else {
		return 0
	}
}

// onlyThree 同类型3条比较  --7
func (sevenJoker *comSevenJoker) onlyThree() (result int) {
	// 存放单牌的slice
	cardSizeSlice1 := make([]byte, len(sevenJoker.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenJoker.cardSizeMap2))
	// 存放3条的面值
	var three1 byte
	var three2 byte

	i := 0

	for k, v := range sevenJoker.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker1 == 3 {
			three1 = k
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range sevenJoker.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker2 == 3 {
			three2 = k
		} else {
			cardSizeSlice2[i] = k
			i++
		}
	}
	if three1 > three2 {
		return 1
	} else if three1 < three2 {
		return 2
	} else {
		// 比较各自单牌中最大的两个
		result = fire.SingleCardSizeCom(2, cardSizeSlice1, cardSizeSlice2)
		return
	}
}

// TwoPair 同类型两对比较 --8
func (sevenJoker *comSevenJoker) TwoPair() (result int) {
	// 记录对子的slice
	pairs1 := make([]byte, 3)
	pairs2 := make([]byte, 3)
	// 记录单牌的slice
	vals1 := make([]byte, 3)
	vals2 := make([]byte, 3)

	j := 0
	i := 0
	for k, v := range sevenJoker.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker1 == 2 {
			pairs1[i] = k
			i++
		} else {
			vals1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range sevenJoker.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v+sevenJoker.joker2 == 2 {
			pairs2[i] = k
			i++
		} else {
			vals2[j] = k
			j++
		}
	}
	// 对 对子 逆序列排序
	pairs1 = fire.QuickSortByte(pairs1)
	pairs2 = fire.QuickSortByte(pairs2)

	// 因为可能出现3个对子，选择选择最大两个比较
	for i := 0; i < 2; i++ {
		if pairs1[i] > pairs2[i] {
			return 1
		} else if pairs1[i] < pairs2[i] {
			return 2
		}
	}

	// 对单牌序列排序
	vals1 = fire.QuickSortByte(vals1)
	vals2 = fire.QuickSortByte(vals2)
	// 如果出现3个对子，让最小的对子跟最大的单牌比较
	if vals1[0] < pairs1[2] {
		vals1[0] = pairs1[2]
	}
	if vals2[0] < pairs2[2] {
		vals2[0] = pairs2[2]
	}
	// 比较单牌最大
	if vals1[0] > vals2[0] {
		return 1
	} else if vals1[0] < vals2[0] {
		return 2
	}
	return 0
}

// OnePair 同类型一对比较 --9
func (sevenJoker *comSevenJoker) OnePair() (result int) {
	cardSizeSlice1 := make([]byte, len(sevenJoker.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenJoker.cardSizeMap1))
	var val1 byte
	var val2 byte
	i := 0

	for k, v := range sevenJoker.cardSizeMap1 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			val1 = k
			continue
		}
		cardSizeSlice1[i] = k
		i++
	}
	i = 0
	for k, v := range sevenJoker.cardSizeMap2 {
		k = fire.SizeTranByte(k)
		if v == 2 {
			val2 = k
			continue
		}

		cardSizeSlice2[i] = k
		i++
	}

	cardSizeSlice1 = fire.QuickSortByte(cardSizeSlice1)
	cardSizeSlice2 = fire.QuickSortByte(cardSizeSlice2)

	// 癞子组成的对子
	if val1 == 0 {
		val1 = cardSizeSlice1[0]
	}
	if val2 == 0 {
		val2 = cardSizeSlice2[0]
		// fmt.Println("val22222222")
	}

	if val1 > val2 {
		return 1
	} else if val1 < val2 {
		return 2
	}

	comLen := 3
	// 单牌一个个对比
	for i := 0; i < comLen; i++ {
		if cardSizeSlice1[i+sevenJoker.joker1] > cardSizeSlice2[i+sevenJoker.joker2] {
			return 1
		} else if cardSizeSlice1[i+sevenJoker.joker1] < cardSizeSlice2[i+sevenJoker.joker2] {
			return 2
		}
	}

	return 0
}

// SingleCard 同类型单牌比较  --10
func (sevenJoker *comSevenJoker) SingleCard() (result int) {
	// 存放单牌的slice
	cardSizeSlice1 := make([]byte, len(sevenJoker.cardSizeMap1))
	cardSizeSlice2 := make([]byte, len(sevenJoker.cardSizeMap2))
	i := 0
	for k := range sevenJoker.cardSizeMap1 {
		cardSizeSlice1[i] = fire.SizeTranByte(k)
		i++
	}
	i = 0
	for k := range sevenJoker.cardSizeMap2 {
		cardSizeSlice2[i] = fire.SizeTranByte(k)
		i++
	}
	// 对比各自最大的5张
	result = fire.SingleCardSizeCom(5, cardSizeSlice1, cardSizeSlice2)
	return
}

// PokerMan 7张加癞子的主函数
func PokerMan() {
	filePath := "../resources/seven_cards_with_ghost.result.json"
	alices, bobs, results := fire.ReadFile(filePath)
	t1 := time.Now()
	k := 0
	for i := 0; i < len(alices); i++ {

		result := -1
		// 先判断各自的牌型
		cardType1, cardSizesMap1, cardColorMap1, max1, joker1 := judgmentGroup([]byte(alices[i]))
		cardType2, cardSizesMap2, cardColorMap2, max2, joker2 := judgmentGroup([]byte(bobs[i]))
		if cardType1 < cardType2 {
			result = 1
		} else if cardType1 > cardType2 {
			result = 2
		} else {

			csj := &comSevenJoker{
				cardSizeMap1: cardSizesMap1,
				cardSizeMap2: cardSizesMap2,
				resCard1:     max1,
				resCard2:     max2,
				joker1:       joker1,
				joker2:       joker2,
			}
			// 同类型比较
			switch cardType1 {
			case 1:
				// 同花顺
				result = csj.straightFlush()
			case 3:
				// 4条
				result = csj.fourCom()
			case 4:
				// 3带2
				result = csj.threeAndTwo()
			case 5:
				// 同花
				result = csj.onlyFlush(cardColorMap1, cardColorMap2, []byte(alices[i]), []byte(bobs[i]))
			case 6:
				// 顺子
				result = csj.OnlyShunZi()
			case 7:
				// 3条
				result = csj.onlyThree()
			case 8:
				// 两对
				result = csj.TwoPair()
			case 9:
				// 一对
				result = csj.OnePair()
			case 10:
				// 单牌
				result = csj.SingleCard()

			}
		}
		if result != results[i] {
			k++
			fmt.Println("[", k, "]"+"判断有误，alice=", alices[i], " bob=", bobs[i], "  我的结果:", result, "  文档的结果：", results[i])
		}
	}

	t2 := time.Now()
	fmt.Println("time--->", t2.Sub(t1))

}
