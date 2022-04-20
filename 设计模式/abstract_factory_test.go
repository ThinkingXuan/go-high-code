package abstract_factory

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("abstract_factory:", ProduceFruitAndEat)
}

func ProduceFruitAndEat(t *testing.T) {
	var factory FruitFactory
	var apple, banana, orange Fruit

	// 创建苹果工厂，生产苹果，吃苹果
	factory = &AppleFactory{}
	apple = factory.CreateFruit()
	apple.Eat()

	//创建香蕉工厂，生产香蕉，吃香蕉
	factory = &BananaFactory{}
	banana = factory.CreateFruit()
	banana.Eat()

	//创建橘子工厂，生产橘子，吃橘子
	factory = &OrangeFactory{}
	orange = factory.CreateFruit()
	orange.Eat()
}
