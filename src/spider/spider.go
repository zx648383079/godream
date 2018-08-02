package main

import (
	"net/url"
)

type Spider struct {
	base_url url
}

func (s *Spider) download(url string, c chan int)  {
	
}

func (s *Spider) findUrl(html string)  {
	
}

func (s *Spider) addUrl(url string)  {
	
}

func startNew(url string)  {
	spider := Spider{url};
	
}