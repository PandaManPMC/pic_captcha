import pmc from 'pmc-react'

// Cache Testing
// import GoCaptcha from '../cache'
import {useClickHandler} from "../hooks/useClickHandler";

function ClickTextCapt() {
  const handler = useClickHandler({
    getApi: "/api/go-captcha-data/click-basic",
    checkApi: "/api/go-captcha-check-data/click-basic"
  })

    return (
        <pmc.CaptchaClick
            data={handler.data}
            config={{
              width: 300,
              height: 220,
              dotSize: 24,
            }}
            events={{
              close: handler.closeEvent,
              refresh: handler.refreshEvent,
              confirm: handler.confirmEvent,
            }}
        />
    );
}

export default ClickTextCapt;
