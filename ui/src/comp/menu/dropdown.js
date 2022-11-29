import React, { useEffect } from 'react';

export default function Dropdown(props) {

    return (
        <li className="pure-menu-item pure-menu-has-children pure-menu-allow-hover">
            <a href="#" className="pure-menu-link">{props.name}</a>
            <ul className="pure-menu-children">
                <li className="pure-menu-item">
                    {props.items.map((item) => (
                        <a className="pure-menu-link"
                            href="#"
                            key={item}
                            onClick={props.onShow}
                        >{item}</a>
                    ))}
                </li>
            </ul>
        </li>

    );
}






