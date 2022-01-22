package socialHelper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/models"
	"net/http"
)

type Social struct {
	InstagramToken string
	InstagramID    string
}

const InstagramApi = "https://graph.instagram.com"

func NewSocial(config config.Social) *Social {
	return &Social{InstagramToken: config.InstagramToken, InstagramID: config.InstagramID}
}

func (i *Social) buildInstagramURL() string {
	return fmt.Sprintf("%s/%s/media?fields=id,media_url,media_type,thumbnail_url,permalink,caption&access_token=%s", InstagramApi, i.InstagramID, i.InstagramToken)
}

func (i *Social) GetWebFeed() models.Socials {
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, i.buildInstagramURL(), nil)
	if err != nil {
		log.Println(err)
	}
	c := &http.Client{}
	resp, err := c.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	data := models.SocialData{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	for i := range data.Socials {
		data.Socials[i].SetMediaSRC()
	}
	return data.Socials
}
