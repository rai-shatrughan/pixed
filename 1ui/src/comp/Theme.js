import React from "react";
import logoDark from '../assets/theme-dark.gif';
import logoLight from '../assets/theme-light.gif';

function setTheme(themeName) {
  localStorage.setItem('theme', themeName);
  document.documentElement.className = themeName;
}

export default class Theme extends React.Component {

  constructor(props) {
    super(props);
    this.toggleTheme = this.toggleTheme.bind(this);
    this.state = { dark: false };
  }

  componentDidMount() {
    (localStorage.getItem('theme') === 'theme-light') ? setTheme('theme-light') : setTheme('theme-dark');
  }

  toggleTheme() {
    if (localStorage.getItem('theme') === 'theme-dark') {
        setTheme('theme-light');
        localStorage.setItem('theme', 'theme-light');
        this.setState({dark: false});
    } else {
        setTheme('theme-dark');
        localStorage.setItem('theme', 'theme-dark');
        this.setState({dark: true});
    }
  }

  render() {
    return (        
      <button className="pure-menu-heading pure-menu-link theme-button" onClick={this.toggleTheme}>
        { this.state.dark ?
          <img className="theme-icon" title="ToggleTheme" alt="ToggleTheme" width="10" height="10" src={logoDark} /> :
          <img className="theme-icon" title="ToggleTheme" alt= "ToggleTheme" width="10" height="10" src={logoLight} /> 
        }
      </button> 
    );
  }
}
