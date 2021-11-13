package full_outer_join

import (
	"io/ioutil"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	input1, _ := ioutil.ReadFile(f1Path)
	input2, _ := ioutil.ReadFile(f2Path)

	lines1 := strings.Split(string(input1), "\n")
	lines2 := strings.Split(string(input2), "\n")

	freqLines := make(map[string]int, len(lines1)+len(lines2))
	for i := 0; i < len(lines1)+len(lines2); i++ {
		if i < len(lines1) {
			freqLines[lines1[i]]++
		} else {
			freqLines[lines2[i-len(lines1)]]++
		}
	}

	uniqueLines := make([]string, 0, len(freqLines))
	for k, v := range freqLines {
		if v < 2 {
			uniqueLines = append(uniqueLines, k)
		}
	}

	sort.Strings(uniqueLines)
	output := strings.Join(uniqueLines, "\n")
	ioutil.WriteFile(resultPath, []byte(output), 0644)
}
