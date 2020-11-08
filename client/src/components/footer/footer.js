import React from "react";
import { Link } from "react-router-dom";
import csses from "./footer.module.css";


const Footer = () => {

    let logoutStr = "Not logged in";

/*
    console.log("footer:", userData);
    try {
        logoutStr = userData[0].User_name + " [Logout]";
    } catch (err) {
    }
*/
    return (
        <div className={csses.footer}>
            <nav>
                <ul>
                    <li className={csses.footerleft}>
                        <Link to="/">Datetime </Link>
                    </li>
                    <li className={csses.footerleft}>
                        <Link to="/">User name</Link>
                    </li>
                    <li className={csses.footerleft}>
                        <Link to="/">Client name</Link>
                    </li>
                    <li className={csses.footerleft}>
                        <Link to="/">Working minutes</Link>
                    </li>
                    <li className={csses.footerright}>
                        <Link to="/logout">{logoutStr}</Link>
                    </li>
                </ul>
            </nav>

        </div>
    )
}

export default Footer;