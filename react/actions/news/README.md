# News

News is a simple action that searches news articles using either Google
or [NewsAPI](https://newsapi.org) and summarizes them.

To use NewsAPI, you need to [register](https://newsapi.org/register) and set
the `CHLOE_REACT_NEWSAPI_TOKEN` environment. If the token is not set, the action will fallback to
Google.

## Configuration

| Environment Variable              | Default Value | Description                                  | Options                                  |
|-----------------------------------|---------------|----------------------------------------------|------------------------------------------|
| CHLOE_REACT_NEWSAPI_MAX_RESULTS   | 5             | Maximum number of NewsAPI results to analyze |                                          |
| CHLOE_REACT_NEWS_SOURCE           | google        | News source to use for news prompts          | google<br/>newsapi                       |
| CHLOE_REACT_NEWSAPI_TOKEN         |               | NewsAPI token                                |                                          |
| CHLOE_REACT_NEWSAPI_SORT_STRATEGY | relevancy     | NewsAPI sort strategy                        | publishedAt<br/>relevancy<br/>popularity |
