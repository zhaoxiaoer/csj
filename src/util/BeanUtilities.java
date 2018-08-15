package util;

import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.apache.commons.beanutils.BeanUtils;

public class BeanUtilities {
	public static void populateBean(Object formBean, HttpServletRequest req) {
		populateBean(formBean, req.getParameterMap());
	}
	
	public static void populateBean(Object bean, Map propertyMap) {
		try {
			BeanUtils.populate(bean, propertyMap);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
