function sum(a){
    return function (b){
        return a + b
    }
}

function multi(a){
    return b => a * b 
}

console.log(sum(1)(2))
console.log(multi(3)(2))