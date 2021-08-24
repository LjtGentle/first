/*
扑克牌52张，花色黑桃spades，红心hearts，方块diamonds，草花clubs各13张，2-10，J，Q，K，A

Face：即2-10，J，Q，K，A，其中10用T来表示。

Color：即S(spades)、H(hearts)、D(diamonds)、C(clubs)

用 Face字母+小写Color字母表示一张牌，比如As表示黑桃A，其中A为牌面，s为spades，即黑桃，Ah即红心A，以此类推。
现分别给定任意两手牌各有5张，例如：AsAhQsQhQc，vs KsKhKdKc2c，请按德州扑克的大小规则来判断双方大小。

代码要求有通用性，可以任意输入一手牌或几手牌来进行比较。


5张
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Poker struct {
	Alice  string `json:"alice"`
	Bob    string `json:"bob"`
	Result int    `json:"result"`
}

type Match struct {
	Matches []Poker `json:"matches"`
}

// 定义一个牌型
// 1.皇家同花顺
// 2.同花顺
// 3.四条
// 4.满堂彩（葫芦，三带二）
// 5.同花
// 6.顺子
// 7.三条
// 8.两对
// 9.一对
// 10.单张大牌
type CardType int

// 把数据读取出来 分别放在切片中
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
		//fmt.Printf("k=%#v,v=%#v\n",k,v)
		alices[k] = v.Alice
		bobs[k] = v.Bob
		results[k] = v.Result
	}
	return
}

// PressIntoPosition 检索出现的次数入坑map
func PressIntoPosition(cardSize string)(resMap map[int32]int) {
	//初始化空间
	resMap = make(map[int32]int,5)
	//遍历入坑
	for _, value := range cardSize {
		_, ok := resMap[value]
		if ok {
			resMap[value] += 1
		}else {
			resMap[value] = 1
		}
	}
	return
}

// 分开牌型的新方法  -->传进来的是 byte
func JudgMentGroupNew(card []byte)(cardType CardType) {
	cardColorMap := make(map[byte]int,5)
	cardSizeMap := make(map[byte]int,5)
	// 扫描牌 分别放好大小，花色
	for i, v := range card {
		if i%2 ==0 {
			// 大小
			if _,ok := cardSizeMap[v]; ok {
				cardSizeMap[v] +=1
			}else {
				cardSizeMap[v] = 1
			}
		}else {
			if _,ok := cardColorMap[v]; ok {
				cardColorMap[v] +=1
			}else {
				cardColorMap[v] = 1
			}
		}
	}
	// 获取map的长度
	sizeLen := len(cardSizeMap)
	colorLen := len(cardColorMap)

	if colorLen < 5 {
		// 非同花
		switch sizeLen {
		case 4:
			// 一对
			cardType = 9
			return
		case 2 : // 3带2  或是 4带1
			// 遍历map value
			for _,v:=range cardSizeMap {
				if v == 4{
					cardType = 3
					return
				}
			}
			cardType = 4
			return
		case 3 :
			// 3条 或是 两对
			for _,v :=range cardSizeMap {
				if v == 3 {
					cardType = 7
					return
				}
			}
			cardType = 8
			return
		case 5:
			// 单牌或是顺子


		}

	}else {
		// 同花 或是 同花顺
	}


return
}

// 根据手牌判断 牌型
func JudgmentGroup(cards string) (cardType CardType, cardSizes, cardColors []string) {
	// 遍历分开手牌的大小和颜色
	cardSizes = make([]string, 5)
	cardColors = make([]string, 5)
	for i, v := range cards {
		// fmt.Println("v=",string(v))
		if i%2 != 0 {
			cardColors[i/2] = string(v)
		} else {
			cardSizes[i/2] = string(v)
		}
	}

	// 先判断是否是同花
	var s, h, d, c int
	for _, v := range cardColors {

		switch v {
		case "h": // 红心
			h++
		case "s": // 黑桃
			s++
		case "d": // 方块
			d++
		case "c": // 草花
			c++
		default:
			fmt.Println("无法解析的花色")
			os.Exit(-1)
		}
	}
	// 同花
	if s == 5 || h == 5 || d == 5 || c == 5 {
		// 同花后判断是不是顺子

		// 判断是不是顺子
		isShun := IsShunZi(cardSizes)

		if isShun == true {
			// 是 就同花顺
			cardType = 1
		} else {
			// 否 就同花
			cardType = 5
		}
		return

	}

	// 不是同花
	// 先判断是不是顺子
	isShun := IsShunZi(cardSizes)
	if isShun == true {
		// 顺子出函数
		cardType = 6
		return
	}

	// 新判断哪张重复，重复多少次
	cardSize := StringSliceToString(cardSizes)
	reMap := PressIntoPosition(cardSize)
	k3 := 0
	k2 := 0
	for _,v := range reMap {
		if v==4 {
			cardType = 3
			return
		}
		if v == 3 {
			k3 += 1
			continue
		}
		if v == 2 {
			k2 += 1
			continue
		}
	}
	if k3==1 && k2==1 {
		cardType = 4
		return
	}
	if k3==1 && k2==0 {
		cardType = 7
		return
	}
	if k2 == 2 {
		cardType = 8
		return
	}
	if k2 == 1 {
		cardType = 9
		return
	}
	cardType = 10
	return



	//  判断重复张数
	// var sum []int
	sums := make([]int, len(cardSizes)/2+1)
	for i := 0; i < len(cardSizes)/2+1; i++ {
		for _, v := range cardSizes {
			// fmt.Printf("v111=%#v--->",v)
			if v == cardSizes[i] {
				sums[i] ++
			}
		}
		// fmt.Println("sum=",sums[i])
	}
	k1 := 0 // 记录三个
	k2 = 0 // 记录两个
	for _, v := range sums {
		if v >= 4 {
			// 为四条出函数
			cardType = 3
			return
		}
		if v == 3 {
			k1++
		}
		if v == 2 {
			k2++
		}
	}
	// fmt.Printf("k1=%#v,k2=%#v\n",k1,k2)
	// 三带二 出函数
	if k1 >= 1 && k2 >= 1 {
		cardType = 4
		return
	}
	// 三条
	if k1 >= 1 {
		cardType = 7
		return
	}
	// 两对
	if k2 >= 2 {
		cardType = 8
		return
	}
	// 一对
	for i := len(cardSizes)/2 + 1; i < len(cardSizes); i++ {
		for _, v := range cardSizes {
			// fmt.Printf("v111=%#v--->",v)
			if v == cardSizes[i] {
				sums[i/2] ++
			}
		}
		// fmt.Println("sum=",sums[i/2])
	}
	for _, v := range sums {
		if v > 2 {
			cardType = 9
			return
		}
	}
	// 单张大牌
	cardType = 10
	return
}

// 判断是不是顺子
func IsShunZi(cardSize []string) (shunZi bool) {
	shunZi = false
	saves := make([]string, 14)
	for _, v := range cardSize {
		switch v {
		case "2":
			saves[1] = v
		case "3":
			saves[2] = v
		case "4":
			saves[3] = v
		case "5":
			saves[4] = v
		case "6":
			saves[5] = v
		case "7":
			saves[6] = v
		case "8":
			saves[7] = v
		case "9":
			saves[8] = v
		case "T":
			saves[9] = v
		case "J":
			saves[10] = v
		case "Q":
			saves[11] = v
		case "K":
			saves[12] = v
		case "A":
			saves[13] = v
			saves[0] = v
		default:
			fmt.Println("无法解析的扑克牌")
		}
	}
	// 判断数组是否连续
	sum := 0
	for _, v := range saves {
		if v != "" {
			sum++
		} else {
			sum = 0
		}
		if sum >= 5 {
			// break
			shunZi = true
			return
		}
	}
	return
}

// 下面是同类型的比较 ---------

// 面值转换
func tranNumFace(num uint8) (f32 float32) {
	switch num {
	case 48:
		f32 = 0
	case 50:
		f32 = 2
	case 51:
		f32 = 3
	case 52:
		f32 = 4
	case 53:
		f32 = 5
	case 54:
		f32 = 6
	case 55:
		f32 = 7
	case 56:
		f32 = 8
	case 57:
		f32 = 9
	case 84:
		f32 = 10
	case 74:
		f32 = 11
	case 81:
		f32 = 12
	case 75:
		f32 = 13
	case 65:
		f32 = 14
	default:
		fmt.Println("无法解析的花色---退出程序 num = ",num)
		//os.Exit(-1)
	}
	return
}

// 花色转换
func tranNumColor(num uint8) (f32 float32) {
	switch num {
	case 100:
		f32 = 0.0
	case 99:
		f32 = 0.1
	case 104:
		f32 = 0.2
	case 115:
		f32 = 0.3
	}
	return
}

// 牌string转成float32切片, 参数带花色的 并且返回一个有顺序的
func TranNums(card string) (f32s []float32) {
	// 转换
	f32s = make([]float32, len(card)/2)
	for i := 0; i < len(card); i += 2 {
		f32s[i/2] = tranNumFace(card[i]) /*+tranNumColor(card[i+1])*/
		if card[i] == 0{
			fmt.Println("TranNums() card[i] ==0")
		}
	}
	// 排序
	// fmt.Println("f32s---->",f32s)
	f32s = QuickSortFloat32(f32s)
	// fmt.Println("after--f32s->",f32s)
	return
}

func TranNumsNoColor(card string) (f32s []float32) {
	f32s = make([]float32, len(card))
	for i := 0; i < len(card); i++ {
		f32s[i] = tranNumFace(card[i])
		// if card[i] == 0 {
		// 	fmt.Println("TranNumsNoColor card[i] = 0")
		// }
	}
	f32s = QuickSortFloat32(f32s)
	return
}

// 快排  -->倒序
func QuickSortFloat32(f32s []float32) []float32 {
	if len(f32s) <= 1 {
		return f32s
	}
	splitdata := f32s[0]           // 第一个数据
	low := make([]float32, 0, 0)   // 比我小的数据
	hight := make([]float32, 0, 0) // 比我大的数据
	mid := make([]float32, 0, 0)   // 与我一样大的数据
	mid = append(mid, splitdata)   // 加入一个
	for i := 1; i < len(f32s); i++ {
		if f32s[i] > splitdata {
			low = append(low, f32s[i])
		} else if f32s[i] < splitdata {
			hight = append(hight, f32s[i])
		} else {
			mid = append(mid, f32s[i])
		}
	}
	low, hight = QuickSortFloat32(low), QuickSortFloat32(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

func SingleCardCompareSizNoColor(cards ...string) (result int) {

	mapf32s := make(map[int][]float32, len(cards))

	// 分别把传进来的很多副手牌各自排序好-->[]float32
	for i, v := range cards {
		// fmt.Printf("i=%#v,v=%#v\n",i,v)
		mapf32s[i] = make([]float32, len(v))
		mapf32s[i] = TranNumsNoColor(v)
	}
	if len(cards) == 2 {
		for i := 0; i < len(mapf32s[0]); i++ {
			if mapf32s[0][i] > mapf32s[1][i] {
				return 1
			} else if mapf32s[0][i] < mapf32s[1][i] {
				return 2
			}
		}
		return 0
	}
	// 比较两幅以上....
	return -1
}

// 单张大牌的比较  -----1前大 2后大  0一样大 -1出错  传进去的参数面值+花色
func SingleCardCompareSize(cards ...string) (result int) {
	// fmt.Println("=========>len(cards)=",len(cards))
	mapf32s := make(map[int][]float32, len(cards))

	// 分别把传进来的很多副手牌各自排序好-->[]float32
	for i, v := range cards {
		// fmt.Printf("i=%#v,v=%#v\n",i,v)
		mapf32s[i] = make([]float32, len(v))
		mapf32s[i] = TranNums(v)
	}
	if len(cards) == 2 {
		for i := 0; i < len(mapf32s[0]); i++ {
			if mapf32s[0][i] > mapf32s[1][i] {
				return 1
			} else if mapf32s[0][i] < mapf32s[1][i] {
				return 2
			}
		}
		return 0
	}
	// 比较两幅以上....
	return -1
}

//  传进来的一个string 返回重复的面值切片，并且将重复的面值置0
// func findPair(card string)(res []uint8) {
// 	res := make([]uint8,1)
// }

// 一对比大小
/*则两张牌中点数大的赢，如果对牌都一样，
则比较另外三张牌中大的赢，如果另外三张牌中较大的也一样则比较第二大的和第三大的，如果所有的牌都一样，则平分彩池。*/
func aPairCom(cardLen int, card ...string) (result int) {
	// 先找出对子
	pairs := make([]uint8, cardLen)
	for index, _ := range card {
		//fmt.Printf("--->begin----card[%#v]=%#v\n", index, card[index])
		b := []byte(card[index])
	label:
		for i := 0; i < cardLen/2+2; i += 2 {
			for j := i + 2; j < cardLen; j += 2 {
				if b[i] == b[j] {
					// 获取对子
					pairs[index] = card[index][i]
					//fmt.Println("pairs[index]=", pairs[index])
					b[i] = 48 // 0
					b[j] = 48
					break label
				}
			}
		}
		card[index] = string(b)
		//fmt.Printf("--->after----card[%#v]=%#v\n", index, card[index])
	}
	// 传进来的参数两个
	if len(card) == 2 {
		val1 := tranNumFace(pairs[0])
		val2 := tranNumFace(pairs[1])
		if val1 > val2{
			result = 1
			return
		} else if val1 < val2 {
			result = 2
			return
		} else {
			// 对子一样，比较剩余牌的值大小
			// 把对子移除
			result = SingleCardCompareSize(card[0], card[1])
			return
		}
	}

	return
}

// 两对
/*两对
两对点数相同但两两不同的扑克和随意的一张牌组成。
平手牌：如果不止一人抓大此牌相，牌点比较大的人赢，如果比较大的牌点相同，那么较小牌点中的较大者赢，
如果两对牌点相同，那么第五张牌点较大者赢（起脚牌）。如果起脚牌也相同，则平分彩池。*/
func twoPairCom(cardLen int, card ...string) (result int) {
	pairs := make([][]byte, len(card))
	for index, _ := range card {
		pairs[index] = make([]byte, cardLen/4) // 进来的牌是带花色的
		// fmt.Println("pairs=",len(pairs),"  ",len(pairs[0]))
		b := []byte(card[index])
		k := 0
	label:
		for i := 0; i < cardLen/2+2; i += 2 {
			for j := i + 2; j < cardLen; j += 2 {
				if b[i] == b[j] {
					// 获取对子
					pairs[index][k] = card[index][i]
					//fmt.Printf("pairs[%#v][%#v]=%#v\n", index, k, string(pairs[index][k]))
					b[i] = 48 // 0
					b[j] = 48
					k++
					if k == 2 {
						break label
					}
				}
			}
		}
		//fmt.Println()
		card[index] = string(b)
	}
	//fmt.Println("--->pairs[0]--->", string(pairs[0]))
	// fmt.Println("card=",card)

	// 传进来两幅牌比较
	if len(card) == 2 {
		// 先对子比大小  --相当于单张牌比大小
		result = SingleCardCompareSizNoColor(string(pairs[0]), string(pairs[1]))
		//fmt.Println("string(pairs[0])=", string(pairs[0]), "  string(pairs[1])=", string(pairs[1]), "res=", result)
		if result != 0 {
			return
		}
		// 光光比较对子比较不出来
		result = SingleCardCompareSize(card[0], card[1])
		return
	}

	return
}

/*
三条
由三张相同点数和两张不同点数的扑克组成 。
平手牌：如果不止一人抓到此牌，则三张牌中最大点数者赢局，
如果三张牌都相同，比较第四张牌，必要时比较第五张，点数大的人赢局。如果所有牌都相同，则平分彩池。
*/
func onlyThreeCom(cardLen int, cardSize ...string) (result int) {
	pokerMaps := make([]map[int32]int, len(cardSize))
	// 把
	for index, value := range cardSize {
		pokerMaps[index] = make(map[int32]int, cardLen)
		for _, v := range value {
			// fmt.Println("i=",i," v=",v)
			mv, ok := pokerMaps[index][v]
			if ok {
				// fmt.Println("mv=",mv)
				mv++
				pokerMaps[index][v] = mv
				// fmt.Println("pokerMaps[index][v]=",pokerMaps[index][v])
			} else {
				pokerMaps[index][v] = 1
			}
		}
		// fmt.Println("---------------")

	}
	// fmt.Println("mapslice=",pokerMaps)
	if len(cardSize) == 2 {
		// 然后遍历map 找出三条是哪个数值
		// var val1 int32
		// var val2 int32
		vals := make([]int32, 2)
		for k, v := range pokerMaps[0] {
			if v == 3 {
				vals[0] = k
				break
			}
		}
		for k, v := range pokerMaps[1] {
			if v == 3 {
				vals[1] = k
				break
			}
		}
		// fmt.Println("val1=",val1," val2=",val2)
		if vals[0] > vals[1] {
			return 1
		} else if vals[0] < vals[1] {
			return 2
		} else {
			// 花色比较不了,继续比较单牌  ---
			result = SingleCardCompareSizNoColor(cardSize[0], cardSize[1])
			//fmt.Printf("cardSize[0]=%#v,cardSize[1]=%#v\n", cardSize[0], cardSize[1])
			return
		}

	}

	return
}

func onlyShunZi(cards ...string) (result int) {

	if len(cards) == 2 {
		result =  SingleCardCompareSize(cards[0],cards[1])
	}

	return
}

func onlySameFlower(cards ... string)(result int){
	if len(cards) == 2 {
		result =  SingleCardCompareSize(cards[0],cards[1])
	}
	return
}

func threeAndTwo(cardLen int,cardSizes ... string)(result int) {
	pokerMaps := make([]map[int32]int, len(cardSizes))
	// 把
	for index, value := range cardSizes {
		pokerMaps[index] = make(map[int32]int, cardLen)
		for _, v := range value {
			mv, ok := pokerMaps[index][v]
			if ok {
				mv++
				pokerMaps[index][v] = mv
			} else {
				pokerMaps[index][v] = 1
			}
		}
	}
	if len(cardSizes) == 2 {

		val3s := make([]int32, 2)
		val2s := make([]int32,2)
		for k, v := range pokerMaps[0] {
			if v == 3 {
				val3s[0] = k
			}else if v == 2{
				val2s[0] = k
			}
		}
		for k, v := range pokerMaps[1] {
			if v == 3 {
				val3s[1] = k
			}else if v == 2 {
				val2s[1] = k
			}
		}
		 f11  := tranNumFace(uint8(val3s[0]))
		 f12  := tranNumFace(uint8(val3s[1]))
		 f21  := tranNumFace(uint8(val2s[0]))
		 f22  := tranNumFace(uint8(val2s[1]))
		if f11 > f12 {
			return 1
		} else if f11 < f12 {
			return 2
		}else if f21 >f22{
			return 1
		}else if f21 < f22 {
			return 2
		}else {
			return 0
		}

	}
	return


}

func fourCom(cardLen int,cardSizes ... string)(result int) {
	pokerMaps := make([]map[int32]int, len(cardSizes))
	// 把
	for index, value := range cardSizes {
		pokerMaps[index] = make(map[int32]int, cardLen)
		for _, v := range value {
			mv, ok := pokerMaps[index][v]
			if ok {
				mv++
				pokerMaps[index][v] = mv
			} else {
				pokerMaps[index][v] = 1
			}
		}
	}
	//fmt.Println("pokerMaps=",pokerMaps)
	if len(cardSizes) == 2 {

		val3s := make([]int32, 2)
		val2s := make([]int32,2)
		for k, v := range pokerMaps[0] {
			if v == 4 {
				val3s[0] = k
			}else if v == 1{
				val2s[0] = k
			}
		}
		for k, v := range pokerMaps[1] {
			if v == 4 {
				val3s[1] = k
			}else if v == 1 {
				val2s[1] = k
			}
		}
		f11  := tranNumFace(uint8(val3s[0]))
		f12  := tranNumFace(uint8(val3s[1]))
		f21  := tranNumFace(uint8(val2s[0]))
		f22  := tranNumFace(uint8(val2s[1]))
		if f11 > f12 {
			return 1
		} else if f11 < f12 {
			return 2
		}else if f21 >f22{
			return 1
		}else if f21 < f22 {
			return 2
		}else {
			return 0
		}

	}
	return
}

// 同花顺
func straightFlush(cardLen int, cardSizes ... string)(result int){
	if len(cardSizes) == 2 {
		result =  SingleCardCompareSizNoColor(cardSizes[0],cardSizes[1])
	}
	return
}

func StringSliceToString(sliceStr []string) (str string) {
	for _, v := range sliceStr {
		str += v
	}
	return str
}



func PokerMan() {
	file := "/home/weilijie/chromeDown/match_result.json"
	alices := make([]string, 1024)
	bobs := make([]string, 1024)
	results := make([]int, 1024)
	alices, bobs, results = ReadFile(file)
	// return
	t1 := time.Now()
	// for i:=0; i < len(alices); i ++ {
	// 	fmt.Printf("alices[%#v]=%#v\n",i,alices[i])
	// 	fmt.Printf("bobs[%#v]=%#v\n",i,bobs[i])
	// 	fmt.Printf("results[%#v]=%#v\n",i,results[i])
	// }

	k := 0
	for i := 0; i < len(alices); i++ {
		result := -1
		val1, cardSizes1, _ := JudgmentGroup(alices[i])
		val2, cardSizes2, _ := JudgmentGroup(bobs[i])
		cardSize1 := StringSliceToString(cardSizes1)
		cardSize2 := StringSliceToString(cardSizes2)
		if val1 < val2 {
			result = 1
		} else if val1 > val2 {
			result = 2
		} else {
			// 牌型处理相同的情况
			// ...

			switch val1 {
			case 10:
				// 同类型下的单张大牌比较
				result = SingleCardCompareSize(alices[i], bobs[i])
			case 9:
				// 同类型的一对
				result = aPairCom(10, alices[i], bobs[i])
			case 8:
				// 同类型两对
				result = twoPairCom(10, alices[i], bobs[i])
			case 7:
				// 同类型三条
				result = onlyThreeCom(10, cardSize1, cardSize2)
			case 6:
				// 同类型顺子
				result = onlyShunZi(alices[i],bobs[i])
			case 5:
				// 同类型同花
				result = onlySameFlower(alices[i],bobs[i])
			case 4 :
				// 同类型3带2
				result = threeAndTwo(5,cardSize1,cardSize2)
			case 3:
				// 同类型四条
				result = fourCom(5,cardSize1,cardSize2)
			case 1:
				result = straightFlush(5,cardSize1,cardSize2)
			}

			// 最后比较结果
		}

		if result != results[i] {
			k++
			fmt.Printf("[%#v]判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n",k, alices[i], bobs[i],results[i],result)
		} else {
			//fmt.Println("判断正确222222")
		}
	}
	t2 := time.Now()
	fmt.Println("timetime--->",t2.Sub(t1))
}

func main() {
	// file := "/home/weilijie/chromeDown/match_result.json"
	// ReadFile(file)

	// JudgmentGroup("AsAhQsQhQc")
	// str1 := "23456789"
	// fmt.Println("[]byte(str1)",[]byte(str1))
	// str2 := "T"
	// fmt.Println("[]byte(str2)",[]byte(str2))
	// str3 := "JQKA"
	// fmt.Println("[]byte(str3)",[]byte(str3))
	// str := []string {"2"}
	// isres := IsShunZi(str)
	// fmt.Println("isres=",isres)
	//
	// type1 ,_,_ := JudgmentGroup("KsKhKdKc2c")
	// fmt.Println("type1=",type1)

	// result := SingleCardCompareSize(5,"AhKdKc2cKs","2cKsKhKdKc")
	// fmt.Println("---->result=",result)
	// str4 := "shcd"
	// fmt.Println("[]byte(str4)=",[]byte(str4))
	//
	// var int32 int32
	// int32 = 105*1000
	// fmt.Println("int_32=",int32)
	// fs :=[]float32 {88.88,11.0,99.2,1.0,5.4,88.2,77,65,20,3}
	// res := make([]float32,10)
	// res =QuickSortFloat32(fs)
	// fmt.Printf("fs=%#v\nres=%#v\n",fs,res)

	// PokerMan()

	// str := "0"
	// fmt.Println("str=",[]byte(str))
	// result := aPairCom(10,"2d4s4h3s5h","4h2h5s3c4d")
	// fmt.Println("result=",result)

	// res := twoPairCom(10,"2d2s4d4s5h","2c3d2h3cJh")
	// fmt.Println("res=",res)

	// res := onlyThreeCom(5,"666JK","76656")
	// fmt.Println("res=",res)
	// slice := make([]string, 5)
	// str := "A"
	// for i := 0; i < len(slice); i++ {
	// 	slice[i] = str
	// }
	//
	// result := StringSliceToString(slice)
	// fmt.Println("slice=", slice)
	// fmt.Println("result=", result)

	// res := fourCom(5,"KKKK3","AAAA2")
	// fmt.Println("---->res=",res)
	t1 :=time.Now()
	PokerMan()
	t2 := time.Now()
	fmt.Println("time----",t2.Sub(t1).Seconds())

}
