package gateway

import (
	"os"
	"bufio"
)

type Contacts struct {
	Path string
}

func (c *Contacts) FetchNumbers() []string {
    fp, err := os.Open(c.Path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	var ret []string
    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
		ret = append(ret, scanner.Text())
    }

	return ret
}
