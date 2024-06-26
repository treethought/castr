package api

import (
	"context"
	"fmt"
)

type FeedRequest struct {
	FeedType   string
	FID        uint64
	FilterType string
	ParentURL  string
	FIDs       []uint64
	Cursor     string
	Limit      uint64
	ViewerFID  uint64
}

func (r *FeedRequest) opts() []RequestOption {
	var opts []RequestOption
	if r.FeedType != "" {
		opts = append(opts, WithQuery("feed_type", r.FeedType))
	}
	if r.FilterType != "" {
		opts = append(opts, WithQuery("filter_type", r.FilterType))
	}
	if r.ParentURL != "" {
		opts = append(opts, WithQuery("parent_url", r.ParentURL))
	}
	if r.FIDs != nil {
		for _, fid := range r.FIDs {
			opts = append(opts, WithQuery("fids", fmt.Sprintf("%d", fid)))
		}
	}
	if r.FeedType == "following" {
		opts = append(opts, WithQuery("fid", fmt.Sprintf("%d", r.FID)))
	}
	if r.Cursor != "" {
		opts = append(opts, WithQuery("cursor", r.Cursor))
	}
	if r.Limit != 0 {
		opts = append(opts, WithQuery("limit", fmt.Sprintf("%d", r.Limit)))
	}
	if r.ViewerFID != 0 {
		opts = append(opts, WithQuery("viewer_fid", fmt.Sprintf("%d", r.ViewerFID)))
	}

	return opts
}

type FeedResponse struct {
	Casts []*Cast
}

func (c *Client) GetFeed(r *FeedRequest) (*FeedResponse, error) {
	path := "/feed"
	opts := r.opts()
	var resp FeedResponse
	if err := c.doRequestInto(context.TODO(), path, &resp, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}
