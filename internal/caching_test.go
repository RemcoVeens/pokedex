package internal_test

import (
	"testing"
	"time"

	// Rename the imported package to internal for clarity in the test file
	"github.com/RemcoVeens/pokedex/internal"
)

// TestReapLoop verifies that old entries are correctly removed by the reapLoop
// after the configured time interval has passed.
func TestReapLoop(t *testing.T) {
	// Use a short, controlled interval for testing purposes (100 milliseconds).
	const reapInterval = 100 * time.Millisecond

	// 1. Initialize the cache. The reapLoop starts immediately.
	c := internal.NewCache(reapInterval)

	// 2. Add an entry (A) that is meant to be reaped.
	const keyA = "entry-A"
	dataA := []byte("old-data")
	c.Add(keyA, dataA)

	// 3. Wait for longer than the reap interval.
	// This ensures the entry ages past the limit AND the reapLoop gets at least one tick.
	time.Sleep(reapInterval + 50*time.Millisecond)

	// Wait a tiny bit more to ensure the reapLoop goroutine has completed its cleanup cycle.
	time.Sleep(10 * time.Millisecond)

	// 4. Check if entry A was reaped.
	_, foundA := c.Get(keyA)
	if foundA {
		t.Errorf("FAIL: Entry %s was NOT reaped after %v elapsed.", keyA, reapInterval)
	}

	// 5. Add a new entry (B) after the reaping has (likely) just finished.
	// This entry should remain in the cache until the next reap cycle.
	const keyB = "entry-B"
	dataB := []byte("new-data")
	c.Add(keyB, dataB)

	// Check immediately that B is present.
	data, foundB := c.Get(keyB)
	if !foundB || string(data) != string(dataB) {
		t.Errorf("FAIL: Entry %s was not found immediately after setting it.", keyB)
	}

	// 6. Wait for the next reap cycle to complete.
	t.Logf("Waiting for second reap cycle (%v) to clear entry B...", reapInterval)
	time.Sleep(reapInterval + 50*time.Millisecond)
	time.Sleep(10 * time.Millisecond) // Wait for execution

	// 7. Check if entry B was reaped in the second cycle.
	_, foundBFinal := c.Get(keyB)
	if foundBFinal {
		t.Errorf("FAIL: Entry %s was NOT reaped after the second interval elapsed.", keyB)
	}
}
