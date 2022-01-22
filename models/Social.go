package models

type Social struct {
	Type          SocialType `json:"type"`
	Caption       string     `json:"caption"`
	Permalink     string     `json:"permalink"`
	MediaUrlSnake string     `json:"media_url"`
	MediaUrlCamel string     `json:"mediaUrl"`
	ThumbnailUrl  string     `json:"thumbnail_url"`
	MediaType     MediaType  `json:"media_type"`
}

func (i *Social) SetMediaSRC() {
	if i.MediaType == MediaTypeImage {
		i.MediaUrlCamel = i.MediaUrlSnake
	}
	if i.MediaType == MediaTypeVideo {
		i.MediaUrlCamel = i.ThumbnailUrl
	}
}

type Socials []*Social

type SocialData struct {
	Socials Socials `json:"data"`
}

type SocialType string

const (
	Instagram SocialType = "Instagram"
)

type MediaType string

const (
	MediaTypeImage         MediaType = "IMAGE"
	MediaTypeVideo         MediaType = "VIDEO"
	MediaTypeCarouselAlbum MediaType = "CAROUSEL_ALBUM"
)
