package search

import (
	"log"
)

// 검색 결과를 저장할 Result 구조체
type Result struct {
	Field string
	Content string
}

/*
go convention
하나의 메소드만 선언하고 있다면 이름은 er로 끝나야 한다
여러개의 메소드라면 관련된 이름을 사용

인터페이스를 구현하는 사용자정의 타입의 경우, 모든 메소드를 구현해야 함
 */
// Matcher 인터페이스는 새로운 검색 타입을 구현할 때 필요한 동작을 정의
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// 고루틴으로써 호출됨
// 개별 피드 타입에 대한 검색을 동시에 수행
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan <- *Result) {

	// 지정된 검색기를 이용해 검색을 수행
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 검색 결과를 channel에 기록
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
