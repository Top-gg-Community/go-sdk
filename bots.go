package dbl

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Bot struct {
	// The id of the bot
	ID string `json:"id"`

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
	Vanity *string `json:"vanity"`

	// The amount of upvotes the bot has
	Points int `json:"points"`
}

type GetBotsPayload struct {
	Limit  int
	Offset int
	Search map[string]string
	Sort   string
	Fields []string
}

type GetBotsResult struct {
	Results []*Bot `json:"results"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Total   int    `json:"total"`
}

type BotStats struct {
	// The amount of servers the bot is in (may be empty)
	ServerCount int `json:"server_count"`

	// The amount of servers the bot is in per shard. Always present but can be empty
	Shards []int `json:"shards"`

	// The amount of shards a bot has (may be empty)
	ShardCount int `json:"shard_count"`
}

type checkResponse struct {
	Voted int `json:"voted"`
}

type BotStatsPayload struct {
	// The amount of servers the bot is in per shard.
	Shards []int `json:"shards"`

	// The zero-indexed id of the shard posting. Makes server_count set the shard specific server count
	ShardID int `json:"shard_id"`

	// The amount of shards the bot has
	ShardCount int `json:"shard_count"`
}

func (c *DBLClient) GetBots(filter *GetBotsPayload) (*GetBotsResult, error) {
	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := http.NewRequest("GET", BaseURL+"bots", nil)

	if filter != nil {
		q := req.URL.Query()

		if filter.Limit != 0 {
			q.Add("limit", strconv.Itoa(filter.Limit))
		}

		if filter.Offset != 0 {
			q.Add("offset", strconv.Itoa(filter.Offset))
		}

		if len(filter.Search) != 0 {
			tStack := make([]string, 0)

			for f, v := range filter.Search {
				tStack = append(tStack, f+": "+v)
			}

			q.Add("search", strings.Join(tStack, " "))
		}

		if filter.Sort != "" {
			q.Add("sort", filter.Sort)
		}

		if len(filter.Fields) != 0 {
			q.Add("fields", strings.Join(filter.Fields, ","))
		}

		req.URL.RawQuery = q.Encode()
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := readBody(res)

	if err != nil {
		return nil, err
	}

	bots := &GetBotsResult{}

	err = json.Unmarshal(body, bots)

	if err != nil {
		return nil, err
	}

	return bots, nil
}

func (c *DBLClient) GetBot(botID string) (*Bot, error) {
	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	res, err := c.client.Get(BaseURL + "bots/" + botID)

	if err != nil {
		return nil, err
	}

	body, err := readBody(res)

	if err != nil {
		return nil, err
	}

	bot := &Bot{}

	err = json.Unmarshal(body, bot)

	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (c *DBLClient) GetVotes(botID string) ([]*User, error) {
	if c.token == "" {
		return nil, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := http.NewRequest("GET", BaseURL+"bots/"+botID+"/votes", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := readBody(res)

	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)

	err = json.Unmarshal(body, users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *DBLClient) HasUserVoted(botID, userID string) (bool, error) {
	if c.token == "" {
		return false, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return false, ErrLocalRatelimit
	}

	req, err := http.NewRequest("GET", BaseURL+"bots/"+botID+"/check", nil)

	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", c.token)

	q := req.URL.Query()

	q.Add("userId", userID)

	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)

	if err != nil {
		return false, err
	}

	body, err := readBody(res)

	if err != nil {
		return false, err
	}

	cr := &checkResponse{}

	err = json.Unmarshal(body, cr)

	if err != nil {
		return false, err
	}

	return cr.Voted == 1, nil
}

func (c *DBLClient) GetBotStats(botID string) (*BotStats, error) {
	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	res, err := c.client.Get(BaseURL + "bots/" + botID + "/stats")

	if err != nil {
		return nil, err
	}

	body, err := readBody(res)

	if err != nil {
		return nil, err
	}

	botStats := &BotStats{}

	err = json.Unmarshal(body, botStats)

	if err != nil {
		return nil, err
	}

	return botStats, nil
}

func (c *DBLClient) PostBotStats(botID string, payload BotStatsPayload) error {
	if !c.limiter.Allow() {
		return ErrLocalRatelimit
	}

	encoded, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", BaseURL+"bots/"+botID+"/stats", bytes.NewBuffer(encoded))

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.token)

	res, err := c.client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return ErrRequestFailed
	}

	return nil
}
