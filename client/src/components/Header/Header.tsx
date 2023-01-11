import { Layout as AntLayout } from "antd";

import { Menu } from "components/Menu";

import styles from "./Header.module.css";

const Header = () => {
  return (
    <AntLayout.Header className={styles.header}>
      <Menu className={styles.menu} />
    </AntLayout.Header>
  );
};

export { Header };
