package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}

	pMap := map[byte]byte{}
	sMap := map[byte]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}

func main() {
	s1 := "egg"
	t1 := "add"
	res1 := isIsomorphic(s1, t1)
	fmt.Println(res1)

	s2 := "foo"
	t2 := "bar"
	res2 := isIsomorphic(s2, t2)
	fmt.Println(res2)

	s3 := "paper"
	t3 := "title"
	res3 := isIsomorphic(s3, t3)
	fmt.Println(res3)

}
