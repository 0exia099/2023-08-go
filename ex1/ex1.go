package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineIdx int
	line    string
}

type FindInfo struct {
	fileName string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("올바른 입력이 필요합니다.")
		return
	}
	word := os.Args[1]
	path := os.Args[2]
	findInfos := []FindInfo{}

	findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.fileName)
		fmt.Println("==============================================================================")
		for _, lineInfo := range findInfo.lines {
			leng := len(lineInfo.line)
			switch { // break필요 없음. 밑에것도 실행시키고 싶으면 fallthrough사용
			case leng < 48:
				fmt.Printf("\t%3d : %s\n", lineInfo.lineIdx, lineInfo.line)
			case leng < 96:
				fmt.Printf("\t%3d : %s\n", lineInfo.lineIdx, lineInfo.line[0:48])
				fmt.Printf("\t     %s\n", lineInfo.line[48:])
			case leng < 144:
				fmt.Printf("\t%3d : %s\n", lineInfo.lineIdx, lineInfo.line[0:48])
				fmt.Printf("\t     %s\n", lineInfo.line[48:96])
				fmt.Printf("\t     %s\n", lineInfo.line[96:])
			default:
				fmt.Printf("\t%3d : %s\n", lineInfo.lineIdx, lineInfo.line[0:48])
				fmt.Printf("\t     %s\n", lineInfo.line[48:96])
				fmt.Printf("\t     %s\n", lineInfo.line[96:144])
				fmt.Printf("\t      .....\n")
			}
		}
		fmt.Println("==============================================================================")
		fmt.Println()
	}
}
func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, name := range filelist {
		go FindWordInFile(word, name, ch)
		//findInfos = append(findInfos, FindWordInFiles(word, name))
	}
	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos
}
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}
func FindWordInFiles(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close()
	lowerWord := strings.ToLower(word)
	lineIdx := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, lowerWord) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineIdx, line})
		}
		lineIdx++
	}
	return findInfo
}
func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		ch <- findInfo
		//return findInfo
	}
	defer file.Close()
	lowerWord := strings.ToLower(word)
	lineIdx := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, lowerWord) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineIdx, line})
		}
		lineIdx++
	}
	ch <- findInfo
	//return findInfo
}
