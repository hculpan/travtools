package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hculpan/travtools/cmd/web/templates"
)

func SpecTradeSellHandler(w http.ResponseWriter, r *http.Request) {
	step := r.URL.Query().Get("step")
	if step == "" {
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
		return
	}

	stepNum, err := strconv.Atoi(step)
	if err != nil {
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
		return
	}

	switch stepNum {
	case 2:
		component := templates.DetermineSalePrice()
		component.Render(context.Background(), w)
	default: // also case 1
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
	}
}
