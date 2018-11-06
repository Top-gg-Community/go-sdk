package dbl

import (
	"time"
)

type Bot struct {
	// The id of the bot
	ID uint64 `json:"id"`

	// The username of the bot
	Username string `json:"username"`

	// The discriminator of the bot
	Discriminator string `json:"discriminator"`

	// The avatar hash of the bot's avatar (may be empty)
	Avatar string `json:"avatar"`

	// The cdn hash of the bot's avatar if the bot has none
	DefAvatar string `json:"defAvatar"`

	// The library of the bot
	Library string `json:"lib"`

	// The prefix of the bot
	Prefix string `json:"prefix"`

	// The short description of the bot
	ShortDescription string `json:"shortdesc"`

	// The long description of the bot. Can contain HTML and/or Markdown (may be empty)
	LongDescription string `json:"longdesc"`

	// The tags of the bot
	Tags []string `json:"tags"`

	// The website url of the bot (may be empty)
	Website string `json:"website"`

	// The support server invite code of the bot (may be empty)
	Support string `json:"support"`

	// The link to the github repo of the bot (may be empty)
	Github string `json:"github"`

	// The owners of the bot. First one in the array is the main owner
	Owners []uint64 `json:"owners"`

	// The custom bot invite url of the bot (may be empty)
	Invite string `json:"invite"`

	// The date when the bot was approved
	Date time.Time `json:"date"`

	// The certified status of the bot
	CertifiedBot bool `json:"certifiedBot"`

	// The vanity url of the bot (may be empty)
	Vanity string `json:"vanity"`

	// The amount of upvotes the bot has
	Points uint16 `json:"points"`
}
