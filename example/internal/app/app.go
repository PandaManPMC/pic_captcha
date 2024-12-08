package app

import (
	"log"
	"net/http"
	"path"
	"server-example/internal/cache"
	"server-example/internal/helper"
	"server-example/internal/logic/captdata"
	"server-example/internal/logic/checkdata"
)

func Start() {

	// example: get captcha data
	http.Handle("/api/go-captcha-data/click-basic", CORS(http.HandlerFunc(captdata.GetClickBasicCaptData)))
	http.Handle("/api/go-captcha-data/click-shape", CORS(http.HandlerFunc(captdata.GetClickShapesCaptData)))
	http.Handle("/api/go-captcha-data/slide-basic", CORS(http.HandlerFunc(captdata.GetSlideBasicCaptData)))
	http.Handle("/api/go-captcha-data/slide-region", CORS(http.HandlerFunc(captdata.GetSlideRegionCaptData)))
	http.Handle("/api/go-captcha-data/rotate-basic", CORS(http.HandlerFunc(captdata.GetRotateBasicCaptData)))

	// example: post check data
	http.Handle("/api/go-captcha-check-data/click-basic", CORS(http.HandlerFunc(checkdata.CheckClickData)))
	http.Handle("/api/go-captcha-check-data/click-shape", CORS(http.HandlerFunc(checkdata.CheckClickData)))
	http.Handle("/api/go-captcha-check-data/slide-basic", CORS(http.HandlerFunc(checkdata.CheckSlideData)))
	http.Handle("/api/go-captcha-check-data/slide-region", CORS(http.HandlerFunc(checkdata.CheckSlideData)))
	http.Handle("/api/go-captcha-check-data/rotate-basic", CORS(http.HandlerFunc(checkdata.CheckRotateData)))

	// example: js+html+css
	viewStatic := path.Join(helper.GetPWD(), "/web/native/")
	viewFsh := http.FileServer(http.Dir(viewStatic))
	http.Handle("/example/", http.StripPrefix("/example/", viewFsh))

	cache.RunTimedTask()

	log.Println(">>>> ListenAndServe 0.0.0.0:9001")
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
