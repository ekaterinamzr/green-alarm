import { ComponentProps, FC, useEffect, useMemo, useState } from "react";
import { Select } from "antd";

import { getIncidents } from "../../api/incidents";
import { Incident } from "../../types/incident";

type Props = ComponentProps<typeof Select>;

const IncidentSelect: FC<Props> = (props) => {
  const [isLoading, setIsLoading] = useState(false);
  const [incidents, setIncidents] = useState<Incident[]>([]);

  const options = useMemo(() => {
    return incidents.map((incident) => ({
      label: `${incident.incident_name} ${incident.incident_date} дата`,
      value: incident.id,
    }));
  }, [incidents]);

  const fetchIncidents = async () => {
    setIsLoading(true);

    try {
      const newIncidents = await getIncidents();
      setIncidents(newIncidents);
    } catch {
      /* empty */
    }

    setIsLoading(false);
  };

  useEffect(() => {
    void fetchIncidents();
  }, []);

  return (
    <Select
      placeholder="Выберите инцидент"
      options={options}
      disabled={isLoading}
      loading={isLoading}
      {...props}
    />
  );
};

export { IncidentSelect };
