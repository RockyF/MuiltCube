/**
 * Created by lenovo on 14-4-4.
 */
/// <reference path="async.d.ts" />

module muiltcube {
	export class LoginPanel {
		onSuccess:any;

		constructor(onSuccess:any) {
			this.onSuccess = onSuccess;
			Utils.loadTPL("tpl/login.html", this.onTplLoaded);
		}

		onTplLoaded = (tpl)=> {
			$("body").append(tpl);

			$("#btnLogin").click(this.onBtnLoginClicked);
		};

		onBtnLoginClicked = ()=> {
			var id = parseInt($("#tiID").val());
			var pwd = $("#tiPwd").val();

			if (id == 0 || pwd == "") {
				return;
			}

			muiltcube.Client.getInstance().addCmdListener(1001, this.on1001Response);
			muiltcube.Client.getInstance().send(1001, {id: id, pwd: pwd});
		};

		on1001Response = (msg:any)=> {
			muiltcube.Client.getInstance().removeCmdListener(1001, this.on1001Response);
			if (msg.result == 0) {
				alert("onLoginSuccess");
				$("#loginPanel").remove();
				this.onSuccess();
			} else {
				alert("onLoginFailed");
			}
		};
	}
	export class RoleCreatePanel {
		onSuccess:any;
		rpc:any;
		colors:any;

		constructor(onSuccess:any) {
			this.onSuccess = onSuccess;

			this.rpc = RPC.getInstance();

			async.parallel({
					one: this.one,
					two: this.two
				},
				function (err, results) {

				});
		}

		one = (callback)=> {
			Utils.loadTPL("tpl/role_create.html", function (tpl) {
				$("body").append(tpl);
				callback();
			});
		};

		two = (callback)=> {
			this.rpc.execute("Common", "getColorList", {id: 1001}, function (msg) {
				if(msg.result == 0){
					this.colors = msg.body;
				}
				var colorSelectBox = $("#colorSelectBox");
				var liTpl = $("<li style='width: 20px;height: 20px;'></li>");
				for (var i = 0; i < this.colors.length; i++) {
					var o = colors[i];
					var li = liTpl.clone();
					console.log("#" + o.value);
					li.css("background-color", "#" + o.value);
					colorSelectBox.append(li);
				}
				callback();
			});
		};
	}
}