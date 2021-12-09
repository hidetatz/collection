package collection

import (
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	al := NewArrayList[int]()
	MustEqual(t, []int{}, al.values)
	al.Add(1)
	MustEqual(t, []int{1}, al.values)
	al.Add(2)
	MustEqual(t, []int{1, 2}, al.values)
}

func TestArrayList_AddAll(t *testing.T) {
	al := NewArrayList[int]()
	al.AddAll([]int{1})
	MustEqual(t, []int{1}, al.values)

	al.AddAll([]int{2, 3, 4})
	MustEqual(t, []int{1, 2, 3, 4}, al.values)

	al.AddAll([]int{5})
	MustEqual(t, []int{1, 2, 3, 4, 5}, al.values)

	al.AddAll([]int{})
	MustEqual(t, []int{1, 2, 3, 4, 5}, al.values)
}

func TestArrayList_AddAt(t *testing.T) {
	al := NewArrayList[int]()

	err := al.AddAt(1, 1)
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAt(0, 1))
	MustEqual(t, []int{1}, al.values)

	err = al.AddAt(3, 2)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.AddAt(2, 2)
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAt(1, 2))
	MustEqual(t, []int{1, 2}, al.values)

	MustBeNil(t, al.AddAt(1, 3))
	MustEqual(t, []int{1, 3, 2}, al.values)

	MustBeNil(t, al.AddAt(0, 4))
	MustEqual(t, []int{4, 1, 3, 2}, al.values)
}

func TestArrayList_AddAllAt(t *testing.T) {
	al := NewArrayList[int]()

	err := al.AddAllAt(1, []int{1})
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAllAt(0, []int{1}))
	MustEqual(t, []int{1}, al.values)

	err = al.AddAllAt(2, []int{2, 3, 4})
	MustBeErr(t, err, ErrInvalidIndex)

	MustBeNil(t, al.AddAllAt(1, []int{2, 3, 4}))
	MustEqual(t, []int{1, 2, 3, 4}, al.values)

	MustBeNil(t, al.AddAllAt(2, []int{5, 6}))
	MustEqual(t, []int{1, 2, 5, 6, 3, 4}, al.values)
}

func TestArrayList_Clear(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2})
	al.Clear()
	MustEqual(t, []int{}, al.values)

	al.AddAll([]int{3, 4})
	MustEqual(t, []int{3, 4}, al.values)

	al.Clear()
	MustEqual(t, []int{}, al.values)
}

func TestArrayList_Get(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3})

	_, err := al.Get(-1)
	MustBeErr(t, err, ErrInvalidIndex)

	_, err = al.Get(3)
	MustBeErr(t, err, ErrInvalidIndex)

	ret, err := al.Get(2)
	MustBeNil(t, err)
	MustEqual(t, ret, 3)
}

func TestArrayList_IsEmpty(t *testing.T) {
	al := NewArrayList[int]()
	MustEqual(t, true, al.IsEmpty())

	al.Add(1)
	MustEqual(t, false, al.IsEmpty())

	al.Clear()
	MustEqual(t, true, al.IsEmpty())
}

func TestArrayList_Len(t *testing.T) {
	al := NewArrayList[int]()
	MustEqual(t, 0, al.Len())

	al.Add(1)
	MustEqual(t, 1, al.Len())

	al.AddAll([]int{2, 3, 4})
	MustEqual(t, 4, al.Len())
}

func TestArrayList_Set(t *testing.T) {
	al := NewArrayList[int]()

	err := al.Set(-1, 1)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.Set(1, 1)
	MustBeErr(t, err, ErrInvalidIndex)

	al.AddAll([]int{1, 2, 3})

	err = al.Set(3, 4)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.Set(2, 4)
	MustBeNil(t, err)
	MustEqual(t, []int{1, 2, 4}, al.values)
}

func TestArrayList_RemoveAt(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3})

	err := al.RemoveAt(-1)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.RemoveAt(3)
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.RemoveAt(1)
	MustBeNil(t, err)
	MustEqual(t, []int{1, 3}, al.values)
}

func TestArrayList_RemoveRange(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})

	err := al.RemoveRange(2, 1) // from must be less than to
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.RemoveRange(-1, 1) // must not be negative
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.RemoveRange(1, 7) // must not be too large values
	MustBeErr(t, err, ErrInvalidIndex)

	err = al.RemoveRange(1, 1)
	MustBeNil(t, err)
	MustEqual(t, []int{1, 2, 3, 4, 5, 6}, al.values)

	err = al.RemoveRange(1, 3)
	MustBeNil(t, err)
	MustEqual(t, []int{1, 4, 5, 6}, al.values)

	err = al.RemoveRange(0, 4)
	MustBeNil(t, err)
	MustEqual(t, []int{}, al.values)
}

func TestArrayList_RemoveIf(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})

	al.RemoveIf(func(i, v int) bool { return v%2 == 0 })
	MustEqual(t, []int{1, 3, 5}, al.values)
}

func TestArrayList_Slice(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	s := al.Slice()
	MustEqual(t, []int{1, 2, 3, 4, 5, 6}, al.values)
	MustEqual(t, s, al.values)
}

func TestArrayList_SubList(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	n, err := al.SubList(2, 1) // from must be less than to
	MustBeErr(t, err, ErrInvalidIndex)

	n, err = al.SubList(-1, 1) // must not be negative
	MustBeErr(t, err, ErrInvalidIndex)

	n, err = al.SubList(1, 7) // must not be too large values
	MustBeErr(t, err, ErrInvalidIndex)

	n, err = al.SubList(1, 1)
	MustBeNil(t, err)
	MustEqual(t, []int{}, n.values)

	n, err = al.SubList(1, 3)
	MustBeNil(t, err)
	MustEqual(t, []int{2, 3}, n.values)
}

func TestArrayList_ReplaceAll(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	al.ReplaceAll(func(v int) int { return v * 2 })
	MustEqual(t, []int{2, 4, 6, 8, 10, 12}, al.values)
}

func TestArrayList_Map(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	r := al.Map(func(v int) int { return v * 2 })
	MustEqual(t, []int{2, 4, 6, 8, 10, 12}, r.values)
	MustEqual(t, []int{1, 2, 3, 4, 5, 6}, al.values) // original list must not be changed
}

func TestArrayList_ForEach(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	buff := []int{}
	al.ForEach(func(index, v int) {
		buff = append(buff, index*v)
	})
	MustEqual(t, []int{0, 2, 6, 12, 20, 30}, buff)
	MustEqual(t, []int{1, 2, 3, 4, 5, 6}, al.values) // original list must not be changed
}

func TestArrayList_Filter(t *testing.T) {
	al := NewArrayList[int]()

	al.AddAll([]int{1, 2, 3, 4, 5, 6})
	r := al.Filter(func(index, v int) bool {
		return v%2 == 0
	})
	MustEqual(t, []int{2, 4, 6}, r.values)
	MustEqual(t, []int{1, 2, 3, 4, 5, 6}, al.values) // original list must not be changed
}