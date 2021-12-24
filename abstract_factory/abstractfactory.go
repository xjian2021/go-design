package abstract_factory

import "fmt"

type ParentsInterface interface {
	GoToWork()
}

type SonInterface interface {
	Study()
}

//HomeFactory 抽象工厂模式
//把一系列工厂之间相关的关系进行抽象则为抽象工厂模式
//抽象工厂模式是在产品间存在联系或者关系时才需要，如果是一些相互独立的产品则无需使用抽象工厂模式
type HomeFactory interface {
	CreateParents() ParentsInterface
	CreateSon() SonInterface
}

type JJHome struct {}

func (j *JJHome) CreateParents() ParentsInterface {
	return &JJParents{}
}

func (j *JJHome) CreateSon() SonInterface {
	return &JJ{}
}

type JJParents struct {}

func (j *JJParents) GoToWork() {
	fmt.Printf("jj 父母去工作\n")
}

type JJ struct {}

func (j *JJ) Study() {
	fmt.Printf("jj 去学习\n")
}

type XjianHome struct {}

func (x *XjianHome) CreateParents() ParentsInterface {
	return &XjianParents{}
}

func (x *XjianHome) CreateSon() SonInterface {
	return &Xjian{}
}

type XjianParents struct {}

func (x *XjianParents) GoToWork() {
	fmt.Printf("Xjian 父母去工作\n")
}

type Xjian struct {}

func (x *Xjian) Study() {
	fmt.Printf("Xjian 去学习\n")
}