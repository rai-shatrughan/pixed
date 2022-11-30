import { useEffect } from "react";
import { SunIcon, MoonIcon } from "@primer/octicons-react";

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
                <SunIcon size={16} fill="#f00" /> :
                <MoonIcon size={16} fill="#f00" />
            }
        </button>
    );
}
