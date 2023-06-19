import * as action from './admin_action.mjs'

window.action = action

window.loadContent = function(page) {
    if (page == "getNewbook") {
        document.getElementById("content").innerHTML = `
        <div class="center">
            <div class="input-box">
              <label for="name">书名</label>
              <br>
              <input type="text" id="Bookname">   
            </div>

            <div class="input-box">
              <label for="name">作者名</label>
              <br>
              <input type="text" id="authorname">   
            </div>

            <div class="input-box">
              <label for="name">封面地址</label>
              <br>
              <input type="text" id="coverUrl">   
            </div>
            <div class="input-box">
              <label for="name">价格</label>
              <br>
              <input type="text" id="price">   
            </div>
            <br>
            <div class="input-box">
              <label for="name">出版社</label>
              <br>
              <input type="text" id="publisher">   
            </div>

            <div class="input-box">
              <label for="name">出版日期</label>
              <br>
              <input type="text" id="publishDate">   
            </div>
            <div class="input-box">
              <label for="name">ISBN</label>
              <br>
              <input type="text" id="isbn">   
            </div>
            <div class="input-box">
              <label for="name">库存量</label>
              <br>
              <input type="text" id="storage">   
            </div>

            <div class="input-box">  
              <label for="text">简介</label>
              <br>
            <textarea id="introduction"></textarea>
            </div>
</div>
            <button type="button" id="search" onclick=action.getNewbookInfo()>云搜索</button>
                <button type="button" onclick=action.putNewbook()>提交</button>
                `;
    }
    if (page == "BOOKlist") {
        document.getElementById("content").innerHTML = ` 
                <div id="book_list"> </div>   
                <div id="pagination"> 
                </div>  
                `
        action.searchAll();
    }
    if (page === "adminReaderInfo") {
        document.getElementById("content").innerHTML = `  
                <div id='addArea'>
                <button type="button" id="add" onclick=action.showPoptwo(this)>增添读者信息</button>
                <div class="popup" style="display: none;">
                <input type="text" placeholder="姓名" id="name"> 
                <input type="text" placeholder="学号" id="studentNo">   
                <input type="text" placeholder="性别" id="gender">  
                <input type="text" placeholder="手机号" id="phone">        
                <button type="button" onclick=action.addReaderinfo(this)>提交</button>      
                </div>
                </div>
                <div id="reader_list"> </div>
                <div id="pagination"> 
                </div>  
                `;
        action.searchReaderlist();
    }
    if (page == 'borrow_list') {
        document.getElementById("content").innerHTML = `
                <div id="book_list"> </div>   
                <div id="pagination"> 
                </div> 
                `;
        action.searchBorrowlist();
    }

    if (page == "reserve_list") {
        document.getElementById("content").innerHTML = `
                <div id="book_list"> </div>   
                <div id="pagination"> 
                </div> 
                `;
        action.searchrReservelist();
    }
    if (page == "change") {
        document.getElementById("content").innerHTML = `
                <form>  
                <input id="old" placeholder="旧密码">            
                <input id="new" placeholder="请输入新密码">
                <input id="check" placeholder="请确认新密码">
                <button onclick="window.action.changePassword()">修改密码</button>
                </form>
                `;
    }
    if (page == 'logout') {
        localStorage.removeItem('token');
        localStorage.removeItem('page');
        localStorage.removeItem('item');
        window.location.href = "admin_login.html";
    }
}
