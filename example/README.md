<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="https://raw.githubusercontent.com/PandaManPMC/public_src/ddfa2ac7f8200914de759d18391ff303e3d9ab55/pmcad/logo.svg"/>
<h1 style="margin: 0; padding: 0">Go Captcha</h1>
</div>

<h3 style="text-align: center ">服务端例子</h3>

<br/>

>
<p style="text-align: center">
Go Captcha 是行为式验证码，支持文本/图形点选、滑动/拖拽、旋转等验证模式。
</p>



<br/>

- GoCaptcha：[https://github.com/wenlng/go-captcha](https://github.com/wenlng/go-captcha)
- Javascript 原生库：[https://github.com/wenlng/go-captcha-jslib](https://github.com/wenlng/go-captcha-jslib)
- Vue Package：[https://github.com/wenlng/go-captcha-vue](https://github.com/wenlng/go-captcha-vue)
- React Package：[https://github.com/wenlng/go-captcha-react](https://github.com/wenlng/go-captcha-react)
- Angular Package：[https://github.com/wenlng/go-captcha-angular](https://github.com/wenlng/go-captcha-angular)
- Svelte Package：[https://github.com/wenlng/go-captcha-svelte](https://github.com/wenlng/go-captcha-svelte)
- Solid Package：[https://github.com/wenlng/go-captcha-solid](https://github.com/wenlng/go-captcha-solid)
- ...

<br/>

浏览器打开: http://127.0.0.1:9001/example

<br/>

### 目录结构
```text
|-internal
  |-app
  |-cache
  |-logic
    |-captdata
      |-click_basic.go      // 生成文本点选验证码数据
      |-click_shape.go      // 生成图形点选验证码数据
      |-rotate_basic.go     // 生成旋转验证码数据
      |-slide_basic.go      // 生成滑动验证码数据
      |-slide_region.go     // 生成拖拽验证码数据
    |-checkdata
      |-click.go            // 点选验证码校验
      |-rotate.go           // 旋转验证码校验
      |-slide.go            // 滑动/拖拽验证码校验
|-web
  |-native                // Javascript 原生实例
  |-bower                 // Javascript 原生实例，bower工具管理
  |-vue                   // Vue3 实例
  |-vue2                  // Vue2 实例
  |-react                 // React 实例
  |-angular               // Angular 实例
  |-solid                 // Solid 实例
  |-svelte                // Svelte 实例
```

<br/>

