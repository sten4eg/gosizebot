package tg

type (
	Updates struct {
		Ok          bool      `json:"ok"`
		Result      *[]Update `json:"result"`
		ErrorCode   *int      `json:"error_code"`
		Description *string   `json:"description"`
	}
	Update struct {
		Id          int64        `json:"update_id"`
		Message     *Message     `json:"message"`
		InlineQuery *InlineQuery `json:"inline_query"`
	}
	Message struct {
		Id       int64
		From     *UserShort       `json:"from"`
		Chat     Chat             `json:"chat"`
		Date     int64            `json:"date"`
		Text     string           `json:"text"`
		Entities *[]MessageEntity `json:"entities"`
	}
	UserShort struct {
		Id           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	}
	Chat struct {
		Id        int64  `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	}
	MessageEntity struct {
		Offset int64  `json:"offset"`
		Length int64  `json:"length"`
		Type   string `json:"type"`
	}
	InlineQuery struct {
		Id   string          `json:"id"`
		From InlineQueryFrom `json:"from"`
	}
	InlineQueryFrom struct {
		Id           int     `json:"id"`
		IsBot        bool    `json:"is_bot"`
		FirstName    string  `json:"first_name"`
		Username     string  `json:"username"`
		LanguageCode string  `json:"language_code"`
		IsPremium    bool    `json:"is_premium"`
		ChatType     *string `json:"chat_type"`
		Query        string  `json:"query"`
		Offset       string  `json:"offset"`
	}
	InlineQueryResultArticle struct {
		Type                string              `json:"type"`
		Id                  string              `json:"id"`
		Title               string              `json:"title"`
		InputMessageContent InputMessageContent `json:"input_message_content"`
	}
	InputMessageContent struct {
		MessageText string `json:"message_text"`
	}
)
