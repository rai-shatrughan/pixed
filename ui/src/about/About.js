import React, { useState } from "react";
import "./style/index.css";
import "purecss/build/pure-min.css"
import "purecss/build/grids-responsive-min.css";
import {
    Header,
    Theme,
    getTheme,
    CompMap
} from "./comp/CompList";

export default function About() {
    const [show, setShow] = useState("Home");
    const [theme, setTheme] = useState(getTheme());
    const VisibleComponent = CompMap[show];

    function headerClicked(e) {
        e.preventDefault();
        setShow(e.target.innerText);
    };

    return (
        <div>
            <div id="div-menu">
                <Header
                    onClicked={headerClicked}
                />

                <Theme
                    theme={theme}
                    onClicked={() =>
                        setTheme(getTheme())
                    }
                />
            </div>

            <div id="div-center">
                <VisibleComponent theme={theme} />
            </div>
        </div>
    );
}

