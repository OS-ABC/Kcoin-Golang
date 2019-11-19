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
    button.innerHTML="下一步";

    var popUp = document.getElementById("import"); 
    popUp.style.position= "absolute";
    popUp.style.zIndex="100";
    popUp.style.width = "100%"; 
    popUp.style.height = "100%"; 
    popUp.style.visibility = "visible"; 
} 

function hidePopup(){
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

    var popUp = document.getElementById("import"); 
    popUp.style.visibility = "hidden";    
} 

function next_step(){
    if(step_num >=1 && step_num < 3)
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
            button.innerHTML="&nbsp确&nbsp&nbsp定&nbsp";
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
            button.innerHTML="下一步";
            break;
        default:
            break;
    }
}

function showImg(input) {
    var file = input.files[0];
    var url = window.URL.createObjectURL(file)
    document.getElemtById('upload_image').src=url
}