# ReAct: Google

Google is a simple action that searches Google and, using the `CHLOE_REACT_GOOGLE_MAX_RESULTS` top
results, forwards then to
the [scrape](https://github.com/kamushadenes/chloe/blob/main/react/actions/scrape/) action for
further processing.

By default, the action will scrape Google search results for the given query. If
both `CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_ID` and `CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_API_KEY` are set,
the action will use the [Custom Search API](https://developers.google.com/custom-search/v1/overview)
instead, which is way more stable.

In order to use the Custom Search API, you need to create
a [Custom Search Engine](https://cse.google.com/cse/all), which you can set to search specific
websites or the entire web.

## Configuration

| Environment Variable                     | Default Value | Description                                 | Options          |
|------------------------------------------|---------------|---------------------------------------------|------------------|
| CHLOE_REACT_GOOGLE_MAX_RESULTS           | 4             | Maximum number of Google results to analyze | Between 1 and 10 |
| CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_ID      | -             | Custom search ID to use for Google searches |                  |
| CHLOE_REACT_GOOGLE_CUSTOM_SEARCH_API_KEY | -             | API key to use for Google searches          |                  |
