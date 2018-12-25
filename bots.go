package dbl

import (
	"bytes"
	"encoding/json"
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
	Owners []string `json:"owners"`

	// The custom bot invite url of the bot (may be empty)
	Invite string `json:"invite"`

	// The date when the bot was approved
	Date time.Time `json:"date"`

	// The certified status of the bot
	CertifiedBot bool `json:"certifiedBot"`

	// The vanity url of the bot (deprecated) (may be empty)
	Vanity string `json:"vanity"`

	// The monthly amount of upvotes the bot has (undocumented)
	MonthlyPoints int `json:"monthlyPoints"`

	// The amount of upvotes the bot has
	Points int `json:"points"`

	// The GuildID for the donate bot (undocumented) (may be empty)
	DonateBotGuildID string `json:"donatebotguildid"`

	// The amount of servers the bot is in (undocumented)
	ServerCount int `json:"server_count"`

	// Server affiliation ("Servers this bot is in" field) (undocumented)
	GuildAffiliation []string `json:"guilds"`

	// The amount of servers the bot is in per shard. Always present but can be empty (undocumented)
	Shards []int `json:"shards"`
}

type GetBotsPayload struct {
	// The amount of bots to return. Max. 500
	// Default 50
	Limit int

	// Amount of bots to skip
	// Default 0
	Offset int

	// Field search filter
	Search map[string]string

	// The field to sort by. Prefix with "-" to reverse the order
	Sort string

	// A list of fields to show
	Fields []string
}

type GetBotsResult struct {
	// Slice of Bot pointers of matching bots
	Results []*Bot `json:"results"`

	// The limit used
	Limit int `json:"limit"`

	// The offset used
	Offset int `json:"offset"`

	// The length of the results array
	Count int `json:"count"`

	// The total number of bots matching your search
	// Not limited by "limit" field
	Total int `json:"total"`
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

	// The zero-indexed id of the shard posting. Makes server_count set the shard specific server count (optional)
	ShardID int `json:"shard_id"`

	// The amount of shards the bot has (optional)
	ShardCount int `json:"shard_count"`
}

// Information about different bots with an optional filter parameter
//
// Use nil if no option is passed
func (c *DBLClient) GetBots(filter *GetBotsPayload) (*GetBotsResult, error) {
	if c.token == "" {
		return nil, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := c.createRequest("GET", "bots", nil)

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

	body, err := c.readBody(res)

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

// Information about a specific bot
func (c *DBLClient) GetBot(botID string) (*Bot, error) {
	if c.token == "" {
		return nil, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := c.createRequest("GET", "bots/"+botID, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := c.readBody(res)

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

// Use this endpoint to see who have upvoted your bot
//
// Requires authentication
//
// IF YOU HAVE OVER 1000 VOTES PER MONTH YOU HAVE TO USE THE WEBHOOKS AND CAN NOT USE THIS
func (c *DBLClient) GetVotes(botID string) ([]*User, error) {
	if c.token == "" {
		return nil, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := c.createRequest("GET", "bots/"+botID+"/votes", nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := c.readBody(res)

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

// Use this endpoint to see who have upvoted your bot in the past 24 hours. It is safe to use this even if you have over 1k votes.
//
// Requires authentication
func (c *DBLClient) HasUserVoted(botID, userID string) (bool, error) {
	if c.token == "" {
		return false, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return false, ErrLocalRatelimit
	}

	req, err := c.createRequest("GET", "bots/"+botID+"/check", nil)

	if err != nil {
		return false, err
	}

	q := req.URL.Query()

	q.Add("userId", userID)

	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)

	if err != nil {
		return false, err
	}

	body, err := c.readBody(res)

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

// Information about a specific bot's stats
func (c *DBLClient) GetBotStats(botID string) (*BotStats, error) {
	if c.token == "" {
		return nil, ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return nil, ErrLocalRatelimit
	}

	req, err := c.createRequest("GET", "bots/"+botID+"/stats", nil)

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := c.readBody(res)

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

// Post your bot's stats
//
// Requires authentication
//
// If your bot is unsharded, pass in server count as the only item in the slice
func (c *DBLClient) PostBotStats(botID string, payload BotStatsPayload) error {
	if c.token == "" {
		return ErrRequireAuthentication
	}

	if !c.limiter.Allow() {
		return ErrLocalRatelimit
	}

	encoded, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	req, err := c.createRequest("POST", "bots/"+botID+"/stats", bytes.NewBuffer(encoded))

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
