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
  "log"

  "github.com/carlosstrand/go-sitemap"
)

func main() {
  sitemap := NewSitemap([]*SitemapItem{
    {
      Loc: "https://my-site.com/about",
      LastMod: time.Now(),
      ChangeFreq: "daily",
      Priority: 0.8,
    },
    {
      Loc: "https://my-site.com/contact",
      LastMod: time.Now(),
      ChangeFreq: "monthly",
      Priority: 0.2,
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
  "log"

  "github.com/carlosstrand/go-sitemap"
)

func main() {
  sitemapIdx := NewSitemapIndex([]*SitemapIndexItem{
    {
      Loc: "https://my-site.com/sitemap_1.xml",
      LastMod: time.Now(),
    },
    {
      Loc: "https://my-site.com/sitemap_2.xml",
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
