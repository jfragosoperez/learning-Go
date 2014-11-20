//Author: Jonathan Fragoso Pérez

// Skip list data structure.
//
// Skip lists use a probabilistic approach in order to perform
// search/insert/removal operations in O(logn), and outperform
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



//Skip List
type SkipList struct{
	decisionCoin coin 
	numberOfLevels, numberOfElements int 
	head *skipListNode
}

func New() *SkipList {
	return new(SkipList)
}

// List of operations supported by a skip list
type skipListOps interface{
	Add(value interface{}) bool
	AddAll(values []interface{}) bool
	Remove(value interface{}) bool
	HasElement(value interface{}) bool
	Len() int
	IsEmpty() bool
}

// when we insert a new node, we have to calculate the level of replication
// of this node in upper level lists
func (skipList *SkipList) calculateInsertionReplicationLevel()int {
	numberOfLevels := 0
	for {
		if !skipList.decisionCoin.flip().isTails() {
			numberOfLevels++
		}
		break
	}
	return numberOfLevels
}

// Adds desired element, if was not inserted into the skip list. 
// returns true if the element is in the list.
func (skipList *SkipList) Add(value interface{}) bool{
	return true
}

// Adds all the elements into the skip list.
// returns true if the elements are now in the list.
func (skipList *SkipList) AddAll(values []interface{}) bool{
	return true
}

// Removes the desired element of the skip list.
// returns true if the value has been removed from the list, false 
// if the element was not contained in the list.
func(skipList *SkipList) Remove(value interface{}) bool{
	return true
}

// Searches the desired value and returns true if is contained in the list.
func(skipList *SkipList) HasElement(value interface{}) bool{
	return true
}

// Returns the number of elements contained in the list.
// NOTE: this number corresponds to the size of the list
// of nodes that contains all the values, so it does not take into
// account the replicated elements in other levels of the list.
func(skipList *SkipList) Len() int{
	return skipList.numberOfElements
}

// Describes if the skip list is empty.
func(skipList *SkipList) IsEmpty() bool{
	return skipList.head == nil
}

//A Skip List is composed of skip list nodes
type skipListNode struct{
	right *skipListNode
	below *skipListNode
	data interface{}
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