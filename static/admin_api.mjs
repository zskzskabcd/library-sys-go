axios.interceptors.request.use(function(config) {
    const token = getToken();
    config.headers.Authorization = `${token}`;
    return config;
});

export const getToken = () => {
    return localStorage.getItem('token')
}

const API = {
    Book: {
        getBooks: config => axios.get('/api/v1/book/list', config),

    },

    Search: {
        book: (BookName) => axios.get('/api/v1/book/search', {
            params: {
                keyword: BookName,
                start: 1,
                count: 10
            }
        }),

        readerlist: config => axios.get('/api/v1/reader/list', config),

        reservelist: config => axios.get('/api/v1/reservation/list', config),

        borrowlist: config => axios.get('/api/v1/lending/list',config)
    },

    Put: {
        book: config => axios.post('/api/v1/book', config),

        reader: config => axios.post('/api/v1/reader',config)
    },

    Delete: {
        book: (bookid) => axios.delete('/api/v1/book', {
            params: {
                id: bookid
            }
        }),

        reader: (readerid) => axios.delete('/api/v1/reader', {
            params: {
                id: readerid
            }
        })
    },

    Admin: {
        changePassword: (oldpassword, newpassword) => axios.post('/api/v1/reader/password', {
            newPassword: newpassword,
            oldPassword: oldpassword
        })
    }

}

export default API