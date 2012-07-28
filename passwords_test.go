package passwords

import (
	"testing"
	_ "fmt"
)

func TestCompute(t * testing.T) {
	password := Compute("garrydanger", 3)

	a := ComputeWithSalt("garrydanger", 3, password.Salt)
	if a.Hash != password.Hash {
		t.Fatal("nonderministic hash detected")
	}

}
