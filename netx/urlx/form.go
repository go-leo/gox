package urlx

import "net/url"

func FormFromMap(m map[string]string) url.Values {
	if m == nil {
		return nil
	}
	form := url.Values{}
	for key, value := range m {
		form.Add(key, value)
	}
	return form
}
