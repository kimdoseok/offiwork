import React from "react";
import { Link } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import csses from "./document.module.css";


const Document = () => {

    let logoutStr = "Not logged in";

    /*
    console.log("document:", userData);
    try {
        logoutStr = userData[0].User_name + " [Logout]";
    } catch (err) {
    }
*/

const body = (
    <Container>
        <Row>
            <Col>
                Documents
            </Col>
        </Row>
    </Container>
)

    return (
        body
    )
}

export default Document;