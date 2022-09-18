import React, { useState } from "react";
import { createRoot } from "react-dom/client";
import "./style/index.css";
import "purecss/build/base-min.css";
import "purecss/build/grids-responsive-min.css";
import { 
    Header, 
    Theme, 
    getTheme,
    Menus,
    CompMap 
} from "./comp/CompList";

function Index() {
    const [show, setShow] = useState("Home");
    const [theme, setTheme] = useState(getTheme());
    const VisibleComponent = CompMap[show];

    function headerClicked(menu, e){
        e.preventDefault();
        setShow(menu);
    };
        
        return (
            <div>
                <div className="pure-menu pure-menu-horizontal">
                    {Menus.map((menu) => (
                        <Header
                            menu={menu}
                            key={menu}
                            onClicked={(e) => headerClicked(menu, e)}
                        />
                    ))}

                    <Theme
                        theme={theme}
                        onClicked={() => 
                            setTheme(getTheme())
                        }
                    />
                </div>

                <VisibleComponent theme={theme} />
            </div>
        );
    }

const container = document.getElementById("root");
const root = createRoot(container);
root.render(<Index />);
