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
			title, link := readTitleAndLink(fname)
			data.WriteString("\n|" + title + "|[" + fn[1] + "](" + link + ")|[GO](./src/" + fname + ")|")
		}
	}

	content := data.String()
	ioutil.WriteFile("./README.md", []byte(content), 0644)
}

func readTitleAndLink(fname string) (string, string) {
	fi, _ := os.Open("./src/" + fname)

	defer fi.Close()

	br := bufio.NewReader(fi)
	title := ""
	link := ""
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data := string(a)
		if strings.Index(data, "@title") > 0 {
			title = data[10:]
		}
		if strings.Index(data, "@see") > 0 {
			link = data[8:]
			break
		}
	}
	return title, link
}
