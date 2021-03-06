package sitemapper

import (
	"bytes"
	"encoding/xml"
)

const (
	xmlns      = "http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlnsxhtml = "http://www.w3.org/1999/xhtml"
)

type url struct {
	Location string
}

//Mapper is the main representation of the sitemapper
type Mapper struct {
	URLs []url
}

type xmlURL struct {
	XMLName  xml.Name `xml:"url"`
	Location xmlLoc   `xml:"loc"`
}

type xmlLoc struct {
	String string `xml:",chardata"`
}

type xmlURLSet struct {
	XMLNS      string   `xml:"xmlns,attr"`
	XMLNSXHTML string   `xml:"xmlns:xhtml,attr"`
	XMLName    xml.Name `xml:"urlset"`
	URLs       []xmlURL `xml:"url"`
}

//NewMapper returns a new "Mapper" instance used to generate your sitemap
func NewMapper() *Mapper {
	m := &Mapper{
		URLs: []url{},
	}
	return m
}

//AddURL adds a new url to the mapper
func (m *Mapper) AddURL(u string) {
	nu := url{
		Location: u,
	}
	m.URLs = append(m.URLs, nu)
}

//GenerateXML returns the generated sitemap xml, and an error if applicable
func (m *Mapper) GenerateXML() ([]byte, error) {
	urls := []xmlURL{}
	for _, x := range m.URLs {
		xu := urlToXMLUrl(x)
		urls = append(urls, xu)
	}
	set := xmlURLSet{
		URLs:       urls,
		XMLNS:      xmlns,
		XMLNSXHTML: xmlnsxhtml,
	}
	b, err := xml.Marshal(set)
	if err != nil {
		return []byte{}, err
	}
	buffer := bytes.Buffer{}
	buffer.WriteString(xml.Header)
	buffer.Write(b)
	return buffer.Bytes(), nil
}

func urlToXMLUrl(in url) xmlURL {
	nu := xmlURL{
		Location: xmlLoc{
			String: in.Location,
		},
	}
	return nu
}
