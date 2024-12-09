import React from 'react'
import pmc from 'pmc-react'
// Cache Testing
// import GoCaptcha from '../cache'
import {useClickHandler} from "../hooks/useClickHandler";

function ClickShapeCapt() {

  const handler = useClickHandler({
    getApi: "/api/go-captcha-data/click-shape",
    checkApi: "/api/go-captcha-check-data/click-shape"
  })

  return (
    <pmc.CaptchaClick
      config={{
        width: 300,
        height: 220,
        dotSize: 24,
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

export default ClickShapeCapt;
