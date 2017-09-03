package main

import (
	"log"
	"os"

	_ "go-in-action/ch2/matchers"  // 직접사용하지 않더라도 import를 유지, rss.go의 init() 때문에
	"go-in-action/ch2/search"
)

/*
├── data
│   └── data.json  -- 데이터 피드를 가지고 있는 파일
├── main.go        -- 프로그램의 진입점
├── matchers
│   └── res.go     -- RSS 피드 검색기를 구현한 코드
└── search
    ├── default.go -- 데이터 검색을 위한 기본적인 검색기 코드
    ├── feed.go    -- JSON 데이터 파일을 읽기 위한 코드
    ├── match.go   -- 서로 다른 종류의 검색기를 지원하기 위한 인터페이스
    └── search.go  -- 검색을 수행하는 주요 로직이 구현 파일
 */

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
