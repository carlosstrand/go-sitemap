package sitemap

import (
	"encoding/xml"
	"fmt"
	"time"
)

// xmlSitemap is the XML representation of <urlset>...</urlset> used in sitemap
type xmlSitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []*xmlSitemapItem    `xml:"url"`
	Xmlns     string   `xml:"xmlns,attr"`
}

// xmlSitemapItem is the XML representation of <url> in <sitemap>
type xmlSitemapItem struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

// Sitemap is the structure used to create new sitemap
type Sitemap struct {
	items []*SitemapItem
	options * Options
}

type SitemapItem struct {
	Loc        string
	LastMod time.Time
	ChangeFreq string
	Priority   float32
}

func isValidItem(i * SitemapItem) bool {
	if !validateURL(i.Loc) {
		_ = fmt.Errorf("[validation error] Invalid URL: %s\n", i.Loc)
		return false
	}
	if !validateChangeFreq(i.ChangeFreq) {
		_ = fmt.Errorf("[validation error] Invalid ChangeFreq: %s\n", i.ChangeFreq)
		return false
	}
	return true
}

// NewSitemap creates a new Sitemap
func NewSitemap(items []*SitemapItem, opts * Options) *Sitemap {
	s := Sitemap{
		items: items,
		options: opts,
	}
	if s.options == nil {
		s.options = NewOptions()
	}
	return &s
}

func (s * Sitemap) ToXMLString() (string, error) {
	itemsXML := make([]*xmlSitemapItem, len(s.items))
	for idx, i := range s.items {
		if s.options.validate && !isValidItem(i) {
			return "", ErrValidation
		}
		itemsXML[idx] = &xmlSitemapItem {
			Loc: i.Loc,
			LastMod: formatTime(i.LastMod),
			ChangeFreq: i.ChangeFreq,
			Priority: fmt.Sprintf("%.1f", i.Priority),
		}
	}
	siXML := xmlSitemap{
		URL: itemsXML,
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	result, err := xmlMarshal(s.options, siXML)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s * Sitemap) AddItem(loc string, lastMod time.Time, changeFreq string, priority float32) {
	s.items = append(s.items, &SitemapItem{Loc: loc, LastMod: lastMod, ChangeFreq: changeFreq, Priority: priority})
}

func (s * Sitemap) RemoveItem(idx int) {
	s.items = append(s.items[:idx], s.items[idx+1:]...)
}