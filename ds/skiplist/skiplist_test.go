//Author: Jonathan Fragoso Pérez

//Testing Skip List

package skiplist

import(
	"testing"
)

func TestNewSkipList(t *testing.T){
	newSkipList := NewIntSkipList(NON_DICREASING, 32)
	if newSkipList.Len() != 0 {
		t.Errorf("A new skip list should contain 0 elements")
	}
	if !newSkipList.IsEmpty() {
		t.Errorf("A new skip list should be empty")
	}
}

func TestAdd(t *testing.T){
    newSkipList := NewIntSkipList(NON_DICREASING, 32)
    newSkipList.Add(2)
}

func TestRemove(t *testing.T){
	
}

func TestAddNil(t *testing.T){
	newSkipList := NewIntSkipList(NON_DICREASING, 32)

	defer func() {
    	if error := recover(); error == nil {
    		t.Errorf("adding nil element should have panicked")
    	} 
    }()

    newSkipList.Add(nil)
}

func TestSearchNil(t *testing.T){
	newSkipList := NewIntSkipList(NON_DICREASING, 32)

	defer func() {
    	if error := recover(); error == nil {
    		t.Errorf("searching nil element should have panicked")
    	} 
    }()

    newSkipList.HasElement(nil)
}

func TestRemoveNil(t *testing.T){
	newSkipList := NewIntSkipList(NON_DICREASING, 32)

	defer func() {
    	if error := recover(); error == nil {
    		t.Errorf("trying to remove nil element should have panicked")
    	} 
    }()

    newSkipList.Remove(nil)
}

func BenchmarkSkipList(b *testing.B){
	//b.N
}