package shard

import (
	"testing"
)

func assertSurrounding(t *testing.T, r *Remotes, pos, pred, succ byte) {
	surr := r.surrounding([]byte{pos})
	if surr.Predecessor.Pos[0] != pred {
		t.Errorf("Wrong Predecessor for %v in %v, wanted %v but got %v", pos, r, pred, surr.Predecessor)
	}
	if surr.Successor.Pos[0] != succ {
		t.Errorf("Wrong Successor for %v in %v, wanted %v but got %v", pos, r, succ, surr.Successor)
	}
}

func TestRemotes(t *testing.T) {
	r := &Remotes{}
	r.add(Remote{[]byte{0}, "a"})
	r.add(Remote{[]byte{1}, "b"})
	r.add(Remote{[]byte{2}, "c"})
	r.add(Remote{[]byte{3}, "d"})
	r.add(Remote{[]byte{4}, "e"})
	r.add(Remote{[]byte{6}, "f"})
	r.add(Remote{[]byte{7}, "g"})
	assertSurrounding(t, r, 0, 0, 1)
	assertSurrounding(t, r, 1, 1, 2)
	assertSurrounding(t, r, 2, 2, 3)
	assertSurrounding(t, r, 3, 3, 4)
	assertSurrounding(t, r, 4, 4, 6)
	assertSurrounding(t, r, 5, 4, 6)
	assertSurrounding(t, r, 6, 6, 7)
	assertSurrounding(t, r, 7, 7, 0)
}