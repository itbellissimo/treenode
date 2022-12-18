package treenode

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// TreeString present tree in string line
//
// Example:
// 3,5,1,6,7,4,2,null,null,null,null,null,null,9,8
//
// present tree:
//
//			          3
//		  	    5           1
//		  6         7    4        2
//	                            9    8
type TreeString struct {
	s        string
	treeNode *TreeNode
}

// treeSlice
type treeSlice []NilInt

// normalize helper function. It checks string on empty null tree nodes, and add its.
// 8,3,10,1,6,null,14,null,null,4,7,13, where 13 is the value of the last node level 3
// After normalize
// 8,3,10,1,6,null,14,null,null,4,7,nul,nul,13
func (tl *treeSlice) normalize(start int, lvl int) {
	levelIndex0 := 0
	for l := 0; l < lvl; l++ {
		levelIndex0 += int(math.Pow(2.0, float64(l)))
	}

	limit := int(math.Pow(2.0, float64(lvl)))

	for i := start; i < len(*tl); i++ {

		if (*tl)[i].null {
			posAddNull := levelIndex0 + limit + (i-levelIndex0)*2
			if len(*tl) > posAddNull {
				// Making space for the new element
				*tl = append(*tl, NewNil(), NewNil())
				copy((*tl)[posAddNull+2:], (*tl)[posAddNull:])
				(*tl)[posAddNull] = NewNil()
				(*tl)[posAddNull+1] = NewNil()

				tl.normalize(i+1, lvl)
				return
			} else {
				return
			}
		}

		if i-levelIndex0 == limit {
			levelIndex0 = i
			limit = 2 * limit
			lvl++
		}
	}

	return
}

// treeNode
// [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8] [0:1] 0
// [5,1,6,7,4,2,null,null,null,null,null,null,9,8] [0:2] 1
// [6,7,4,2,null,null,null,null,null,null,9,8] [0:4] 2
// [null,null,null,null,null,null,9,8] [0:8] 3
//
//			          3
//		  	    5           1
//		  6         7    4        2
//	                            9    8
func (tl *treeSlice) treeNode() (*TreeNode, error) {
	if (*tl)[0].null {
		return nil, errors.New("wrong root value")
	}

	lvl := 1
	tmp := (*tl)[1:]

	byLvl := make(map[int][]NilInt)

	byLvl[0] = make([]NilInt, 1, 1)
	byLvl[0][0] = (*tl)[0]

	// 3,10,1,6,null,14,null,null,4,7,13
	// 1,6,null,14,null,null,4,7,13
	// null,null,4,7,13
	for len(tmp) != 0 {
		end := int(math.Pow(2.0, float64(lvl)))
		if end >= len(tmp) {
			end = len(tmp)
		}
		values := tmp[:end]
		tmp = tmp[end:]

		byLvl[lvl] = make([]NilInt, 0, len(values))
		for i := 0; i < len(values); i = i + 2 {
			if i+1 < len(values) {
				byLvl[lvl] = append(byLvl[lvl], values[i], values[i+1])
			} else {
				byLvl[lvl] = append(byLvl[lvl], values[i])
			}
		}

		lvl++
	}

	lvl = len(byLvl) - 1
	nodeByLvl := make(map[int][]*TreeNode)

	for n := lvl; n >= 0; n-- {
		nodeByLvl[n] = make([]*TreeNode, len(byLvl[n]), len(byLvl[n]))
		for index, v := range byLvl[n] {
			var (
				l *TreeNode
				r *TreeNode
			)

			if n <= lvl {
				if len(nodeByLvl[n+1])-1 >= index*2 && nodeByLvl[n+1][index*2] != nil {
					l = nodeByLvl[n+1][index*2]
				}
				if len(nodeByLvl[n+1])-1 >= index*2+1 && nodeByLvl[n+1][index*2+1] != nil {
					r = nodeByLvl[n+1][index*2+1]
				}
			}

			if v.null {
				nodeByLvl[n][index] = nil
			} else {
				nodeByLvl[n][index] = &TreeNode{
					Val:   v.value,
					Left:  l,
					Right: r,
				}
			}
		}
	}

	return nodeByLvl[0][0], nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NilInt struct {
	value int
	null  bool
}

func (n *NilInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(x int) NilInt {
	return NilInt{x, false}
}

func NewNil() NilInt {
	return NilInt{0, true}
}

func NewTreeString(s string) *TreeString {
	return &TreeString{s: s}
}

func (ts *TreeString) toTreeSlice() (treeSlice, error) {
	sl := strings.Split(ts.s, ",")

	treeVals := make(treeSlice, 0, len(sl))
	for _, v := range sl {
		if v == "null" {
			treeVals = append(treeVals, NewNil())
			continue
		}

		if vInt, err := strconv.Atoi(v); err != nil {
			return nil, err
		} else {
			treeVals = append(treeVals, NewInt(vInt))
		}
	}

	return treeVals, nil
}

func (ts *TreeString) GetTreeNode() (*TreeNode, error) {
	treeVals, err := ts.toTreeSlice()
	if err != nil {
		return nil, err
	}

	// check and fill empty tree nodes
	treeVals.normalize(0, 0)

	return treeVals.treeNode()
}
