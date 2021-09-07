package com.sangupta.redefine;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.Collection;

import org.eclipse.jetty.server.Request;
import org.eclipse.jetty.server.handler.AbstractHandler;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.sangupta.jerry.constants.HttpMimeType;
import com.sangupta.jerry.constants.HttpStatusCode;
import com.sangupta.redefine.model.ComponentDef;

import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;

public class HttpHandler extends AbstractHandler {
	
	private static final Gson GSON = new GsonBuilder().setPrettyPrinting().create();

	@Override
	public void handle(String target, Request baseRequest, HttpServletRequest request, HttpServletResponse response) throws IOException, ServletException {
		// add our custom headers - this happens first otherwise
		// Jetty adds its own headers
		response.addHeader("Server", "redefine");
		
		handleIncomingRequest(request, response);
		baseRequest.setHandled(true);
	}
	
	private void handleIncomingRequest(HttpServletRequest request, HttpServletResponse response) throws IOException {
		String uri = request.getRequestURI();
		
		if("/components.json".equals(uri)) {
			sendComponentsList(request, response);
			return;
		}

		// we can't handle the request
		response.sendError(HttpStatusCode.NOT_FOUND);
	}

	private void sendComponentsList(HttpServletRequest request, HttpServletResponse response) throws IOException {
		Collection<ComponentDef> components = RedefineMain.COMPONENT_MAP.values();
//		Collections.sort(components);
		
		String json = GSON.toJson(components);
		byte[] bytes = json.getBytes(StandardCharsets.UTF_8);
		response.setContentType(HttpMimeType.JSON);
		response.setContentLengthLong(bytes.length);
		response.getOutputStream().write(bytes, 0, bytes.length);
	}

}
