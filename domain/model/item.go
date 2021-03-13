package model

// Item - item collection
type Item struct {
	ID         string         `firestore:"id" json:"id"`
	Snippet    ItemSnippet    `firestore:"snippet" json:"snippet,omitempty"`
	Statistics ItemStatistics `firestore:"statistics" json:"statistics,omitempty"`
	Author     ItemAuthor     `firestore:"author" json:"author"`
}

// ItemAuthor - item author
type ItemAuthor struct {
	GoogleID string        `firestore:"googleId" json:"-"`
	Profile  SocialProfile `firestore:"-" json:"profile"`
}

// ItemSnippet - item datails
type ItemSnippet struct {
	Thumbnails Thumbnails `firestore:"thumbnails" json:"thumbnails,omitempty"`
	Title      string     `firestore:"title" json:"title,omitempty" validate:"required,max=40"`
	MusicTitle string     `firestore:"musicTitle" json:"musicTitle,omitempty" validate:"required,max=40"`
	Summary    string     `firestore:"summary" json:"summary,omitempty" validate:"max=200"`
	Price      int        `firestore:"price" json:"price,omitempty" validate:"min=0,max=10000000"`
	Tags       []string   `firestore:"tags" json:"tags,omitempty" validate:"unique"`
}

// Thumbnails - item thumbnails
type Thumbnails struct {
	YouTube string  `firestore:"youtube" json:"youtube"`
	Images  []Image `firestore:"images" json:"images"`
}

// Image - thumbnail image
type Image struct {
	Path string `firestore:"path" json:"path"`
}

// ItemStatistics - item statistics
type ItemStatistics struct {
	Rating        float64 `firestore:"rating" json:"rating,omitempty"`
	DownloadCount int     `firestore:"downloadCount" json:"downloadCount,omitempty"`
}
