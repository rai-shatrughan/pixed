import { useEffect } from "react";
import logoDark from '../../assets/theme-dark.gif';
import logoLight from '../../assets/theme-light.gif';

export function setTheme(themeName) {
    localStorage.setItem('theme', themeName);
    document.documentElement.className = themeName;
}

export function getTheme() {
    return localStorage.getItem('theme');
}

export function toggleTheme() {
    if (getTheme() === 'dark') {
        setTheme('light');
    } else {
        setTheme('dark');
    }
}

export default function Theme(props) {

    useEffect(() => {
        setTheme(props.theme);
    });

    return (
        <button
            className="theme-button"
            onClick={() => {
                toggleTheme();
                props.onClicked();
            }}>

            {props.theme === 'dark' ?
                <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoDark} /> :
                <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoLight} />
            }
        </button>
    );
}
