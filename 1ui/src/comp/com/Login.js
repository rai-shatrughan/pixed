import React from 'react';

export default function Login(props) {

    return (
        <form className="pure-form">
            <fieldset>
                <legend>Login to explore</legend>
                <input type="email" placeholder="Email"></input>
                <input type="password" placeholder="Password"></input>
                <button type="submit" className="pure-button pure-button-primary">Sign In</button>
            </fieldset>
        </form>
    );
}