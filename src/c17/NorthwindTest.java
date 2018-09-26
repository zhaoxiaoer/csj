package c17;

import java.sql.*;

public class NorthwindTest {
	public static void main(String[] args) {
		String driver = "com.mysql.cj.jdbc.Driver";
		String url = "jdbc:mysql://127.0.0.1:3306/ztest";
		String username = "";
		String password = "";
		
		showEmployeeTable(driver, url, username, password);
	}
	
	public static void showEmployeeTable(String driver,
			String url, String username, String password) {
		try {
			// Load database driver if it's not already loaded.
			Class.forName(driver);
			Connection connection = DriverManager.getConnection(url, username, password);
			Statement statement = connection.createStatement();
			String query = "SELECT * FROM employee";
			ResultSet resultSet = statement.executeQuery(query);
			while (resultSet.next()) {
				System.out.println(resultSet.getString("firstname"));
				System.out.println(resultSet.getString("lastname"));
			}
			connection.close();
		} catch (ClassNotFoundException cnfe) {
			cnfe.printStackTrace();
		} catch (SQLException sqle) {
			sqle.printStackTrace();
		}
	}
}
