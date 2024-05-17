package langdetect

type config struct {
	whitelist                   map[Lang]struct{}
	blacklist                   map[Lang]struct{}
	reliableConfidenceThreshold float64
}

func (c config) isWhitelisted(lang Lang) bool {
	if c.whitelist == nil {
		return true
	}
	_, ok := c.whitelist[lang]
	return ok
}

func (c config) isBlacklisted(lang Lang) bool {
	if c.blacklist == nil {
		return false
	}
	_, ok := c.blacklist[lang]
	return ok
}

func (c config) isReliable(confidence float64) bool {
	return confidence > c.reliableConfidenceThreshold
}

var DefaultConfig = config{
	whitelist:                   nil,
	blacklist:                   nil,
	reliableConfidenceThreshold: RELIABLE_CONFIDENCE_THRESHOLD,
}

type Option func(*config)

func WithWhitelist(langs ...Lang) Option {
	return func(c *config) {
		if c.whitelist == nil {
			c.whitelist = make(map[Lang]struct{})
		}
		for _, lang := range langs {
			c.whitelist[lang] = struct{}{}
		}
	}
}

func WithBlacklist(langs ...Lang) Option {
	return func(c *config) {
		if c.blacklist == nil {
			c.blacklist = make(map[Lang]struct{})
		}
		for _, lang := range langs {
			c.blacklist[lang] = struct{}{}
		}
	}
}

func WithReliableConfidenceThreshold(threshold float64) Option {
	return func(c *config) {
		c.reliableConfidenceThreshold = threshold
	}
}
