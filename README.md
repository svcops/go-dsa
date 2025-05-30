# Golang DSA

> Data Structures and Algorithms in Golang

## 排序

[冒泡排序](pkg/sorter/bubble)

[插入排序](pkg/sorter/insert)

[快速排序](pkg/sorter/quick)

- [单路快排](pkg/sorter/quick/oneway.go)
- [双路快排](pkg/sorter/quick/twoway.go)

## 查找

### 二分搜索树

[二分搜索树的接口定义](pkg/tree/bst/basic.go)

[非平衡二分搜索树的实现](pkg/tree/bst/unbalanced)

[avl树的实现](pkg/tree/bst/avl)

### 堆

[最大堆的接口定义](pkg/tree/maxheap/basic.go)

[定长最大堆的实现](pkg/tree/maxheap/fixedlength/impl.go)

## 图

### 图的定义

[权重图的定义](pkg/graph/definition.go)

[稀疏图的实现](pkg/graph/sparse/graph.go)

[稠密图的实现](pkg/graph/dense/graph.go)

### 图中的算法

[无向图求联通分量个数](pkg/graph/algo/components)

[最小生成树 LazyPrim](pkg/graph/algo/mst/lazyprim/lazyprim.go)

[最小生成树 Prim](pkg/graph/algo/mst/prim/prim.go)

[单源最短路径 dijkstra](pkg/graph/algo/dijkstra)

[环分析](pkg/graph/algo/cycles/cycleAnalyzer.go)

[求图中的环](pkg/graph/algo/ring)

- [有向图求环](pkg/graph/algo/cycles/directed.go)
- [无向图求环](pkg/graph/algo/cycles/undirected.go)

[深度遍历求两点之前的所有路径](pkg/graph/algo/dfs/dfs.go)

## 动态规划

[one hundred](pkg/dp/hundred/solution.go)

[kmp](pkg/dp/kmp/kmp.go)

- todo : `createNext`
