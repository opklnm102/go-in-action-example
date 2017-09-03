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

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string `xml:"url"`
		Title   string `xml:"title"`
		Link    string `xml:"link"`
	}

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

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (matcher rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	document, err := matcher.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}

	}

	return results, nil
}

func (matcher rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
