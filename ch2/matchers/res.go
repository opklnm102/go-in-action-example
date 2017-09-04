/*
RSS 피드를 처리하기 위한 검색기
향후 matchers package에 json, csv 파일등을 읽을 수 있는 검색기 추가 가능
 */

package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"go-in-action/ch2/search"
)

// RSS 문서 내의 item 태그에 정의된 필드에 대응하는 필드 선언
type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		Pubdate     string `xml:"pubDate"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string  `xml:"link"`
		GUID        string  `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}

	// RSS 문서 내의 image 태그에 정의된 필드에 대응하는 필드 선언
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string `xml:"url"`
		Title   string `xml:"title"`
		Link    string `xml:"link"`
	}

	// RSS 문서 내의 channel 태그에 정의된 필드에 대응하는 필드 선언
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string `xml:"title"`
		Description    string `xml:"description"`
		Link           string `xml:"link"`
		PubDate        string `xml:"pubDate"`
		LastBuildDate  string `xml:"lastBuildDate"`
		TTL            string `xml:"ttl"`
		Language       string `xml:"language"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster      string `xml:"webMaster"`
		Image          image `xml:"image"`
		Item           []item `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)

// Matcher 인터페이스를 구현하는 rssMatcher 타입 선언
// 관리해야 할 상태가 없기 때문에 빈 구조체 사용
type rssMatcher struct{}

// 검색기를 등록
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// 지정된 문서에서 검색어를 검색
func (matcher rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	// nil로 초기화된 result 타입의 슬라이스 선언
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// 검색할 데이터 조회
	document, err := matcher.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {

		// 제목에서 검색어를 검색
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장
		if matched {
			results = append(results, &search.Result{ // append(값을 덧붙일 슬라이스, 추가하고자 하는 값)
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		// 상세 내용에서 검색어를 검색
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}

	}

	return results, nil
}

// HTTP GET으로 RSS 피드를 요청한 후 결과를 디코딩
// 메소드 이름 시작이 소문자 -> 비공개 메소드
func (matcher rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// RSS 문서 조회
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 함수 리턴시 close response stream
	defer resp.Body.Close()

	// 올바른 응답 확인
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// RSS 피드 문서를 구조체 타입으로 디코드
	// 호출함수가 에러를 판단할 것이기 때문에 이 함수에서 에러 처리 X
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
