import {handleCurrency} from "./currencyhandler.js";
import {handleQuote} from "./quotehandler.js";
import {handleStocks} from "./stockshandler.js";
import {handleCrypto} from "./cryptohandler.js";
import {handleWeather} from "./weatherhandler.js";
import {handleGeo} from "./geolocationhandler.js";

let socket = new WebSocket("ws://localhost:4000/ws")
console.log("try connect websocket")

socket.onopen = () => {
    console.log("successfully connected");
    socket.send("Hi from client")
}

socket.onclose = (event) => {
    console.log("socket close connection: ", event)
}

socket.onmessage = (msg) => {
    let apiInfo = JSON.parse(msg.data)
    if (apiInfo.crypto != undefined){
        handleCrypto(apiInfo.crypto)
    }
    if (apiInfo.currency != undefined){
        handleCurrency(apiInfo.currency)
    }
    if (apiInfo.geo != undefined){
        handleGeo(apiInfo.geo)
    }
    if (apiInfo.quote != undefined){
        handleQuote(apiInfo.quote)
    }
    if (apiInfo.stocks != undefined){
        handleStocks(apiInfo.stocks)
    }
    if (apiInfo.weather != undefined){
        handleWeather(apiInfo.weather)
    }
    console.log(msg)
}

socket.onerror = (error) => {
    console.log("socket error: ", error);
}
