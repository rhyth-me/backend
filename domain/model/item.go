package model

// Item - item
type Item struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailURL  string `json:"thumbnailUrl"`
	MusicTitle    string `json:"musicTitle"`
	Rating        int    `json:"rating"`
	Price         int    `json:"price"`
	DownloadCount int    `json:"downloadCount"`
}
