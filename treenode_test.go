package treenode

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

// TestNormalize calls normalize
func TestNormalize(t *testing.T) {

	type normalizeTest struct {
		input  treeSlice
		result treeSlice
	}

	var normalizeTests = []normalizeTest{
		{
			input: treeSlice{
				NewInt(3),
				NewInt(5),
				NewInt(1),
				NewInt(6),
				NewInt(2),
				NewInt(9),
				NewInt(8),
				NewNil(),
				NewNil(),
				NewInt(7),
				NewInt(4),
			},
			result: treeSlice{
				NewInt(3),
				NewInt(5),
				NewInt(1),
				NewInt(6),
				NewInt(2),
				NewInt(9),
				NewInt(8),
				NewNil(),
				NewNil(),
				NewInt(7),
				NewInt(4),
			},
		},
		// 8,3,10,1,6,null,14,null,null,4,7,13
		// 8,3,10,1,6,null,14,null,null,4,7,nul,nul,13
		{
			input: treeSlice{
				NewInt(8),
				NewInt(3),
				NewInt(10),
				NewInt(1),
				NewInt(6),
				NewNil(),
				NewInt(14),
				NewNil(),
				NewNil(),
				NewInt(4),
				NewInt(7),
				NewInt(13),
			},
			result: treeSlice{
				NewInt(8),
				NewInt(3),
				NewInt(10),
				NewInt(1),
				NewInt(6),
				NewNil(),
				NewInt(14),
				NewNil(),
				NewNil(),
				NewInt(4),
				NewInt(7),
				NewNil(),
				NewNil(),
				NewInt(13),
			},
		},
		{
			input: treeSlice{
				NewInt(1),
				NewInt(2),
				NewInt(3),
			},
			result: treeSlice{
				NewInt(1),
				NewInt(2),
				NewInt(3),
			},
		},
	}

	for _, test := range normalizeTests {
		//input := make(treeSlice, 0, len(test.input))
		//copy(input, test.input)
		test.input.normalize(0, 0)

		assert.Equal(t, len(test.input), len(test.result))

		for i, v := range test.input {
			assert.ObjectsAreEqualValues(v, test.result[i])
			assert.Equal(t, v.null, test.result[i].null)
			assert.Equal(t, v.value, test.result[i].value)
		}
	}
}

// TestToTreeSlice calls toTreeSlice
func TestToTreeSlice(t *testing.T) {

	type toTreeSliceTest struct {
		input  *TreeString
		result treeSlice
		err    error
	}

	var toTreeSliceTests = []toTreeSliceTest{
		{
			input: NewTreeString("1,2,3"),
			result: treeSlice{
				NewInt(1),
				NewInt(2),
				NewInt(3),
			},
		},
		{
			input: NewTreeString("1,null,2,null,3,null,4,null,5"),
			result: treeSlice{
				NewInt(1),
				NewNil(),
				NewInt(2),
				NewNil(),
				NewInt(3),
				NewNil(),
				NewInt(4),
				NewNil(),
				NewInt(5),
			},
		},
		{
			input:  NewTreeString("1, null,null"),
			result: nil,
			err: &strconv.NumError{
				Func: "Atoi",
				Num:  " null",
				Err:  nil,
			},
		},
		{
			input: NewTreeString("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"),
			result: treeSlice{
				NewInt(3),
				NewInt(5),
				NewInt(1),
				NewInt(6),
				NewInt(7),
				NewInt(4),
				NewInt(2),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewInt(9),
				NewInt(8),
			},
		},
	}

	for _, test := range toTreeSliceTests {
		result, err := test.input.toTreeSlice()

		if err != nil {
			assert.Equal(t, test.err != nil, true)
		}

		assert.Equal(t, len(result), len(test.result))
		assert.Equal(t, result, test.result)
	}
}

// TestGetTreeNode calls GetTreeNode
func TestGetTreeNode(t *testing.T) {

	type GetTreeNodeTest struct {
		input  *TreeString
		result *TreeNode
		err    error
	}

	var getTreeNodeTests = []GetTreeNodeTest{
		{
			input: NewTreeString("1,2,3"),
			result: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			input: NewTreeString("1,null,2,null,3,null,4,null,5"),
			result: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:  4,
							Left: nil,
							Right: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
		},
		{
			input: NewTreeString("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"),
			result: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
	}

	for _, test := range getTreeNodeTests {
		tn, err := test.input.GetTreeNode()

		assert.Equal(t, test.err, err)
		assert.Equal(t, test.result, tn)
		if !reflect.DeepEqual(test.result, tn) {
			assert.Failf(t, "TestGetTreeNode fail", "Different Trees: %v , %v", *test.result, *tn)
		}
	}
}
