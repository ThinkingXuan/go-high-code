package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s world peace!", m.Player.Name, m.Monster.Name)
}

//func main() {
//	//代码量少，结构不复杂的情况下，上面的实现方式确实没什么问题。
//	//但是项目庞大到一定程度，结构之间的关系变得非常复杂的时候，
//	//这种手动创建每个依赖，然后将它们组装起来的方式就会变得异常繁琐，
//	//并且容易出错。这个时候勇士wire出现了！
//	monster := NewMonster()
//	player := NewPlayer("dj")
//	mission := NewMission(player, monster)
//
//	mission.Start()
//
//}
