package cycles

import (
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
			analyzer.dfsFindCycles(vertex, visited, make([]string, 0), make(map[string]bool), result)
		}
	}
	return result
}

func (analyzer *cycleAnalyzer) dfsFindCycles(vertex string, visited map[string]bool, path []string, marked map[string]bool, result *AnalyzerResult) {
	visited[vertex] = true
	path = append(path, vertex)
	marked[vertex] = true

	log.Printf("开始遍历顶点: %s, 当前路径: %v", vertex, path)

	edges := analyzer.g.Adj(vertex)
	for to, _ := range edges {
		if !visited[to] {
			// 路径副本
			pathSnapshot := make([]string, len(path))
			copy(pathSnapshot, path)
			log.Printf("顶点 %s 的邻接点 %s 未被访问，继续深度遍历", vertex, to)
			analyzer.dfsFindCycles(to, visited, pathSnapshot, marked, result)
		} else {
			// 当map的key不存在的时候，它的返回值为这个类型的默认返回值
			if marked[to] == true {
				log.Printf("顶点 %s 的邻接点 %s 找到环，并记录", vertex, to)
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
				pathSnapshot := make([]string, len(path))
				copy(pathSnapshot, path)
				analyzer.dfsFindCycles(to, visited, pathSnapshot, marked, result)
			}
		}
	}
	marked[vertex] = false
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
	// todo
}
