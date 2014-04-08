<?php
require("DB.class.php");
/**
 * Created by IntelliJ IDEA.
 * User: RockyF
 * Date: 2014/4/8
 * Time: 1:17
 */

class Common{
	private $db;
	function __construct(){
		$this->db = new DB();
	}

	function getColorList(){
		$result = $this->db->get_all("select * from skin_color");
		return $result;
	}
}