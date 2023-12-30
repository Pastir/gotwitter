package gotwitter

// Manage Tweets https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets

type (
	Geo struct {
		// Optional
		// Place ID being attached to the Tweet for geo location.
		PlaceId string `json:"place_id"`
	}

	Media struct {
		// Optional
		// A list of Media IDs being attached to the Tweet.
		// This is only required if the request includes the tagged_user_ids.
		MediaIds []string `json:"media_ids,omitempty"`

		// Optional
		// A list of User IDs being tagged in the Tweet with Media.
		// If the user you're tagging doesn't have photo-tagging enabled, their names won't show up in the list of
		// tagged users even though the Tweet is successfully created.
		TaggedUserId []string `json:"tagged_user_id,omitempty"`
	}

	Poll struct {
		// Optional
		// Duration of the poll in minutes for a Tweet with a poll.
		// This is only required if the request includes poll.options.
		DurationMinute int `json:"duration_minute,omitempty"`

		// Optional
		// A list of poll options for a Tweet with a poll.
		// For the request to be successful it must also include duration_minutes too
		Options []string `json:"options,omitempty"`
	}

	Reply struct {
		// Optional
		// A list of User IDs to be excluded from the reply Tweet thus removing a user from a thread.
		ExcludeReplyUserIds []string `json:"exclude_reply_user_ids,omitempty"`

		// Optional
		// Tweet ID of the Tweet being replied to.
		// Please note that in_reply_to_tweet_id needs to be in the request if exclude_reply_user_ids is present.
		InReplyToTweetId string `json:"in_reply_to_tweet_id,omitempty"`
	}

	Tweet struct {
		// Optional
		// Tweets a link directly to a Direct Message conversation with an account.
		DirectMessageDeepLink string `json:"direct_message_deep_link,omitempty"`

		// Optional
		// Allows you to Tweet exclusively for Super Followers.
		ForSuperFollowersOnly bool `json:"for_super_followers_only,omitempty"`

		// Optional
		// A JSON object that contains location information for a Tweet.
		// You can only add a location to Tweets if you have geo enabled in your profile settings.
		// If you don't have geo enabled, you can still add a location parameter in your request body,
		// but it won't get attached to your Tweet
		Geo *Geo `json:"geo,omitempty"`

		// Optional
		// A JSON object that contains media information being attached to created Tweet.
		// This is mutually exclusive from Quote Tweet ID and Poll.
		Media *Media `json:"media,omitempty"`

		// Optional
		// A JSON object that contains options for a Tweet with a poll.
		// This is mutually exclusive from Media and Quote Tweet ID.
		Poll *Poll `json:"poll,omitempty"`

		// Optional
		// Link to the Tweet being quoted.
		QuoteTweetId string `json:"quote_tweet_id,omitempty"`

		// Optional
		// A JSON object that contains information of the Tweet being replied to.
		Reply *Reply `json:"reply,omitempty"`

		// Optional
		// Settings to indicate who can reply to the Tweet.
		// Options include "mentionedUsers" and "following". If the field isn’t specified, it will default to everyone.
		ReplySettings string `json:"replySettings,omitempty"`

		// Optional
		// Text of the Tweet being created. This field is required if media.media_ids is not present.
		Text string `json:"text,omitempty"`
	}
)

func (t *Tweet) Create() {

}
