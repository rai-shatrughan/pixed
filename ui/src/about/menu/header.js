import React from 'react';
import Simple from './simple';
import Dropdown from './dropdown';
import { Menus } from "../comp/CompList";

export default function Header({ onClicked }) {

    return (
        <div className="pure-menu pure-menu-horizontal">
            <ul className="pure-menu-list">
                <Simple
                    name="Home"
                    onShow={onClicked} />
                <Dropdown
                    name="About"
                    onShow={onClicked}
                    items={Menus} />

            </ul>
        </div>

    );
}






