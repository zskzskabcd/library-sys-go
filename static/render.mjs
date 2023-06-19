import * as action from './action.mjs'
window.action = action

function renderReservations(reservations) {
    // 渲染数据
    let listHTML = "";

    reservations.data.forEach((reservation) => {
        listHTML += `
            <li data-reservationid="${reservation.id}">
            <p class="item-image"><img src="${reservation.book.cover}" /></p> 
            <p>书名${reservation.book.title}</p>
            <p class="ISBN">ISBN：${reservation.book.isbn}</p>
            <p>读者名：${reservation.reader.name}</p>
            <p>预约时间：${reservation.updatedAt}</p>
            <button onclick="action.cancelReserve(this)">取消预约</button>
            </li>
            `;
    });

    document.getElementById("book_list").innerHTML = listHTML;
}

function renderList(books) {
    let listHTML = "";

    books.data.forEach((book) => {
        listHTML += `
            <li data-bookid="${book.id}">  
            <div class="item-image"><img src="${book.cover}"></div> 
            <i class="fas fa-chevron-right" onclick="toggleDetails(this)"></i>
            <p class="book-title">书籍名：${book.title}</p>
            <p class="author">作者：${book.author}</p>    
            <p class="publisher">出版社：${book.publisher}</p>
            <p class="publish-date">发行时间：${book.publishDate}</p>
            <p class="ISBN">ISBN：${book.isbn}</p>
            <button class="reserve" onclick="action.reserveBook(this)">预约</button>
            <button class="borrow" onclick="action.borrow(this)">借阅</button>
            </li>
            <div class="details" style="display: none;">
            <!-- 详细信息 -->
            </div>
            `;
    });

    document.getElementById("book_list").innerHTML = listHTML;
}


function renderBooks(data) {
    const books = data.data;

    let listHTML = "";

    books.data.forEach((book) => {
        if (book.status == 1) {
            listHTML += `
            <li data-bookid="${book.book.id}"> 
            <p>书名：${book.book.title}</p>       
            <p>出版社：${book.book.publisher}</p>       
            <p>读者名：${book.reader.name}</p>       
            <p>应归还日期:${book.returnTime}</p>  
            <button onclick="action.returnBook(this)">还书</button>
            </li>
            `;
        }
    });

    document.getElementById("book_list").innerHTML = listHTML;
}

function renderPagination(page) {
    const total = localStorage.getItem('total');
    const totalPages = Math.ceil(localStorage.getItem("total") / 10);

    let paginationHTML = "";

    // 上一页按钮    
    paginationHTML += `<button `;
    if (page === 1) paginationHTML += `disabled `;
    paginationHTML += `onclick="action.changePage('prev')">`;
    paginationHTML += `上一页</button>`;

    // 下一页按钮   
    paginationHTML += `<button `;
    if (page === totalPages) paginationHTML += `disabled `;
    paginationHTML += `onclick="action.changePage('next')">`;
    paginationHTML += `下一页</button>`;

    document.getElementById("pagination").innerHTML = paginationHTML;
}

export { renderReservations, renderList, renderBooks, renderPagination }