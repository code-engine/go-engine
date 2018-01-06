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
