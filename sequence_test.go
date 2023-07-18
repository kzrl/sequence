package sequence

import (
	"sync"
	"testing"
)



// TestNextStart tests that we get 1 after the first call.
func TestNextStart(t *testing.T) {
	Reset()
	got := Next()
	want := 1
	if got != want {
		t.Fatalf("Expect first call of Next() == %d, got Next() == %d", want, got)
	}
}

// TestNextThousand tests that we get 1000 after 1000 calls.
func TestNextThousand(t *testing.T) {
	Reset()
	var got int
	calls := 1000
	want := calls
	for i := 0; i < calls; i++ {
		got = Next()
	}
	
	if got != want {
		t.Fatalf("Expect 1000th call of Next() == %d, got Next() == %d", want, got)
	}
}


// TestNextThreadSafe tests we can call Next() from goroutines safely.
// run with go test -race
func TestNextThreadSafe(t *testing.T) {
	Reset()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			Next()
		}()

	}
	wg.Wait()
	got := Next()
	want := 11
	if got != want {
		t.Fatalf("Expect 11th call of Next() == %d, got Next() == %d", want, got)
	}
}

// TestReset omitted due to time constraints. You get the idea.
func TestReset(t *testing.T) {
}	
