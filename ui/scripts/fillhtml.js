export function FillApiItemGrowth(value, prevPeriod, lastUpdated, addDollar) {
    let item = document.createElement('div');
    item.className="api-item"
    let objName = document.createElement('p');
    let objGrowth = document.createElement('p');
    let dollar = ``;
    if (addDollar){
        dollar = `&#36;`;
    }
    let title = `Previos check date: ${prevPeriod} &#10;Last updated: ${lastUpdated}&#10;Previous value: ${value.previous_value}${dollar}`
    objName.innerHTML = `${value.object_change} <span title='${title}' class='in-circle'>&#63;</span>`;

    if (value.growth >= 0){
        objGrowth.innerHTML = `${value.today_value}${dollar} <span title="previous value: ${value.previous_value}${dollar}" class="growth-up">&#x2B9D;${value.growth}% </span>`;
    } else {
        objGrowth.innerHTML = `${value.today_value}${dollar} <span title="previous value: ${value.previous_value}${dollar}" class="growth-down">&#x2B9F;${value.growth}% </span>`;
    }
    //objGrowth.innerHTML = `${value.today_value} <span title="previous value: ${value.previous_value}" class="growth-up">&#x2B9D;${value.growth}% </span>`;
    item.appendChild(objName);
    item.appendChild(objGrowth);
    return item;
}