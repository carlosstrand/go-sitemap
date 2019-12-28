package sitemap

import (
	"testing"
	"time"
)

var NewYorkLoc, _ = time.LoadLocation("America/New_York")
var MockIndexItems = []*SitemapIndexItem{
	{
		Loc: "http://mysite.com/sitemap_1.xml",
		LastMod: time.Date(2019, 10, 10, 8, 0, 0, 0, NewYorkLoc),
	},
	{
		Loc: "http://mysite.com/sitemap_2.xml",
		LastMod: time.Date(2020, 10, 10, 8, 0, 0, 0, NewYorkLoc),
	},
}

func assertIndexItem(t *testing.T, i *SitemapIndexItem, loc string, lastMod time.Time) {
	if i.Loc != loc {
		t.Errorf("Assert loc failed. Expected %s, Received: %s", loc, i.Loc)
	}
	if !i.LastMod.Equal(lastMod) {
		t.Errorf("Assert lastMod failed. Expected %s, Received: %s", lastMod, i.LastMod)
	}
}

func TestNewSitemapIndex(t *testing.T) {
	si := NewSitemapIndex(MockIndexItems, nil)
	for idx, i := range MockIndexItems {
		assertIndexItem(t, i, si.items[idx].Loc, si.items[idx].LastMod)
	}
}

func TestAddItem(t *testing.T) {
	si := NewSitemapIndex(MockIndexItems, nil)
	item := SitemapIndexItem{
		Loc:     "http://mysite.com/sitemap_3.xml",
		LastMod: time.Date(2020, 10, 10, 8, 0, 0, 0, NewYorkLoc),
	}
	si.addItem(item.Loc, item.LastMod)
	if len(si.items) != 3 {
		t.Error("Expect items length to be equal 3")
	}
	assertIndexItem(t, si.items[2], item.Loc, item.LastMod)
}

func TestToXMLString(t *testing.T) {
	si := NewSitemapIndex(MockIndexItems[:1], nil)
	expectedXMLString := `<sitemapindex><sitemap><loc>http://mysite.com/sitemap_1.xml</loc><lastmod>2019-10-10T08:00:00-04:00</lastmod></sitemap></sitemapindex>`
	xmlString, _ := si.toXMLString()
	if xmlString != expectedXMLString {
		t.Error("received xml string is not as expected")
	}
}

func TestToXMLStringWithValidation(t *testing.T) {
	si := NewSitemapIndex([]*SitemapIndexItem{
		{
			Loc: "http//invalid-url.com",
			LastMod: time.Now(),
		},
	}, nil)
	_, err := si.toXMLString()
	if err != ErrValidation {
		t.Error("ErrValidation should be returned")
	}
}