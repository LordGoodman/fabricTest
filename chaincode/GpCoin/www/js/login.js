$(document).ready(function(){
	
	$("#SignIn").click(function(){
		var user = $("#user").val()
		var pass = $("#pass").val()
		if (user!="" && pass !=""){
			if (user == "Luke" && pass == "Luke"){
				//alert("Welcome come Luke!")
				location.href = "./index.html"
				}else{
						alert("user name or password wrong")
					}
		}else{
			alert("invaild input!")
			}
		//this.hide()
		});
		
	$("#Register").click(function(){
			alert("Register temporary unvaluable")
		});
		
});