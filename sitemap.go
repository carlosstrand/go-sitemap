package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"time"
)

const Header = `<?xml version="1.0" encoding="UTF-8"?>\n`

var ErrValidation = errors.New("validation error")

// XMLSitemapIndex is the XML representation of <sitemap>...</sitemap> used in sitemapindex
type XMLSitemapIndexItem struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

// XMLSitemapIndex is the XML representation of <sitemapindex>...</sitemapindex>
type XMLSitemapIndex struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Sitemap []*XMLSitemapIndexItem  `xml:"sitemap"`
}

// XMLSitemapURLSet is the XML representation of <urlset>...</urlset> used in sitemap
type XMLSitemapURLSet struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []*XMLSitemapURL    `xml:"url"`
}

// XMLSitemapURL is the XML representation of <url> in <sitemap>
type XMLSitemapURL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   float32 `xml:"priority"`
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
		si.options = &Options{
			prettyOutput:  false,
			withXMLHeader: true,
			validate: true,
		}
	}
	return &si
}

func (si * SitemapIndex) addItem(loc string, lastMod time.Time) {
	si.items = append(si.items, &SitemapIndexItem{Loc: loc, LastMod: lastMod})
}

func (si * SitemapIndex) toXMLString() (string, error) {
	itemsXML := make([]*XMLSitemapIndexItem, len(si.items))
	for idx, i := range si.items {
		if si.options.validate && !isValidIndexItem(i) {
			return "", ErrValidation
		}
		itemsXML[idx] = &XMLSitemapIndexItem {
			Loc: i.Loc,
			LastMod: formatTime(i.LastMod),
		}
	}
	siXML := XMLSitemapIndex{
		Sitemap: itemsXML,
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

func validateURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func xmlMarshal(options *Options, obj interface{}) (string, error) {
	var result []byte
	var err error
	if options != nil && options.prettyOutput {
		result, err = xml.MarshalIndent(obj, "", "  ")
	} else {
		result, err = xml.Marshal(obj)
	}
	if err != nil {
		return "", nil
	}
	return string(result), nil
}