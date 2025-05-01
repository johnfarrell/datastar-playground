package pages

import (
	. "github.com/delaneyj/gostar/elements"
	"github.com/johnfarrell/datastar-playground/internal/components"
)

func Index(title string) ElementRenderer {
	return Group(
		Text("<!DOCTYPE html>"),
		HTML(
			BODY(
				components.Header(title),
			),
		),
	)
}
