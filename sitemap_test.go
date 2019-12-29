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

var MockItems = []*SitemapItem{
	{
		Loc: "http://mysite.com/about",
		LastMod: time.Date(2019, 10, 10, 8, 0, 0, 0, NewYorkLoc),
		ChangeFreq: "daily",
		Priority: 0.5,
	},
	{
		Loc: "http://mysite.com/contact",
		LastMod: time.Date(2020, 10, 10, 8, 0, 0, 0, NewYorkLoc),
		ChangeFreq: "monthly",
		Priority: 0.5,
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

func assertItem(t *testing.T, i *SitemapItem, loc string, lastMod time.Time, changeFreq string, priority float32) {
	if i.Loc != loc {
		t.Errorf("Assert loc failed. Expected %s, Received: %s", loc, i.Loc)
	}
	if !i.LastMod.Equal(lastMod) {
		t.Errorf("Assert lastMod failed. Expected %s, Received: %s", lastMod, i.LastMod)
	}
	if i.ChangeFreq != i.ChangeFreq {
		t.Errorf("Assert changeFreq failed. Expected %s, Received: %s", changeFreq, i.ChangeFreq)
	}
	if i.Priority != i.Priority {
		t.Errorf("Assert priority failed. Expected %.1f, Received: %.1f", priority, i.Priority)
	}
}

func TestNewSitemapIndex(t *testing.T) {
	si := NewSitemapIndex(MockIndexItems, nil)
	for idx, i := range MockIndexItems {
		assertIndexItem(t, i, si.items[idx].Loc, si.items[idx].LastMod)
	}
}

func TestAddIndexItem(t *testing.T) {
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

func TestRemoveIndexItem(t *testing.T) {
	items := make([]*SitemapIndexItem, len(MockItems))
	copy(items, MockIndexItems)
	s := NewSitemapIndex(items, nil)
	s.removeItem(0)
	if len(s.items) != 1 {
		t.Error("Expect items length to be equal 1")
	}
	assertIndexItem(t, s.items[0], MockIndexItems[1].Loc, MockIndexItems[1].LastMod)
}

func TestSitemapIndexToXMLString(t *testing.T) {
	si := NewSitemapIndex(MockIndexItems[:1], nil)
	expectedXMLString := `<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><sitemap><loc>http://mysite.com/sitemap_1.xml</loc><lastmod>2019-10-10T08:00:00-04:00</lastmod></sitemap></sitemapindex>`
	xmlString, _ := si.toXMLString()
	if xmlString != expectedXMLString {
		t.Error("received xml string is not as expected")
	}
}

func TestSitemapIndexToXMLStringWithValidation(t *testing.T) {
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

func TestNewSitemap(t *testing.T) {
	s := NewSitemap(MockItems, nil)
	for idx, i := range MockItems {
		assertItem(t, i, s.items[idx].Loc, s.items[idx].LastMod, s.items[idx].ChangeFreq, s.items[idx].Priority)
	}
}

func TestAddItem(t *testing.T) {
	s := NewSitemap(MockItems, nil)
	item := SitemapItem{
		Loc:     "http://mysite.com/my-url",
		LastMod: time.Date(2020, 10, 10, 8, 0, 0, 0, NewYorkLoc),
	}
	s.addItem(item.Loc, item.LastMod, "daily", 0.5)
	if len(s.items) != 3 {
		t.Error("Expect items length to be equal 3")
	}
	assertItem(t, s.items[2], item.Loc, item.LastMod, item.ChangeFreq, item.Priority)
}

func TestRemoveItem(t *testing.T) {
	items := make([]*SitemapItem, len(MockItems))
	copy(items, MockItems)
	s := NewSitemap(items, nil)
	s.removeItem(0)
	if len(s.items) != 1 {
		t.Error("Expect items length to be equal 1")
	}
	assertItem(t, s.items[0], MockItems[1].Loc, MockItems[1].LastMod, MockItems[1].ChangeFreq, MockItems[1].Priority)
}

func TestSitemapToXMLString(t *testing.T) {
	si := NewSitemap(MockItems[:1], nil)
	expectedXMLString := `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>http://mysite.com/about</loc><lastmod>2019-10-10T08:00:00-04:00</lastmod><changefreq>daily</changefreq><priority>0.5</priority></url></urlset>`
	xmlString, _ := si.toXMLString()
	if xmlString != expectedXMLString {
		t.Error("received xml string is not as expected")
	}
}
