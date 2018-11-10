package dbl

import (
	"encoding/json"
)

type User struct {
	// The id of the user
	ID string `json:"id"`

	// The username of the user
	Username string `json:"username"`

	// The discriminator of the user
	Discriminator string `json:"discriminator"`

	// The avatar hash of the user's avatar (may be empty)
	Avatar string `json:"avatar"`

	// The cdn hash of the user's avatar if the user has none
	DefAvatar string `json:"defAvatar"`

	// The bio of the user
	Biography string `json:"bio"`

	// The banner image url of the user (may be empty)
	Banner string `json:"banner"`

	Social *Social `json:"social"`

	// The custom hex color of the user (may be empty)
	Color string `json:"color"`

	// The supporter status of the user
	Supporter bool `json:"supporter"`

	// The certified status of the user
	CertifiedDeveloper bool `json:"certifiedDev"`

	// The mod status of the user
	Moderator bool `json:"mod"`

	// The website moderator status of the user
	WebsiteModerator bool `json:"webMod"`

	// 	The admin status of the user
	Admin bool `json:"admin"`
}

type Social struct {
	// The youtube channel id of the user (may be empty)
	Youtube string `json:"youtube"`

	// The reddit username of the user (may be empty)
	Reddit string `json:"reddit"`

	// 	The twitter username of the user (may be empty)
	Twitter string `json:"twitter"`

	// The instagram username of the user (may be empty)
	Instagram string `json:"instagram"`

	// The github username of the user (may be empty)
	Github string `json:"github"`
}

// Information about a particular user
func (c *DBLClient) GetUser(UserID string) (*User, error) {
	res, err := c.client.Get(BaseURL + "users/" + UserID)

	if err != nil {
		return nil, err
	}

	body, err := c.readBody(res)

	if err != nil {
		return nil, err
	}

	user := &User{}

	err = json.Unmarshal(body, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
