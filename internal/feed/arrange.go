package feed

import (
	"context"
	"slices"

	"github.com/mmcdole/gofeed"
)

type Post struct {
    Source string
    *gofeed.Item
}

func GetPost(urls []string, timeseconds int64, ogctx context.Context) ([]*Post, error) {
    iter, iterErr := RetFeeds(urls, timeseconds, ogctx)
    if iterErr != nil {
        return nil, iterErr
    }

    posts := make([]*Post, 0)
    for post := range iter {
        for index := range post.Items {
            posts = append(posts, &Post{
                Source: post.Title,
                Item: post.Items[index],
            })
        }
    }

    slices.SortFunc(posts, func(a *Post, b *Post) int {
        if a.PublishedParsed == nil {
            return 1
        }

        if b.PublishedParsed == nil {
            return -1
        }

        return -1 * a.PublishedParsed.Compare(*b.PublishedParsed)
    })

    return posts, nil
}


