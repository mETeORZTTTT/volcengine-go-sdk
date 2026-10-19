package model

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestChatMediaFileIDJSON(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value interface{}
		want  string
	}{
		{"image file", ChatMessageImageURL{FileID: "file-image", Detail: ImageURLDetailHigh}, `{"file_id":"file-image","detail":"high"}`},
		{"video file", ChatMessageVideoURL{FileID: "file-video"}, `{"file_id":"file-video"}`},
		{"image url", ChatMessageImageURL{URL: "https://example.com/a.png"}, `{"url":"https://example.com/a.png"}`},
		{"video url", ChatMessageVideoURL{URL: "https://example.com/a.mp4"}, `{"url":"https://example.com/a.mp4"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(b) != tc.want {
				t.Fatalf("got %s, want %s", b, tc.want)
			}
		})
	}
	req := CreateChatCompletionRequest{Model: "fixture", Messages: []*ChatCompletionMessage{{Role: "user", Content: &ChatCompletionMessageContent{ListValue: []*ChatCompletionMessageContentPart{{Type: ChatCompletionMessageContentPartTypeVideoURL, VideoURL: &ChatMessageVideoURL{FileID: "file-video"}}}}}}}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"video_url":{"file_id":"file-video"}`) {
		t.Fatalf("file lost in request: %s", b)
	}
}
