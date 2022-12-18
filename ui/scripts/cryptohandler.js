import {FillApiItemGrowth} from "./fillhtml.js";

export function handleCrypto(obj) {
    let crypto = JSON.parse(obj)
    if (crypto != undefined){
        const apicrypto = document.getElementsByClassName("api-crypto")
        // clear previous (need rework)
        let i = 0
        while (apicrypto[0].children.length > 0){
            apicrypto[0].children[i].remove()
        }
        let item = document.createElement('div');
        item.className="api-item-header";
        let cryptoText = document.createElement('p', );
        cryptoText.innerText = "Crypto Prices:";
        item.appendChild(cryptoText)
        apicrypto[0].appendChild(item)

        // convert crypto_rates object to map
        let cryptoPrices = new Map(Object.entries(crypto.crypto_change))
        cryptoPrices.forEach((value, key) => {
            let item = FillApiItemGrowth(value, crypto.prev_period, crypto.last_updated, true)
            apicrypto[0].appendChild(item)
        });
    }
}