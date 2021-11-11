package main

import (
	"sort"
	"testing"
)

func TestWalkDir(t *testing.T) {
	all, err := WalkAllFilesInDir("C:\\Users\\mainj\\Desktop\\v2x\\New folder1")
	if err != nil {
		t.Error("fails")
	}
	sort.Strings(all)

	for _, file := range all {
		println(file)
	}
}
