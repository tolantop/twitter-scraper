package twitterscraper

// GetTrends return list of trends.
func (s *Scraper) GetTrends() ([]string, error) {
	req, err := s.newRequest("GET", "https://twitter.com/i/api/2/guide.json")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("candidate_source", "trends")
	q.Add("include_page_configuration", "false")
	q.Add("entity_tokens", "false")
	req.URL.RawQuery = q.Encode()

	var jsn timeline
	err = s.RequestAPI(req, &jsn)
	if err != nil {
		return nil, err
	}

	var trends []string
	for _, item := range jsn.Timeline.Instructions[1].AddEntries.Entries[1].Content.TimelineModule.Items {
		trends = append(trends, item.Item.ClientEventInfo.Details.GuideDetails.TransparentGuideDetails.TrendMetadata.TrendName)
	}

	return trends, nil
}

// GetTrends wrapper for default Scraper
func GetTrends() ([]string, error) {
	return defaultScraper.GetTrends()
}
