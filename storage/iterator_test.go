package storage

import (
	"io/ioutil"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestIterator(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "goka_storage_TestIterator")
	assertNil(t, err)

	db, err := leveldb.OpenFile(tmpdir, nil)
	assertNil(t, err)

	st, err := New(db)
	assertNil(t, err)

	kv := map[string]string{
		"key-1": "val-1",
		"key-2": "val-2",
		"key-3": "val-3",
	}

	for k, v := range kv {
		assertNil(t, st.Set([]byte(k), []byte(v)))
	}

	assertNil(t, st.SetOffset(777))

	iter, err := st.Iterator()
	assertNil(t, err)
	defer iter.Release()
	count := 0

	// accessing iterator before Next should only return nils
	val, err := iter.Value()
	assertTrue(t, val == nil)
	assertNil(t, err)

	for iter.Next() {
		count++
		key := string(iter.Key())
		expected, ok := kv[key]
		if !ok {
			t.Fatalf("unexpected key from iterator: %s", key)
		}

		val, err := iter.Value()
		assertNil(t, err)
		assertEqual(t, expected, string(val))
	}
	assertEqual(t, count, len(kv))
}
