package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func ReadPuzzle(filePath string) ([]int, []int) {

	file, err := os.OpenFile(filePath, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	c1, c2 := readColumns(file)

	return c1, c2
}

func readColumns(r io.ReadCloser) ([]int, []int) {
	col1, col2 := []int{}, []int{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		vals := strings.Split(line, "   ")
		i1, _ := strconv.ParseInt(vals[0], 10, 64)
		i2, _ := strconv.ParseInt(vals[1], 10, 64)

		col1 = append(col1, int(i1))
		col2 = append(col2, int(i2))

	}

	return col1, col2
}
