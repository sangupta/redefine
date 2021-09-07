package com.sangupta.redefine;

import org.eclipse.jetty.server.Server;

public class HttpServer {
	
	/**
	 * The server instance that we use
	 */
	private static Server HTTP_SERVER;

	public static void start() {
		// add shutdown hook
		Runtime.getRuntime().addShutdownHook(new Thread() {
			
			@Override
			public void run() {
				HttpServer.stop();
			}
			
		});
				
		HTTP_SERVER = new Server(13090);
		HTTP_SERVER.setHandler(new HttpHandler());
		
		try {
			HTTP_SERVER.start();
		} catch (Exception e) {
			System.out.println("Unable to start server.");
			e.printStackTrace();
			return;
		}
		
		try {
			HTTP_SERVER.join();
		} catch (InterruptedException e) {
			// server needs to be stopped
		}
		
		System.out.println("Bye!");
	}
	
	public static void stop() {
		if(HTTP_SERVER == null) {
			return;
		}
		
		try {
			HTTP_SERVER.stop();
		} catch (Exception e) {
			// eat this up
		}
	}
	
}
