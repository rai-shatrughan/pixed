import React from 'react';
import "../style/header.css"

const menus = [
    "Home",
    "Asset",
    "Agent",
    "Metrics",
    "Insights",
    "User",
    "Setting"
];

export default function Header({ onClicked }) {

    return (
        <div className="pure-menu pure-menu-horizontal" id="iot-menu">
            <ul className="pure-menu-list">
                {menus.map((menu) => (
                    < li className="pure-menu-item">
                        <a className="pure-menu-link"
                            href="#"
                            key={menu}
                            onClick={onClicked}
                        >{menu}</a>
                    </li>
                ))}
            </ul>
        </div >

    );
}