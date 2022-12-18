import {FillApiItemGrowth} from "./fillhtml.js";

export function handleStocks(obj) {
    let stocks = JSON.parse(obj)
    if (stocks != undefined){
        const apistocks = document.getElementsByClassName("api-stocks")
        // clear previous (need rework)
        let i = 0
        while (apistocks[0].children.length > 0){
            apistocks[0].children[i].remove()
        }
        let item = document.createElement('div');
        item.className="api-item-header";
        let stocksText = document.createElement('p', );
        stocksText.innerText = "Stocks Prices:";
        item.appendChild(stocksText)
        apistocks[0].appendChild(item)

        // convert stocks_rates object to map
        let stocksPrices = new Map(Object.entries(stocks.stocks_change))
        stocksPrices.forEach((value, key) => {
            let item = FillApiItemGrowth(value, stocks.prev_period, stocks.last_updated, true)
            apistocks[0].appendChild(item)
        });
    }
}