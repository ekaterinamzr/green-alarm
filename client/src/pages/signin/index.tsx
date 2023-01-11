import { Button, Form, Input, message } from "antd";

import { signInUser } from "services/auth";

import { Layout } from "components/Layout";

interface FormValues {
    username: string;
    user_password: string;
  }

const SignInPage = () => {

const [form] = Form.useForm<FormValues>();

const onFinish = async (values: FormValues) => {
    try {
      console.log(values);
      await signInUser({...values}) 
      message.success("Вход выполнен.");
    } catch(e) {
      if (typeof e === "string") {
        console.log("error1");
        console.log(e.toUpperCase()); // works, `e` narrowed to string
    } else if (e instanceof Error) {
      console.log("error2");
      console.log(e.message); // works, `e` narrowed to Error
    }
      console.log("error");
      message.error("Не удалось войти.");
    }
  };

  return (
    <Layout title="Вход">
      <Form form={form} onFinish={onFinish} preserve={false} layout="vertical">
        <Form.Item label="Ник" name="username" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item label="Пароль" name="user_password" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
      </Form>
      <Button type="primary" onClick={form.submit} key="submit">
          Войти
        </Button>
      
      
    </Layout>
  );
};

export default SignInPage;
