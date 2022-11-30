import React from 'react';

export default function Simple(props) {
    return (
        <li className="pure-menu-item">
            <a className="pure-menu-link"
                href="#"
                key={props.name}
                onClick={props.onShow}
            >{props.name}</a>
        </li>
    );
}






