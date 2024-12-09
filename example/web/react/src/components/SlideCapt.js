import React from 'react'
import pmc from 'pmc-react'

// Cache Testing
// import GoCaptcha from '../cache'
import {useSlideHandler} from "../hooks/useSlideHandler";

function SlideCapt() {
  const handler = useSlideHandler({
    getApi: "/api/go-captcha-data/slide-basic",
    checkApi: "/api/go-captcha-check-data/slide-basic"
  })

  return (
    <pmc.CaptchaSlide
      config={{
        width: 300,
        height: 220,
      }}
      data={handler.data}
      events={{
        close: handler.closeEvent,
        refresh: handler.refreshEvent,
        confirm: handler.confirmEvent,
      }}
    />
  );
}

export default SlideCapt;
