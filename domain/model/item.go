package model

// Item - item
type Item struct {
	ID         string         `firestore:"id" json:"id"`
	Snippet    ItemSnippet    `firestore:"snippet" json:"snippet"`
	Statistics ItemStatistics `firestore:"statistics" json:"statistics"`
	Author     User           `firestore:"author" json:"author"`
}

type ItemSnippet struct {
	Title        string   `firestore:"title" json:"title"`
	ThumbnailURL string   `firestore:"thumbnailUrl" json:"thumbnailUrl"`
	MusicTitle   string   `firestore:"musicTitle" json:"musicTitle"`
	Price        int      `firestore:"price" json:"price"`
	Tags         []string `firestore:"tags" json:"tags"`
}

type ItemStatistics struct {
	Rating        float64 `firestore:"rating" json:"rating"`
	DownloadCount int     `firestore:"downloadCount" json:"downloadCount"`
}
