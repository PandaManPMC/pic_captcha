import React from 'react'
import pmc from 'pmc-react'

// Cache Testing
// import GoCaptcha from '../cache'
import {useRotateHandler} from "../hooks/useRotateHandler";

function RotateCapt() {
  const handler = useRotateHandler({
    getApi: "/api/go-captcha-data/rotate-basic",
    checkApi: "/api/go-captcha-check-data/rotate-basic"
  })

  return (
    <pmc.CaptchaRotate
      config={{
        width: 300,
        height: 220,
        size: 220,
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

export default RotateCapt;
