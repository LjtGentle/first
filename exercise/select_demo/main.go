/*
select
*/

package main

//func (tc *NoExecuteTaintManager)worker(worker int,done func(),stopCh<-chan struct{}) {
//
//}


func main() {

	//空select  该goroutine 进入无法被唤醒的永久休眠状态
	//go func() {
	//	select {
	//
	//	}
	//	fmt.Println("end of goroutine")
	//}()
	//time.Sleep(time.Second*10)
	//ch1 := make(chan bool)
	//
	//go func() {
	//	time.Sleep(1*time.Second)
	//	ch1 <- true
	//}()
	//for {
	//	time.Sleep(3*time.Second)
	//	select {
	//	case <- ch1:
	//		fmt.Println("goroutine exit")
	//	default:
	//		fmt.Println("let go " )
	//	}
	//}

	// 只有一个case
	//intChan := make(chan int,1)
	//intChan <- 10 //写入
	//i := <-intChan //取出
	//fmt.Println("i=",i)
	//ch := make(chan bool)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-ch:
	//			fmt.Println("goroutine read from chan ok")
	//		}
	//	}
	//}()
	//for {
	//	time.Sleep(1*time.Second)
	//	ch <- true
	//
	//}

	//ch := make(chan bool)
	//go func() {
	//	for {
	//		time.Sleep(time.Second*3)
	//		ch <- true
	//	}
	//
	//}()
	//
	//for {
	//	time.Sleep(time.Second * 2)
	//	select {
	//	case <-ch:
	//		fmt.Println("read ok")
	//	default:
	//		fmt.Println("read none")
	//	}
	//}

	// select 的优先级实现


}
