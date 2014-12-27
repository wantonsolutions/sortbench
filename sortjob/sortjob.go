package sortjob

import (
	"sort"
)

type SortJob struct {
	Name string
	Size int
	Sort func(sort.Interface)
	Ordering string
}
func New(iname string,isize int,isort func(sort.Interface), iordering string) *SortJob {
	return &SortJob{ Name: iname, Size: isize, Sort: isort, Ordering: iordering }
}
