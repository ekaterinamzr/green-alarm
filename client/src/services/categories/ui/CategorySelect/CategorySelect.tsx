import { ComponentProps, FC, useEffect, useMemo, useState } from "react";
import { Select } from "antd";

import { getCategories } from "../../api/categories";
import { Category } from "../../types/category";

type Props = ComponentProps<typeof Select>;

const CategorySelect: FC<Props> = (props) => {
  const [isLoading, setIsLoading] = useState(false);
  const [categories, setCategories] = useState<Category[]>([]);

  const options = useMemo(() => {
    return categories.map((category) => ({
      label: category.type_name,
      value: category.id,
    }));
  }, [categories]);

  const fetchCategories = async () => {
    setIsLoading(true);

    try {
      const newCategories = await getCategories();
      setCategories(newCategories);
    } catch {
      /* empty */
    }

    setIsLoading(false);
  };

  useEffect(() => {
    void fetchCategories();
  }, []);

  return (
    <Select
      placeholder="Выберите категорию"
      options={options}
      disabled={isLoading}
      loading={isLoading}
      {...props}
    />
  );
};

export { CategorySelect };
