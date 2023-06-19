function prevPage() {
    loadReserveList(page - 1);
    if (page === 2) {
        previousPageButton.removeAttribute("disabled");
    }
}

function nextPage() {
    loadReserveList(page + 1);

    if (page === totalPages - 1) {
        nextPageButton.removeAttribute("disabled");
    }
}

function successReserve(elem) {
    elem.innerText = "已预约";
    elem.disabled = true;
}

function hideLoadingReserve(elem) {
    elem.innerText = "预约";
}

function showLoading(elem) {
    console.log(elem); // 打印 elem参数
    console.log('show loading called!')
    elem.innerText = "处理中..";
}

function successBorrow(elem) {
    console.log('successBorrow called!')
    elem.innerText = "已借阅";
    elem.disabled = true;
}

function alreadyBorrowed(elem) {
    console.log('alreadyBorrowed called!')
    alert("已经借阅过!");
}

function onSuccessBorrow(elem) {
    console.log('onSuccessBorrow called!')
    alert("借阅成功!");
}

function alreadyReserve(elem) {
    console.log('alreadyReserve called!')
    alert("已经预约过!");
}

function onSuccessReserve(elem) {
    console.log('onSuccessReserve called!')
    alert("预约成功!");
}

export { prevPage, nextPage, successReserve, hideLoadingReserve, showLoading, successBorrow, alreadyBorrowed, onSuccessBorrow ,alreadyReserve,onSuccessReserve }