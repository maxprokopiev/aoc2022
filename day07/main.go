package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func Solve2(input string) int {
	root := BuildFs(input)

	return dfs2(root, 30000000-(70000000-root.size), root.size)
}

func Solve1(input string) int {
	root := BuildFs(input)

	return dfs(root)
}

func dfs(node *dir) int {
	total := 0
	for _, c := range node.children {
		if !c.isFile {
			if c.size <= 100000 {
				total += c.size
			}
			total += dfs(c)
		}
	}

	return total
}

func dfs2(node *dir, min int, current int) int {
	for _, c := range node.children {
		if !c.isFile {
			if (c.size >= min) && (c.size < current) {
				current = dfs2(c, min, c.size)
			} else {
				current = dfs2(c, min, current)
			}
		}
	}

	return current
}

type dir struct {
	name     string
	size     int
	isFile   bool
	parent   *dir
	children map[string]*dir
}

func newDir(name string, parent *dir) *dir {
	d := dir{name: name, parent: parent, children: make(map[string]*dir), isFile: false}

	return &d
}

func newFile(name string, size int, parent *dir) *dir {
	f := dir{name: name, parent: parent, size: size, isFile: true}

	return &f
}

var reCd = regexp.MustCompile(`\$ cd (.+)`)
var reDir = regexp.MustCompile(`dir (.+)`)
var reFile = regexp.MustCompile(`(\d+) (.+)`)

func BuildFs(input string) *dir {
	root := dir{name: "/", children: make(map[string]*dir)}
	var cd *dir

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cmd := scanner.Text()
		switch {
		case cmd == "$ cd /":
			cd = &root
		case cmd == "$ cd ..":
			size := cd.size
			cd = cd.parent
			if cd.name != "/" {
				cd.size += size
			}
		case cmd == "$ ls":
			// noop
		case reCd.MatchString(cmd):
			ms := reCd.FindStringSubmatch(cmd)
			cd = cd.children[ms[1]]
		case reDir.MatchString(cmd):
			ms := reDir.FindStringSubmatch(cmd)
			cd.children[ms[1]] = newDir(ms[1], cd)
		case reFile.MatchString(cmd):
			ms := reFile.FindStringSubmatch(cmd)
			size, _ := strconv.Atoi(ms[1])
			cd.children[ms[2]] = newFile(ms[2], size, cd)
			cd.size += size
		}
	}

	for _, c := range root.children {
		if !c.isFile {
			root.size += c.size
		}
	}

	return &root
}

func main() {
}
