# ReAct: Scrape

This action allows you to scrape a website for information. It is a wrapper around
the [go-colly](http://go-colly.org/) library with some additional features and site-specific
rules.

Pages are analyzed, some metadata is extract and a cleaned version of the page, with all HTML
stripped, is returned. In case the page is known, it avoids extracting random text and gets only the
main content of the page.

## Configuration

There are no configuration options for this action.