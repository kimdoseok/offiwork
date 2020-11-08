import React from "react";
import { Link } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import csses from "./message.module.css";


const Message = () => {

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
            <Row>
                <Col>
                Message
                </Col>
            </Row>
        </Container>
    )

    return (
        body
    )
}

export default Message;