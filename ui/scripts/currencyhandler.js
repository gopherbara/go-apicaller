import {FillApiItemGrowth} from "./fillhtml.js";

export function handleCurrency(obj) {
    let currency = JSON.parse(obj)
    if (currency != undefined){
        const apiCurrency = document.getElementsByClassName("api-currency")
        // clear previous (need rework)
        let i = 0
        while (apiCurrency[0].children.length > 0){
            apiCurrency[0].children[i].remove()
        }
        let item = document.createElement('div');
        item.className="api-item-header";
        let fromText = document.createElement('p', );
        let fromCurrency = document.createElement('p');
        fromText.innerText = "Currency rates: ";
        fromCurrency.innerText = "(in " + currency.from_currency + ")"
        item.appendChild(fromText)
        item.appendChild(fromCurrency)
        apiCurrency[0].appendChild(item)

        // convert currency_rates object to map

        let currencyRates = new Map(Object.entries(currency.currency_rates))
        currencyRates.forEach((value, key) => {
            let item = FillApiItemGrowth(value, currency.prev_period, currency.last_updated, false)
            apiCurrency[0].appendChild(item)
        });
    }
}