package process

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
)

func NewProcDir() ProcDir {
	return ProcDir{
		Path: "/proc",
	}
}

type ProcDir struct {
	Path string
}

func (p ProcDir) List() ([]int, error) {
	files, err := ioutil.ReadDir(p.Path)

	if err != nil {
		return []int{}, err
	}

	out := []int{}

	r := regexp.MustCompile(`^\d+$`)

	for _, file := range files {
		if !r.MatchString(file.Name()) {
			continue
		}

		proc, err := strconv.Atoi(file.Name())

		if err != nil {
			return out, err
		}

		out = append(out, proc)
	}

	sort.Ints(out)

	return out, err
}

func (p ProcDir) Exists(pid int) (bool, error) {
	list, err := p.List()

	if err != nil {
		return false, err
	}

	if len(list) <= 0 {
		return false, err
	}

	i := sort.Search(len(list), func(i int) bool { return list[i] == pid })

	if list[i] == pid {
		return true, nil
	}

	return false, nil
}
