package notionarticle

import (
	"context"

	"github.com/jomei/notionapi"
)

const (
	MaxPageSize = 100
)

func NewClient(c *notionapi.Client) *client {
	return &client{
		client: c,
	}
}

type client struct {
	client *notionapi.Client
}

func (c *client) Get(ctx context.Context, pageID string) (*Article, error) {
	page, err := c.getPage(ctx, pageID)
	if err != nil {
		return nil, err
	}

	blocks, err := c.getBlocksRecursive(ctx, pageID)
	if err != nil {
		return nil, err
	}

	return &Article{
		Page:   page,
		Blocks: blocks,
	}, nil
}

func (c *client) getPage(ctx context.Context, pageID string) (notionapi.Page, error) {
	page, err := c.client.Page.Get(ctx, notionapi.PageID(pageID))
	if err != nil {
		return notionapi.Page{}, err
	}
	return *page, nil
}

func (c *client) getBlocksRecursive(ctx context.Context, pageID string) (notionapi.Blocks, error) {
	var blocks notionapi.Blocks
	for cursor := ""; ; {
		r, err := c.client.Block.GetChildren(ctx, notionapi.BlockID(pageID), &notionapi.Pagination{
			StartCursor: notionapi.Cursor(cursor),
			PageSize:    MaxPageSize,
		})
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, r.Results...)

		if !r.HasMore {
			break
		}
		cursor = r.NextCursor
	}

	for _, b := range blocks {
		if b.GetHasChildren() {
			children, err := c.getBlocksRecursive(ctx, b.GetID().String())
			if err != nil {
				return nil, err
			}
			if cPtr := Children(b); cPtr != nil {
				(*cPtr) = children
			}
		}
	}

	return blocks, nil
}
