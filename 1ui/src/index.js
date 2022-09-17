import React from "react";
import { createRoot } from "react-dom/client";
import "./style/index.css";
import "purecss/build/base-min.css";
import "purecss/build/grids-responsive-min.css";
import { Header, Theme } from "./comp/CompList";
import { Menus, CompMap } from "./app/constants";
import { getTheme, toggleTheme } from "./comp/Theme";

class Index extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            show: "Home",
            theme: getTheme(),
        };
    }

    headerClicked = (menu, e) => {
        e.preventDefault();
        this.setState({ show: menu });
    };

    render() {
        const VisibleComponent = CompMap[this.state.show];

        return (
            <div>
                <div className="pure-menu pure-menu-horizontal">
                    {Menus.map((menu) => (
                        <Header
                            menu={menu}
                            key={menu}
                            onClicked={(e) => this.headerClicked(menu, e)}
                        />
                    ))}

                    <Theme
                        theme={this.state.theme}
                        onClicked={() => {
                            toggleTheme();
                            this.setState({ theme: getTheme() });
                        }}
                    />
                </div>

                <VisibleComponent theme={this.state.theme} />
            </div>
        );
    }
}

const container = document.getElementById("root");
const root = createRoot(container);
root.render(<Index />);
