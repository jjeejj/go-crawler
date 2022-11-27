package test

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestChromeDp(t *testing.T) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://pkg.go.dev/time"),
		chromedp.WaitVisible(`details#example-After`),
		chromedp.Click("#example-After", chromedp.NodeVisible),
		chromedp.Value("#example-After textarea", &example),
	)
	if err != nil {
		t.Log(err)
	}
	t.Log(example)
}
