import React from "react";
import { createRoot } from "react-dom/client";
import "./style/index.css";
import "purecss/build/base-min.css";
import "purecss/build/grids-responsive-min.css";
import {
  Header,
  Theme
} from "./comp/CompList";
import { Menus, CompMap } from './app/constants';

class Index extends React.Component {

  constructor(props) {
    super(props);
    this.headerClicked = this.headerClicked.bind(this);
    this.state = { show: "Home" };
  }

  headerClicked(menu, e) {
    e.preventDefault();
    this.setState({ show: menu });
  }

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

          <Theme />

        </div>

        <VisibleComponent />
      </div>
    );
  }
}

const container = document.getElementById("root");
const root = createRoot(container);
root.render(<Index />);
