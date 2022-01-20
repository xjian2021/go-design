# **go-design**
使用golang实现各种设计模式并总结体会。还有一些在leetcode刷题的答案，目前从简单题开始刷起。

## [单例模式](https://github.com/xjian2021/go-design/tree/main/singleton)
`创建型设计模式`， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

## [简单工厂模式](https://github.com/xjian2021/go-design/tree/main/easy_factory)
`创建型设计模式`， 对外只暴露API接口和实例化实现API接口具体对象的构造函数，封装了实现细节。构造函数返回API接口就是简单工厂模式。

## [工厂模式](https://github.com/xjian2021/go-design/tree/main/factory_method)
`创建型设计模式`， 其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。

## [抽象工厂模式](https://github.com/xjian2021/go-design/tree/main/abstract_factory)
`创建型设计模式`， 它能创建一系列相关的对象， 而无需指定其具体类。

## [建造者模式](https://github.com/xjian2021/go-design/tree/main/builder)
`创建型设计模式`， 使你能够分步骤创建复杂对象。 该模式允许你使用相同的创建代码生成不同类型和形式的对象。

## [原型模式](https://github.com/xjian2021/go-design/tree/main/prototype)
`创建型设计模式`， 使你能够复制已有对象， 而又无需使代码依赖它们所属的类。
- 优点
  - 通过克隆对象，而无需与对象的类型相耦合
  - 对于使用者，可以很方便地生成复杂对象
  - 使用克隆的方式来生成新的对象，避免反复运行初始化代码
- 缺点
  - 原型需要对被复制对象进行复杂的初始化，复杂程度取决于对象类型的复杂程度。
  
## [适配器模式](https://github.com/xjian2021/go-design/tree/main/adapter)