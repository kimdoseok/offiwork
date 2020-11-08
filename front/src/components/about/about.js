import React from "react";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import csses from "./about.module.css";

const About = () => {
    const body = (
        <Container>
            <Row>
                <Col>
                    <div className={csses.about}>
                        <h2>About</h2>
                        <p>This site consists of 2 parts.  One is Frontend made with REACT.  It make it possible to build a single page application.  The other is backend that was built with Golang and provides grpc connection to frontend and the other apps. There are not enogh data to build schema but tried to build with sample data.  This will be populated with real data.</p>
                        <h3>Client</h3>
                        <p>This page has select customer and the selected client is stored in session.  This client will be used until a user changes to another client</p>
                        <h3>Activity</h3>
                        <p>All the transactions are managed by this module.  </p>
                        <h3>Documents</h3>
                        <p>All the files are managed by this module. </p>
                        <h3>Messages</h3>
                        <p>A general purpose messaging module</p>
                    </div>
                </Col>
            </Row>
        </Container>

    )    

    return (
        body
    )
}

export default About;