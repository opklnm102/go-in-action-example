# architecture

![architecture](https://github.com/opklnm102/go-in-action-example/blob/master/ch2/architecture.png)

```sh
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
```
