import React from "react";
import "../style/login.css";

export default function Login({ loggedIn }) {
    return (
        <form className="pure-form pure-form-stacked" id="login-form">
            <fieldset>
                <legend id="login-legend">Sign In</legend>
                <input id="email" type="email" placeholder="Email" />
                <input id="password" type="password" placeholder="Password" />
                <label htmlFor="default-remember">
                    <input type="checkbox" id="default-remember" /> Remember me
                </label>
                <button type="submit" className="pure-button pure-button-primary">Login</button>
            </fieldset>
        </form >
    );
}

