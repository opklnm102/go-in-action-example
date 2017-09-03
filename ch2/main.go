package main

import (
	"log"
	"os"

	_ "go-in-action/ch2/matchers"  // 직접사용하지 않더라도 import를 유지, rss.go의 init() 때문에
	"go-in-action/ch2/search"
)

// main()보다 먼저 호출
func init() {
	// stdout으로 로그를 출력하도록 번경
	log.SetOutput(os.Stdout)
}

// 프로그램의 진입점
// main package에 속해있어야 한다
func main() {
	// 지정된 검색어로 검색을 수행
	search.Run("president")
}
