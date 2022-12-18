export function handleGeo(obj){
    let geo = JSON.parse(obj)
    if (geo != undefined){
        let apiGeo = document.getElementsByClassName('api-geolocation')
        let i = 0
        while (apiGeo[0].children.length > 0){
            apiGeo[0].children[i].remove()
        }
        let item = document.createElement('div');
        let textItem = document.createElement('p');
        item.className="api-item-header";
        textItem.innerText='Your geo:'
        item.appendChild(textItem)

        let city = document.createElement('div');
        let cityText = document.createElement('p');
        city.className="api-item";
        cityText.innerText= "City: " + geo.city;
        city.appendChild(cityText)

        let country = document.createElement('div');
        let countryText = document.createElement('p');
        country.className="api-item";
        countryText.innerText= "Country: " + geo.country;
        country.appendChild(countryText)

        let lat = document.createElement('div');
        let latText = document.createElement('p');
        lat.className="api-item";
        latText.innerText= "Lat: " + geo.lat;
        lat.appendChild(latText)

        let lon = document.createElement('div');
        let lonText = document.createElement('p');
        lon.className="api-item";
        lonText.innerText="Lon: " + geo.lon;
        lon.appendChild(lonText)

        apiGeo[0].appendChild(item)
        apiGeo[0].appendChild(city)
        apiGeo[0].appendChild(country)
        apiGeo[0].appendChild(lat)
        apiGeo[0].appendChild(lon)
    }
}