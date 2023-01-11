import { message } from "antd";
import { useEffect, useState } from "react";

import { Map1 } from 'components/Map/';

import { Incident, getIncidents, AddIncidentModal, IncidentsTable, deleteIncident } from "services/incidents";

import { Layout } from "components/Layout";

const IncidentsPage = () => {

  const [isLoading, setIsLoading] = useState(false);
  const [incidents, setIncidents] = useState<Incident[]>([]);

  const fetchIncidents = async () => {
    setIsLoading(true);

    try {
      const newIncidents = await getIncidents();
      setIncidents(newIncidents);
    } catch (error) {
      message.error("Не удалось загрузить список инцидентов.");
      console.error(error);
    }

    setIsLoading(false);
  };


  useEffect(() => {
    void fetchIncidents();
  }, []);


  return (
    <Layout title="Инциденты">

      <Map1 data={incidents}/>
      
    </Layout>
  );
};

export default IncidentsPage;
