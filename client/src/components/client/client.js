import React from "react";
import { Link } from "react-router-dom";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import {splitArray} from "../../library/splitArray";
import csses from "./client.module.css";


const Client = () => {
    const numcols = 4;

    const getCells = (line,i) => {
        let item = [];
        let itemkey = 0;
        for (const [j, cell] of line.entries()) {
            console.log(numcols,i,j);
            itemkey = i*numcols+j;
            item.push(<Col key={itemkey} xl lg={4} sm={6} sm xs={12}>{cell.name}</Col>)
        }
        if (line.length<numcols) {
            for (let j=0; j<numcols-line.length; j++) {
                itemkey += 1;
                item.push(<Col key={itemkey} xl lg={4} sm={6} sm xs={12}></Col>)
            }
        }
        return item;
    }

    let lines = [{name: "A"},{name: "B"},{name: "C"},{name: "D"},{name: "E"},{name: "F"},{name: "G"},{name: "H"},{name: "I"},{name: "J"},];
    const slines = splitArray(lines,numcols);
    let items = [];
    for (const [i, line] of slines.entries()) {
        const cells = getCells(line,i);
        items.push(<Row>{cells}</Row>);
    }
    return (
        <Container>
            {items}
        </Container>
    )
    
}

export default Client;