import React from "react";
import logoDark from '../assets/theme-dark.gif';
import logoLight from '../assets/theme-light.gif';

export function setTheme(themeName) {
  localStorage.setItem('theme', themeName);
  document.documentElement.className = themeName;
}

export function getTheme() {
  return localStorage.getItem('theme');
}

export default class Theme extends React.Component {

  constructor(props) {
    super(props);
    this.toggleTheme = this.toggleTheme.bind(this);
    this.state = { theme: getTheme() };
  }

  componentDidMount() {
    getTheme() === 'light' ? setTheme('light') : setTheme('dark');
  }

  toggleTheme() {
    if (getTheme() === 'dark') {
      setTheme('light');
      this.setState({ theme: getTheme() });
    } else {
      setTheme('dark');
      this.setState({ theme: getTheme() });
    }
  }

  render() {
    return (
      <button className="theme-button" onClick={this.toggleTheme}>
        {this.state.theme === 'dark' ?
          <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoDark} /> :
          <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" src={logoLight} />
        }
      </button>
    );
  }
}
