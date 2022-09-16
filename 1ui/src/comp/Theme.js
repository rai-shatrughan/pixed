import React from "react";
import logoDark from '../assets/theme-dark.gif';
import logoLight from '../assets/theme-light.gif';

export function SetTheme(themeName) {
  localStorage.setItem('theme', themeName);
  document.documentElement.className = themeName;
}

export function GetTheme() {
  return localStorage.getItem('theme');
}

export function ToggleTheme() {
  if (GetTheme() === 'dark') {
    SetTheme('light');
  } else {
    SetTheme('dark');
  }
}

export default class Theme extends React.Component {

  render() {
    return (
      <button className="theme-button" onClick={this.props.onClicked}>
        { this.props.theme === 'dark' ?
          <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoDark} /> :
          <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoLight} />
        }
      </button>
    );
  }
}
