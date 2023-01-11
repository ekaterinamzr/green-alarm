import axios from "axios";

const client = axios.create({
  baseURL: process.env.REACT_APP_API_HOST,
  headers: {Authorization: `Bearer ${localStorage.getItem('jwt')}`},
});

// const mock = new MockAdapter(client, { delayResponse: 2000 });

// mock.onGet("/incidents/").reply(200, [
//   {
//     id: 1,
//     name: "Инцидент 1",
//     date: "21-01-2002",
//     country: "Россия",
//     latitude: 100,
//     longitude: 100.2,
//     publication: "22-12-2022",
//     comment: "Описание 1",
//     status: "1",
//     type: 1,
//     author: 1,
//   },
//   {
//     id: 2,
//     name: "Инцидент 2",
//     date: "12-11-2001",
//     country: "Россия",
//     latitude: 100,
//     longitude: 100.2,
//     publication: "22-12-2022",
//     comment: "Описание 1",
//     status: "1",
//     type: 1,
//     author: 1,
//   },
// ]);

// mock.onGet("/incidents/1/").reply(200, {
//   id: 1,
//   name: "Инцидент 1",
//   date: "21-01-2002",
//   country: "Россия",
//   latitude: 100,
//   longitude: 100.2,
//   publication: "22-12-2022",
//   comment: "Описание 1",
//   status: "1",
//   type: 1,
//   author: 1,
// });

// mock.onGet("/incidents/2/").reply(200, {
//   id: 2,
//   name: "Инцидент 2",
//   date: "12-11-2001",
//   country: "Россия",
//   latitude: 100,
//   longitude: 100.2,
//   publication: "22-12-2022",
//   comment: "Описание 1",
//   status: "1",
//   type: 1,
//   author: 1,
// });

// mock.onDelete(/\/incidents\/\d+/).reply(200);

// mock.onGet("/statuses/").reply(200, [
//   { id: 1, name: "Подтвержден" },
//   { id: 2, name: "Не подтвержден" },
// ]);

// mock.onGet("/types/").reply(200, [
//   { id: 1, name: "Разлив нефти или нефтепродуктов" },
//   { id: 2, name: "Выброс радиоактивных веществ" },
//   { id: 3, name: "Выброс аварийно химически опасных веществ" },
//   { id: 4, name: "Выброс биологически опасных веществ" },
//   { id: 5, name: "Пожар" },
//   { id: 6, name: "Несанкционированная свалка, скопление мусора" },
//   { id: 7, name: "Другие экологические инциденты" },
// ]);

// mock.onGet("/offers/").reply(200, [
//   {
//     incidentId: 0,
//     price: 1000,
//     ownerId: 0,
//     availability: true,
//   },
//   {
//     incidentId: 1,
//     price: 2000,
//     ownerId: 0,
//     availability: false,
//   },
// ]);

export { client };
