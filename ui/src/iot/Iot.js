import React, { useState } from "react"
import "./style/iot.css"
import "purecss/build/pure-min.css"
import "purecss/build/grids-responsive-min.css"
import Login from './comp/Login'
import Theme from './comp/Theme'
import Header from "./comp/Header"
import { MenuMap } from "./comp/Constants"

export default function Iot() {
    const [loggedIn] = useState(false)
    const [activeComp, setActiveComp] = useState("Home")
    const VisibleComponent = MenuMap[activeComp]

    function setVisibleComp(e) {
        e.preventDefault()
        setActiveComp(e.target.innerText)
    }

    if (loggedIn) {
        return (
            <div>
                <Theme />
                <div id="div-menu">
                    <Header onClicked={setVisibleComp} />
                </div>
                <div id="div-center">
                    <VisibleComponent />
                </div>
            </div>
        )
    } else {
        return (
            <div>
                <Theme />
                <Login />
            </div>
        )
    }

}

