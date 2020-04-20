package gateway

import (
	"os"
    "strings"
    "encoding/csv"
    "github.com/yuukichi-nankou/casample-go/domain"
)

type Contacts struct {
	Path string
}

func (c *Contacts) FetchContacts() []domain.Contact {
    fp, err := os.Open(c.Path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    reader := csv.NewReader(fp)
    var ret []domain.Contact
    var line []string
    for {
        line, err = reader.Read()

        if err != nil {
            break
        }

        con := domain.Contact{
            Name: line[0],
            Number: c.formatNumber(line[1]),
        }
        ret = append(ret, con)
    }

	return ret
}

// 080-XXXX-YYYY to +8180XXXXYYYY
func (c *Contacts) formatNumber(n string) string {
    // delete -
    n = strings.Replace(n, "-", "", -1)
    // delete first 0
    n = strings.TrimLeft(n, "0")
    
	return "+81" + n
}