package search // search 폴더의 모든 .go파일은 search package에 속한다

// GOROOT, GOPATH에 정의된 위치를 기준으로 패키지 탐색
import (
	"log"  // logging
	"sync"  //  gorutine 사이의 동기화
)

/*
변수명
	소문자 시작 - package 외부로 노출 X, return으로 받은 것들에 대해 간접 접근 가능
	대문자 시작 - package 외부로 노출, 직접 접근 가능
map
	make()로 런타임에 새성을 요청해야 하는 참조타입
	map의 zero value는 nil이기 때문

모든 변수는 zero value로 초기화
숫자 - 0
문자열 - 빈문자열
boolean - false
포인터 - nil
참조 타입 - 각 기반 타입의 zero value지만, 변수 자체의 zero value는 nil

 */
// 검색을 처리할 검색기의 매핑 정보를 저장할 map
var matchers = make(map[string]Matcher)  // package 수준의 변수

// 검색 로직을 수행할 함수
func Run(searchTerm string) {

	// 검색할 피드의 목록 조회
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 channel을 생성하여 화면에 표시할 검색 결과를 전달받는다
	results := make(chan *Result)

	// 모든 피드를 처리하는 동안 기다릴 Wait Group을 설정
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할 고루틴의 개수를 설정
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행
	for _, feed := range feeds {

		// 검색을 위해 검색기를 조회
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴을 실행
	go func() {
		// 모든 작업이 처리될 때까지 기다린다
		waitGroup.Wait()

		// Display()에게 프로그램을 종료할 수 있음을 알리기 위해 체널을 닫는다
		close(results)
	}()

	// 검색 결과를 화면에 표시
	// 마지막 결과를 표시한 뒤 리턴
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
