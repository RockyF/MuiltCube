<?php
/**
 * Created by IntelliJ IDEA.
 * User: RockyF
 * Date: 2014/4/8
 * Time: 0:15
 */

class RPC {
	function __construct(){
		$this->requestHandler();
	}

	function requestHandler(){
		$json_body = json_decode(file_get_contents("php://input"));
		$class_name = $json_body->className;
		$method_name = $json_body->methodName;

		$class_path = 'class/' . strtolower($class_name) . '.class.php';
		if(file_exists($class_path)){
			require ($class_path);
			if(class_exists($class_name)){
				$obj = new $class_name;
				if(method_exists($obj, $method_name)){
					$this->responseHandler($obj->$method_name($json_body->params));
				}else{
					echo("method is not exist.");
				}
			}else{
				echo("class is not exist.");
			}
		}else{
			echo("class file is not exist.");
		}
	}

	function responseHandler($data){
		echo(json_encode($data));
	}
} 