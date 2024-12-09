import React from 'react'
import pmc from 'pmc-react'

// Cache Testing
// import GoCaptcha from '../cache'
import {useSlideHandler} from "../hooks/useSlideHandler";

function SlideRegionCapt() {
    const handler = useSlideHandler({
            getApi: "/api/go-captcha-data/slide-region",
            checkApi: "/api/go-captcha-check-data/slide-region"
        })

    return (
        <pmc.CaptchaSlideRegion
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

export default SlideRegionCapt;
