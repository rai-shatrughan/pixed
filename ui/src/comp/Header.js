import React from 'react';

export default function Header(props) {

    return (
        <React.Fragment>
            <a className="pure-menu-heading pure-menu-link"
                href={props.menu}
                key={props.menu}
                onClick={props.onClicked}
            >{props.menu}</a>
        </React.Fragment>
    );
}
