
/**
* 分页函数
* pno--页数
* psize--每页显示记录数
* 分页部分是从真实数据行开始，因而存在加减某个常数，以确定真正的记录数
* 纯js分页实质是数据行全部加载，通过是否显示属性完成分页功能
**/
function goPage1(pno,psize){

    var itable = document.getElementById("idData1");
    var num = itable.rows.length;//表格所有行数(所有记录数)
    console.log(num);
    var totalPage = 0;//总页数
    var pageSize = psize;//每页显示行数
    //总共分几页
    if(num/pageSize > parseInt(num/pageSize)){            //判断是否为小数
    totalPage=parseInt(num/pageSize)+1;           }
    else{
    totalPage=parseInt(num/pageSize);
        }
    var currentPage = pno;//当前页数
    var startRow = (currentPage - 1) * pageSize+1;//开始显示的行 31
    var endRow = currentPage * pageSize;//结束显示的行  40
    endRow = (endRow > num)? num : endRow;  //40
    console.log(endRow);
    //遍历显示数据实现分页
    for(var i=1;i<(num+1);i++){
    var irow = itable.rows[i-1];
    if(i>=startRow && i<=endRow){
    irow.style.display = "table-row";
    }else{
    irow.style.display = "none";
    }
    }
    var tempStr="";
    
if(currentPage>1){
    tempStr += "<span class='paging-btn' href=\"#\" onClick=\"goPage1("+(1)+","+psize+")\">首页</span>";
    tempStr += "<span class='paging-btn' href=\"#\" onClick=\"goPage1("+(currentPage-1)+","+psize+")\"><<</span>";
}else{
    tempStr += "<span class='paging-btn'>首页</span>";
    tempStr += "<span class='paging-btn'><<</span>";
}

for(var pageIndex= 1;pageIndex<totalPage+1;pageIndex++){
    tempStr += "<a onclick=\"goPage1("+pageIndex+","+psize+")\"><span class='paging-btn'>"+ pageIndex +"</span></a>";
}

if(currentPage<totalPage){
    tempStr += "<span class='paging-btn' href=\"#\" onClick=\"goPage1("+(currentPage+1)+","+psize+")\">>></span>";
    tempStr += "<span class='paging-btn' href=\"#\" onClick=\"goPage1("+(totalPage)+","+psize+")\">尾页</span>";
}else{
    tempStr += "<span class='paging-btn'>>></span>";
    tempStr += "<span class='paging-btn'>尾页</span>";   
}
    document.getElementById("barcon1").innerHTML = tempStr;
}

//功能相同用于同一个页面第二次分页
function goPage2(pno,psize){

    var itable = document.getElementById("idData2");
    var num = itable.rows.length;//表格所有行数(所有记录数)
    console.log(num);
    var totalPage = 0;//总页数
    var pageSize = psize;//每页显示行数
    //总共分几页
    if(num/pageSize > parseInt(num/pageSize)){            //判断是否为小数
    totalPage=parseInt(num/pageSize)+1;           
    }else{
    totalPage=parseInt(num/pageSize);
    }
    var currentPage = pno;//当前页数
    var startRow = (currentPage - 1) * pageSize+1;//开始显示的行 31
    var endRow = currentPage * pageSize;//结束显示的行  40
    endRow = (endRow > num)? num : endRow;  //40
    console.log(endRow);
    //遍历显示数据实现分页
    for(var i=1;i<(num+1);i++){
    var irow = itable.rows[i-1];
    if(i>=startRow && i<=endRow){
    irow.style.display = "table-row";
    }else{
    irow.style.display = "none";
    }
    }
    var tempStr="";
    
if(currentPage>1){
    tempStr += "<span class='paging-btn' href=\"#barcon2\" onClick=\"goPage2("+(1)+","+psize+")\">首页</span>";
    tempStr += "<span class='paging-btn' href=\"#barcon2\" onClick=\"goPage2("+(currentPage-1)+","+psize+")\"><<</span>";
}else{
    tempStr += "<span class='paging-btn'>首页</span>";
    tempStr += "<span class='paging-btn'><<</span>";
}

for(var pageIndex= 1;pageIndex<totalPage+1;pageIndex++){
    tempStr += "<a href=\"#barcon2\" onclick=\"goPage2("+pageIndex+","+psize+")\"><span class='paging-btn'>"+ pageIndex +"</span></a>";
}

if(currentPage<totalPage){
    tempStr += "<span class='paging-btn' href=\"#barcon2\" onClick=\"goPage2("+(currentPage+1)+","+psize+")\">>></span>";
    tempStr += "<span class='paging-btn' href=\"#barcon2\" onClick=\"goPage2("+(totalPage)+","+psize+")\">尾页</span>";
}else{
    tempStr += "<span class='paging-btn'>>></span>";
    tempStr += "<span class='paging-btn'>尾页</span>";   
}
    document.getElementById("barcon2").innerHTML = tempStr;
}
