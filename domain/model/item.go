package model

// Item - item
type Item struct {
	ID         string         `firestore:"id" json:"id"`
	Snippet    ItemSnippet    `firestore:"snippet" json:"snippet,omitempty"`
	Statistics ItemStatistics `firestore:"statistics" json:"statistics,omitempty"`
	Author     SocialProfile  `firestore:"author" json:"author,omitempty"`
}

type ItemSnippet struct {
	Title        string   `firestore:"title" json:"title,omitempty"`
	ThumbnailURL string   `firestore:"thumbnailUrl" json:"thumbnailUrl,omitempty"`
	MusicTitle   string   `firestore:"musicTitle" json:"musicTitle,omitempty"`
	Price        int      `firestore:"price" json:"price,omitempty"`
	Tags         []string `firestore:"tags" json:"tags,omitempty"`
}

type ItemStatistics struct {
	Rating        float64 `firestore:"rating" json:"rating,omitempty"`
	DownloadCount int     `firestore:"downloadCount" json:"downloadCount,omitempty"`
}
