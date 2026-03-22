// 重要结论: nxn 的棋盘要放 n 个皇后, 就必须每行每列一个皇后
// 更数学: 一个任意摆放了 n 个皇后的 nxn 棋盘, 这种局面是 N 皇后的解的必要条件是每行每列有且仅有一个皇后
// 证明:
// i.  仅有:
//        如果有一行/列有 n 个 (n > 1) 个皇后,
//        则这一行/列上的皇后必然能相互攻击, 与假设矛盾, 否定
// ii. 有:
//        如果某一行/列上没有皇后, 那么 n 个皇后必然分布在其他的 n - 1 行/列
//        于是就必然有一行/列的皇后数量大于 1, 根据 i. 的推论, 与假设矛盾, 否定
func solveNQueens(n int) [][]string {
    // path 是搜索路径
    // 物理意义: x = path[y]
    // 表示纵坐标为 y 的皇后的横坐标为 x
    // 也就是说我们是横着看这个棋盘的
    // 当然也可以竖着看这个棋盘, 答案一定是关于对角线对称的, 此处横着看
	path := make([]int, 0, n)
    answer := [][]string{}

    // 讨论剩余的 remain 个皇后没放的问题
    var backtrack func(int)
    backtrack = func(remain int) {
        if len(path) == n {
            tmp := make([]string, 0, n)
            for _, x := range path {
                // 构造一行
                row := make([]byte, n)
                for i := range n {
                    if i == x {
                        row[i] = 'Q'
                    } else {
                        row[i] = '.'
                    }
                }
                tmp = append(tmp, string(row))
            }
            answer = append(answer, tmp)
        }

        for x := range n {
            if valid(path, x) {
                path = append(path, x)
                backtrack(remain - 1)
                path = path[:len(path)-1]
            }
        }
    }

    backtrack(n)
    return answer
}

func valid(path []int, newX int) bool {
    // 新皇后一定在最后一排
    newY := len(path)
    for y, x := range path {
        dx, dy := newX - x, newY - y
        // dy 不可能为 0
        if dx == dy || dy == -dx || dx == 0 {
            return false
        }
    }

    return true
}
