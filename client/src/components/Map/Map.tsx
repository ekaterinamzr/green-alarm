import React, {useRef} from 'react';
import { YMaps, Map, Placemark } from "@pbe/react-yandex-maps";

import { Incident } from "../../services/incidents/types/incident";

interface Props {
    data: Incident[];
}

const Map1: React.FC<Props> = ({ data, ...props }) => {
    const mapRef = useRef();

const defaultState = {
    center: [55.751574, 37.573856],
    zoom: 2,
  };

  return (
    <YMaps query={{ lang: 'en_RU', apikey: '7f8e597a-1f6d-44d3-a9dc-99eab25ff130' }}>
      <Map  width="100%" height="600px" defaultState={defaultState} instanceRef={mapRef}>
      
      {data.map((item, key) => (
        <Placemark
        modules={["geoObject.addon.balloon", "geoObject.addon.hint"]}
        key={key}
        geometry={{
            type: "Point",
            coordinates: [ item.latitude, item.longitude ],
          }}
        properties={{
            // Проставляем цифры иконкам на карте
            // iconContent: key+1,
            hintContent: item.incident_name,
            balloonContent: item.comment,
          }}
        />
        ))
    } 
      
      </Map>
    </YMaps>
  );
};

export {Map1};