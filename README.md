# sitemapper
simple site mapping library for go

## Documentation
[![GoDoc](https://godoc.org/github.com/blasphemy/sitemapper?status.svg)](https://godoc.org/github.com/blasphemy/sitemapper)

## Usage

Super simple usage example:

```go
package main

import (
	"fmt"
	"log"

	"github.com/blasphemy/sitemapper"
)

func main() {
	m := sitemapper.NewMapper()
	m.AddURL("http://yoursite.com/whatever")
	m.AddURL("https://yoursite.com/whatever-article-2")
	x, err := m.GenerateXML()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(x))
}
```

Output for above example:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">
   <url>
      <loc>http://yoursite.com/whatever</loc>
   </url>
   <url>
      <loc>https://yoursite.com/whatever-article-2</loc>
   </url>
</urlset>

```
