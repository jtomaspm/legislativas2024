// ==UserScript==
// @name         Legislativas 2024
// @author       PopAndBoom
// @include      https://www.legislativas2024.mai.gov.pt/resultados/territorio-nacional?local=*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=travian.com
// @grant        none
// @version      2024
// @description  none
// ==/UserScript==

function main(){
    const selDistrito = document.getElementsByClassName("form-control minimal")[0];
    let res = {};
    addOptions(res, selDistrito);

    let distritoIter = iterator(selDistrito);
    loopDistritos(distritoIter, res);
}

function loopDistritos(iter, res) {
    if(!iter.hasNext()) return;
    const distritoOpt = iter.getNext();
    delay(1000).then(()=>{
        const selConcelho = document.getElementsByClassName("form-control minimal")[1];
        let concelhoIter = iterator(selConcelho);
        delay(1000).then(()=>{loopConcelhos(concelhoIter, res)})
            .then(()=>delay(40000).then(()=>loopDistritos(iter, res)));

    });
}

function loopConcelhos(iter, res) {
    console.log(res);
    if(!iter.hasNext()) return;
    const concelhoOpt = iter.getNext();
    delay(1000).then(()=>{
        res[concelhoOpt.value] = concelhoOpt.innerText.trim();
        const selFreguesia = document.getElementsByClassName("form-control minimal")[2];
        addOptions(res, selFreguesia);
        delay(1000).then(()=>loopConcelhos(iter, res));
    });
}

function iterator(sel){
    class Iterator {
        constructor(sele){
            this.selector = sele;
            this.current = this.selector.options[this.selector.selectedIndex];
        }
        getNext(){
            this.selector.selectedIndex = this.selector.selectedIndex + 1;
            this.current = this.selector.options[this.selector.selectedIndex];
            this.selector.dispatchEvent(new Event('change'));
            return this.current;
        }
        hasNext(){
            return this.selector.selectedIndex < this.selector.options.length + 1;
        }
    }
    return new Iterator(sel);
}

function addOptions(res, sel){
    let temp = Array.from(sel.children).map(e => [e.innerText, e.value]);
    for(let i = 1; i < temp.length; i++){
        res[temp[i][1]] = temp[i][0].trim();
    }
}

function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}


//RUN in localidades
delay(2000).then(()=> {
    document.getElementsByClassName("nav-link w-100 text-center py-0")[1].click();
    delay(1000).then(()=>main());
});