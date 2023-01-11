import { ComponentProps, FC } from "react";
import { Button, Table } from "antd";

import { Incident } from "../../types/incident";

interface Props extends Omit<ComponentProps<typeof Table<Incident>>, "columns"> {
  data: Incident[];
  onDelete?: (arg0: number) => void;
}

const IncidentsTable: FC<Props> = ({ data, onDelete, ...props }) => {
  data.map(el => {
    const date = /*el.incident_date.format("DD-MM-YYYY")*/ new Date(el.incident_date)
    el.incident_date_printable = date.getDate()+"/"+ (+date.getMonth() + 1) +"/"+date.getFullYear()
    // el.incident_date_printable = date
})

if (localStorage.getItem("jwt") === null || localStorage.getItem("role") == "3") { 
  return (
    <Table
      dataSource={data}
      rowKey="id"
      pagination={false}
      scroll={{ x: true }}
      {...props}
    >
      <Table.Column title="Название" dataIndex="incident_name" />
      <Table.Column title="Описание" dataIndex="comment" />
      <Table.Column title="Страна" dataIndex="country" />
      <Table.Column title="Дата" dataIndex="incident_date_printable"/>
      
    </Table>
  );
} 

  return (

    <Table
      dataSource={data}
      rowKey="id"
      pagination={false}
      scroll={{ x: true }}
      {...props}
    >
      <Table.Column title="Название" dataIndex="incident_name" />
      <Table.Column title="Описание" dataIndex="comment" />
      <Table.Column title="Страна" dataIndex="country" />
      <Table.Column title="Дата" dataIndex="incident_date_printable"/>
      
      <Table.Column
        title="Действия"
        render={(_, record: Incident) => (
          <Button type="primary" danger onClick={() => onDelete?.(record.id)}>
            Удалить
          </Button>
        )}
      />
    </Table>
  );
};

export { IncidentsTable };
