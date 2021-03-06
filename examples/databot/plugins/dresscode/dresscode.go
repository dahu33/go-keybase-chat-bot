package dresscode

import (
	"encoding/csv"
	"os"
	"fmt"
	"time"
	"strings"
	"hash/fnv"
)

const timezone = "Asia/Tokyo"

var location *time.Location

type Dresscodes struct {
	Styles []string
}

func LoadDresscodes(filename string) Dresscodes {
	location, _ = time.LoadLocation(timezone)
	d := Dresscodes{Styles:[]string{}}
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return d
	}
	csv := csv.NewReader(file)
	for row, err := csv.Read(); row != nil; row, err = csv.Read() {
		if err != nil {
			fmt.Println(err.Error())
			return d
		}
		style := strings.TrimSpace(row[0])
		d.Styles = append(d.Styles, style)
	}
	fmt.Printf("Loaded %d dresscodes\n", len(d.Styles)) 
	return d
}

func (d *Dresscodes) RespondToDresscode(msg string) string {
	// what day is it?
	datestr := time.Now().In(location).Format("Mon Jan 2 2006")
	if strings.Contains(datestr, "Oct 31") {
		return "Today's dress code is HALLOWEEN!!!!"
	}
	// compute the hash
	h := hash(datestr)
	// modulo t
	idx := h % len(d.Styles)
	return fmt.Sprintf("Today's dress code is %v.", d.Styles[idx])
}

func hash(s string) int {
	h := fnv.New32a()
        h.Write([]byte(s))
        return int(h.Sum32())
}
