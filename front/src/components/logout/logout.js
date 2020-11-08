import React from "react";
import { Redirect } from "react-router-dom";

const Logout = () => {

    /*
    let setUserCtx = useContext(SetUserContext)
    setUserCtx("", "");
    */
    return (<Redirect to="/login" />)
}

export default Logout;