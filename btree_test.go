package btree

import (
    "testing"
//    "math"
    )

func TestCases(t *testing.T) {
    tree := NewTree()
    if tree.Count != 1 {
        t.Errorf("expecting 1 node")
    }
    tree.Insert(Int(1));
    if !tree.Has(Int(1)) {
        t.Errorf("tree should have an element with key 1")
    }        
}
