import { ComponentProps, FC } from "react";
import { useNavigate } from "react-router-dom";
import { Menu as AntMenu, MenuProps } from "antd";
import { ItemType } from "antd/es/menu/hooks/useItems";

type Props = ComponentProps<typeof AntMenu>;

let ITEMS: ItemType[]

const Menu: FC<Props> = (props) => {
  if (localStorage.getItem("jwt") === null) {
    ITEMS = [
      { key: "incidents", label: "Список" },
      { key: "map", label: "Карта" },
      { key: "signin", label: "Вход" },
      { key: "signup", label: "Регистрация" },
    ];
  } else {
    ITEMS = [
      { key: "incidents", label: "Список" },
      { key: "map", label: "Карта" },
      { key: "signout", label: "Выход" },
    ];
  }

  const navigate = useNavigate();

  const onClick: MenuProps["onClick"] = (e) => {
    switch (e.key) {
      case "incidents":
        navigate("/incidents/");
        break;
      case "map":
        navigate("/map/");
        break;
      case "signin":
        navigate("/signin/");
        break;
      case "signup":
        navigate("/signup/");
        break;
      case "signout":
        localStorage.removeItem('jwt');
        localStorage.removeItem('role');
        break;
    }
  };

  return (
    <AntMenu
      mode="horizontal"
      items={ITEMS}
      selectedKeys={[]}
      onClick={onClick}
      {...props}
    />
  );
};

export { Menu };
