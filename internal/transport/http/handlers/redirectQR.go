package handlers

import (
	"net/http"

	"github.com/houseme/mobiledetect/ua"
)

// HRedirectQR ...
func (h *Handler) HRedirectQR(rw http.ResponseWriter, r *http.Request) {
	ua := ua.New(r.UserAgent())

	if ua.Platform() == "iPhone" {
		http.Redirect(rw, r, "https://apps.apple.com/ru/app/humo-online/id1242252363", http.StatusSeeOther)
	} else {
		http.Redirect(rw, r, "https://play.google.com/store/apps/details?id=tj.humo.online&pli=1", http.StatusSeeOther)
	}
}
