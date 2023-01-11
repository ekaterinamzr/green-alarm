import { Button, DatePicker, Form, Input, message, Modal } from "antd";
import { ComponentProps, FC } from "react";
import { Moment } from "moment";

import { CategorySelect } from "services/categories";

import { createIncident } from "../../api/incidents";

type Props = ComponentProps<typeof Modal>;

interface FormValues {
  incident_name: string;
  incident_type: string;
  incident_date: Moment;
  country: string;
  latitude: string;
  longitude: string;
  comment: string;
  colorId: string;
}

const AddIncidentModal: FC<Props> = ({ onCancel, ...props }) => {
  const [form] = Form.useForm<FormValues>();

  const onFinish = async (values: FormValues) => {
    try {
      console.log(values);
      await createIncident({...values, incident_status: '1', incident_type: values.incident_type.toString(), /*latitude: '12', longitude: '12', country: 'feff', comment: 'jefje'*/})
      message.success("Инцидент успешно добавлен.");
    } catch {
      message.error("Не удалось добавить инцидент.");
    }
  };

  return (
    <Modal
      title="Добавление инцидента"
      footer={[
        <Button type="primary" onClick={form.submit} key="submit">
          Добавить
        </Button>,
        <Button onClick={onCancel} key="cancel">
          Отменить
        </Button>,
      ]}
      onCancel={onCancel}
      destroyOnClose
      {...props}
    >
      <Form form={form} onFinish={onFinish} preserve={false} layout="vertical">
        <Form.Item
          label="Тип"
          name="incident_type"
          rules={[{ required: true }]}
        >
          <CategorySelect />
        </Form.Item>
        <Form.Item label="Название" name="incident_name" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item
          label="Дата"
          name="incident_date"
          rules={[{ required: true }]}
        >
          <DatePicker style={{ width: "100%" }} />
        </Form.Item>
        <Form.Item
          label="Страна"
          name="country"
          rules={[{ required: true }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Ширина"
          name="latitude"
          rules={[{ required: true }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Долгота"
          name="longitude"
          rules={[{ required: true }]}
        >
          <Input />
        </Form.Item>
        <Form.Item label="Описание" name="comment">
          <Input.TextArea />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export { AddIncidentModal };
