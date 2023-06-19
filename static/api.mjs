axios.interceptors.request.use(function (config) {
  const token = getToken();
  config.headers.Authorization = `${token}`;
  return config;
});

export const getToken = () => {
  return localStorage.getItem("token");
};

const API = {
  Book: {
    getBooks: (config) => axios.get("/api/v1/book/list", config),

    getBookdetails: (bookId) =>
      axios.get("/api/v1/book/get", {
        params: {
          id: bookId,
        },
      }),
  },

  Borrow: {
    borrowBook: (bookId, days) =>
      axios.post("/api/v1/lending/book", {
        bookId,
        days: days,
      }),
  },

  Return: {
    getBorrowwedbooks: (config) =>
      axios.get("/api/v1/lending/listByReader", config),

    returnBook: (config) => axios.post("/api/v1/return/book", config),
  },

  Reservation: {
    saveReservation: (bookId, days) =>
      axios.post("/api/v1/reservation/save", {
        bookId,
        retain: days,
      }),
    getReservations: (config) =>
      axios.get("/api/v1/reservation/reader/list", config),

    cancelReserve: (reservedid) =>
      axios.post("/api/v1/reservation/cancel", {
        id: reservedid,
      }),
  },

  User: {
    changePassword: (oldpassword, newpassword) =>
      axios.post("/api/v1/reader/password", {
        newPassword: newpassword,
        oldPassword: oldpassword,
      }),
  },
};

export default API;
