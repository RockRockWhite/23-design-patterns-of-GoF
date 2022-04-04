package Lazy

import (
    "fmt"
    "sync"
    "time"
)

// 懒汉模式
type lazy struct {
}

var instance *lazy
var mu sync.Mutex
var once sync.Once

func GetInstance() *lazy {
    // DCL懒汉式 Double Check Lock
    // 第一次判断是否要加锁
    if instance == nil {
        mu.Lock()
        defer mu.Unlock()
    }

    // 第二次判断是否要创建对象
    if instance == nil {
        // 模拟创建对象前的一些初始化操作 该处要解决线程不安全问题
        time.Sleep(3 * time.Second)

        instance = &lazy{}
        fmt.Println("init once")
    }

    return instance
}

func GetInstanceOnce() *lazy {
    // 使用sync.Once确保只执行一次
    once.Do(func() {
        // 模拟创建对象前的一些初始化操作 该处要解决线程不安全问题
        time.Sleep(3 * time.Second)

        instance = &lazy{}
        fmt.Println("init once")
    })

    return instance
}

func (lazy) Nothing() {

}
