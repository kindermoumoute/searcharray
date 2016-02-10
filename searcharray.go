package searcharray

type SearchType int
type SearchResult int

// Enumerated constants are created using the iota enumerator
const (
	LessThan SearchType = iota
	LessThanEquals
	Equals
	GreaterThanEquals
	GreaterThan
)

const (
	NotFound SearchResult = iota
	FoundExact
	FoundGreater
	FoundLess
)

const (
	DESCENDING = iota
	ASCENDING
)

var SuccessResult = map[SearchType]SearchResult{LessThan: FoundLess, LessThanEquals: FoundLess, GreaterThanEquals: FoundGreater, GreaterThan: FoundGreater}

/* Search an array of sorted numbers.
*
* items		: An array of sorted ints, with no duplicates
* n_items	: Number of elements in the items array
* ascending	: non-zero if the array is sorted in ascending order
* key		: the key to search for
* srchType	: the type of match to find
*
* This function finds the element in the array
* that best fits the search criteria. It returns
* the match type and the index of the matching item.
*
*
* Assumptions
* -----------
*	The items are sorted
*	Items will be non-NULL
*	There are no duplicate items
*	n_items will be > 0
 */
func Search(items []int, n_items int, ascending int, key int, srchType SearchType) (index int, sr SearchResult) {

	// Default result values
	index = -1
	sr = NotFound

	// Ascending : +1
	// Descending : -1
	way := ascending*2 - 1

	// Dichotomy starts at n_items - 1
	mask := (n_items - 1)
	modulo := n_items % 2

	// Current index
	i := 0
	for {
		i += mask
		if key < items[i] && ascending == ASCENDING || key > items[i] && ascending == DESCENDING {
			i -= mask
		}

		// Shift the mask
		mask >>= 1

		// Adjustment if n_items is odd
		if mask == 0 && modulo == 1 {
			mask = 1
			modulo = 0
		}

		// Popular conditions
		equal := items[i] == key
		noMask := mask == 0
		lastItem := i == n_items-1

		// Works with 3 conditions :
		// 1- Returns FoundExact cases (for Equals, GreaterThanEquals and LessThanEquals)
		// 2- Returns the next/previous greater/lower or NotFound (for LessThan, GreaterThan, GreaterThanEquals and LessThanEquals)
		// 3- Returns NotFound (for Equals)

		if equal && (srchType != GreaterThan && srchType != LessThan) {
			return i, FoundExact
		} else if (equal || noMask || lastItem) && (srchType != Equals) {
			if (srchType == LessThan || srchType == LessThanEquals) && items[i] >= key {
				i -= way
			} else if (srchType == GreaterThan || srchType == GreaterThanEquals) && items[i] <= key {
				i += way
			}
			if i >= 0 && i < n_items {
				return i, SuccessResult[srchType]
			} else {
				return
			}
		} else if noMask || lastItem {
			return
		}
	}
}

/*
* LessThan
* --------
*  Finds the largest item which is less than the key.
*  It returns FoundLess if a match is found, NotFound
*  if no match is found.
*
* LessThanEquals
* --------------
*  Finds the item which is equal to the key, or the
*  largest item which is less than the key. Returns
*  FoundExact if an item that exactly matches the key
*  is found, FoundLess if a non-exact match is found
*  and NotFound if no match is found.
*
* Equals
* ------
*  Finds an item which is equal to the key. Returns
*  FoundExact if an item if found, NotFound otherwise.
*
* GreaterThanEquals
* -----------------
*  Finds the item which is equal to the key, or the
*  smallest item which is greater than the key. Returns
*  FoundExact if an item that exactly matches the key
*  is found, FoundGreater if a non-exact match is found
*  and NotFound if no match is found.
*
* GreaterThan
* -----------
*  Finds the smallest item which is greater than the
*  key. Returns FoundGreater if a match if found, NotFound
*  if no match is found.
 */
