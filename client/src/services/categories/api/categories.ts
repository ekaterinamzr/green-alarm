import { client } from "services/api";

import { Category } from "../types/category";

const getCategories = async () => {
  const { data } = await client.get<Category[]>("/types");

  return data;
};

export { getCategories };
