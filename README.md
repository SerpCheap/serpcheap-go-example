# serpcheap-go example

Minimal example using the official [serp.cheap Go SDK](https://github.com/SerpCheap/serpcheap-go).

## Quickstart

```bash
go mod download
SERPCHEAP_API_KEY=your-key go run .
```

The example runs a single search for `best running shoes` (US), prints each
organic result, and shows a billing/cache stats line.

## Page scraping

`scrape/main.go` opts the same search into page-content scraping by setting
`Scrape` with `ScrapeOptions{RenderJS: true, Screenshot: true, TopN: 3}`. It then
prints each result's scraped `Content` and `ScreenshotURL` (or the per-result
`ScrapeError` when a page can't be fetched).

```bash
SERPCHEAP_API_KEY=your-key go run ./scrape
```

This requires a `serpcheap-go` release with scrape support (the `ScrapeOptions`
type and the `Content` / `ScreenshotURL` / `ScrapeError` organic fields), which is
not yet published or tagged — pin that version once it ships.

Get an API key at [serp.cheap](https://serp.cheap).
