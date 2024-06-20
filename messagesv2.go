package openai

import (
	"context"
	"fmt"

	"net/http"
)

type ImageFileV2 struct {
	FileID string `json:"file_id"`
	Detail string `json:"detail,omitempty"`
}

type ImageURLV2 struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"`
}

type MessageContentV2 struct {
	Type      string       `json:"type"`
	Text      string       `json:"text,omitempty"`
	ImageFile *ImageFileV2 `json:"image_file,omitempty"`
	ImageURL  *ImageURLV2  `json:"image_url,omitempty"`
}

type MessageRequestV2 struct {
	Role    string             `json:"role"`
	Content []MessageContentV2 `json:"content"`
}

// CreateMessageV2 creates a new message.
func (c *Client) CreateMessageV2(ctx context.Context, threadID string, request MessageRequestV2) (msg Message, err error) {
	urlSuffix := fmt.Sprintf("/threads/%s/%s", threadID, messagesSuffix)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request),
		withBetaAssistantVersion(c.config.AssistantVersion))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &msg)
	return
}
