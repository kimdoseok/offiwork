import React from "react";
import { Link } from "react-router-dom";
import csses from "./activityheader.module.css";
import Container from 'react-bootstrap/container';

const ActivityHeader = () => {
    return (
        <div className={csses.activityheader}>
            <Container flex>
            <nav>
                <ul>
                    <li className={csses.activityheaderleft}>
                        <Link to="/">Home</Link>
                    </li>
                </ul>
            </nav>
            </Container>
        </div >
    )
}

export default ActivityHeader;