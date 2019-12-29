package sitemap

import (
	"encoding/xml"
	"time"
)

const Header = `<?xml version="1.0" encoding="UTF-8" ?>` + "\n"

func NewOptions() *Options {
	return &Options{
		prettyOutput:  false,
		withXMLHeader: false,
		validate: true,
	}
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
	out := string(result)
	if options != nil && options.withXMLHeader {
		out = Header + out
	}
	return out, nil
}