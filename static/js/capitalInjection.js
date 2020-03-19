function showPopup(){
    // var page = document.getElementById("injection");
    // page.style.display="block";
    // var button = document.getElementById("commit");
    // button.innerHTML="确认";
    var popUp = document.getElementById("injection");
    popUp.style.position= "absolute";
    popUp.style.zIndex="100";
    popUp.style.width = "100%";
    popUp.style.height = "100%";
    popUp.style.visibility = "visible";

}

function hidePopup(){
    var popUp = document.getElementById("injection");
    popUp.style.visibility = "hidden";
}


