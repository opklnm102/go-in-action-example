package search

/*
빈 구조체
	type의 값이 생성될 때 0바이트의 메모리 할당
	타입은 필요하지만, 타입의 상태를 관리할 필요가 없는 경우 유용
 */
// 기본 검색기를 구현한 defaultMatcher 타입
type defaultMatcher struct {
}

// 기본 검색기를 등록
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 기본 검색기의 동작을 구현
// defaultMatcher에 대한 value receiver를 선언
// 메소드 선언부에 value receiver 사용하면 지정된 receiver type에만 연결된다
// 메소드 호출을 위해 receiver type의 값이나, 포인터 중 어떤 것이든 컴파일러는 메소드 호출에 필요한 참조 역참조를 알아서 수행
/*
// defaultMatcher 타입의 value receiver를 이용해 메소드 선언
func (m defaultMatcher) Search(feed *Feed, searchTerm string)

// defaultMatcher 타입의 포인터를 선언
dm := new(defaultMatcher)

// 메소드를 호출하면 컴파일러가 알아서 dm 포인터에 대한 dereference(역참조)를 수행
dm.Search(feed, "test")

// defaultMatcher 타입의 pointer receiver를 이용해 메소드 선언
func (m *defaultMatcher) Search(feed *Feed, searchTerm string)

// defaultMatcher 타입의 값 선언
var dm defaultMatcher

// 메소드를 호출하면 컴파일러가 알아서 dm 값에 대한 reference(참조)를 수행
dm.Search(feed, "test")


대부분의 메소드는 실행 과정에서 값의 상태를 조작해야 하는 경우가 많으므로 pointer receiver를 이용해 선언하는 것을 권장
defaultMatcher 타입의 경우 value receiver를 이용한 이유는 defaultMatcher 타입의 값을 생성할 때 메모리를 소비하지 않아도 되기 때문
조작할 상태가 없는데 pointer를 이용해서 메소드를 선언할 이유가 없다

value, pointer로 메소드를 호출할 때와 달리 인터페이스 타입 값을 통해 메소드를 호출하면 다른 규칙 적용
pointer receiver를 이용해 선언된 메소드는 인터페이스 타입에 대한 pointer를 통해서만 호출 가능
value receiver를 이용해 선언된 메소드는 값, 포인터 변수 모두를 통해 호출 가능

// defaultMatcher 타입의 pointer receiver를 이용해 메소드 선언
func (m *defaultMatcher) Search(feed *Feed, searchTerm string)

// 인터페이스 타입 값을 통해 메소드 호출
var dm defaultMatcher
var matcher Matcher = dm  // 인터페이스 타입의 값을 대입
matcher.Search(feed, "test")  // 값에 의한 인터페이스 메소드 호출

> go build
cannot use dm (type defaultMatcher) as type Matcher in assignment

// defaultMatcher 타입의 value receiver를 이용해 메소드 선언
func (m defaultMatcher) Search(feed *Feed, searchTerm string)

// 인터페이스 타입 값을 통해 메소드 호출
var dm defaultMatcher
var matcher Matcher = &dm  // 인터페이스 타입의 포인터를 대입
matcher.Search(feed, "test")  // 포인터에 의한 인터페이스 메소드 호출

> go build
Build Successful
 */

/*
defaultMatcher 타입의 값과 포인터는 인터페이스 구현 조건 만족
defaultMatcher 타입의 값과 포인터는 Matcher 타입의 값에 대입하거나 Matcher 타입의 값을 매개변수로 정의하는 함수에 전달할 수 있다
 */
func (matcher defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
