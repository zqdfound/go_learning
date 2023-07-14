package main

import "fmt"

//func main() {
//	ch := make(chan int, 1000)
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//	}()
//	go func() {
//		for {
//			a, err := <-ch
//			if !err {
//				fmt.Println("Close")
//				return
//			}
//			fmt.Println(a)
//		}
//	}()
//	close(ch)
//	fmt.Println("END")
//	time.Sleep(time.Second * 100)
//}

//const N = 10
//
//var wg = &sync.WaitGroup{}
//
//func main() {
//	for i := 0; i < N; i++ {
//		wg.Add(1)
//		go func() {
//			println(i)
//			defer wg.Done()
//		}()
//	}
//	wg.Wait()
//}

//type AbstractBanker interface {
//	doBusi()
//}
//
//type SaveBanker struct {
//}
//
//func (saveBanker *SaveBanker) doBusi() {
//	fmt.Println("banker doing saving")
//}
//
//type TransBanker struct {
//}
//
//func (transBanker *TransBanker) doBusi() {
//	fmt.Println("banker doing transfering")
//}
//
//func BankerDoing(banker AbstractBanker) {
//	banker.doBusi()
//}
//func main() {
//	BankerDoing(&SaveBanker{})
//	BankerDoing(&TransBanker{})
//}

// ===== >   抽象层  < ========
//type Car interface {
//	Run()
//}
//
//type Driver interface {
//	Drive(car Car)
//}
//
//// ===== >   实现层  < ========
//type BenZ struct {
//	//...
//}
//
//func (benz *BenZ) Run() {
//	fmt.Println("Benz is running...")
//}
//
//type Bmw struct {
//	//...
//}
//
//func (bmw *Bmw) Run() {
//	fmt.Println("Bmw is running...")
//}
//
//type Zhang_3 struct {
//	//...
//}
//
//func (zhang3 *Zhang_3) Drive(car Car) {
//	fmt.Println("Zhang3 drive car")
//	car.Run()
//}
//
//type Li_4 struct {
//	//...
//}
//
//func (li4 *Li_4) Drive(car Car) {
//	fmt.Println("li4 drive car")
//	car.Run()
//}
//
//// ===== >   业务逻辑层  < ========
//func main() {
//	//张3 开 宝马
//	var bmw Car
//	bmw = &Bmw{}
//
//	var zhang3 Driver
//	zhang3 = &Zhang_3{}
//
//	zhang3.Drive(bmw)
//
//	//李4 开 奔驰
//	var benz Car
//	benz = &BenZ{}
//
//	var li4 Driver
//	li4 = &Li_4{}
//
//	li4.Drive(benz)
//}
//模拟组装2台电脑，
//--- 抽象层 ---有显卡Card  方法display，有内存Memory 方法storage，有处理器CPU 方法calculate
//--- 实现层层 ---有 Intel因特尔公司 、产品有(显卡、内存、CPU)，有 Kingston 公司， 产品有(内存3)，有 NVIDIA 公司， 产品有(显卡)
//--- 逻辑层 ---1. 组装一台Intel系列的电脑，并运行，2. 组装一台 Intel CPU  Kingston内存 NVIDIA显卡的电脑，并运行

type Card interface {
	display()
}
type Memory interface {
	storage()
}
type Cpu interface {
	calculate()
}

// ==================
type intelCard struct {
}
type intelMemory struct {
}
type intelCpu struct {
}

type kinstoneMemory struct {
}
type nvidiaCard struct {
}

func (intel *intelCard) display() {
	fmt.Println("intel Card display")
}
func (intel *intelMemory) storage() {
	fmt.Println("intel Memory storage")
}
func (intel *intelCpu) calculate() {
	fmt.Println("intel Cpu calculate")
}

func (kingston *kinstoneMemory) storage() {
	fmt.Println("kingston Memory storage")
}

func (nvidia *nvidiaCard) display() {
	fmt.Println("nvidia Card display")
}

type Computer struct {
	card   Card
	memory Memory
	cpu    Cpu
}

func (computer *Computer) computerRun() {
	computer.card.display()
	computer.memory.storage()
	computer.cpu.calculate()
}
func main() {
	computer := &Computer{
		card:   &intelCard{},
		memory: &intelMemory{},
		cpu:    &intelCpu{},
	}
	computer.computerRun()
}
