/**
* Created by lenovo on 14-4-4.
*/
var muiltcube;
(function (muiltcube) {
    var LoginPanel = (function () {
        function LoginPanel(onSuccess) {
            var _this = this;
            this.onTplLoaded = function (tpl) {
                $("body").append(tpl);

                $("#btnLogin").click(_this.onBtnLoginClicked);
            };
            this.onBtnLoginClicked = function () {
                var id = parseInt($("#tiID").val());
                var pwd = $("#tiPwd").val();

                if (id == 0 || pwd == "") {
                    return;
                }

                muiltcube.Client.getInstance().addCmdListener(1001, _this.on1001Response);
                muiltcube.Client.getInstance().send(1001, { id: id, pwd: pwd });
            };
            this.on1001Response = function (msg) {
                muiltcube.Client.getInstance().removeCmdListener(1001, _this.on1001Response);
                if (msg.result == 0) {
                    alert("onLoginSuccess");
                    $("#loginPanel").remove();
                    _this.onSuccess();
                } else {
                    alert("onLoginFailed");
                }
            };
            this.onSuccess = onSuccess;
            Utils.loadTPL("tpl/login.html", this.onTplLoaded);
        }
        return LoginPanel;
    })();
    muiltcube.LoginPanel = LoginPanel;
    var RoleCreatePanel = (function () {
        function RoleCreatePanel(onSuccess) {
            this.onTplLoaded = function (tpl) {
                $("body").append(tpl);
                //$("#btnLogin").click(this.onBtnLoginClicked);
            };
            this.onSuccess = onSuccess;
            Utils.loadTPL("tpl/role_create.html", this.onTplLoaded);
        }
        return RoleCreatePanel;
    })();
    muiltcube.RoleCreatePanel = RoleCreatePanel;
})(muiltcube || (muiltcube = {}));