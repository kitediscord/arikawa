package discord

import (
	"fmt"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/utils/json/enum"
)

// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	// ID is the id of the message.
	ID MessageID `json:"id"`
	// ChannelID is the id of the channel the message was sent in.
	ChannelID ChannelID `json:"channel_id"`
	// GuildID is the id of the guild the message was sent in.
	GuildID GuildID `json:"guild_id,omitempty"`

	// Type is the type of message.
	Type MessageType `json:"type"`

	// Flags are the MessageFlags.
	Flags MessageFlags `json:"flags"`

	// TTS specifies whether the was a TTS message.
	TTS bool `json:"tts"`
	// Pinned specifies whether the message is pinned.
	Pinned bool `json:"pinned"`

	// MentionEveryone specifies whether the message mentions everyone.
	MentionEveryone bool `json:"mention_everyone"`
	// Mentions contains the users specifically mentioned in the message.
	//
	// The user objects in the mentions array will only have the partial
	// member field present in MESSAGE_CREATE and MESSAGE_UPDATE events from
	// text-based guild channels.
	Mentions []GuildUser `json:"mentions"`
	// MentionRoleIDs contains the ids of the roles specifically mentioned in
	// the message.
	MentionRoleIDs []RoleID `json:"mention_roles"`
	// MentionChannels are the channels specifically mentioned in the message.
	//
	// Not all channel mentions in a message will appear in mention_channels.
	// Only textual channels that are visible to everyone in a lurkable guild
	// will ever be included. Only crossposted messages (via Channel Following)
	// currently include mention_channels at all. If no mentions in the message
	// meet these requirements, the slice will be empty.
	MentionChannels []ChannelMention `json:"mention_channels,omitempty"`

	// Author is the author of the message.
	//
	// The author object follows the structure of the user object, but is only
	// a valid user in the case where the message is generated by a user or bot
	// user. If the message is generated by a webhook, the author object
	// corresponds to the webhook's id, username, and avatar. You can tell if a
	// message is generated by a webhook by checking for the webhook_id on the
	// message object.
	Author User `json:"author"`

	// Content contains the contents of the message.
	Content string `json:"content"`

	// Timestamp specifies when the message was sent
	Timestamp Timestamp `json:"timestamp,omitempty"`
	// EditedTimestamp specifies when this message was edited.
	//
	// IsValid() will return false, if the messages hasn't been edited.
	EditedTimestamp Timestamp `json:"edited_timestamp,omitempty"`

	// Attachments contains any attached files.
	Attachments []Attachment `json:"attachments"`
	// Embeds contains any embedded content.
	Embeds []Embed `json:"embeds"`
	// Reactions contains any reactions to the message.
	Reactions []Reaction `json:"reactions,omitempty"`
	// Components contains any attached components.
	Components ContainerComponents `json:"components,omitempty"`

	// Used for validating a message was sent
	Nonce string `json:"nonce,omitempty"`

	// WebhookID contains the ID of the webhook, if the message was generated
	// by a webhook.
	WebhookID WebhookID `json:"webhook_id,omitempty"`

	// Activity is sent with Rich Presence-related chat embeds.
	Activity *MessageActivity `json:"activity,omitempty"`
	// Application is sent with Rich Presence-related chat embeds.
	Application *MessageApplication `json:"application,omitempty"`

	// ApplicationID contains the ID of the application, if this message was
	// generated by an interaction or an application-owned webhook.
	ApplicationID AppID `json:"application_id,omitempty"`

	// Reference is the reference data sent with crossposted messages and
	// inline replies.
	Reference *MessageReference `json:"message_reference,omitempty"`
	// ReferencedMessage is the message that was replied to. If not present and
	// the type is InlinedReplyMessage, the backend couldn't fetch the
	// replied-to message. If null, the message was deleted. If present and
	// non-null, it is a message object
	ReferencedMessage *Message `json:"referenced_message,omitempty"`

	// Interaction is the interaction that the message is in response to.
	// This is only present if the message is in response to an interaction.
	Interaction *MessageInteraction `json:"interaction,omitempty"`

	// Stickers contains the sticker "items" sent with the message.
	Stickers []StickerItem `json:"sticker_items,omitempty"`
}

// URL generates a Discord client URL to the message. If the message doesn't
// have a GuildID, it will generate a URL with the guild "@me".
func (m Message) URL() string {
	var guildID = "@me"
	if m.GuildID.IsValid() {
		guildID = m.GuildID.String()
	}

	return fmt.Sprintf(
		"https://discord.com/channels/%s/%s/%s",
		guildID, m.ChannelID.String(), m.ID.String(),
	)
}

type MessageType uint8

// https://discord.com/developers/docs/resources/channel#message-object-message-types
const (
	DefaultMessage MessageType = iota
	RecipientAddMessage
	RecipientRemoveMessage
	CallMessage
	ChannelNameChangeMessage
	ChannelIconChangeMessage
	ChannelPinnedMessage
	GuildMemberJoinMessage
	NitroBoostMessage
	NitroTier1Message
	NitroTier2Message
	NitroTier3Message
	ChannelFollowAddMessage
	_
	GuildDiscoveryDisqualifiedMessage
	GuildDiscoveryRequalifiedMessage
	GuildDiscoveryGracePeriodInitialWarning
	GuildDiscoveryGracePeriodFinalWarning
	// ThreadCreatedMessage is a new message sent to the parent GuildText
	// channel, used to inform users that a thread has been created. It is
	// currently only sent in one case: when a GuildPublicThread is created
	// from an older message (older is still TBD, but is currently set to a
	// very small value). The message contains a message reference with the
	// GuildID and ChannelID of the thread. The content of the message is the
	// name of the thread.
	ThreadCreatedMessage
	InlinedReplyMessage
	ChatInputCommandMessage
	// ThreadStarterMessage is a new message sent as the first message in
	// threads that are started from an existing message in the parent channel.
	// It only contains a message reference field that points to the message
	// from which the thread was started.
	ThreadStarterMessage
	GuildInviteReminderMessage
	ContextMenuCommand
)

type MessageFlags enum.Enum

// NullMessage is the JSON null value of MessageFlags.
const NullMessage MessageFlags = enum.Null

// https://discord.com/developers/docs/resources/channel#message-object-message-flags
const (
	// CrosspostedMessage specifies whether the message has been published to
	// subscribed channels (via Channel Following).
	CrosspostedMessage MessageFlags = 1 << iota
	// MessageIsCrosspost specifies whether the message originated from a
	// message in another channel (via Channel Following).
	MessageIsCrosspost
	// SuppressEmbeds specifies whether to not include any embeds when
	// serializing the message.
	SuppressEmbeds
	// SourceMessageDeleted specifies whether the source message for the
	// crosspost has been deleted (via Channel Following).
	SourceMessageDeleted
	// UrgentMessage specifies whether the message came from the urgent message
	// system.
	UrgentMessage
	// MessageHasThread specifies whether the message has an associated thread
	// with the same id as the message
	MessageHasThread
	// EphemeralMessage specifies whether the message is only visible to
	// the user who invoked the Interaction
	EphemeralMessage
	// MessageLoading specifies whether the message is an Interaction Response
	// and the bot is "thinking"
	MessageLoading
)

// StickerItem contains partial data of a Sticker.
//
// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type StickerItem struct {
	// ID is the ID of the sticker.
	ID StickerID `json:"id"`
	// Name is the name of the sticker.
	Name string `json:"name"`
	// FormatType is the type of sticker format.
	FormatType StickerFormatType `json:"format_type"`
}

// StickerURLWithType returns the URL to the emoji's image.
//
// Supported ImageTypes: PNG
func (s StickerItem) StickerURLWithType(t ImageType) string {
	return "https://cdn.discordapp.com/stickers/" + t.format(s.ID.String())
}

// https://discord.com/developers/docs/resources/channel#message-object-message-sticker-structure
type Sticker struct {
	// ID is the ID of the sticker.
	ID StickerID `json:"id"`
	// PackID is the ID of the pack the sticker is from.
	PackID StickerPackID `json:"pack_id,omitempty"`
	// Name is the name of the sticker.
	Name string `json:"name"`
	// Description is the description of the sticker.
	Description string `json:"description"`
	// Tags is a comma-delimited list of tags for the sticker. To get the list
	// as a slice, use TagList.
	Tags string `json:"tags"`
	// Type is the type of sticker.
	Type StickerType `json:"type"`
	// FormatType is the type of sticker format.
	FormatType StickerFormatType `json:"format_type"`
	// Available specifies whether this guild sticker can be used, may be false due to loss of Server Boosts.
	Available bool `json:"available,omitempty"`
	// GuildID is the id of the guild that owns this sticker.
	GuildID GuildID `json:"guild_id,omitempty"`
	// User is the user that uploaded the guild sticker
	User *User `json:"user,omitempty"`
	// SortValue is the standard sticker's sort order within its pack.
	SortValue *int `json:"sort_value,omitempty"`
}

// CreatedAt returns a time object representing when the sticker was created.
func (s Sticker) CreatedAt() time.Time {
	return s.ID.Time()
}

// PackCreatedAt returns a time object representing when the sticker's pack
// was created.
func (s Sticker) PackCreatedAt() time.Time {
	return s.PackID.Time()
}

// TagList splits the sticker tags into a slice of strings. Each tag will have
// its trailing space trimmed.
func (s Sticker) TagList() []string {
	tags := strings.Split(s.Tags, ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}
	return tags
}

// StickerURLWithType returns the URL to the emoji's image.
//
// Supported ImageTypes: PNG
func (s Sticker) StickerURLWithType(t ImageType) string {
	return "https://cdn.discordapp.com/stickers/" + t.format(s.ID.String())
}

type StickerType int

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
const (
	// StandardSticker is an official sticker in a pack, part of Nitro or in a removed purchasable pack.
	StandardSticker StickerType = iota + 1
	// GuildSticker is a sticker uploaded to a boosted guild for the guild's members.
	GuildSticker
)

type StickerFormatType uint8

// https://discord.com/developers/docs/resources/channel#message-object-message-sticker-format-types
const (
	StickerFormatPNG    = 1
	StickerFormatAPNG   = 2
	StickerFormatLottie = 3
)

// https://discord.com/developers/docs/resources/channel#channel-mention-object
type ChannelMention struct {
	// ChannelID is the ID of the channel.
	ChannelID ChannelID `json:"id"`
	// GuildID is the ID of the guild containing the channel.
	GuildID GuildID `json:"guild_id"`
	// ChannelType is the type of channel.
	ChannelType ChannelType `json:"type"`
	// ChannelName is the name of the channel.
	ChannelName string `json:"name"`
}

type GuildUser struct {
	User
	Member *Member `json:"member,omitempty"`
}

//

// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	// Type is the type of message activity.
	Type MessageActivityType `json:"type"`
	// PartyID is the party_id from a Rich Presence event.
	PartyID string `json:"party_id,omitempty"`
}

type MessageActivityType uint8

// https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
const (
	JoinMessage MessageActivityType = iota + 1
	SpectateMessage
	ListenMessage
	JoinRequestMessage
)

//

// https://discord.com/developers/docs/resources/channel#message-object-message-application-structure
type MessageApplication struct {
	// ID is the id of the application.
	ID AppID `json:"id"`
	// CoverID is the id of the embed's image asset.
	CoverID string `json:"cover_image,omitempty"`
	// Description is the application's description.
	Description string `json:"description"`
	// Icon is the id of the application's icon.
	Icon string `json:"icon"`
	// Name is the name of the application.
	Name string `json:"name"`
}

// CreatedAt returns a time object representing when the message application
// was created.
func (m MessageApplication) CreatedAt() time.Time {
	return m.ID.Time()
}

// MessageReference is used in four situations:
//
// Crosspost messages
//
// Messages that originated from another channel (IS_CROSSPOST flag). These
// messages have all three fields, with data of the original message that was
// crossposted.
//
// Channel Follow Add messages
//
// Automatic messages sent when a channel is followed into the current channel
// (type 12). These messages have the ChannelID and GuildID fields, with data
// of the followed announcement channel.
//
// Pin messages
//
// Automatic messages sent when a message is pinned (type 6). These messages
// have MessageID and ChannelID, and GuildID if it is in a guild, with data
// of the message that was pinned.
//
// Replies
//
// Messages replying to a previous message (type 19). These messages have
// MessageID, and ChannelID, and GuildID if it is in a guild, with data of the
// message that was replied to. The ChannelID and GuildID will be the
// same as the reply.
//
// Replies are created by including a message_reference when sending a message.
// When sending, only MessageID is required.
// https://discord.com/developers/docs/resources/channel#message-object-message-reference-structure
type MessageReference struct {
	// MessageID is the id of the originating message.
	MessageID MessageID `json:"message_id,omitempty"`
	// ChannelID is the id of the originating message's channel.
	ChannelID ChannelID `json:"channel_id,omitempty"`
	// GuildID is the id of the originating message's guild.
	GuildID GuildID `json:"guild_id,omitempty"`
}

//

// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	// ID is the id of the originating interaction.
	ID InteractionID `json:"id"`
	// Type is the type of the originating interaction.
	Type InteractionDataType `json:"type"`
	// Name is the name of the application command that was invoked with the
	// originating interaction.
	Name string `json:"name"`
	// User is the user who invoked the originating interaction.
	User User `json:"user"`
	// Member is the member who invoked the originating interaction in
	// the guild.
	Member *Member `json:"member,omitempty"`
}

//

// https://discord.com/developers/docs/resources/channel#attachment-object
type Attachment struct {
	// ID is the attachment id.
	ID AttachmentID `json:"id"`
	// Filename is the name of file attached.
	Filename string `json:"filename"`
	// Description is the attachment's description. It is a maximum of 1024
	// characters long.
	Description string `json:"description,omitempty"`
	// ContentType is the media type of file.
	ContentType string `json:"content_type,omitempty"`
	// Size is the size of file in bytes.
	Size uint64 `json:"size"`

	// URL is the source url of file.
	URL URL `json:"url"`
	// Proxy is the a proxied url of file.
	Proxy URL `json:"proxy_url"`

	// Height is the height of the file, if it is an image.
	Height uint `json:"height,omitempty"`
	// Width is the width of the file, if it is an image.
	Width uint `json:"width,omitempty"`
	// Ephemeral is whether this attachment is ephemeral. Ephemeral attachments
	// will automatically be removed after a set period of time. Ephemeral
	// attachments on messages are guaranteed to be available as long as
	// the message itself exists.
	Ephemeral bool `json:"ephemeral,omitempty"`
}

//

// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	// Count is the amount of times the emoji has been used to react.
	Count int `json:"count"`
	// Me specifies whether the current user reacted using this emoji.
	Me bool `json:"me"`
	// Emoji contains emoji information.
	Emoji Emoji `json:"emoji"`
}
