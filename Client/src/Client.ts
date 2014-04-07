/**
 * Created by lenovo on 14-3-31.
 */
module muiltcube {
	export class Client {
		private static _instance:Client;
		public static getInstance():any {
			if (!this._instance) {
				this._instance = new Client();
			}

			return this._instance;
		}

		wsServer:string;
		websocket:WebSocket;
		callbackMap:any;

		init(wsServer:string) {
			this.wsServer = wsServer;
			this.callbackMap = [];
		}

		start() {
			this.websocket = new WebSocket(this.wsServer);

			this.websocket.onopen = this.onOpen;
			this.websocket.onclose = this.onClose;
			this.websocket.onmessage = this.onMessage;
			this.websocket.onerror = this.onError;
		}

		send(cmd:number, body:any) {
			var msg:any = new Object();
			msg.cmd = cmd;
			msg.body = body;
			this.websocket.send(JSON.stringify(msg));
		}

		sendData(content:string) {
			this.websocket.send(content);
		}

		addCmdListener(cmd:number, callback:any) {
			if (!this.callbackMap[cmd]) {
				this.callbackMap[cmd] = [];
			}

			this.callbackMap[cmd].push(callback);
		}

		removeCmdListener(cmd:number, callback:any) {
			var callbacks:any = this.callbackMap[cmd];
			if (callbacks) {
				for (var i = 0, len = callbacks.length; i < len; i++) {
					if (callback = callbacks[i]) {
						callbacks.splice(i, 1);
					}
				}
			}
		}

		getState() {
			return this.websocket.readyState;
		}

		onOpen=(event)=>{
			console.log(Utils.stringFormat('connect to {0} success!', this.wsServer));
		}

		onClose=(event)=>{
			console.log('connect closed!');
		}

		onMessage=(event)=>{
			console.log(Utils.stringFormat('get data: {0}', event.data));
			var msg:any = JSON.parse(event.data);
			var cmd:number = msg.cmd;

			var callbacks:any = this.callbackMap[cmd];
			if (callbacks) {
				for (var i = 0, len = callbacks.length; i < len; i++) {
					callbacks[i].call(this, msg.body);
				}
			}
		}

		onError=(event)=>{
			console.log('socket error!');
		}
	}

	export class RPC{
		private static _instance:RPC;
		public static getInstance():any {
			if (!this._instance) {
				this._instance = new RPC();
			}

			return this._instance;
		}

		gateWay:string;
		init(gateWay:string){
			this.gateWay = gateWay;
		}

		execute=(objectName:string, params:any, callback:any)=>{
			$.ajax({
				url:this.gateWay,
				method:"POST",
				data:JSON.stringify(params),
				beforeSend:function(request){
					request.setRequestHeader("objectName", objectName);
				},
				success:function(data){
					if(callback){
						callback(data);
					}
				}
			});
		}
	}
}