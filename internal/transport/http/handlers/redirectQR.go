package handlers

import (
	"net/http"
	"runtime"
)

// HRedirectQR ...
func (h *Handler) HRedirectQR(rw http.ResponseWriter, r *http.Request) {
	if runtime.GOOS == "ios" {
		http.Redirect(rw, r, "https://apps.apple.com/ru/app/humo-online/id1242252363", http.StatusSeeOther)
	} else {
		http.Redirect(rw, r, "https://play.google.com/store/apps/details?id=tj.humo.online&pli=1", http.StatusSeeOther)
	}

}
