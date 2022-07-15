package tests

import (
	"fmt"
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestTodo(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		t.Fatalf("could not goto: %v", err)
	}
	entries, err := page.QuerySelectorAll(".athing")
	if err != nil {
		t.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("td.title > a")
		if err != nil {
			t.Fatalf("could not get title element: %v", err)
		}
		title, err := titleElement.TextContent()
		if err != nil {
			t.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		t.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		t.Fatalf("could not stop Playwright: %v", err)
	}
}
