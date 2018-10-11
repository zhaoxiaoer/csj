package v2c2;

import java.util.*;
import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class ReportFilter implements Filter {
	public void init(FilterConfig config) throws ServletException {
		
	}
	
	public void doFilter(ServletRequest request, 
			ServletResponse response,
			FilterChain chain)
		throws ServletException, IOException {
		HttpServletRequest req = (HttpServletRequest) request;
		System.out.println(req.getRemoteHost() +
				" tried to access " +
				req.getRequestURI() +
				" on " + new Date() + ".");
		chain.doFilter(request, response);
	}
	
	public void destroy() {
		
	}
}
