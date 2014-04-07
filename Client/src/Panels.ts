/**
 * Created by lenovo on 14-4-4.
 */

module muiltcube{
	export class LoginPanel{
		onSuccess:any;

		constructor(onSuccess:any){
			this.onSuccess = onSuccess;
			Utils.loadTPL("tpl/login.html", this.onTplLoaded);
		}

		onTplLoaded=(tpl)=>{
			$("body").append(tpl);

			$("#btnLogin").click(this.onBtnLoginClicked);
		};

		onBtnLoginClicked=()=>{
			var id = parseInt($("#tiID").val());
			var pwd = $("#tiPwd").val();

			if(id == 0 || pwd == ""){
				return;
			}

			muiltcube.Client.getInstance().addCmdListener(1001, this.on1001Response);
			muiltcube.Client.getInstance().send(1001, {id:id, pwd:pwd});
		};

		on1001Response=(msg:any)=>{
			muiltcube.Client.getInstance().removeCmdListener(1001, this.on1001Response);
			if(msg.result == 0){
				alert("onLoginSuccess");
				$("#loginPanel").remove();
				this.onSuccess();
			}else{
				alert("onLoginFailed");
			}
		};
	}
	export class RoleCreatePanel{
		onSuccess:any;

		constructor(onSuccess:any){
			this.onSuccess = onSuccess;
			Utils.loadTPL("tpl/role_create.html", this.onTplLoaded);
		}

		onTplLoaded=(tpl)=>{
			$("body").append(tpl);

			//$("#btnLogin").click(this.onBtnLoginClicked);
		};

	}
}