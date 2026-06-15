// Page-scraping example for the serp.cheap Go SDK.
//
// Requires a serpcheap-go release with scrape support (the ScrapeOptions type and
// the Content/ScreenshotURL/ScrapeError organic fields). That version is not yet
// published/tagged — pin it once it ships.
package main

import (
	"context"
	"fmt"
	"os"

	serpcheap "github.com/serpcheap/serpcheap-go"
)

func main() {
	key := os.Getenv("SERPCHEAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Set SERPCHEAP_API_KEY to run this example, e.g.:")
		fmt.Fprintln(os.Stderr, "  SERPCHEAP_API_KEY=your-key go run ./scrape")
		os.Exit(1)
	}

	client, err := serpcheap.New(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create client:", err)
		os.Exit(1)
	}

	res, err := client.Search(context.Background(), serpcheap.SearchParams{
		Q:  "best running shoes",
		GL: "us",
		Scrape: &serpcheap.ScrapeOptions{
			RenderJS:   true,
			Screenshot: true,
			TopN:       3,
		},
	})
	if err != nil {
		if apiErr, ok := serpcheap.AsError(err); ok {
			fmt.Fprintf(os.Stderr, "[%s] %s\n", apiErr.Code, apiErr.Message)
		} else {
			fmt.Fprintln(os.Stderr, "search failed:", err)
		}
		os.Exit(1)
	}

	for _, r := range res.Organic {
		fmt.Printf("%d. %s — %s\n", r.Position, r.Title, r.Link)
		if r.ScrapeError != "" {
			fmt.Printf("   scrape failed: %s\n", r.ScrapeError)
			continue
		}
		if r.Content != "" {
			fmt.Printf("   content: %s\n", r.Content)
		}
		if r.ScreenshotURL != "" {
			fmt.Printf("   screenshot: %s\n", r.ScreenshotURL)
		}
	}

	if res.Stats != nil {
		fmt.Printf("\nbalance=%d cost=%d cached=%t\n", res.Stats.Balance, res.Stats.Cost, res.Stats.Cached)
	}
}
