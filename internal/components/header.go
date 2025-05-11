package components

import . "github.com/delaneyj/gostar/elements"

func Header(title string) ElementRenderer {
	return HEADER(
		TITLE().TextF("%s", title),
		DIV().CLASS("header").TextF("Hello %s", title))
}
