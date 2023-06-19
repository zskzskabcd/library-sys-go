import * as action from "./action.mjs";

window.action = action;

window.loadContent = function (page) {
  if (page == "borrow") {
    document.getElementById("content").innerHTML = `
                <div class="form-group"><input type="text" id="book" placeholder="请输入书名查询"><button onclick="action.search()">查询</button> </div>  
                <div id="book_list"> </div>     
                `;
  }
  if (page === "return") {
    document.getElementById("content").innerHTML = `  
                 <div id="book_list"> </div> 
                 <div id="pagination"> 
                </div>  
                `;
    action.searchBorrowwedBooks();
  }
  if (page == "BOOKlist") {
    document.getElementById("content").innerHTML = ` 
                <div id="book_list"> </div>   
                 <div id="pagination"> 
                </div>  
                `;
    action.searchAll();
  }
  if (page == "reserve_list") {
    document.getElementById("content").innerHTML = `
                <div id="book_list"> </div>   
                <div id="pagination"> 
                </div>  
                `;
    action.loadReserveList();
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
  if (page == "logout") {
    localStorage.removeItem("token");
    window.location.href = "login.html";
  }
};
