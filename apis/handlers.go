package apis

import (
	"net/http"
	"net/url"

	"github.com/aryahadii/readan/simplifier"

	"github.com/sirupsen/logrus"
)

func simplifyGet(w http.ResponseWriter, r *http.Request) {
	websiteRawURL, ok := r.URL.Query()["url"]
	if !ok {
		logrus.Infoln("request does not contain `url` field")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	websiteURL, err := url.Parse(websiteRawURL[0])
	if err != nil {
		logrus.WithError(err).Infoln("requested url is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	simplifiedDocument, err := simplifier.GetMercury().SimplifyHTML(websiteURL)
	if err != nil {
		logrus.WithError(err).Errorln("cannot simplify")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(simplifiedDocument))
}

func simplifyPost(w http.ResponseWriter, r *http.Request) {
}
