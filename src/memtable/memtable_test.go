package memtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupSuite() *SkipList {
	skipList := NewMemTable(StringComparer{}, 1)

	// Return a function to teardown the test
	return skipList
}

func TestMemTablePut_Insert(t *testing.T) {
	skipList := setupSuite()

	skipList.Put("1", "1-val")

	val := skipList.Get("1")

	assert.NotNil(t, val)

	assert.Equal(t, *val, "1-val")
}

func TestMemTablePut_Update(t *testing.T) {
	skipList := setupSuite()

	skipList.Put("1", "1-val")
	skipList.Put("1", "new")

	val := skipList.Get("1")

	assert.NotNil(t, val)

	assert.Equal(t, *val, "new")
}
