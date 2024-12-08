package captdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server-example/internal/cache"
	"server-example/internal/helper"

	"github.com/PandaManPMC/pic_captcha/assets/resources/images"
	"github.com/PandaManPMC/pic_captcha/assets/resources/shapes"
	"github.com/PandaManPMC/pic_captcha/captcha/base/option"
	"github.com/PandaManPMC/pic_captcha/captcha/click"
)

var shapeCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 3, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 3}),
		click.WithRangeThumbBgDistort(1),
		click.WithIsThumbNonDeformAbility(true),
	)

	// shape
	// click.WithUseShapeOriginalColor(false) -> Random rewriting of graphic colors
	shapeMaps, err := shapes.GetShapes()
	if err != nil {
		log.Fatalln(err)
	}

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	// set resources
	builder.SetResources(
		click.WithShapes(shapeMaps),
		click.WithBackgrounds(imgs),
	)

	shapeCapt = builder.MakeWithShape()
}

// GetClickShapesCaptData .
func GetClickShapesCaptData(w http.ResponseWriter, r *http.Request) {
	captData, err := shapeCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    1,
			"message": "gen captcha data failed",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	var masterImageBase64, thumbImageBase64 string
	masterImageBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    1,
			"message": "base64 data failed",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	thumbImageBase64, err = captData.GetThumbImage().ToBase64()
	if err != nil {
		bt, _ := json.Marshal(map[string]interface{}{
			"code":    1,
			"message": "base64 data failed",
		})
		_, _ = fmt.Fprintf(w, string(bt))
		return
	}

	dotsByte, _ := json.Marshal(dotData)
	key := helper.StringToMD5(string(dotsByte))
	cache.WriteCache(key, dotsByte)

	bt, _ := json.Marshal(map[string]interface{}{
		"code":         0,
		"captcha_key":  key,
		"image_base64": masterImageBase64,
		"thumb_base64": thumbImageBase64,
	})
	_, _ = fmt.Fprintf(w, string(bt))
}
