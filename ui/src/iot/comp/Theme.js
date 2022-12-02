import { React, useState } from "react";
import { SunIcon, MoonIcon } from "@primer/octicons-react";
import '../style/theme.css'

export default function Theme() {
    const [theme, setTheme] = useState("dark");

    function toggleThemeState() {
        if (theme === "dark") {
            setTheme("light")
            document.documentElement.className = "light";
        } else {
            setTheme("dark")
            document.documentElement.className = "dark";
        }
    }

    return (
        <button className="theme-button" onClick={toggleThemeState}>
            {theme === "dark" ?
                <SunIcon size={16} fill="#f00" /> :
                <MoonIcon size={16} fill="#f00" />
            }
        </button>
    );
}
