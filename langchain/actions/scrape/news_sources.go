package scrape

import "strings"

var newsSources = map[string]NewsSource{
	"G1": {
		Domains:           []string{"g1.globo.com"},
		AuthorSelector:    "p.content-publication-data__from",
		TitleSelector:     "h1.content-head__title",
		SummarySelector:   "h2.content-head__subtitle",
		ContentSelector:   "div.mc-article-body",
		ParagraphSelector: "p.content-text__container",
		CreatedAtSelector: "time[itemprop=datePublished]",
		UpdatedAtSelector: "time[itemprop=dateModified]",
		TimeFormat:        "2006-01-02T15:04:05.000Z",
	},
	"Terra": {
		Domains:           []string{"terra.com.br"},
		AuthorSelector:    "p.content-publication-data__from",
		TitleSelector:     "div.article__header__headline",
		SummarySelector:   "div.article__header__subtitle",
		ContentSelector:   "div.article__content--body",
		ParagraphSelector: "p.text",
		CreatedAtSelector: "meta[itemprop=datePublished]",
		UpdatedAtSelector: "meta[itemprop=dateModified]",
		TimeFormat:        "2006-01-02T15:04:05-07:00",
	},
	"CNN Brasil": {
		Domains:           []string{"cnnbrasil.com.br"},
		AuthorSelector:    "p.author__name",
		TitleSelector:     "h1.post__title",
		SummarySelector:   "p.post__excerpt",
		ContentSelector:   "div.post__content",
		ParagraphSelector: "div.__post_content p",
	},
	"Estadao": {
		Domains:           []string{"estadao.com.br"},
		AuthorSelector:    "span.authors-names",
		TitleSelector:     "h1.cover-titulo",
		SummarySelector:   "h2",
		ContentSelector:   "div.news-body",
		ParagraphSelector: "div.news-body p",
	},
	"InfoMoney": {
		Domains:           []string{"infomoney.com.br"},
		AuthorSelector:    "div.single__author-info span.typography__body--5",
		TitleSelector:     "div.single__title h1",
		SummarySelector:   "div.single__excerpt",
		ContentSelector:   "div.single__content",
		ParagraphSelector: "div.single__content p",
		CreatedAtSelector: "time.published",
	},
	"New York Times": {
		Domains:           []string{"nytimes.com"},
		AuthorSelector:    "span.byline-author",
		TitleSelector:     "h1",
		SummarySelector:   "p#article-summary",
		ContentSelector:   "section[name=articleBody]",
		ParagraphSelector: "section[name=articleBody] p",
		CreatedAtSelector: "time",
		TimeFormat:        "2006-01-02T15:04:05-07:00",
	},
	"CNN": {
		Domains:           []string{"cnn.com"},
		AuthorSelector:    "div.byline__names",
		TitleSelector:     "h1#maincontent",
		ContentSelector:   "div.article__content",
		ParagraphSelector: "div.article__content p",
	},
	"The Wall Street Journal": {
		Domains:           []string{"wsj.com"},
		AuthorSelector:    "div.article-byline",
		TitleSelector:     "h1",
		SummarySelector:   "h2",
		ContentSelector:   "section",
		ParagraphSelector: "section p",
		CreatedAtSelector: "time",
	},
	"The Washington Post": {
		Domains:           []string{"washingtonpost.com"},
		AuthorSelector:    "div[data-qa=author-byline]",
		TitleSelector:     "h1#main-content",
		SummarySelector:   "h2",
		ContentSelector:   "div.grid-body",
		ParagraphSelector: "div.article-body",
	},
	"BBC": {
		Domains:           []string{"bbc.com"},
		AuthorSelector:    "div.[data-componenet=byline-block]",
		TitleSelector:     "h1#main-heading",
		SummarySelector:   "p.story-body__introduction",
		ContentSelector:   "article",
		ParagraphSelector: "article div[data-component=text-block] p",
		CreatedAtSelector: "time",
		TimeFormat:        "2006-01-02T15:04:05.000Z",
	},
	"The Guardian": {
		Domains:           []string{"theguardian.com"},
		TitleSelector:     "div[data-gu-name=headline]",
		ContentSelector:   "div#maincontent",
		ParagraphSelector: "div#maincontent p",
	},
}

type NewsSource struct {
	Domains           []string
	AuthorSelector    string
	TitleSelector     string
	SummarySelector   string
	ContentSelector   string
	ParagraphSelector string
	CreatedAtSelector string
	UpdatedAtSelector string
	TimeFormat        string
}

func GetNewsSource(domain string) *NewsSource {
	for _, newsSource := range newsSources {
		for _, newsSourceDomain := range newsSource.Domains {
			if strings.HasSuffix(domain, newsSourceDomain) {
				return &newsSource
			}
		}
	}

	return nil
}
