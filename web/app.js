const apiUrl = "http://localhost:8080/api/"

function getValue() {
    var select = document.getElementById("select-type");
    var value = select.value;
    // console.log(value);
    return value;
}

async function fetchIncidents(incidentType) {
    try {
        if (incidentType == 0) {
            var url = "incidents"
        } else {
            var url = `incidents?type=${incidentType}`
        }

        const response = await fetch(`${apiUrl}${url}`);

        if (!response.ok) {
            throw new Error(`Failed to fetch incidents: ${response.status}`)
        }

        return await response.json();
    } catch (e) {
        console.log(e);
    }
}

function incidentElement(incident) {
    const anchorElement = document.createElement('a');
    const url = "incidents/"
    anchorElement.setAttribute("href", `${apiUrl}${url}${incident.id}`)
    anchorElement.setAttribute("target", "_blank");
    anchorElement.innerText = incident.incident_name;

    const incidentTitleElement = document.createElement("h3");
    incidentTitleElement.appendChild(anchorElement);

    return incidentTitleElement;
}

function listIncidents(incidentContainerElementId) {
    incidentType = getValue()
    const incidentContainerElement = document.getElementById(incidentContainerElementId);
    if (!incidentContainerElement) {
        return;
    }
    fetchIncidents(incidentType).then(incidents => {
        if (!incidents) {
            incidentContainerElement.innerHTML = "Инциденты не найдены";
            return;
        }
        console.log(incidents)
        for (const incident of incidents) {
            incidentContainerElement.appendChild(incidentElement(incident));
        }
    }).catch(e => {
        console.log(e);
    })
}

// Функция ymaps.ready() будет вызвана, когда
// загрузятся все компоненты API, а также когда будет готово DOM-дерево.

function initMap() {
    ymaps.ready(init);
}

function init() {
    // Создание карты.
    var myMap = new ymaps.Map("map", {
            // Координаты центра карты.
            // Порядок по умолчанию: «широта, долгота».
            // Чтобы не определять координаты центра карты вручную,
            // воспользуйтесь инструментом Определение координат.
            center: [55.76, 37.64],
            // Уровень масштабирования. Допустимые значения:
            // от 0 (весь мир) до 19.
            zoom: 2,
            controls: ['smallMapDefaultSet']
        }, {
            searchControlProvider: 'yandex#search'
        }

    );

    fetchIncidents(incidentType).then(incidents => {
        if (!incidents) {
            return;
        }
        for (const incident of incidents) {
            myMap.geoObjects
                .add(new ymaps.Placemark([incident.latitude, incident.longitude], {
                    balloonContent: incident.incident_name
                }, {
                    preset: 'islands#icon',
                    iconColor: '#0095b6'
                }))
        }
    }).catch(e => {
        console.log(e);
    })
}