package spider

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/wudizhangzhi/HupuApp"
)

type Region string

const (
	NBA  Region = "nba"
	CBA  Region = "cba"
	Vote Region = "湿乎乎" // 湿乎乎
)

func (r Region) String() string {
	switch r {
	case NBA:
		return "all-nba"
	case CBA:
		return "all-cba"
	case Vote:
		return "vote"
	default:
		return "all-nba"
	}
}

// 对应帖子列表的selector
func (r Region) GetBBSListSelector() string {
	switch r {
	case NBA:
		return "div.text-list-model > div > div > div > div.t-info"
	case CBA:
		return "div.text-list-model > div > div > div > div.t-info"
	case Vote:
		return "div.bbs-sl-web-post > ul > li > div > div.post-title"
	default:
		return "div.bbs-sl-web-post > ul > li > div > div.post-title"
	}
}

type BBS struct {
	Region   Region `comment:"领域"`
	Uid      string `comment:"帖子id"`
	Title    string `comment:"标题"`
	Href     string `comment:"链接"`
	Label    string `comment:"标签"`
	ReplyCnt int    `comment:"回复"`
	LightCnt int    `comment:"亮了"`
	Content  string `comment:"内容"`
}

type User struct {
	Uid      string
	Nickname string
}

type Comment struct {
	Uid       string
	Content   string `comment:"内容"`
	Location  string
	ReplyTime time.Time
	LightCnt  int
	Nickname  string
}

func GetBBSList(region Region, page int) ([]BBS, error) {
	url := "https://bbs.hupu.com/" + region.String()
	if page > 1 {
		url += fmt.Sprintf("-%d", page)
	}

	bbsList := make([]BBS, 0)
	resp, err := SpiderClient.R().
		Get(url)
	if err != nil {
		return nil, err
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return bbsList, nil
	}
	// Find the bbs items
	doc.Find(region.GetBBSListSelector()).
		Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			selection := s.Find("div > a")
			title := selection.Text()
			href, _ := selection.Attr("href")
			lightCntS := s.Find("div.t-info > span.t-lights").Text()
			replyCntS := s.Find("div.t-info > span.t-replies").Text()
			label := s.Find("div.t-label > a").Text()
			fmt.Printf("No.%d: 标题:%s 亮:%s 回复:%s\n", i*page+1, title, lightCntS, replyCntS)

			uid := regexp.MustCompile(`\d+`).FindString(href)
			lightCnt, _ := strconv.Atoi(lightCntS)
			replyCnt, _ := strconv.Atoi(replyCntS)
			bbsList = append(bbsList, BBS{
				Uid:      uid,
				Title:    title,
				LightCnt: lightCnt,
				ReplyCnt: replyCnt,
				Href:     href,
				Label:    label,
			})
		})
	return bbsList, nil
}

func (bbs *BBS) GetDetail() (string, error) {
	url := "https://bbs.hupu.com/" + bbs.Uid + ".html"
	resp, err := SpiderClient.R().
		Get(url)
	if err != nil {
		return bbs.Content, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return bbs.Content, err
	}
	content := ""
	doc.Find("div.thread-content-detail > p").
		Each(func(i int, s *goquery.Selection) {
			content += s.Text() + "\n"
		})
	return bbs.Content, nil
}

func (bbs *BBS) GetComments(page int) ([]Comment, error) {
	url := "https://bbs.hupu.com/" + bbs.Uid + fmt.Sprintf("-%d", page) + ".html"
	comments := make([]Comment, 0)
	resp, err := SpiderClient.R().
		Get(url)
	if err != nil {
		return comments, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return comments, nil
	}
	doc.Find("div.post-reply-list-wrapper").Each(func(i int, s *goquery.Selection) {
		content := s.Find("div.thread-content-detail > p").Text()
		nickname := s.Find("div.user-base-info > a").Text()
		replyTimeS := s.Find("div.user-base-info > span.post-reply-list-user-info-top-time").Text()
		replyTime, _ := HupuApp.GetTimeFromString(replyTimeS)
		location := s.Find("div.user-base-info > span.post-reply-list-user-info-user-location").Text()
		location = strings.Trim(location, "发布于")
		lightCntS := s.Find("div.post-reply-list-operate > div.light > span").Text()
		lightCntS = regexp.MustCompile(`\d+`).FindString(lightCntS)
		lightCnt, _ := strconv.Atoi(lightCntS)
		// fmt.Printf("评论: %s\n", content)
		comments = append(comments, Comment{
			Content:   content,
			Location:  location,
			ReplyTime: replyTime,
			Nickname:  nickname,
			LightCnt:  lightCnt,
		})
	})
	for _, comment := range comments {
		fmt.Printf("评论: %+v\n", comment)
	}
	return comments, nil
}
