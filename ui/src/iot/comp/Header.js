import "../style/header.css"
import { Menus } from './Constants';

export default function Header({ onClicked }) {
    return (
        <div className="pure-menu pure-menu-horizontal" id="iot-menu">
            <ul className="pure-menu-list">
                {Menus.map((menu) => (
                    < li className="pure-menu-item" key={menu}>
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