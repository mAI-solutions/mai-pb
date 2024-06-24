package feed

import (
	"context"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedIter <-chan *gofeed.Feed
func RetFeeds(urls []string, timeseconds int64, ogctx context.Context) (FeedIter, error) {
    newFeedParser := gofeed.NewParser()
    if timeseconds != 0 {
        return retFeedsWithTimeout(newFeedParser, urls, timeseconds, ogctx)
    }

    sentinel := sync.WaitGroup{}
    feedChan := make(chan *gofeed.Feed)
    for index := range urls {
        sentinel.Add(1)
        go func() {
            defer sentinel.Done()
            newFeed, feedErr := newFeedParser.ParseURL(urls[index])
            if feedErr != nil { return }
            feedChan <- newFeed
        }()
    }

    go func(){
        sentinel.Wait()
        close(feedChan)
    }()

    return feedChan, nil
}

func retFeedsWithTimeout(parser *gofeed.Parser, urls []string, timeseconds int64, ogctx context.Context) (FeedIter, error) {
    timeoutContext, cancelTimeout := context.WithTimeout(
        ogctx,
        time.Duration(timeseconds*time.Millisecond.Nanoseconds()),
    )

    sentinel := sync.WaitGroup{}
    feedChan := make(chan *gofeed.Feed)
    for index := range urls {
        sentinel.Add(1)
        go func() {
            defer sentinel.Done()
            newFeed, feedErr := parser.ParseURLWithContext(urls[index], timeoutContext)
            if feedErr != nil { return }
            feedChan <- newFeed
        }()
    }

    go func(){
        sentinel.Wait()
        cancelTimeout()
        close(feedChan)
    }()

    return feedChan, nil
}





