import {FillApiItemGrowth} from "./fillhtml.js";

export function handleWeather(obj) {
    let weatherObj = JSON.parse(obj)
    if (weatherObj != undefined){
        const apiWeather = document.getElementsByClassName("api-weather")
        // clear previous (need rework)
        let i = 0
        while (apiWeather[0].children.length > 0){
            apiWeather[0].children[i].remove()
        }
        

        let item = document.createElement('div');
        item.className="api-item-header";
        let todayText = document.createElement('p', );
        let cityText = document.createElement('p');
        todayText.innerText = "Today`s weather:";
        cityText.innerText = '(' + weatherObj.city + ')'
        item.appendChild(todayText)
        item.appendChild(cityText )
        apiWeather[0].appendChild(item)

        let weatherDifference = new Array(Object.values(weatherObj.weather_difference))

        weatherDifference[0].forEach((value, index) => {
            let item = FillApiItemGrowth(value, weatherObj.prev_period, weatherObj.last_updated, false)
            // let item = document.createElement('div');
            // item.className="api-item"
            // let objChange = document.createElement('p');
            // let change = document.createElement('p');
            // objChange.innerText = value.object_change
            // change.innerText = value.today_value + " (" + value.growth + ")"
            // item.appendChild(objChange)
            // item.appendChild(change)
            apiWeather[0].appendChild(item)
        });

    }
}