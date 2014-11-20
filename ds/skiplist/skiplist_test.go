//Author: Jonathan Fragoso Pérez

//Testing Skip List

package skiplist

import(
	"testing"
)

func TestNewSkipList(t *testing.T){
	newSkipList := New()
	if newSkipList.Len() != 0 {
		t.Errorf("A new skip list should contain 0 elements")
	}
	if !newSkipList.IsEmpty() {
		t.Errorf("A new skip list should be empty")
	}
}

func TestAdd(t *testing.T){

}

func TestRemove(t *testing.T){
	
}

func BenchmarkSkipList(b *testing.B){
	//b.N
}