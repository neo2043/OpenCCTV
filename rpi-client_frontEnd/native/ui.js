function addPlusSign(){
    rootDiv=makeElement("div","class=plusSign")
    rootDiv.appendChild(makeElement("div","class=vertical"))
    rootDiv.appendChild(makeElement("div","class=horizontal"))
    return rootDiv
}

function wrapper(classes,ael){
    return makeElement("div",classes,ael) 
}

