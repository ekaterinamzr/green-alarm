import moment from "moment";

interface Incident {
  id: number;
  incident_name: string;
  incident_date: moment;
  incident_date_printable: string;
  country: string;
  latitude: string;
  longitude: string;
  publication_date: moment;
  comment?: string;
  incident_status: string;
  incident_type: string;
  author: string;
}

export { Incident };
