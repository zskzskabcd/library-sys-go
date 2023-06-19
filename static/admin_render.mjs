import * as action from './action.mjs'
window.action = action

function renderReservations(reservations) {
    // 渲染数据
    let listHTML = "";
    reservations.data.forEach((reservation) => {
        listHTML += `
            <li data-reservationid="${reservation.id}">
            <p>书名：${reservation.book.title}</p>
            <p class="ISBN">ISBN：${reservation.book.isbn}</p>
            <p>读者名：${reservation.reader.name}</p>
            <p>预约时间：${reservation.updatedAt}</p>
            `;
        if (reservation.status == 1) {
            listHTML += `
            <p>当前状态：进行中 </p>
            `
        }
        if (reservation.status == 2) {
            listHTML += `
            <p>当前状态：已借阅 </p>
            `
        }
        if (reservation.status == 3) {
            listHTML += `
            <p>当前状态：已超时 </p>
            `
        }
        if (reservation.status == 4) {
            listHTML += `
            <p>当前状态：已取消 </p>
            `
        }

        listHTML += ` </li> <hr />`
    });
    document.getElementById("book_list").innerHTML = listHTML;

}

function renderList(books) {
    let listHTML = "";
    books.data.forEach((book) => {
        listHTML += `
                <li data-bookid="${book.id}">  
                    <p class="item-image"><img src="${book.cover}"></p> 
                    <i class="fas fa-chevron-right" onclick="toggleDetails(this)"></i>
                    <p class="book-title">${book.title}</p>
                    <p class="author">作者：${book.author}</p>    
                    <p class="publisher">出版社：${book.publisher}</p>
                    <p class="publish-date">发行时间：${book.publishDate}</p>
                    <p class="ISBN">ISBN：${book.isbn}</p>
                    <p id="createdAt_time${book.id}">${book.createdAt}</p>
                    <button class="reserve" onclick="action.showPopup(this)">更新信息</button>
                    <button class="borrow" onclick="action.deleteBookOne(this)">删除书籍</button>
                </li>
                <hr />
            <div class="popup" data-bookid="${book.id}"  style="display: none;">
             <input type="text" placeholder="书名" id="Bookname${book.id}" />   
            <input type="text" placeholder="作者名" id="authorname${book.id}" />  
            <input type="text" placeholder="封面地址" id="coverUrl${book.id}" />  
            <input type="text" placeholder="ISBN" id="isbn${book.id}" />     
            <input type="number" placeholder="价格" id="price${book.id}" />
            <input type="text" placeholder="出版社" id="publisher${book.id}" />
            <input type="text" placeholder="出版日期" id="publishDate${book.id}" /> 
            <input type="number" placeholder="库存量" id="storage${book.id}" />
             <textarea placeholder="简介" id="introduction${book.id}"></textarea>   
            <button type="button" onclick=action.Renewbook(this)>提交</button>      
            </div>
            `;
    });

    document.getElementById("book_list").innerHTML = listHTML;
}

function renderReaderlist(data) {
    const getedReaders = data;
    let listHTML = "";

    getedReaders.data.forEach((reader) => {
        listHTML += `
            <li data-readerid="${reader.id}"> 
            <p>姓名${reader.name}</p>       
            <p>学号${reader.studentNo}</p>       
            <p>性别${reader.gender}</p>       
            <p>电话${reader.phone}</p>  
            <button onclick="action.deleteInfo(this)">删除信息</button>
            </li>
            <hr />
            `;
    });

    document.getElementById("reader_list").innerHTML = listHTML;
}

function renderBorrwolist(data) {
    const books = data.data;

    let listHTML = "";

    books.forEach((book) => {
        listHTML += `
            <li data-bookid="${book.book.id}"> 
            <p>书名：${book.book.title}</p>       
            <p>出版社：${book.book.publisher}</p>       
            <p>读者名：${book.reader.name}</p>       
            <p>应归还日期:${book.returnTime}</p>  
            `;
        if (book.status == 1) {
            listHTML += `
            <p>当前状态：借阅中 </p>
            `
        }
        if (book.status == 2) {
            listHTML += `
            <p>当前状态：已归还 </p>
            `
        }
        listHTML += `
        </li>
        <hr />
        `
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
export { renderReservations, renderList, renderReaderlist, renderBorrwolist, renderPagination }