import API from './api.mjs'
import * as utils from './utils.mjs'
import * as render from './render.mjs'

function searchBorrowwedBooks() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 1);

    API.Return.getBorrowwedbooks({
        params: {
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        localStorage.setItem("total", total);
        render.renderBooks(res)
        render.renderPagination(1)
    }).catch(err => {
        console.log(err);
    })
}

function search() {
    const searchedBook = document.getElementById("book").value;

    API.Book.getBooks({
        params: {
            keyword: searchedBook,
            page: 1,
            size: 10
        }
    }).then(res => {
        render.renderList(res.data, 1, 10);
    }).catch(err => {
        console.log(err);
    });
}

function searchAll() {

    localStorage.setItem("page", 1);
    localStorage.setItem("item", 2);
    API.Book.getBooks({
        params: {
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        localStorage.setItem("total", total);
        render.renderList(res.data, 1, 10);
        render.renderPagination(1);
    });
}

function borrow(elem) {
    const bookId = elem.parentNode.dataset.bookid;

    API.Borrow.borrowBook(parseInt(bookId), 30)
        .then(res => {
            if (res.data.code === 200) {
                utils.onSuccessBorrow(elem);
            } else {
                utils.alreadyBorrowed(elem);
            }
        })
        .catch(err => {
            console.log(err);
        });
}

function returnBook(btn) {
    const bookId = btn.parentNode.dataset.bookid;
    console.log(bookId)
    API.Return.returnBook({
            BookId: parseInt(bookId)
        })
        .then(res => {
            if (res.data.code == 200) {
                alert("还书成功！")
                btn.parentNode.remove();
            }
            else{
                alert(res.data.msg)
            }
        })
        .catch(err => {
            console.log(err);
        });
}

function reserveBook(elem) {
    const bookId = elem.parentNode.dataset.bookid;

    API.Reservation.saveReservation(parseInt(bookId), 7)
        .then(res => {
            if (res.data.code === 200) {
                utils.onSuccessReserve(elem);
            } else {
                utils.alreadyReserve(elem);
            }
        })
        .catch(err => {
            console.error("预约失败!", err);
        });
}

function loadReserveList() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 3);

    API.Reservation.getReservations({
        params: {
            status: 1,
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        localStorage.setItem("total", total);
        render.renderReservations(res.data, 1, 10);
        render.renderPagination(1);
    });
}

function cancelReserve(elem) {
    const ReserveId = elem.parentNode.dataset.reservationid;

    API.Reservation.cancelReserve(parseInt(ReserveId))
        .then(res => {
            // 取消预约成功
            elem.parentNode.remove();
            alert('预约已取消');
        })
        .catch(err => {
            alert('取消预约失败');
        })
}

function changePassword() {
    const newPassword = document.getElementById("new").value;
    const oldPassword = document.getElementById("old").value;
    const checkPassword = document.getElementById("check").value;
    console.log(oldPassword, newPassword, checkPassword);
    console.log(oldPassword)
    if (!oldPassword && !newPassword && !checkPassword) {
        alert("请完善信息！");
        return;
    }
    if (newPassword != checkPassword) {
        alert("两次输入的密码不一致!");
        return;
    }
    API.User.changePassword(oldPassword, newPassword)
        .then(res => {
            if (res.data.code === 200) {
                alert("修改密码成功!")

            } else {
                alert(res.data.msg)
            }
        })
        .catch(err => {
            alert("修改密码失败,请重试!")
        })
}

function changePage(type) {
    var page = localStorage.getItem('page');
    const item = localStorage.getItem('item')

    const totalPages = Math.ceil(localStorage.getItem("total") / 10);

    if (type === 'prev') {
        page = page - 1;
    } else {
        page = page + 1;
    }

    if (page < 1) page = 1;
    if (page > totalPages) page = totalPages;
    console.log(page)
    localStorage.setItem("page", page);

    if (item == 2) {
        API.Book.getBooks({
            params: {
                page: page,
                size: 10
            }
        }).then(res => {
            render.renderList(res.data);
            render.renderPagination(page);
        });
    }

    if (item == 3) {
        API.Reservation.getReservations({
            params: {
                status: 1,
                page: page,
                size: 10
            }
        }).then(res => {
            render.renderReservations(res.data);
            render.renderPagination(page);
        });
    }

}





export { search, searchAll, borrow, returnBook, reserveBook, loadReserveList, searchBorrowwedBooks, cancelReserve, changePassword, changePage }