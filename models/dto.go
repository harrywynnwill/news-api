package models

import "encoding/xml"

type ReutersDTO struct {
	XMLName    xml.Name `xml:"rss"`
	Text       string   `xml:",chardata"`
	Feedburner string   `xml:"feedburner,attr"`
	Version    string   `xml:"version,attr"`
	Channel    struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text   string `xml:",chardata"`
			Atom10 string `xml:"atom10,attr"`
			Rel    string `xml:"rel,attr"`
			Type   string `xml:"type,attr"`
			Href   string `xml:"href,attr"`
		} `xml:"link"`
		Description string `xml:"description"`
		Image       struct {
			Text   string `xml:",chardata"`
			Title  string `xml:"title"`
			Width  string `xml:"width"`
			Height string `xml:"height"`
			Link   string `xml:"link"`
			URL    string `xml:"url"`
		} `xml:"image"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Copyright     string `xml:"copyright"`
		Info          struct {
			Text string `xml:",chardata"`
			URI  string `xml:"uri,attr"`
		} `xml:"info"`
		EmailServiceId     string `xml:"emailServiceId"`
		FeedburnerHostname string `xml:"feedburnerHostname"`
		FeedFlare          []struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Src  string `xml:"src,attr"`
		} `xml:"feedFlare"`
		Item []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Category string `xml:"category"`
			PubDate  string `xml:"pubDate"`
			OrigLink string `xml:"origLink"`
		} `xml:"item"`
	} `xml:"channel"`
}


type BBCDTO struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Dc      string   `xml:"dc,attr"`
	Content string   `xml:"content,attr"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"media,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Image       struct {
			Text  string `xml:",chardata"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Generator     string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		Copyright     string `xml:"copyright"`
		Language      string `xml:"language"`
		Ttl           string `xml:"ttl"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}
