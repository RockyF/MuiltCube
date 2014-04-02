/**
 * Created by lenovo on 14-3-31.
 */
class Utils{
	public static stringFormat(format:string, ...args){
		var result:string = format;
		for(var i = 0, len = args.length; i < len; i++){
			result = result.replace('{' + (i) + '}', args[i]);
		}

		return result;
	}
}