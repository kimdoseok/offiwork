const splitArray = (inarr, numcols) => {
    let outarr = [];
    let rowarr = [];
    for (let i=0;i<inarr.length;i++) {
        if (i%numcols==0 && i>0) {
            outarr.push(rowarr);
            rowarr = [] ;     
        } 
        rowarr.push(inarr[i]);
    }
    if (rowarr.length>0) {
        outarr.push(rowarr);
    }
    return outarr;
}

export {splitArray}