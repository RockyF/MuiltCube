/**
 * Created by RockyF on 14-3-31.
 */
/// <reference path="jquery.d.ts" />
/// <reference path="Client.ts" />
/// <reference path="Panels.ts" />
/// <reference path="Utils.ts" />

$(function(){
	muiltcube.Main.getInstance().start();
});

var wsuri = "ws://127.0.0.1:8888";
module muiltcube{
	export class Main{
		private static _instance:Main;
		public static getInstance():any {
			if (!this._instance) {
				this._instance = new Main();
			}

			return this._instance;
		}

		client:any;
		rpc:any;

		start(){
			this.client = muiltcube.Client.getInstance();
			this.client.init(wsuri);
			this.client.start();

			this.rpc = RPC.getInstance();
			this.rpc.init("http://localhost/WebServer/gateWay.php");

			this.rpc.execute("User", {id:1001}, function(data){
				console.log(data);
			});

			this.showLoginPanel();
		}

		showLoginPanel(){
			new LoginPanel(this.onLoginSuccess);
		}

		onLoginSuccess=()=>{
			new RoleCreatePanel(this.onLoginSuccess);
		}
	}
}