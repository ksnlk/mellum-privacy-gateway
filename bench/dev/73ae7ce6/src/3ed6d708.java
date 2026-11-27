package com.pj;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import cn.dev33.satoken.SaManager;

/**
 * Sa-Token 整合 WebSocket 鉴权示例 
 * @author click33
 *
 */
@SpringBootApplication
public class SaTokenWebSocketSpringApplication {

	/*
	 * 1、访问登录接口，拿到会话Token：
	 * 		http://localhost:8081/acc/doLogin?name=zhang&pwd=123456
	 * 
	 * 2、找一个WebSocket在线测试页面进行连接，
	 * 		例如：
	 * 			https://www.bejson.com/httputil/websocket/
	 * 		然后连接地址：
	 * 			ws://localhost:8081/ws-connect?satoken=51918c58-863b-3b25-cbb5-8b68634d6cbe
	 */
	
	public static void main(String[] args) {
		SpringApplication.run(SaTokenWebSocketSpringApplication.class, args); 
		System.out.println("\n启动成功：Sa-Token配置如下：" + SaManager.getConfig());
	}
	
}
