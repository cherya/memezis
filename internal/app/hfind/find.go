package hfind

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

type HStore struct {
	store map[Hash]struct{}
}

type Match struct {
	Hash  Hash
	Score int
}

type Matches []Match

func (m Matches) Len() int           { return len(m) }
func (m Matches) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matches) Less(i, j int) bool { return m[i].Score < m[j].Score }

type Hash uint64

func (h Hash) String() string {
	return fmt.Sprintf("%x", uint64(h))
}

func NewHStore(hashes []Hash) *HStore {
	h := make(map[Hash]struct{})
	for _, hash := range hashes {
		h[hash] = struct{}{}
	}
	return &HStore{store: h}
}


func (h *HStore) Add(p Hash) {
	h.store[p] = struct{}{}
}

func (h *HStore) AddHexString(p string) error {
	b, err := FromString(p)
	if err != nil {
		return err
	}
	h.store[b] = struct{}{}
	return nil
}

func FromString(s string) (Hash, error) {
	u, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, errors.New("string must be hex of uint64")
	}

	return Hash(u), nil
}

func (h *HStore) FindKNN(hash Hash, k int) Matches {
	res := make(Matches, 0)

	for v := range h.store {
		sim := hamming(hash, v)
		res = insertSort(res, Match{
			Hash:  v,
			Score: sim,
		})
		if len(res) > k {
			res = res[:k]
		}
	}

	return res
}

// Hamming returns the normalized similarity value.
// (https://en.wikipedia.org/wiki/Hamming_distance)
func hamming(hash1, hash2 Hash) int {
	x := (hash1 ^ hash2) & ((1 << uint64(64)) - 1)
	tot := 0
	for x != 0 {
		tot += 1
		x &= x - 1
	}
	return tot
}

func insertSort(data []Match, el Match) []Match {
	index := sort.Search(len(data), func(i int) bool { return data[i].Score > el.Score })
	data = append(data, Match{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}
