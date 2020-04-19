package question3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IPstats struct {
	ipstats [][]string
	logpath string
}

func NewIPstats(logpath string) *IPstats {
	return &IPstats{logpath: logpath}
}

func (i *IPstats) SetIpstats() error {
	ipstats := [][]string{}
	f, err := os.Open(i.logpath)
	if err != nil {
		return err
	}
	br := bufio.NewReader(f)
	for {
		line, _, eof := br.ReadLine()
		if eof == io.EOF {
			break
		}
		logline := strings.Fields(strings.TrimSpace(string(line)))
		if len(logline) != 4 {
			return errors.New("Log format is not correct!")
		}
		ipstats = append(ipstats, logline)
	}
	i.ipstats = ipstats
	return nil
}

func (i *IPstats) SetLogpath(logpath string) *IPstats {
	i.logpath = logpath
	return i
}

func (i *IPstats) CountIP() map[string]int {
	result := make(map[string]int)
	for m := 0; m < len(i.ipstats); m++ {
		result[i.ipstats[m][0]]++
	}
	return result
}

func (i *IPstats) CountRetcode() map[string]int {
	result := make(map[string]int)
	for m := 0; m < len(i.ipstats); m++ {
		result[i.ipstats[m][2]]++
	}
	return result
}

func (i *IPstats) UrlStatistics() map[string]map[string]int64 {
	result := make(map[string]map[string]int64)
	for m := 0; m < len(i.ipstats); m++ {
		flowbytes, err := strconv.ParseInt(i.ipstats[m][3], 0, 0)
		if err != nil {
			continue
		}
		if _, ok := result[i.ipstats[m][1]]; !ok {
			result[i.ipstats[m][1]] = make(map[string]int64)
			result[i.ipstats[m][1]][i.ipstats[m][0]] = flowbytes
		} else {
			result[i.ipstats[m][1]][i.ipstats[m][0]] += flowbytes
		}
	}
	return result
}

func IpAnalysisMain() {
	logpath := "question3/access.log"
	ipstats := NewIPstats(logpath)
	err := ipstats.SetIpstats()
	if err != nil {
		fmt.Println(err)
		return
	}

	a := ipstats.CountIP()
	b := ipstats.CountRetcode()
	c := ipstats.UrlStatistics()

	fmt.Printf("Counting the times of each IP in log %s\n", logpath)
	fmt.Printf("%-16s %-6s\n", "IP", "COUNTS")
	for key, value := range a {
		fmt.Printf("%-16s %-6d\n", key, value)
	}
	fmt.Print("\n\n")
	fmt.Printf("Counting the times of each ReturnCode in log %s\n", logpath)
	fmt.Printf("%-10s  %-6s\n", "ReturnCode", "COUNTS")
	for key, value := range b {
		fmt.Printf("%-10s  %-6d\n", key, value)
	}
	fmt.Print("\n\n")
	fmt.Printf("Counting the volume of each IP group by url in log %s\n", logpath)
	fmt.Printf("%-20s %-16s %-6s\n", "URL", "IP", "COUNTS")
	for key, value := range c {
		fmt.Printf("%-20s ", key)
		flag := 0
		for subkey, subvalue := range value {
			if flag == 0 {
				fmt.Printf("%-16s %-6d\n", subkey, subvalue)
			} else {
				fmt.Printf("%-20s %-16s %-6d\n", " ", subkey, subvalue)
			}
			flag++
		}
	}
}
