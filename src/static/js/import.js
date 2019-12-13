var step_num = 1;

function showPopup(){ 
    step_num = 1; 
    var step = document.getElementById("step2");
    step.className = "";
    step = document.getElementById("step3");
    step.className = "";
    var page = document.getElementById("myprojects");
    page.style.display="block";
    page = document.getElementById("init_message");
    page.style.display="none";
    page = document.getElementById("config_project");
    page.style.display="none";
    var button = document.getElementById("btnNextStep");
    button.style.display="block";
    button = document.getElementById("btnconfirm");
    button.style.display="none";

    var popUp = document.getElementById("import"); 
    popUp.style.position= "absolute";
    popUp.style.zIndex="100";
    popUp.style.width = "100%"; 
    popUp.style.height = "100%"; 
    popUp.style.visibility = "visible"; 
} 

function hidePopup(){
    var popUp = document.getElementById("import"); 
    popUp.style.visibility = "hidden";    
} 

function next_step(){
    if(step_num === 1) {            //在第一步输入url后，点击下一步时，进行判断，验证其权限
        var isProjOwner = false;    //表示仓库是否属于此用户，默认为false
        var orgMember = false;      //表示仓库是否属于用户所在的组织，暂且默认为false，待后端写好再实现
        var projUrl = document.getElementById("projectUrl").value;    //获取用户输入的URL
        var temp = projUrl.split('/');

        //用cookie获取用户的用户名
        var userName = getCookie("userName");
        if (userName==null) {
            //此时浏览器中cookie已经到期，需要重新登录
        }

        //用斜杠‘/’分割项目url，则倒数第二项为用户名。两个用户名相等，则项目属于该用户
        if(userName === temp[temp.length-2] ||
            userName === temp[temp.length-2].replace("git@github.com:", ""))    //使用SSH
            isProjOwner = true;
        //如果两个条件都不满足，则提示错误并返回
        if(!isProjOwner && !orgMember) {
            alert('请检查输入的URL:您不是此项目的所有者，且不属于此项目的组织。（当前仅支持https）');
            return;
        }
    }

    if(step_num >= 1 && step_num < 3)
        step_num += 1;
    switch(step_num){
        case 2:
            var step = document.getElementById("step2");
            step.className = "active";
            var page = document.getElementById("myprojects");
            page.style.display="none";
            page = document.getElementById("init_message");
            page.style.display="block";
            break;
        case 3:
            var step = document.getElementById("step3");
            step.className = "active";
            var page = document.getElementById("init_message");
            page.style.display="none";
            page = document.getElementById("config_project");
            page.style.display="block";
            var button = document.getElementById("btnNextStep");
            button.style.display="none";
            button = document.getElementById("btnconfirm");
            button.style.display="block";
            break;
        default:
            break;
    } 
}

function back_step(){
    if(step_num >1 && step_num <= 3)
        step_num -= 1;
    switch(step_num){
        case 1:
            var step = document.getElementById("step2");
            step.className = "";
            var page = document.getElementById("myprojects");
            page.style.display="block";
            page = document.getElementById("init_message");
            page.style.display="none";
            break;
        case 2:
            var step = document.getElementById("step3");
            step.className = "";
            var page = document.getElementById("init_message");
            page.style.display="block";
            page = document.getElementById("config_project");
            page.style.display="none";
            var button = document.getElementById("btnNextStep");
            button.style.display="block";
            button = document.getElementById("btnconfirm");
            button.style.display="none";

            break;
        default:
            break;
    }
}

function showImg(input) {
    var file = input.files[0];
    var url = window.URL.createObjectURL(file);
    document.getElemtById('upload_image').src=url;
}

function show_joined(){   
    var popUp = document.getElementById("joined_projects");
    popUp.style.display = "block";
    popUp = document.getElementById("managed_projects"); 
    popUp.style.display = "none";
    document.getElementById("join").style.backgroundColor = 'white';
    document.getElementById("manage").style.backgroundColor = '#bfbfbf';
}

function show_managed(){
    var popUp = document.getElementById("managed_projects"); 
    popUp.style.display = "block";
    popUp = document.getElementById("joined_projects"); 
    popUp.style.display = "none";
    document.getElementById("join").style.backgroundColor = '#bfbfbf';
    document.getElementById("manage").style.backgroundColor = 'white';
}

function setUrl(url) {
    var projectUrl = document.getElementById("projectUrl");
    projectUrl.value = url;
    next_step();
}

function getCookie(cookieKey){
    var arrcookie = document.cookie.split("; ");
    //遍历匹配
    for ( var i = 0; i < arrcookie.length; i++) {
        var arr = arrcookie[i].split("=");
        if (arr[0] == cookieKey){
            return arr[1];
        }
    }
    return null;
}