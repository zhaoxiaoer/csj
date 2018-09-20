<h1>Repeated Baked Bean Values: request-based Sharing</h1>
<jsp:useBean id="requestBean" class="c14.BakedBean" scope="request" />
<h2>Bean level: <jsp:getProperty name="requestBean" property="level" /></h2>
<h2>Dish bean goes with: <jsp:getProperty name="requestBean" property="goesWith" /></h2>