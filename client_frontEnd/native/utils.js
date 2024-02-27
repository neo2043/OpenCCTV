function addChildren(parent,children){
    children.forEach((child)=>{
        parent.appendChild(child)
    })
    return parent
}

function makeElement(element,attributes,ael,childPara){
    if(!element){
        return false
    }
    const tempElement = document.createElement(element)
    if(attributes){
        attributes.split(";").forEach((element) => {
            diff=element.split("=")[0]
            if (diff==="class"){
                element.substring("class=".length,element.length).split(" ").forEach((ele)=>{
                    if(ele!==" "){
                        tempElement.classList.add(ele)
                    }
                })
            }
            if (diff==="id"){
                element.substring("id=".length,element.length).split(" ").forEach((ele)=>{
                    if(ele!==" "){
                        tempElement.id=(ele)
                    }
                })
            }
            if(diff==="style"){
                style=""
                element.substring("style=".length,element.length).split(" ").forEach((ele)=>{
                    if(ele!==" "){
                        style=style+ele+";"
                    }
                })
                tempElement.style=style
            }
        });
    }
    if(ael){
        tempElement.addEventListener(ael.type,ael.func)
    }
    if(childPara){
        tempElement.appendChild(document.createTextNode(childPara));
    }
    return tempElement
}