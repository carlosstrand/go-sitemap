# go-sitemap

Go package for manipulate golang object of sitemap and generate an ready-to-use XML.

- [X] Following sitemaps.org XML format
- [X] Validate URLs and ChangeFreq
- [X] Add Header (see in options)
- [X] Tests


### Quickstart


#### Generate Sitemap

```go
package main

import (
  "fmt"
  "github.com/carlosstrand/go-sitemap"
  "log"
)

func main() {
	sitemap := NewSitemap([]*SitemapItem{
		{
			Loc: "https://my-site.com/about",
			LastMod: time.Now(),
		},
		{
			Loc: "https://my-site.com/contact",
			LastMod: time.Now(),
		},
	}, nil)
	
	xmlResult, err := sitemap.toXMLString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xmlResult)
}

```

#### Generate Sitemap Index

```go
package main

import (
  "fmt"
  "github.com/carlosstrand/go-sitemap"
  "log"
)

func main() {
	sitemapIdx := NewSitemapIndex([]*SitemapIndexItem{
		{
			Loc: "https://my-site.com/about",
			LastMod: time.Now(),
		},
		{
			Loc: "https://my-site.com/contact",
			LastMod: time.Now(),
		},
	}, nil)
	
	xmlResult, err := sitemapIdx.toXMLString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xmlResult)
}

```