package question2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FindTOPmain() {
	fpath:="question2/ihaveadream.txt"
	counts, err := countLetter(fpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(findTOP(counts, 10))
}

func countLetter(filepath string) (map[string]int, error) {
	letters := make(map[string]int)
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		abyte, eof := br.ReadByte()
		if eof == io.EOF {
			break
		}
		if abyte >= 65 && abyte <= 90 || abyte >= 97 && abyte <= 122 {
			letters[string(abyte)]++
		}
	}
	return letters, err
}

//You can specify how many tops you wanna look over.
func findTOP(letters map[string]int, top int) string {
	sli := []string{}

	for key, _ := range letters {
		sli = append(sli, key)
	}

	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if letters[sli[j]] < letters[sli[j+1]] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}

	//In case of the specific "TOP" exceeding the amount of letters
	if top > len(sli) {
		top = len(sli)
	}

	result := make([]string, top+1)

	result[0] = fmt.Sprintf("%-6s | %-8s | %-8s", "RANK", "LETTER", "COUNTS")
	for i := 0; i < top; i++ {
		result[i+1] = fmt.Sprintf("%-6d | %-8s | %-8d", i+1, sli[i], letters[sli[i]])
	}
	return strings.Join(result, "\n")
}