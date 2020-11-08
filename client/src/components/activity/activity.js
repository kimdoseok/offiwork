import React from "react";
import { Link } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import ActivityHeader from "../activityheader/activityheader";
import csses from "./activity.module.css";


const Activity = () => {

    let logoutStr = "Not logged in";

    /*
    console.log("client:", userData);
    try {
        logoutStr = userData[0].User_name + " [Logout]";
    } catch (err) {
    }
*/
    const body = (
        <Container>
            <ActivityHeader/>
            <Row>
                <Col>
                    Activity
                </Col>
            </Row>
        </Container>
    );

    return (
        body
    )
}

export default Activity;