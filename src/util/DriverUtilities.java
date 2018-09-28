package util;

import java.util.*;

public class DriverUtilities {
	public static final String MSACCESS = "MSACCESS";
	public static final String MYSQL = "MYSQL";
	public static final String ORACLE = "ORACLE";
	
	protected static Map driverMap = new HashMap();
	
	public static void loadDrivers() {
		String vendor, description, driverClass, url;
		DriverInfoBean info = null;
		
		// MSAccess
		vendor = MSACCESS;
		description = "MS Access 4.0";
		driverClass = "sun.jdbc.odbc.JdbcOdbcDriver";
		url = "jdbc:odbc:[$dbName]";
		info = new DriverInfoBean(vendor, description, driverClass, url);
		addDriverInfoBean(info);
		
		// MySQL
		vendor = MYSQL;
		description = "MySQL Connector/J 3.0";
		driverClass = "com.mysql.cj.jdbc.Driver";
		url = "jdbc:mysql://[$host]:3306/[$dbName]";
		info = new DriverInfoBean(vendor, description, driverClass, url);
		addDriverInfoBean(info);
		
		// Oracle
		vendor = ORACLE;
		description = "Oracle9i Database";
		driverClass = "oracle.jdbc.driver.OracleDriver";
		url = "jdbc:oracle:thin:@[$host]:1521:[$dbName]";
		info = new DriverInfoBean(vendor, description, driverClass, url);
		addDriverInfoBean(info);
		
		// Add info on your database here...
	}
	
	public static void addDriverInfoBean(DriverInfoBean info) {
		driverMap.put(info.getVendor().toUpperCase(), info);
	}
	
	public static boolean isValidVendor(String vendor) {
		DriverInfoBean info = (DriverInfoBean) driverMap.get(vendor.toUpperCase());
		return info != null;
	}
	
	public static String makeURL(String host, String dbName, String vendor) {
		DriverInfoBean info = (DriverInfoBean) driverMap.get(vendor.toUpperCase());
		if (info == null) {
			return null;
		}
		StringBuffer url = new StringBuffer(info.getUrl());
		DriverUtilities.replace(url, "[$host]", host);
		DriverUtilities.replace(url, "[$dbName]", dbName);
		return url.toString();
	}
	
	public static String getDriver(String vendor) {
		DriverInfoBean info = (DriverInfoBean) driverMap.get(vendor.toUpperCase());
		if (info == null) {
			return null;
		} else {
			return info.getDriverClass();
		}
	}
	
	private static void replace(StringBuffer buffer, String match, String value) {
		int index = buffer.toString().indexOf(match);
		if (index > 0) {
			buffer.replace(index, index + match.length(), value);
		}
	}
}
