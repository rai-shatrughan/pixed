import React, { useState } from "react";
import "./style/iot.css";
import "purecss/build/pure-min.css"
import "purecss/build/grids-responsive-min.css";
import Login from './comp/Login'
import Home from './comp/Home'

export default function Iot() {
    const [loggedIn, setLoggedIn] = useState(true);



    if (loggedIn) {
        return (<Home />)
    } else {
        return (<Login />)
    }

}

