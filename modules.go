package main

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

var startText = `<b>🚀 Don't Miss Out on Free Signals! 📡</b>

<i>✅ Join our <b>free signals channel</b> today and stay ahead of the game. Click the buttons below to join and start receiving valuable insights right away!</i>

<b>🌟 Thank you!</b> `

var startButton = gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{
			gotgbot.InlineKeyboardButton{Text: "💰 Gold Signals", CallbackData: "query#gold_signals"},
			gotgbot.InlineKeyboardButton{Text: "💱 Forex Signals", CallbackData: "query#forex_signals"},
		},
		{
			gotgbot.InlineKeyboardButton{Text: "₿ Crypto Signals", CallbackData: "query#crypto_signals"},
			gotgbot.InlineKeyboardButton{Text: "📈 Stock Signals", CallbackData: "query#stock_signals"},
		},
	},
}

func registerHandlers(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", start))
	d.AddHandler(handlers.NewCallback(callbackquery.Prefix("query"), handleCallbackQuery))
	d.AddHandler(handlers.NewCallback(callbackquery.Equal("home"), handleHomeCallback))
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveChat.Id
	_, err := b.SendMessage(chatId, startText, &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: startButton,
	})
	return err

}

func handleCallbackQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.CallbackQuery
	querydata := strings.Split(query.Data, "#")
	query.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "⚡ Fetching Data. Please wait ...",
	})
	switch querydata[1] {
	case "gold_signals":
		handleGoldSignals(b, query)

	case "forex_signals":
		handleForexSignals(b, query)

	case "crypto_signals":
		handleCryptoSignals(b, query)

	case "stock_signals":
		handleStockSignals(b, query)

	default:
		return nil

	}

	return nil

}

func handleGoldSignals(b *gotgbot.Bot, q *gotgbot.CallbackQuery) {
	text := ` 
<b>💰 Gold Signals</b>

🔹 Daily 3 - 6 GOLD Signals,
🔹 Expert analysis on gold price movements,


<b>🔸 Don't miss! Join right now using button below! </b>
`
	sendMessageWithChannelWithButton(b, q, text, "https://t.me/+Jil6ArNaTd4yZjA1")

}

func handleForexSignals(b *gotgbot.Bot, q *gotgbot.CallbackQuery) {
	text := `
	<b>💱 Forex Signals</b>

	🔹 Real-time forex trading alerts,
	🔹 Exclusive signals for major currency pairs,


	<b>🔸 Don't miss! Join right now using button below!	</b>
	`

	sendMessageWithChannelWithButton(b, q, text, "https://t.me/+-qRWp_QWJgBjNzA9")

}

func handleCryptoSignals(b *gotgbot.Bot, q *gotgbot.CallbackQuery) {

	text := `
	<b>₿ Crypto Signals</b>

	🔹 Latest updates on top cryptocurrencies,
	🔹 Insights on altcoin trading opportunities,


	<b>🔸 Don't miss! Join right now using button below! </b>
		
	`
	sendMessageWithChannelWithButton(b, q, text, "https://t.me/+trZ-UNf6bSM0N2Y1")
}

func handleStockSignals(b *gotgbot.Bot, q *gotgbot.CallbackQuery) {
	text := `
	<b>📈 Stock Signals</b>

	🔹 Timely stock market notifications,
	🔹 Recommendations for blue-chip stocks and more.


	<b>🔸 Don't miss! Join right now using button below! </b>
	`
	sendMessageWithChannelWithButton(b, q, text, "https://t.me/+xcTc7pbCcwY2ZWY1")

}

func sendMessageWithChannelWithButton(b *gotgbot.Bot, q *gotgbot.CallbackQuery, text string, channelLink string) error {
	_, _, err := q.Message.EditText(b, text, &gotgbot.EditMessageTextOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					gotgbot.InlineKeyboardButton{Text: "➡️ Join Now", Url: channelLink},
				},
				{
					gotgbot.InlineKeyboardButton{Text: "🔙 Back", CallbackData: "home"},
				},
			},
		},
	})

	return err
}

func handleHomeCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.CallbackQuery
	query.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "⚡ Fetching Data. Please wait ...",
	})

	_, _, err := query.Message.EditText(
		b, startText, &gotgbot.EditMessageTextOpts{
			ParseMode:   "html",
			ReplyMarkup: startButton,
		},
	)
	return err
}
