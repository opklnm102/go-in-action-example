package search

import (
	"encoding/json"
	"os"
)

const dataFile = "ch2/data/data.json"

// json형식을 사용하려면 디코딩을 통해 구조체의 슬라이스로 변환해야 한다
// 피드를 처리할 정보를 표현하는 구조체
type Feed struct {
	Name string `json:"site"`  // json 필드를 구조체의 변수에 매핑
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 피드 데이터 파일을 읽어 구조체로 변환
func RetrieveFeeds() ([]*Feed, error) {
	// 파일을 연다
	file, err := os.Open(dataFile)  // File 구조체 포인터, error
	if err != nil {
		return nil, err
	}

	/*
	defer
		함수가 리턴된 직후에 실행될 작업을 예약
		panic에 의해 종료되더라도 반드시 실행
		defer로 파일을 여는 코드 주변에 닫으면 가독성 향상
	 */
	// 함수가 리턴될 때 열어둔 파일을 닫는다
	defer file.Close()

	// 파일을 읽어 Feed 구조체의 포인터의 슬라이스로 변환
	var feeds []*Feed
	/*
	func (dec *Decoder) Decode(v interface{}) error
	interface{}
		Go에서 특별하게 취급하는 타입
		reflect 패키지를 이용한 reflection 지원이 가능한 타입
	 */
	err = json.NewDecoder(file).Decode(&feeds)

	// 호출 함수가 오류를 처리할 수 있으므로 오류 처리는 하지 않는다
	return feeds, err
}
