package util;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.parsers.ParserConfigurationException;

import org.w3c.dom.Document;
import org.w3c.dom.Element;
import org.w3c.dom.Node;
import org.w3c.dom.NodeList;
import org.xml.sax.SAXException;

public class DriverUtilities2 extends DriverUtilities {
	public static final String DEFAULT_FILE = "drivers.xml";
	
	public static void loadDrivers() {
		DriverUtilities2.loadDrivers(DEFAULT_FILE);
	}
	
	public static void loadDrivers(String filename) {
		File file = new File(filename);
		try {
			InputStream in = new FileInputStream(file);
			DocumentBuilderFactory builderFactory = DocumentBuilderFactory.newInstance();
			DocumentBuilder builder = builderFactory.newDocumentBuilder();
			Document document = builder.parse(in);
			document.getDocumentElement().normalize();
			Element rootElement = document.getDocumentElement();
			NodeList driverElements = rootElement.getElementsByTagName("driver");
			for (int i = 0; i < driverElements.getLength(); i++) {
				Node node = driverElements.item(i);
				DriverInfoBean info = DriverUtilities2.createDriverInfoBean(node);
				if (info != null) {
					addDriverInfoBean(info);
				}
			}
		} catch (FileNotFoundException fnfe) {
			fnfe.printStackTrace();
		} catch (ParserConfigurationException pce) {
			pce.printStackTrace();
		} catch (IOException ioe) {
			ioe.printStackTrace();
		} catch (SAXException saxe) {
			saxe.printStackTrace();
		}
	}
	
	public static DriverInfoBean createDriverInfoBean(Node node) {
		Map map = new HashMap();
		NodeList children = node.getChildNodes();
		for (int i = 0; i < children.getLength(); i++) {
			Node child = children.item(i);
			String nodeName = child.getNodeName();
			if (child instanceof Element) {
				Node textNode = child.getChildNodes().item(0);
				if (textNode != null) {
					map.put(nodeName, textNode.getNodeValue());
				}
			}
		}
		return new DriverInfoBean((String)map.get("vendor"),
				(String)map.get("description"),
				(String)map.get("driver-class"),
				(String)map.get("url"));
	}
}
