import { client } from "services/api";

import { Incident } from "../types/incident";

const getIncidents = async () => {
  const { data } = await client.get<Incident[]>("/incidents");

  return data;
};

const getIncident = async (id: number) => {
  const { data } = await client.get<Incident>(`/incidents/${id}`);

  return data;
};

const createIncident = async (incident: Partial<Incident>) => {
  await client.post("/incidents", incident);
};

const deleteIncident = async (id: number) => {
  await client.delete(`/incidents/${id}`);
};

export { getIncidents, createIncident, getIncident, deleteIncident };
