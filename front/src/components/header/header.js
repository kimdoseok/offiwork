import React from "react";
import { Link } from "react-router-dom";
import csses from "./header.module.css";

const Header = () => {
    return (
        <div className={csses.header}>
            <nav>
                <ul>
                    <li className={csses.headerleft}>
                        <Link to="/">Home</Link>
                    </li>
                    <li className={csses.headerleft}>
                        <Link to="/client">Clients</Link>
                    </li>
                    <li className={csses.headerleft}>
                        <Link to="/activity">Activity</Link>
                    </li>
                    <li className={csses.headerleft}>
                        <Link to="/document">Documents</Link>
                    </li>
                    <li className={csses.headerleft}>
                        <Link to="/message">Messages</Link>
                    </li>
                    <li className={csses.headerleft}>
                        <Link to="/about">About</Link>
                    </li>
                    <li className={csses.headerright}>
                        <Link to="/login">Login</Link>
                    </li>
                </ul>
            </nav>
        </div >
    )
}

export default Header;