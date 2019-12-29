package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

var ErrValidation = errors.New("validation error")

// xmlSitemapIndex is the XML representation of <sitemap>...</sitemap> used in sitemapindex
type xmlSitemapIndexItem struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

// xmlSitemapIndex is the XML representation of <sitemapindex>...</sitemapindex>
type xmlSitemapIndex struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Sitemap []*xmlSitemapIndexItem  `xml:"sitemap"`
	Xmlns     string   `xml:"xmlns,attr"`
}

type Options struct {
	prettyOutput bool
	withXMLHeader bool
	validate bool
}

// SitemapIndex is the structure used to create new sitemap index
type SitemapIndex struct {
	items []*SitemapIndexItem
	options * Options
}

type SitemapIndexItem struct {
	Loc string
	LastMod time.Time
}

// NewSitemapIndex creates a new Sitemap Index
func NewSitemapIndex(items []*SitemapIndexItem, opts * Options) *SitemapIndex {
	si := SitemapIndex{
		items: items,
		options: opts,
	}
	if si.options == nil {
		si.options = NewOptions()
	}
	return &si
}

func (si * SitemapIndex) AddItem(loc string, lastMod time.Time) {
	si.items = append(si.items, &SitemapIndexItem{Loc: loc, LastMod: lastMod})
}

func (si * SitemapIndex) RemoveItem(idx int) {
	si.items = append(si.items[:idx], si.items[idx+1:]...)
}

func (si * SitemapIndex) ToXMLString() (string, error) {
	itemsXML := make([]*xmlSitemapIndexItem, len(si.items))
	for idx, i := range si.items {
		if si.options.validate && !isValidIndexItem(i) {
			return "", ErrValidation
		}
		itemsXML[idx] = &xmlSitemapIndexItem {
			Loc: i.Loc,
			LastMod: formatTime(i.LastMod),
		}
	}
	siXML := xmlSitemapIndex{
		Sitemap: itemsXML,
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	result, err := xmlMarshal(si.options, siXML)
	if err != nil {
		return "", err
	}
	return result, nil
}

func isValidIndexItem(i * SitemapIndexItem) bool {
	if !validateURL(i.Loc) {
		_ = fmt.Errorf("[validation error] Invalid URL: %s\n", i.Loc)
		return false
	}
	return true
}