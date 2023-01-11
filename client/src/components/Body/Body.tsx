import { Layout as AntLayout } from "antd";
import { ComponentProps, FC } from "react";

type Props = ComponentProps<typeof AntLayout.Content>;

const Body: FC<Props> = ({ children, ...props }) => {
  return <AntLayout.Content {...props}>{children}</AntLayout.Content>;
};

export { Body };
