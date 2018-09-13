/**
据说著名犹太历史学家 Josephus有过以下的故事：
在罗马人占领乔塔帕特后，39 个犹太人与Josephus及他的朋友躲到一个洞中，39个犹太人决定宁愿死也不要被敌人抓到，
于是决定了一个自杀方式，41个人排成一个圆圈，由第1个人开始报数，每报数到第3人该人就必须自杀，然后再由下一个重
新报数，直到所有人都自杀身亡为止。然而Josephus 和他的朋友并不想遵从。首先从一个人开始，越过k-2个人（因为第一个
人已经被越过），并杀掉第k个人。接着，再越过k-1个人，并杀掉第k个人。这个过程沿着圆圈一直进行，直到最终只剩下一个
人留下，这个人就可以继续活着
 * Description: 
 * User: 1067
 * Date: 2018-09-13
 * Time: 10:56
 */

package main

import (
	"container/ring"
	"fmt"
)

type player struct {
	position int  //位置
	isAlive  bool //是否存活
}

func main() {
	const (
		playerNum           = 41 //玩家总人数
		dieNum              = 3  //死亡数字 报到该数 死亡
		firstReportPosition = 4  //第一个报数的位置
	)
	var (
		report = 1 //从那个数开始报数
	)

	//初始化玩家
	r := initPlayer(playerNum)

	//初始化位置
	r = initPosition(r, firstReportPosition)

	//开始
	playGame(r, dieNum, report, playerNum)

	//查看唯一存活者
	showAlive(r)
}

//游戏开始
func playGame(r *ring.Ring, dieNum, report, playerNum int, ) {
	var diePlayCount int
	fmt.Println("Game start!")
	for {
		p := r.Value.(*player)
		if p.isAlive {
			fmt.Print("player ", p.position, " report:", report)
			if report == dieNum {
				fmt.Println("  daied")
				p.isAlive = false
				report = 0
				diePlayCount++
			} else {
				fmt.Println()
			}
			report++
		}

		if diePlayCount == playerNum-1 {
			fmt.Println("Game over!")
			break
		}
		r = r.Next()
	}
}

//初始化第一个报数的位置
func initPosition(r *ring.Ring, position int) *ring.Ring {

	if r.Value.(*player).position != position {
		r = r.Move(position - 1)
	}
	return r
}

//初始化玩家
func initPlayer(playerNum int) *ring.Ring {

	r := ring.New(playerNum)
	for i := 1; i <= playerNum; i++ {
		var p player
		p.position = i
		p.isAlive = true
		r.Value = &p
		r = r.Next()
	}

	fmt.Println("玩家列表:")
	r.Do(printRing)
	fmt.Println()
	return r
}

func printRing(v interface{}) {
	fmt.Print(v.(*player).position, " ")
}

//show 唯一存活的人
func showAlive(r *ring.Ring) {
	for i := 1; i <= r.Len(); i++ {
		if r.Value.(*player).isAlive {
			fmt.Println("player ", r.Value.(*player).position, " I'm alive , i am the king!")
		}
		r = r.Next()
	}
}
