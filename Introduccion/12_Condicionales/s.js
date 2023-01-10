isPalindrome = (texto) => {
    let textoF = texto.toUpperCase();
    return textoF.split('').reverse().join('') == texto.toUpperCase()
}

console.log(isPalindrome('Kayak'));