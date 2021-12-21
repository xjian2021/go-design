package singleton

import "sync"

type Singleton struct{}

var (
	singleton *Singleton
	once      sync.Once
)

//GetInstance 懒惰式的单例模式(只有需要时才去实例化)
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
