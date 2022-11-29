import React, { useState } from "react";
import { createRoot } from "react-dom/client";
import "./style/index.css";
import "purecss/build/grids-responsive-min.css";
import "purecss/build/pure.css"
import {
    Header,
    Theme,
    getTheme,
    CompMap
} from "./comp/CompList";

function Index() {
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

const container = document.getElementById("root");
const root = createRoot(container);
root.render(<Index />);
