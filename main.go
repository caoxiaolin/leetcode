package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	createReadme()
}

func createReadme() {
	data := bytes.Buffer{}
	data.WriteString(`[![Build Status](https://www.travis-ci.org/caoxiaolin/leetcode.svg?branch=master)](https://www.travis-ci.org/caoxiaolin/leetcode)
[![codecov.io](https://codecov.io/github/caoxiaolin/leetcode/coverage.svg?branch=master)](https://codecov.io/github/caoxiaolin/leetcode?branch=master)

# leetcode

| # | Title | Solution |
|---| ----- | -------- |`)
	path := "./src/"
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		fname := f.Name()
		fn := strings.Split(fname, ".")
		_, err := strconv.Atoi(fn[0])
		if err == nil && strings.Index(fn[1], "_test") == -1 {
			link := readLink(fname)
			data.WriteString("\n|" + fn[0] + "|[" + fn[1] + "](" + link + ")|[GO](./src/" + fname + ")|")
		}
	}

	content := data.String()
	ioutil.WriteFile("./README.md", []byte(content), 0644)
}

func readLink(fname string) string {
	fi, _ := os.Open("./src/" + fname)

	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data := string(a)
		if strings.Index(data, "@see") > 0 {
			return data[8:]
		}
	}
	return ""
}
