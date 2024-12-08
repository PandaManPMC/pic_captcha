import React from 'react'
import './App.css';
// import 'antd/dist/antd.min.css';
import "@arco-design/web-react/dist/css/arco.css";
import ClickTextCapt from "./components/ClickTextCapt";
import ClickShapeCapt from "./components/ClickShapeCapt";
import SlideCapt from "./components/SlideCapt";
import SlideRegionCapt from "./components/SlideRegionCapt";
import RotateCapt from "./components/RotateCapt";
import {useGithub} from "./hooks/useGithub";

function App() {
  useGithub()

  return (
    <div className="container">
      <div style={{marginTop: "40px", display: "flex", alignItems: "center", justifyContent: "center"}}><ClickTextCapt/></div>
      <div style={{marginTop: "40px", display: "flex", alignItems: "center", justifyContent: "center"}}><ClickShapeCapt/></div>
      <div style={{marginTop: "40px", display: "flex", alignItems: "center", justifyContent: "center"}}><SlideCapt/></div>
      <div style={{marginTop: "40px", display: "flex", alignItems: "center", justifyContent: "center"}}><SlideRegionCapt/></div>
      <div style={{marginTop: "40px", display: "flex", alignItems: "center", justifyContent: "center"}}><RotateCapt/></div>

      <div className="cap-tip">
        <a className="github-button" href="https://github.com/PandaManPMC/pic_captcha" data-size="large">Github</a>
      </div>
    </div>
  );
}

export default App;
