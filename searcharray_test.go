package searcharray

import (
	"fmt"
	"testing"
)

type arrayTest struct {
	in     []int
	order  int
	result []searchArrayTest
}

type searchArrayTest struct {
	key   int
	typ   SearchType
	ret   SearchResult
	index int
}

/*
* Examples
* --------
*  Given the input array [0, 2, 4, 6, 8] (ascending order)
*
*  +-----+-------------------+--------------+-------+
*  | Key | Type              | Returns      | Index |
*  +-----+-------------------+--------------+-------+
*  | -1  | LessThanEquals    | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  0  | LessThan          | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  0  | Equals            | FoundExact   | 0     |
*  +-----+-------------------+--------------+-------+
*  |  1  | Equals            | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  2  | GreaterThanEquals | FoundExact   | 1     |
*  +-----+-------------------+--------------+-------+
*  |  2  | GreaterThan       | FoundGreater | 2     |
*  +-----+-------------------+--------------+-------+
*
*  Given the input array [8, 6, 4, 2, 0] (descending order)
*
*  +-----+-------------------+--------------+-------+
*  | Key | Type              | Returns      | Index |
*  +-----+-------------------+--------------+-------+
*  | -1  | LessThan          | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  0  | LessThan          | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  4  | LessThanEquals    | FoundExact   | 2     |
*  +-----+-------------------+--------------+-------+
*  |  8  | Equals            | FoundExact   | 0     |
*  +-----+-------------------+--------------+-------+
*  |  5  | GreaterThanEquals | FoundGreater | 1     |
*  +-----+-------------------+--------------+-------+
*  |  2  | GreaterThanEquals | FoundExact   | 3     |
*  +-----+-------------------+--------------+-------+
*  |  8  | GreaterThan       | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*  |  9  | GreaterThan       | NotFound     | X     |
*  +-----+-------------------+--------------+-------+
*
 */
var golden = []arrayTest{

	{[]int{-5, -2, 4, 6, 8, 10, 11, 14, 17, 18, 19, 20, 34}, ASCENDING,
		[]searchArrayTest{
			{-2, Equals, FoundExact, 1},
			{11, Equals, FoundExact, 6},
			{34, Equals, FoundExact, 12},
			{35, Equals, NotFound, -1},
			{21, Equals, NotFound, -1},
			{-6, LessThan, NotFound, -1},
			{11, LessThan, FoundLess, 5},
			{16, LessThan, FoundLess, 7},
			{35, LessThan, FoundLess, 12},
			{-5, LessThan, NotFound, -1},
			{-6, GreaterThan, FoundGreater, 0},
			{10, GreaterThan, FoundGreater, 6},
			{34, GreaterThan, NotFound, -1},
			{35, GreaterThan, NotFound, -1},
			{6, LessThanEquals, FoundExact, 3},
			{34, LessThanEquals, FoundExact, 12},
			{-5, LessThanEquals, FoundExact, 0},
			{35, LessThanEquals, FoundLess, 12},
			{15, LessThanEquals, FoundLess, 7},
			{-6, LessThanEquals, NotFound, -1},
			{6, GreaterThanEquals, FoundExact, 3},
			{34, GreaterThanEquals, FoundExact, 12},
			{-5, GreaterThanEquals, FoundExact, 0},
			{-6, GreaterThanEquals, FoundGreater, 0},
			{15, GreaterThanEquals, FoundGreater, 8},
			{35, GreaterThanEquals, NotFound, -1},
		}},

	{[]int{-4, -6, -16, -31}, DESCENDING,
		[]searchArrayTest{
			{-4, Equals, FoundExact, 0},
			{-31, Equals, FoundExact, 3},
			{-32, Equals, NotFound, -1},
			{-17, Equals, NotFound, -1},
			{-31, LessThan, NotFound, -1},
			{-32, LessThan, NotFound, -1},
			{-6, LessThan, FoundLess, 2},
			{-2, LessThan, FoundLess, 0},
			{-6, GreaterThan, FoundGreater, 0},
			{-40, GreaterThan, FoundGreater, 3},
			{-4, GreaterThan, NotFound, -1},
			{2, GreaterThan, NotFound, -1},
		}},

	{[]int{0, 2, 4, 6, 8}, ASCENDING,
		[]searchArrayTest{
			{-1, LessThanEquals, NotFound, -1},
			{0, LessThan, NotFound, -1},
			{0, Equals, FoundExact, 0},
			{1, Equals, NotFound, -1},
			{2, GreaterThanEquals, FoundExact, 1},
			{2, GreaterThan, FoundGreater, 2},
		}},

	{[]int{8, 6, 4, 2, 0}, DESCENDING,
		[]searchArrayTest{
			{-1, LessThan, NotFound, -1},
			{0, LessThan, NotFound, -1},
			{4, LessThanEquals, FoundExact, 2},
			{8, Equals, FoundExact, 0},
			{5, GreaterThanEquals, FoundGreater, 1},
			{2, GreaterThanEquals, FoundExact, 3},
			{8, GreaterThan, NotFound, -1},
			{9, GreaterThan, NotFound, -1},
		}},
}

var (
	typeString   = map[SearchType]string{LessThan: "LessThan\t", LessThanEquals: "LessThanEquals", Equals: "Equals\t", GreaterThanEquals: "GreaterThanEquals", GreaterThan: "GreaterThan\t"}
	resultString = map[SearchResult]string{NotFound: "NotFound", FoundExact: "FoundExact", FoundGreater: "FoundGreater", FoundLess: "FoundLess"}
	orderString  = map[int]string{DESCENDING: "descending", ASCENDING: "ascending"}
	lineString   = "+-------+-----------------------+---------------+-------+\n"
)

func TestGolden(t *testing.T) {
	for _, aTest := range golden {
		fmt.Printf("Given the input array  %v (%s order)\n%s| Key\t| Type\t\t\t| Returns\t| Index\t|\n%s", aTest.in, orderString[aTest.order], lineString, lineString)
		for _, test := range aTest.result {
			index, result := Search(aTest.in, len(aTest.in), aTest.order, test.key, test.typ)
			if test.ret == result && (index == -1 || test.index == index) {
				fmt.Printf("|  %d\t| %s\t| %s\t| %v\t|\n%s", test.key, typeString[test.typ], resultString[result], index, lineString)
			} else {

				t.Fatalf("\n|  %d\t| %s\t| %s\t| %v\t|\n", test.key, typeString[test.typ], resultString[result], index)
			}
		}
	}
}
