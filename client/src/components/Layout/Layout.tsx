import { ComponentProps, FC, ReactNode } from "react";
import { Layout as AntLayout, Space, Typography } from "antd";

import { Header } from "components/Header";
import { Body } from "components/Body";

import styles from "./Layout.module.css";

interface Props extends Omit<ComponentProps<typeof AntLayout>, "title"> {
  title?: ReactNode;
}

const Layout: FC<Props> = ({ title, children }) => {
  return (
    <AntLayout className={styles.layout}>
      <Header />
      <Body className={styles.body}>
        <Space direction="vertical" size="middle" className={styles.space}>
          {title && <Typography.Title>{title}</Typography.Title>}
          {children}
        </Space>
      </Body>
    </AntLayout>
  );
};

export { Layout };
