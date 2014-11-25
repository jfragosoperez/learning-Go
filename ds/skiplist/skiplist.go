//Author: Jonathan Fragoso Pérez

// Skip list data structure.
//
// Skip lists use a probabilistic approach in order to perform
// search/insert/removal operations in O(lgn), and outperform
// better than several balanced trees.

// Running time:
// a) Average case of executing the operations (happens almost always):
// 		(log n) for all -> search, add, remove
// b) Worst case:
// 		n for all -> search, add, remove
//

//REFERENCES:
//	-"A Skip List Cookbook" - Pugh, W. - 1990
//
//	-"Skip Lists: A probabilistic alternative to balanced trees" - Pugh, W. - 1990
//
//  -Introduction to algorithms, MIT lecture slides -
//  http://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-046j-introduction-to-algorithms-sma-5503-fall-2005/video-lectures/lecture-12-skip-lists/lec12.pdf
//
//  -Introduction to algorithms, MIT video lecture -
//  http://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-046j-introduction-to-algorithms-sma-5503-fall-2005/video-lectures/lecture-12-skip-lists/
//
 

package skiplist

import (
	"math/rand"
)

//A Skip List is composed of skip list nodes
type skipListNode struct{
	right []*skipListNode
	value interface{}
}

//to know if one node has one next
func (node *skipListNode) hasNext() bool {
	return len(node.right) == 0 
}

// next node in the skip 
func (node *skipListNode) next() *skipListNode {
	if node.hasNext() {
		return nil
	}
	return node.right[0]
}

type compareResult int

const(
	//case equal (onEl == otherEl)
	EQUAL compareResult = 1 + iota
	//case we need to go on (in non-dicreasing list will be when onEl < otherEl)
	TRUE
	//case we have passed (in non-dicreasing list will be when onEl > otherEl)
	FALSE
)

//Skip List
type SkipList struct{
	compare func(oneEl, otherEl interface{}) compareResult
	decisionCoin coin  
	length int
	head *skipListNode
	maxLevel int
}

//the order in which the elements need to be sorted.
type Order int

const(
	NON_DICREASING Order = 1 + iota 
	DICREASING
)

//default max level of the skip list
const DefaultMaxLevel = 28

func NewIntSkipList(order Order, maxLevel int) *SkipList {
	return newSkipList(func(oneEl, otherEl interface{}) compareResult {
		if oneEl.(int) == otherEl.(int) {
			return EQUAL
		}
		if (order == NON_DICREASING && oneEl.(int) < otherEl.(int)) || 
			(order == DICREASING && oneEl.(int) > otherEl.(int)) {
			return TRUE
		} else {
			return FALSE
		}
	}, maxLevel)
}

func newSkipList(compare func(oneEl, otherEl interface{}) compareResult, maxLevel int) *SkipList {
	return &SkipList{
		compare: compare,
		head: &skipListNode{
			right: []*skipListNode{nil},
		},
		maxLevel: maxLevel,
	}
}

// List of operations supported by a skip list
type skipListOps interface{
	Add(value interface{}) bool
	AddAll(values []interface{}) bool
	Remove(value interface{}) bool
	HasElement(value interface{}) bool
	Len() int
	IsEmpty() bool
	GetMaxLevel() int
}

// Adds desired element, if was not inserted into the skip list. 
// returns the value inserted and a confirmation to know that the value 
// has been inserted correctly.
func (skipList *SkipList) Add(value interface{}) (v interface{}, ok bool){
	if value == nil {
		panic("cannot insert nil elements")
	}
	return value, true
}

// Removes the element of the skip list.
// returns true if the value has been removed from the list, false 
// if the element was not contained in the list.
func(skipList *SkipList) Remove(value interface{}) (v interface{}, ok bool){
	if value == nil {
		panic("cannot remove nil elements")
	}
	return value, true
}

// Searches the desired value and returns true if is contained in the list.
func(skipList *SkipList) HasElement(value interface{}) bool{
	if value == nil {
		panic("cannot search nil elements")
	}

	current := skipList.head
	depth := len(current.right) - 1

	for i := depth; i >= 0; i-- {
		for current.right[i] != nil && (skipList.compare(current.right[i].value, value) == TRUE) {
			current = current.right[i]
		}
		if current.right[i] != nil && (skipList.compare(current.right[i].value, value) == EQUAL) {
			return true
		}
	}
	return false
}

// Returns the number of elements contained in the list.
// NOTE: this number corresponds to the size of the list
// of nodes that contains all the values, so it does not take into
// account the replicated elements in other levels of the list.
func(skipList *SkipList) Len() int{
	return skipList.length
}

// Describes if the skip list is empty.
func(skipList *SkipList) IsEmpty() bool{
	return skipList.Len() == 0
}

// returns the max level of the list
func(skipList *SkipList) GetMaxLevel() int{
	return skipList.maxLevel
}

// when we insert a new node, we have to calculate the level of replication
// of this node in upper level lists
func (skipList *SkipList) calculateInsertionReplicationLevel()int {
	numberOfLevels := 0
	for {
		if !skipList.decisionCoin.flip().isTails() {
			numberOfLevels++
		}else {
			break
		}	
	}
	return numberOfLevels
}

// This coin will be used to make the decision of replicating nodes
// to next levels once we add a new element.
type coin struct { }

func(c coin) flip() coinResult{
	if randomValue := random(0,1); randomValue < 0.5 {
		return TAILS
	} else {
		return HEADS
	}
}

func random(min, max float64) float64 {
  return rand.Float64() * (max - min) + min
}

//Coin Result is the result we get from flipping a coin
type coinResult int

const(
	HEADS coinResult = 1 + iota 
	TAILS
)

func (result coinResult) isTails() bool {
	return result == TAILS 
}

func (result coinResult) isHeads() bool {
	return result == HEADS
}