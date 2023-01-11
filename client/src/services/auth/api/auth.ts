import { client } from "services/api";

import { User } from "../types/auth";

const signUpUser = async (user: Partial<User>) => {
  await client.post("/auth/sign-up", user);
};

const signInUser = async (user: Partial<User>) => {
    await client.post("/auth/sign-in", user).then(function (response) {
      console.log(response.data)
      localStorage.setItem('jwt', response.data.token);
      localStorage.setItem('role', response.data.role);
    })
    .catch(function (error) {
      console.log(error);
    });
  };

export { signUpUser, signInUser };
