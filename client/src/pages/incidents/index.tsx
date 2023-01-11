import { Button, message } from "antd";
import { useEffect, useState } from "react";

import { Incident, getIncidents, AddIncidentModal, IncidentsTable, deleteIncident } from "services/incidents";

import { Layout } from "components/Layout";

const IncidentsPage = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

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

  const onDelete = async (id: number) => {
    setIsLoading(true);

    try {
      await deleteIncident(id);
      message.success("Инцидент успешно удалён.");

      await fetchIncidents();
    } catch (error) {
      message.error("Не удалось удалить инцидент.");
      console.error(error);
    }

    setIsLoading(false);
  };

  useEffect(() => {
    void fetchIncidents();
  }, []);


  if (localStorage.getItem("jwt") === null) { 
    return (
      <Layout title="Инциденты">
        
        <IncidentsTable loading={isLoading} data={incidents} onDelete={onDelete} />
        <AddIncidentModal open={isModalOpen} onCancel={() => setIsModalOpen(false)} />
      </Layout>
    );
  } 

  return (
    <Layout title="Инциденты">
      <Button type="primary" onClick={() => setIsModalOpen(true)}>
        Добавить
      </Button>
      
      <IncidentsTable loading={isLoading} data={incidents} onDelete={onDelete} />
      <AddIncidentModal open={isModalOpen} onCancel={() => setIsModalOpen(false)} />
    </Layout>
  );
};

export default IncidentsPage;
