package main

import (
	"testing"
)

// 테스트 코드는 파일명이 테스트할 파일이름_test.go여야 한다
// testing 패키지를 import해야한다
// 테스트 코드는 func TestXxxx(t *testing.T)형태여야 한다
// 함수명은 반드시 Test로 시작. 뒤에는 첫글자 대문자로
// 터미널에서 go test실행
func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81, but the result is %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be 9, but the result is %d", rst)
	}
}
