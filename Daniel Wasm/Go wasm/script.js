const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then((result) => {go.run(result.instance);})

    function display(){
        document.getElementById("displaySection").innerHTML = 
            goDisplay(document.getElementById("userWord").value.toUpperCase());
    }

    function CaR(id){
        goCaR(id);
    }

    function reset(){
        const myElement = document.getElementById("displaySection");
        for (const child of myElement.children) {
            goReset(child);
          }
    }
