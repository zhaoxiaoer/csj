package v2c2;

import java.util.*;
import javax.servlet.*;

public class ContextReporter implements ServletContextListener {
	public void contextInitialized(ServletContextEvent event) {
		System.out.println("Context created on " + new Date() + ".");
	}
	public void contextDestroyed(ServletContextEvent event) {
		System.out.println("Context destroyed on " + new Date() + ".");
	}
}
