package imgur

// Comment is an imgur comment
type Comment struct {
	ID         int       `json:"id"`          // The ID for the comment
	ImageID    string    `json:"image_id"`    //The ID of the image that the comment is for
	Comment    string    `json:"comment"`     // The comment itself.
	Author     string    `json:"author"`      // Username of the author of the comment
	AuthorID   int       `json:"author_id"`   // The account ID for the author
	OnAlbum    bool      `json:"on_album"`    // If this comment was done to an album
	AlbumCover string    `json:"album_cover"` // The ID of the album cover image, this is what should be displayed for album comments
	Ups        int       `json:"ups"`         //	Number of upvotes for the comment
	Downs      int       `json:"downs"`       // The number of downvotes for the comment
	Points     float32   `json:"points"`      // the number of upvotes - downvotes
	Datetime   int       `json:"datetime"`    // Timestamp of creation, epoch time
	ParentID   int       `json:"parent_id"`   // If this is a reply, this will be the value of the comment_id for the caption this a reply for.
	Deleted    bool      `json:"deleted"`     // Marked true if this caption has been deleted
	Vote       string    `json:"vote"`        // The current user's vote on the comment. null if not signed in or if the user hasn't voted on it.
	Children   []Comment `json:"children"`    // All of the replies for this comment. If there are no replies to the comment then this is an empty set.
}
