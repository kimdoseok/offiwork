import React from "react";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import csses from "./home.module.css";

const Home = () => {
    let body = (
        <Container>

        <Row className={csses.home}>
            <Col className={csses.homeleftcol} xs lg="4" sm="12">
                List of Clients
            </Col>
            <Col xs lg="8" sm="12">
                Task List
            </Col>
        </Row>
        <Row>
            <Col>
            <img src={process.env.PUBLIC_URL + '/logo192.png'} alt="logo" />;
            </Col>
        </Row>
        </Container>
    )
/*
    const userdataCtx = useContext(UserContext);
    console.log("home:", userdataCtx, userdataCtx.length);
    if (userdataCtx === undefined || userdataCtx.length !== 1) {
        body = (<Redirect to="/login" />)
    }
*/
    return (
        body
    )
}

export default Home;