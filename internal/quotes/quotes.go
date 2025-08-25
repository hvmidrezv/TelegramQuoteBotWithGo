package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ZenQuote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func FetchQuote() (ZenQuote, error) {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		return ZenQuote{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var quotes []ZenQuote
	if err := json.Unmarshal(body, &quotes); err != nil {
		return ZenQuote{}, err
	}
	if len(quotes) == 0 {
		return ZenQuote{}, fmt.Errorf("empty response")
	}
	return quotes[0], nil
}

func HtmlEscape(s string) string {
	repl := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return repl.Replace(s)
}

func CurrentTheme() (emoji string, title string) {
	h := time.Now().Hour()
	switch {
	case h >= 5 && h < 11:
		return "☀️", "Morning Quote"
	case h >= 11 && h < 17:
		return "🔆", "Midday Quote"
	case h >= 17 && h < 21:
		return "🌆", "Evening Quote"
	default:
		return "🌙", "Night Quote"
	}
}

func FormatQuote(q ZenQuote) string {
	e, title := CurrentTheme()
	quote := HtmlEscape(q.Q)
	author := HtmlEscape(q.A)
	t := time.Now().Format("15:04") // ساعت:دقیقه

	return fmt.Sprintf(
		`%s <b>%s</b><br>
“%s”<br>
— <i>%s</i><br>
🕒 <i>%s</i>`,
		e, title, quote, author, t,
	)
}
