export function handleQuote(obj){
    let quote = JSON.parse(obj)
    document.getElementById("api-quote-value").innerText = quote.quote
}