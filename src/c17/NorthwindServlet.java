package c17;

import java.io.*;
import java.sql.*;

import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class NorthwindServlet extends HttpServlet {
	public void doGet(HttpServletRequest request, HttpServletResponse response)
		throws ServletException, IOException {
		response.setContentType("text/html");
		PrintWriter out = response.getWriter();
		String title = "Northwind Results";
		out.print(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\"><center>\n" +
				"<h1>" + title + "</h1>\n");
		showTable(out);
		out.println("</center></body></html>");
	}
	
	public void showTable(PrintWriter out) {
		try {
			Class.forName("com.mysql.cj.jdbc.Driver");
			String url = "jdbc:mysql://127.0.0.1:3306/ztest";
			String username = "root";
	        String password = "";
			Connection connection = DriverManager.getConnection(url, username, password);
			out.println("<ul>");
			DatabaseMetaData dbMetaData = connection.getMetaData();
			String productName = dbMetaData.getDatabaseProductName();
			String productVersion = dbMetaData.getDatabaseProductVersion();
			out.println("  <li><b>Database:</b> " + productName + "</li>\n" +
					"  <li><b>Version:</b> " + productVersion + "</li>\n" +
					"</ul>");
			Statement statement = connection.createStatement();
			ResultSet resultSet = statement.executeQuery("SELECT * FROM employee");
			out.println("<table border=1>");
			ResultSetMetaData resultSetMetaData = resultSet.getMetaData();
			int columnCount = resultSetMetaData.getColumnCount();
			out.println("<tr>");
			for (int i = 1; i <= columnCount; i++) {
				out.println("<th>" + resultSetMetaData.getColumnName(i) + "</th>");
			}
			out.println("</tr>");
			while (resultSet.next()) {
				out.println("<tr>");
				for (int i = 1; i <= columnCount; i++) {
					out.println("  <td>" + resultSet.getString(i) + "</td>");
				}
				out.println("</tr>");
			}
			out.println("</table>");
			connection.close();
		} catch (ClassNotFoundException cnfe) {
			cnfe.printStackTrace();
		} catch (SQLException sqle) {
			sqle.printStackTrace();
		}
	}
}
