package imgur

// Comment is an imgur comment
type Comment struct {
	// The ID for the comment
	ID int `json:"id"`
	//The ID of the image that the comment is for
	ImageID string `json:"image_id"`
	// The comment itself.
	Comment string `json:"comment"`
	// Username of the author of the comment
	Author string `json:"author"`
	// The account ID for the author
	AuthorID int `json:"author_id"`
	// If this comment was done to an album
	OnAlbum bool `json:"on_album"`
	//	Number of upvotes for the comment
	Ups int `json:"ups"`
	// The ID of the album cover image, this is what should be displayed for album comments
	AlbumCover string `json:"album_cover"`
	// The number of downvotes for the comment
	Downs int `json:"downs"`
	// the number of upvotes - downvotes
	Points float32 `json:"points"`
	// Timestamp of creation, epoch time
	Datetime int `json:"datetime"`
	// If this is a reply, this will be the value of the comment_id for the caption this a reply for.
	ParentID int `json:"parent_id"`
	// Marked true if this caption has been deleted
	Deleted bool `json:"deleted"`
	// The current user's vote on the comment. null if not signed in or if the user hasn't voted on it.
	Vote string `json:"vote"`
	// All of the replies for this comment. If there are no replies to the comment then this is an empty set.
	Children []Comment `json:"children"`
}
