// Package hundred
// @Description: 有100块钱，现在有面额1，2，5元的零钱，100块有多少种找零方式

package hundred

const unused = true

//
// CombinationNum
//  @Description: 获取组合的个数
//  @param i
//  @return int
//
func CombinationNum(i int) int {
	if i <= 0 {
		return 0
	}
	com := getCom(i, make(map[int]*combination))
	return len(com.nc)
}

//
// getCom
//  @Description: getCombination 或者组合个数
//  @param i
//  @return combination
//
func getCom(i int, cache map[int]*combination) *combination {
	if i == 1 {
		cn1cache, ok := cache[1]
		if ok {
			return cn1cache
		}
		set := make(map[int]bool)
		set[1] = unused
		cn1 := &combination{
			total: 1,
			nc:    set,
		}
		cache[1] = cn1
		return cn1
	}

	if i == 2 {

		cn2cache, ok := cache[2]
		if ok {
			return cn2cache
		}

		set := make(map[int]bool)
		set[1] = unused
		set[2] = unused

		cn2 := &combination{
			total: 2,
			nc:    set,
		}
		cache[2] = cn2
		return cn2
	}

	var com1, com2 *combination

	contains1, ok1 := cache[i-1]
	if ok1 {
		com1 = merge(i, getCom(1, cache), contains1)
	} else {
		com1 = merge(i, getCom(1, cache), getCom(i-1, cache))
	}

	contains2, ok2 := cache[i-2]
	if ok2 {
		com2 = merge(i, getCom(1, cache), contains2)
	} else {
		com2 = merge(i, getCom(2, cache), getCom(i-2, cache))
	}

	if i < 5 {
		cnLt5cache, okLt5 := cache[i]
		if okLt5 {
			return cnLt5cache
		}
		cnLt5 := distinct(i, com1, com2)
		cache[i] = cnLt5
		return cnLt5
	}

	if i == 5 {
		cn5cache, ok5 := cache[5]
		if ok5 {
			return cn5cache
		}

		cn5 := distinct(i, com1, com2)
		cn5.nc[1] = unused
		cache[5] = cn5
		return cn5
	}

	cnGt5cache, okGt5 := cache[i]
	if okGt5 {
		return cnGt5cache
	}

	var com5 *combination
	contains5, ok5 := cache[i-5]
	if ok5 {
		com5 = merge(i, getCom(5, cache), contains5)
	} else {
		com5 = merge(i, getCom(5, cache), getCom(i-5, cache))
	}

	cnGt5 := multiDistinct(i, com1, com2, com5)
	cache[i] = cnGt5
	return cnGt5
}

func multiDistinct(total int, cs ...*combination) *combination {
	resMap := make(map[int]bool)
	for _, c := range cs {
		for n := range c.nc {
			resMap[n] = unused
		}
	}
	return &combination{
		total: total,
		nc:    resMap,
	}
}

func distinct(total int, a, b *combination) *combination {
	resMap := make(map[int]bool)
	aNc, bNc := a.nc, b.nc
	for aN := range aNc {
		resMap[aN] = unused
	}

	for bN := range bNc {
		resMap[bN] = unused
	}

	return &combination{
		total: total,
		nc:    resMap,
	}
}

//
//
//  @Description: merge ,利用set的特性
//  @param total
//  @param a
//  @param b
//  @return combination
//
func merge(total int, a, b *combination) *combination {
	resMap := make(map[int]bool)
	aNc, bNc := a.nc, b.nc
	for aN := range aNc {
		for bN := range bNc {
			resMap[aN+bN] = unused
		}
	}
	return &combination{
		total: total,
		nc:    resMap,
	}
}

//
//  combination
//  @Description: nc : 组合个数 当set使用
//  @Description: nc : number of combinations
//
type combination struct {
	total int
	nc    map[int]bool
}
