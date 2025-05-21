package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	cache := NewCache(interval)

	for _, tc := range cases {
		cache.Add(tc.key, tc.val)
		actual, ok := cache.Get(tc.key)
		if !ok {
			t.Errorf("expected to find key %s", tc.key)
			continue
		}
		if string(actual) != string(tc.val) {
			t.Errorf("expected value %s, got %s", string(tc.val), string(actual))
		}
	}
}

func TestGetNonExistent(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	_, ok := cache.Get("nonexistentkey")
	if ok {
		t.Errorf("expected not to find key 'nonexistentkey'")
	}
}

func TestReapLoop(t *testing.T) {
	const baseInterval = 50 * time.Millisecond
	const waitMultiplier = 3

	cache := NewCache(baseInterval)

	key1 := "reapKey1"
	val1 := []byte("reapData1")
	cache.Add(key1, val1)

	time.Sleep(baseInterval * waitMultiplier)

	_, ok1 := cache.Get(key1)
	if ok1 {
		t.Errorf("expected key %s to be reaped, but it was found", key1)
	}

	key2 := "reapKey2"
	val2 := []byte("reapData2")
	cache.Add(key2, val2)

	time.Sleep(baseInterval / 2)

	val2Actual, ok2 := cache.Get(key2)
	if !ok2 {
		t.Errorf("expected key %s to be found, but it was not", key2)
	}
	if string(val2Actual) != string(val2) {
		t.Errorf("expected value %s for key %s, got %s", string(val2), key2, string(val2Actual))
	}

	key3 := "reapKey3"
	val3 := []byte("reapData3")
	cache.Add(key3, val3)

	sleepToReapKey2ButNotKey3 := (baseInterval / 2) + (5 * time.Millisecond)
	time.Sleep(sleepToReapKey2ButNotKey3)

	_, ok2AfterWait := cache.Get(key2)
	if ok2AfterWait {
		t.Errorf("expected key %s to be reaped after longer wait, but it was found", key2)
	}

	val3Actual, ok3 := cache.Get(key3)
	if !ok3 {
		t.Errorf("expected key %s to be found, but it was not", key3)
	}
	if string(val3Actual) != string(val3) {
		t.Errorf("expected value %s for key %s, got %s", string(val3), key3, string(val3Actual))
	}
}

func TestCacheConcurrency(t *testing.T) {
	const interval = 100 * time.Millisecond // Interval for the cache
	cache := NewCache(interval)

	numGoroutines := 50
	numOpsPerGoroutine := 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(gID int) {
			defer wg.Done()
			for j := 0; j < numOpsPerGoroutine; j++ {
				key := fmt.Sprintf("key-%d-%d", gID, j)
				val := []byte(fmt.Sprintf("val-%d-%d", gID, j))

				cache.Add(key, val)

				retrievedVal, ok := cache.Get(key)
				if !ok {
				} else if string(retrievedVal) != string(val) {
					t.Errorf("goroutine %d: expected value %s for key %s, got %s", gID, string(val), key, string(retrievedVal))
				}

				otherKeyIndex := (j + 1) % numOpsPerGoroutine
				otherGID := (gID + 1) % numGoroutines
				otherKey := fmt.Sprintf("key-%d-%d", otherGID, otherKeyIndex)
				cache.Get(otherKey)
			}
		}(i)
	}

	wg.Wait()

	t.Log("TestCacheConcurrency completed. Run with -race flag to detect race conditions.")
}
