package lcs

import "sort"

//split the compared string to []rune
//record the pos and val
type comparedValPos struct {
	val rune
	pos int
}
type lcsTargetPosInfo struct {
	slicePos int
	rawPos   int
}
type comparedValPosSlice []comparedValPos

func (b comparedValPosSlice) Len() int { return len(b) }
func (b comparedValPosSlice) Less(i, j int) bool {
	if b[i].val < b[j].val {
		return true
	} else if b[i].val > b[j].val {
		return false
	} else {
		return !(b[i].pos < b[j].pos)
	}
}
func (b comparedValPosSlice) Swap(i, j int) { b[i], b[j] = b[j], b[i] }


func lcsSort(target string) comparedValPosSlice {
	var slice comparedValPosSlice

	runeSlice := []rune(target)
	for i := 0; i < len(runeSlice); i++ {
		var l comparedValPos
		l.pos = i
		l.val = runeSlice[i]
		slice = append(slice, l)
	}
	sort.Sort(slice)
	return slice
}

func findMatchList(ch rune, slice comparedValPosSlice, left int, right int, start *int) (matchLen int) {
	var middle, matchedLen1, matchedLen2 int
	var start1, start2 int
	if left > right {
		matchLen = 0
		return
	} else if left == right {
		if ch == slice[left].val {
			*start = left
			matchLen = 1
			return
		}
		matchLen = 0
		return
	}

	middle = (left + right) >> 1
	if slice[middle].val < ch {
		matchedLen1 = findMatchList(ch, slice, middle+1, right, &start1)
	} else if slice[middle].val > ch {
		matchedLen1 = findMatchList(ch, slice, left, middle-1, &start1)
	} else {
		matchedLen1 = findMatchList(ch, slice, left, middle-1, &start1)
		matchedLen2 = findMatchList(ch, slice, middle+1, right, &start2) + 1
		if matchedLen1 == 0 {
			start1 = middle
		}
		matchedLen1 += matchedLen2
	}
	*start = start1
	matchLen = matchedLen1
	return
}

func matchListLcs(rawStr string, slice comparedValPosSlice) (matchedSlice []lcsTargetPosInfo) {
	var start int
	runeSlice := []rune(rawStr)
	for i := 0; i < len(runeSlice); i++ {
		matchLen := findMatchList(runeSlice[i], slice, 0, len(slice)-1, &start)
		for k := 0; k < matchLen; k++ {
			var l lcsTargetPosInfo
			l.slicePos = slice[start+k].pos
			l.rawPos = i
			matchedSlice = append(matchedSlice, l)
		}
	}
	return
}
func findPos(l []int, currLen int, value int) int {
	left := 0
	right := currLen - 1
	middle := 0
	for left <= right {
		middle = (left + right) >> 1
		if l[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
func lis(slice []lcsTargetPosInfo, incSeq []int) int {
	if len(slice) == 0 {
		return 0
	}
	L := make([]int, len(slice))
	M := make([]int, len(slice))
	prev := make([]int, len(slice))
	L[0] = slice[0].slicePos
	M[0] = 0
	prev[0] = -1
	currLen := 1
	for i := 1; i < len(slice); i++ {
		pos := findPos(L, currLen, slice[i].slicePos)
		L[pos] = slice[i].slicePos
		M[pos] = i
		if pos > 0 {
			prev[i] = M[pos-1]
		} else {
			prev[i] = -1
		}
		if pos+1 > currLen {
			currLen++
		}
	}

	pos := M[currLen-1]
	for i := currLen - 1; i >= 0 && pos != -1; i-- {
		incSeq[i] = slice[pos].rawPos
		pos = prev[pos]
	}
	return currLen

}
//Lcs func
//int return the length of  lcs
//string return a possible lcs
func Lcs(rawStr string, targetStr string) (int ,string) {
	 slice :=  lcsSort(targetStr)
	 return calcLcsUsingLis(rawStr,slice)
}

func calcLcsUsingLis(rawStr string, slice comparedValPosSlice) (int, string) {
	var maxIncreaseSequencerLen int
	if len(rawStr) > slice.Len() {
		maxIncreaseSequencerLen = slice.Len()

	} else {
		maxIncreaseSequencerLen = len(rawStr)
	}
	lcsPosSlice := matchListLcs(rawStr, slice)
	increaseSequnceSlice := make([]int, maxIncreaseSequencerLen)
	ret := lis(lcsPosSlice, increaseSequnceSlice)
	var result []rune
	runeSlice := []rune(rawStr)
	for i := 0; i < ret; i++ {
		result = append(result, runeSlice[increaseSequnceSlice[i]])
	}

	return ret, string(result)
}
