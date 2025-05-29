package cycles

import (
	"fmt"
	"go-dsa/pkg/graph"
	"log"
	"sort"
)

type cycleAnalyzer struct {
	g     graph.Graph
	debug bool
}

func CreateCycleAnalyzer(g graph.Graph, debug bool) *cycleAnalyzer {
	return &cycleAnalyzer{
		g:     g,
		debug: debug,
	}
}

func (analyzer *cycleAnalyzer) FindCycles() *AnalyzerResult {
	visited := make(map[string]bool)
	result := &AnalyzerResult{
		direct:   analyzer.g.IsDirect(),
		cycles:   [][]string{},
		cycleZip: make(map[string]bool),
	}
	vertices := analyzer.g.Vertexes()
	for vertex := range vertices {
		if !visited[vertex] {
			analyzer.dfsFindCycles(vertex, visited, make([]string, 0), make(map[string]bool), result, 1)
		}
	}
	return result
}

func (analyzer *cycleAnalyzer) dfsFindCycles(vertex string, visited map[string]bool, path []string, marked map[string]bool,
	result *AnalyzerResult, depth int) {
	visited[vertex] = true
	path = append(path, vertex)
	marked[vertex] = true

	indent := fmt.Sprintf("%d  ", depth)
	for range depth {
		indent += " - "
	}

	if analyzer.debug {
		log.Printf("%s开始遍历顶点: %s, 当前路径: %v", indent, vertex, path)
	}

	edges := analyzer.g.Adj(vertex)
	for to, _ := range edges {
		if !visited[to] {
			// 路径副本
			pathSnapshot := make([]string, len(path))
			copy(pathSnapshot, path)
			if analyzer.debug {
				log.Printf("%s顶点 %s 的邻接点 %s 未被访问，继续深度遍历", indent, vertex, to)
			}
			analyzer.dfsFindCycles(to, visited, pathSnapshot, marked, result, depth+1)
		} else {
			// 当map的key不存在的时候，它的返回值为这个类型的默认返回值
			if marked[to] == true {
				if analyzer.debug {
					log.Printf("%s顶点 %s 的邻接点 %s 找到环，并记录", indent, vertex, to)
				}
				// 发现环,记录环
				cycle := make([]string, 0)
				for index, v := range path {
					if v == to {
						// 找到环的起点
						cycle = append(cycle, path[index:]...) // 从起点到当前路径的所有节点
						result.addCycle(cycle)
						break
					}
				}
			} else {
				// important: 如果邻接点已经被访问过，但没有被标记为在当前路径中，则继续深度遍历
				pathSnapshot := make([]string, len(path))
				copy(pathSnapshot, path)
				analyzer.dfsFindCycles(to, visited, pathSnapshot, marked, result, depth+1)
			}
		}
	}
	if analyzer.debug {
		log.Printf("%s出栈 %s 节点，因为不确定是否还有其他点 ??? --> %s", indent, vertex, vertex)
	}
	marked[vertex] = false
	if analyzer.debug {
		log.Printf("%s顶点 %s 的所有邻接点已遍历完毕，回溯到上一级", indent, vertex)
	}
}

type AnalyzerResult struct {
	direct bool // 是否是有向图
	// 所有的环
	cycles [][]string

	// 环排序压缩去重用
	cycleZip map[string]bool
}

func Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, item := range slice {
		result = fn(result, item)
	}
	return result
}

func (result *AnalyzerResult) addCycle(cycle []string) {
	if len(cycle) == 0 {
		return
	}
	// cycle 排序
	cycleCopy := make([]string, len(cycle))
	copy(cycleCopy, cycle)
	sort.Strings(cycleCopy)
	sentence := Reduce(cycleCopy, "", func(acc, word string) string {
		if acc == "" {
			return word
		}
		return acc + " " + word
	})

	if result.cycleZip[sentence] {
		return // 已经存在
	} else {
		result.cycleZip[sentence] = true
		result.cycles = append(result.cycles, cycle)
	}
}

func (result *AnalyzerResult) PrintCycles() {
	if len(result.cycles) == 0 {
		fmt.Println("没有发现环")
		return
	}

	fmt.Printf("发现 %d 个环\n", len(result.cycles))
	for i, cycle := range result.cycles {
		fmt.Printf("环 %d: ", i+1)
		printCycle(result.direct, cycle)
	}
	fmt.Println()
}

func printCycle(direct bool, cycle []string) {
	fmt.Printf("Printing Cycle|Vertex's num = %d|vertexes = %v\n", len(cycle), cycle)

	lineStart := "  "
	if len(cycle) == 2 {
		fmt.Println(lineStart + cycle[0] + " <=> " + cycle[1])
		fmt.Println()
		return
	}
	isEven := len(cycle)%2 == 0
	var mid int
	if isEven {
		mid = len(cycle) / 2
	} else {
		mid = len(cycle)/2 + 1
	}
	// 打印上半部分

	upperLine := lineStart
	up := cycle[0:mid]
	for i, v := range up {
		if i == len(up)-1 {
			upperLine += v
		} else {
			if direct {
				upperLine += v + " -> "
			} else {
				upperLine += v + " - "
			}
		}
	}
	fmt.Println(upperLine)
	// 打印中间部分
	var midStart string
	var midEnd string
	if direct {
		midStart = lineStart + "↑"
		if isEven {
			midEnd = "↓"
		} else {
			midEnd = "↙"
		}
	} else {
		midStart = lineStart + "|"
		if isEven {
			midEnd = "|"
		} else {
			midEnd = "/"
		}
	}
	midLine := midStart
	var sub int
	if isEven {
		sub = 2
	} else {
		if direct {
			sub = 4
		} else {
			sub = 3
		}
	}
	for i := 0; i < len(upperLine)-sub-len(lineStart); i++ {
		midLine += " "
	}
	midLine += midEnd
	fmt.Println(midLine)

	// 打印下半部分
	down := cycle[mid:]
	downLine := lineStart
	for i := len(down) - 1; i >= 0; i-- {
		v := down[i]
		if i == 0 {
			downLine += v
		} else {
			if direct {
				downLine += v + " <- "
			} else {
				downLine += v + " - "
			}
		}
	}
	fmt.Println(downLine)
	fmt.Println()
}
