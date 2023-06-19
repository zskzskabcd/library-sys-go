import API from './admin_api.mjs'
import * as render from './admin_render.mjs'

function getNewbookInfo() {
    const Bookname = document.getElementById("Bookname").value;

    API.Search.book(Bookname)
        .then(res => {
            const data = res.data.data[0];
            console.log(data)
            const nameInput = document.getElementById('authorname');
            nameInput.value = data.author;

            const coverInput = document.getElementById('coverUrl');
            coverInput.value = data.cover;

            const priceInput = document.getElementById('price');
            priceInput.value = data.price;

            const publisherInput = document.getElementById('publisher');
            publisherInput.value = data.publisher;

            const publishDateInput = document.getElementById('publishDate');
            publishDateInput.value = data.publishDate;

            const introductionInput = document.getElementById('introduction');
            introductionInput.value = data.summary;

            const isbnInput = document.getElementById('isbn');
            isbnInput.value = data.isbn;
        })
        .catch(err => {
            console.log(err);
        });
}

function searchAll() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 1);

    API.Book.getBooks({
        params: {
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        localStorage.setItem("total", total);
        render.renderList(res.data);
        render.renderPagination(1);

    });
}

function changePage(type) {
    var page = parseInt(localStorage.getItem('page'))
    const item = localStorage.getItem('item')
    const totalPages = Math.ceil(localStorage.getItem("total") / 10);
    if (type === 'prev') {
        page = page - 1;
    } else {
        page = page + 1;
    }
    if (page < 1) page = 1;
    if (page > totalPages) page = totalPages;
    localStorage.setItem("page", page);

    if (item == 1) {
        API.Book.getBooks({
            params: {
                page: page,
                size: 10
            }
        }).then(res => {
            render.renderList(res.data);
            render.renderPagination(page);
        })
    }

    if (item == 2) {
        API.Search.readerlist({
            params: {
                keyword: null,
                page: page,
                size: 10
            }
        }).then(res => {
            render.renderReaderlist(res.data);
            render.renderPagination(page);
        });
    }

    if (item == 3) {
        API.Search.reservelist({
            params: {
                page: page,
                size: 10
            }
        }).then(res => {
            if (res.data.count == 0) {
                alert("无人正在预约！")
            }
            render.renderReservations(res.data);
            render.renderPagination(page);
        });
    }

    if (item == 4) {
        API.Search.borrowlist({
            params: {
                page: page,
                size: 10
            }
        }).then(res => {
            const total = res.data.count;
            if (total == 0) {
                alert("无借阅记录！")
                return
            }
            render.renderBorrwolist(res.data);
            render.renderPagination(page);
        })
    }
}


function searchReaderlist() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 2);

    API.Search.readerlist({
        params: {
            keyword: null,
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        localStorage.setItem("total", total);
        render.renderReaderlist(res.data, 1, 10);
        render.renderPagination(1);
    });

}

function searchrReservelist() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 3);

    API.Search.reservelist({
        params: {
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        if (res.data.count == 0) {

            alert("无人正在预约！")
            return
        }
        localStorage.setItem("total", parseInt(total));
        render.renderReservations(res.data);
        render.renderPagination(1);
    });
}

function searchBorrowlist() {
    localStorage.setItem("page", 1);
    localStorage.setItem("item", 4);

    API.Search.borrowlist({
        params: {
            page: 1,
            size: 10
        }
    }).then(res => {
        const total = res.data.count;
        if (total == 0) {
            alert("无借阅记录！")
            return
        }
        localStorage.setItem("total", total);
        render.renderBorrwolist(res.data);
        render.renderPagination(1);
    })

}

function putNewbook() {
    const Bookname = document.getElementById("Bookname").value;
    const nameInput = document.getElementById('authorname').value;
    const coverInput = document.getElementById('coverUrl').value;
    const priceInput = document.getElementById('price').value;
    const publisherInput = document.getElementById('publisher').value;
    const publishDateInput = document.getElementById('publishDate').value;
    const introductionInput = document.getElementById('introduction').value;
    const isbnInput = document.getElementById('isbn').value;
    const stockInput = document.getElementById('storage').value;

    API.Put.book({
            author: nameInput,
            cover: coverInput,
            isbn: isbnInput,
            price: parseInt(priceInput),
            publishDate: publishDateInput,
            publisher: publishDateInput,
            stock: parseInt(stockInput),
            summary: introductionInput,
            title: Bookname
        }).then(res => {
            if (res.data.code === 200) {
                alert("添加成功！");
            }
        })
        .catch(err => {
            console.log(err);
        });
}

function addReaderinfo() {
    const nameInput = document.getElementById("name").value;
    const studentNoInput = document.getElementById('studentNo').value;
    const genderInput = document.getElementById('gender').value;
    const phoneInput = document.getElementById('phone').value;

    console.log(nameInput)

    API.Put.reader({
            gender: genderInput,
            name: nameInput,
            phone: phoneInput,
            studentNo: parseInt(studentNoInput)
        }).then(res => {
            if (res.data.code === 200) {
                alert("添加成功！");
            }
        })
        .catch(err => {
            console.log(err);
        });
}

function deleteBookOne(elem) {
    const bookId = elem.parentNode.dataset.bookid;

    API.Delete.book(parseInt(bookId))
        .then(res => {
            if (res.data.code === 200) {
                alert("删除成功！");
                elem.parentNode.remove();
            } else {
                alert(res.data.msg)
            }
        })
        .catch(err => {
            console.log(err);
        });
}


function deleteInfo(elem) {
    const readerID = elem.parentNode.dataset.readerid;
    API.Delete.reader(parseInt(readerID))
        .then(res => {
            if (res.data.code === 200) {
                elem.parentNode.remove();
            } else {
                alert("还有未归还的书籍，不能更新！")
                return 0
            }
        })
        .catch(err => {
            console.log(err);
        });
}

function showPopup(elem) {
    const bookId = elem.closest('li').dataset.bookid;
    const popup = document.querySelector(`.popup[data-bookid="${bookId}"]`);


    if (popup) {
        popup.style.display = "block";
    }
}

function showPoptwo(elem) {
    const popup = document.getElementById('add').nextElementSibling
    popup.style.display = 'block'
}


function Renewbook(elem) {
    const bookId = elem.parentNode.dataset.bookid
    const li = document.querySelector(`li[data-bookid="${bookId}"]`)


    const Bookname = document.getElementById("Bookname" + bookId).value;
    const nameInput = document.getElementById('authorname' + bookId).value;
    const coverInput = document.getElementById('coverUrl' + bookId).value;
    const priceInput = document.getElementById('price' + bookId).value;
    const publisherInput = document.getElementById('publisher' + bookId).value;
    const publishDateInput = document.getElementById('publishDate' + bookId).value;
    const introductionInput = document.getElementById('introduction' + bookId).value;
    const isbnInput = document.getElementById('isbn' + bookId).value;
    const stockInput = document.getElementById('storage' + bookId).value;
    const timeInput = document.getElementById('createdAt_time' + bookId).textContent;
    console.log(document.getElementById('createdAt_time' + bookId))
    console.log(timeInput)

    API.Put.book({
            author: nameInput,
            cover: coverInput,
            id: parseInt(bookId),
            isbn: isbnInput,
            price: parseInt(priceInput),
            publishDate: publishDateInput,
            publisher: publishDateInput,
            stock: parseInt(stockInput),
            summary: introductionInput,
            createdAt: timeInput,
            title: Bookname
        }).then(res => {
            if (res.data.code === 200) {
                alert("更新成功！");
            }
        })
        .catch(err => {
            console.log(err);
        });
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
    API.Admin.changePassword(oldPassword, newPassword)
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

export { getNewbookInfo, putNewbook, deleteBookOne, searchBorrowlist, searchAll, changePassword, showPopup, Renewbook, searchReaderlist, deleteInfo, searchrReservelist, changePage, showPoptwo, addReaderinfo }