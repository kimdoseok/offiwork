import React, { useState } from "react";

import csses from "./login.module.css";

const Login = () => {

    const [userid, setUserId] = useState("");
    const [passwd, setPasswd] = useState("");

    const clearForm = () => {
        setUserId("");
        setPasswd("");
    }

    const loginForm = () => {
        console.log("loginform:")
  //      setUserCtx(userid, passwd);
    }

    const updateUserId = (event) => {
//        setUserId(event.target.value);
    }

    const updatePasswd = (event) => {
//        setPasswd(event.target.value);
    }

    let body = (
        <div className={csses.login}>
            <div className={csses.loginform} >
                <div>UserID: <input type="text" id="userid" name="userid" onChange={updateUserId} /></div>
                < div > Password: <input type="password" id="passwd" name="passwd" onChange={updatePasswd} /></div >
                <div>
                    <button onClick={() => { loginForm() }}>Login</button>
                    <button onClick={() => { clearForm() }}>Clear</button>
                </div>
                <p className={csses.message}>Please use 882966 / c@sLdGgxI[sE|aJ as a user ID and password pair.</p>
            </div >
        </div >
    )
/*
    console.log("login:", userdataCtx);

    if (userdataCtx !== undefined && userdataCtx.length === 1) {
        body = (
            <Redirect to="/home" />
        )
    }
*/
    return (
        body
    )
}

export default Login;