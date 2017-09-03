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
	feeds, err := RetrieveFeeds()  // 여러개의 리턴값을 가질 수 있다
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 channel을 생성하여 화면에 표시할 검색 결과를 전달받는다
	/*
	channel
	고루틴 사이의 데이터 통신에 사용될 특정 타입의 값들을 위한 queue를 구현하고 있다
	안전한 통신을 위해 기본적으로 동기화 알고리즘 내장
	 */
	results := make(chan *Result)

	// 모든 피드를 처리하는 동안 기다릴 WaitGroup을 설정
	// WaitGroup은 카운팅 세마포어 - 고루틴의 실행이 종료될 때마다 개수를 하나씩 줄여나간다
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할 고루틴의 개수를 설정
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행
	for _, feed := range feeds {

		// 검색을 위해 검색기를 조회
		// 피드의 종류에 따라 각기 다른 matcher 구현체를 호출하기 위해 interface 사용
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행
		/*
		go의 모든 변수들은 값에 의해 전달
		포인터 변수 - 메모리상의 주소를 가리키는 값을 가진다
		 */
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)

			/*
			익명함수의 매개변수로 전달되지 않았어도 사용 가능
			closure 덕분에 함수 내에 매개변수로 전달되지 않은 변수들에 대해 접근 가능
			익명함수를 둘러싼 외부함수의 범위에 선언된 변수들에 직접 접근 가능
			 */
			waitGroup.Done()  // waitGroup의 카운팅 감소
		}(matcher, feed)
	}

	// 모든 작업이 완료되었는지를 모니터링할 고루틴을 실행
	go func() {
		// 모든 작업이 처리될 때까지 기다린다
		waitGroup.Wait()  // waitGroup의 카운터가 0이 될 때까지 실행 중단

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
