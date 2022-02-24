/*IR BUSCAR POR ID A IMAGEM*/
let see =document.getElementById('image-see');
//SETAR UM EVENTO PARA ESSA IMAGEM
see.onclick=function(){
    //ir buscar o input
let x=document.getElementById("inputpass")
    //se a classe for see passar para not see e passar o campo para tipo texto senao passar para pass e por a imagem see
    if (see.className=="see") {
        see.className="not-see";
        x.type="text";
        see.setAttribute("src","assets/notsee.png");
    }else{
    
        see.className="see";
        x.type="password"
        see.setAttribute("src","assets/see.png")

    }

}
/*IR BUSCAR O LINK POR ID*/
let link=document.getElementById('aqui')

/*funcao sobre este link*/
link.onclick= async function() {
    /*ir buscar os containers */
    let direito=document.getElementById('direito')
    let esquerdo=document.getElementById('esquerdo')
    let imagem=document.getElementById('imagem')
    let change=document.getElementById('change')
    /*tirar a animacao do aparecimento dos labels e afins */
    document.getElementById('h1esq').style.transition='none'
    document.getElementById('h2esq').style.transition='none'
    document.getElementById('p1esq').style.transition='none'
    document.getElementById('p2esq').style.transition='none'
    document.getElementById('p3esq').style.transition='none'
    /*tirar animacao do container */
    link.style.transition='none'
    /*esconder o esquerdo e a imagem*/
    esquerdo.style.visibility="hidden"
    imagem.style.visibility="hidden"
    /*mover o direito e dar 1 segundo de transicao*/
    direito.style.position="relative"
    direito.style.transition="all 1s"
    /*mover para a direita*/
    direito.style.right="15%"
    /*fazer stretch para a direita*/
    direito.style.marginRight="50%"
    /*tornar o display do esquerdo como none para nao ocupar espaco */
    esquerdo.style.display="none"
    /*fazer um sleep ate a animacao da parte um acabar*/
    await new Promise(r => setTimeout(r, 1000));
    /*tornar o direito static dnv*/
    direito.style.position="static"
    /*tirar a margin*/
    direito.style.marginRight="0%"
    /*coloca-lo na posicao desejada*/
    direito.style.gridColumnStart="1"
    direito.style.gridColumnEnd="2"
    direito.style.gridRowStart="1"
    direito.style.gridRowEnd="3"
    /*fazer novamente um sleep de 1s*/
    await new Promise(r => setTimeout(r, 1000));
    /*retirar transicoes e margins*/
    direito.style.transition="none"
    direito.style.marginRight="none"
    /*establecer a imagem nova */
    imagem.setAttribute("src","assets/registo-livros.png")
    /*por a imagem como visivel*/
    imagem.style.visibility="visible"
    /*colocar o contentor novo */
    change.style.gridColumnStart="2"
    change.style.gridColumnEnd="3"
    change.style.gridRowStart="1"
    change.style.gridRowEnd="3"
    /*coloca-lo com o display correto */
    change.style.display="flex"
    /*torna-lo visivel */
    change.style.visibility="visible"
    

}

let eyeone=document.getElementById('image-seeone');
//SETAR UM EVENTO PARA ESSA IMAGEM
eyeone.onclick=function(){
    //ir buscar o input
let x=document.getElementById("inputpasschone")
    //se a classe for see passar para not see e passar o campo para tipo texto senao passar para pass e por a imagem see
    if (eyeone.className=="see") {
        eyeone.className="not-see";
        x.type="text";
        eyeone.setAttribute("src","assets/notsee.png");
    }else{
    
        eyeone.className="see";
        x.type="password"
        eyeone.setAttribute("src","assets/see.png")

    }
}

let eyetwo=document.getElementById('image-seedois');
//SETAR UM EVENTO PARA ESSA IMAGEM
eyetwo.onclick=function(){
    //ir buscar o input
let x=document.getElementById("inputpasschdois")
    //se a classe for see passar para not see e passar o campo para tipo texto senao passar para pass e por a imagem see
    if (eyetwo.className=="see") {
        eyetwo.className="not-see";
        x.type="text";
        eyetwo.setAttribute("src","assets/notsee.png");
    }else{
    
        eyetwo.className="see";
        x.type="password"
        eyetwo.setAttribute("src","assets/see.png")

    }
}

/*IR BUSCAR O LINK POR ID*/
let linkdois=document.getElementById('aquich')

/*funcao sobre este link*/
linkdois.onclick= async function() {
    /*ir buscar os containers */
    let direito=document.getElementById('direito')
    let esquerdo=document.getElementById('esquerdo')
    let imagem=document.getElementById('imagem')
    let change=document.getElementById('change')
    /*acabar com a transicao*/
    link.style.transition='none'
    /*colocar contentor change invisivel e a imagem tambem */
    change.style.visibility="hidden"
    imagem.style.visibility="hidden"
    /*por posicao relativa e 1s de transicao como antras */
    direito.style.position="relative"
    direito.style.transition="all 1s"
    /* mover para a esquerda*/
    direito.style.left="15%"
    /*dar stretch para a esquerda */
    direito.style.marginLeft="50%"
    /*retirar o tipo de display para nao bugar */
    change.style.display="none"
    /*sleep durante 1s */
    await new Promise(r => setTimeout(r, 1000));
    /*por o direito static novamente */
    direito.style.position="static"
    /*retirar a margin */
    direito.style.marginLeft="0%"
    /*coloca-lo no lugar desejado */
    direito.style.gridColumnStart="2"
    direito.style.gridColumnEnd="3"
    direito.style.gridRowStart="1"
    direito.style.gridRowEnd="3"
    /*1 segundo de sleep */
    await new Promise(r => setTimeout(r, 1000));
    /*retirar transicao e margem do direito */
    direito.style.transition="none"
    direito.style.marginRight="none"
    /*establecer uma imagem nova */
    imagem.setAttribute("src","assets/Login-pc.png")
    /*colocar imagem visivel */
    imagem.style.visibility="visible"
    /*establecer o novo container na posicao desejada */
    esquerdo.style.gridColumnStart="1"
    esquerdo.style.gridColumnEnd="2"
    esquerdo.style.gridRowStart="1"
    esquerdo.style.gridRowEnd="3"
    /*colocar visivel e empacotado*/
    esquerdo.style.display="flex"
    esquerdo.style.visibility="visible"
    

}