package util;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class ConnectionInfoBean {
	private String connectionName;
	private String description;
	private String driver;
	private String url;
	private String username;
	private String password;
	
	public ConnectionInfoBean() {
		
	}
	
	public ConnectionInfoBean(String connectionName,
			String description,
			String driver,
			String url,
			String username,
			String password) {
		setConnectionName(connectionName);
		setDescription(description);
		setDriver(driver);
		setUrl(url);
		setUsername(username);
		setPassword(password);
	}

	public String getConnectionName() {
		return connectionName;
	}

	public void setConnectionName(String connectionName) {
		this.connectionName = connectionName;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public String getDriver() {
		return driver;
	}

	public void setDriver(String driver) {
		this.driver = driver;
	}

	public String getUrl() {
		return url;
	}

	public void setUrl(String url) {
		this.url = url;
	}

	public String getUsername() {
		return username;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}
	
	public Connection getConnection() {
		return getConnection(driver, url, username, password);
	}
	
	public static Connection getConnection(String driver,
			String url,
			String username,
			String password) {
		try {
			Class.forName(driver);
			Connection connection = DriverManager.getConnection(url, username, password);
			return connection;
		} catch (ClassNotFoundException cnfe) {
			cnfe.printStackTrace();
			return null;
		} catch (SQLException sqle) {
			sqle.printStackTrace();
			return null;
		}
	}
}
