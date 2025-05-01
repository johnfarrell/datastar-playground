package components

import . "github.com/delaneyj/gostar/elements"

func Header(title string) ElementRenderer {
	return HEADER(
		TITLE().TextF("%", title),
		DIV().CLASS("header").TextF("Hello %s", title))
}
