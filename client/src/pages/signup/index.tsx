import { Button, Form, Input, message } from "antd";

import { signUpUser } from "services/auth";

import { Layout } from "components/Layout";

interface FormValues {
    first_name: string;
    last_name: string;
    username: string;
    email: string;
    user_password: string;
  }

const SignUpPage = () => {

const [form] = Form.useForm<FormValues>();

const onFinish = async (values: FormValues) => {
    try {
      console.log(values);
      await signUpUser({...values}) 
      message.success("Регистрация прошла успешно.");
    } catch(e) {
      if (typeof e === "string") {
        console.log("error1");
        console.log(e.toUpperCase()); // works, `e` narrowed to string
    } else if (e instanceof Error) {
      console.log("error2");
      console.log(e.message); // works, `e` narrowed to Error
    }
      console.log("error");
      message.error("Не удалось зарегистрироваться.");
    }

  };

  return (
    <Layout title="Регистрация">
      <Form form={form} onFinish={onFinish} preserve={false} layout="vertical">
        <Form.Item label="Имя" name="first_name" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="Фамилия" name="last_name" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="Ник" name="username" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="Почта" name="email" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="Пароль" name="user_password" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
      </Form>
      <Button type="primary" onClick={form.submit} key="submit">
          Зарегистрироваться
        </Button>
      
    </Layout>
  );
};

export default SignUpPage;
