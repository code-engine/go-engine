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

func (p ProcDir) ProcList() ([]int, error) {
	out := []int{}

	pdl, err := p.RawList()

	if err != nil {
		return out, err
	}

	sort.Ints(pdl)

	for _, prc := range pdl {
		out = append(out, prc)
	}

	return out, nil
}

func (p ProcDir) RawList() ([]int, error) {
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

	return out, err
}
