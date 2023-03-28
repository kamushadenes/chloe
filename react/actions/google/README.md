# ReAct: Google

Google is a simple action that searches Google and, using the `CHLOE_REACT_GOOGLE_MAX_RESULTS` top
results, forwards then to
the [scrape](https://github.com/kamushadenes/chloe/blob/main/react/actions/scrape/) action for
further processing.

## Configuration

| Environment Variable           | Default Value | Description                                 | Options |
|--------------------------------|---------------|---------------------------------------------|---------|
| CHLOE_REACT_GOOGLE_MAX_RESULTS | 4             | Maximum number of Google results to analyze |         |
