// package cn
package Week04

func robotSim(commands []int, obstacles [][]int) int {
	POINTBASE := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	//add the obstacles
	obsAll := map[int]map[int]bool{}
	for _, value := range obstacles {
		if _, ok := obsAll[value[0]]; ok {
			obsAll[value[0]][value[1]] = true
		} else {
			obsAll[value[0]] = make(map[int]bool)
			obsAll[value[0]][value[1]] = true
		}
	}
	p := point{
		0, 0,
	}
	directIdx := 0 // 表示方向目前从北开始，即0，1方向
	max := 0
	for _, command := range commands {
		switch command {
		case -1:
			directIdx = (directIdx + 1) % 4
		case -2:
			directIdx = (directIdx + 3) % 4
		default:
			for i := 0; i < command; i++ {
				p.run(POINTBASE[directIdx])
				if _, ok := obsAll[p.X]; ok {
					if okY := obsAll[p.X][p.Y]; okY {
						p.back(POINTBASE[directIdx])
					}
				}
				if p.length() > max {
					max = p.length()
				}
			}
		}
	}
	return max
}

type point struct {
	X int
	Y int
}

func (item *point) length() int {
	return item.X*item.X + item.Y*item.Y

}

func (item *point) run(x point) {
	item.X = item.X + x.X
	item.Y = item.Y + x.Y
}

func (item *point) back(x point) {
	item.X = item.X - x.X
	item.Y = item.Y - x.Y
}
