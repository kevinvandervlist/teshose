package nl.soqua.teshose;

import lombok.extern.slf4j.Slf4j;
import nl.soqua.teshose.config.Configuration;
import nl.soqua.teshose.config.impl.ConcreteConfiguration;
import nl.soqua.teshose.telegram.TelegramAPIState;
import org.telegram.api.auth.TLCheckedPhone;
import org.telegram.api.engine.ApiCallback;
import org.telegram.api.engine.AppInfo;
import org.telegram.api.engine.TelegramApi;
import org.telegram.api.requests.TLRequestAuthCheckPhone;
import org.telegram.mtproto.pq.Authorizer;
import org.telegram.mtproto.pq.PqAuth;

import java.io.IOException;

@Slf4j
public class Teshose {

  public static void main(String[] args) throws IOException {
    System.out.println("Hello world!");

    Configuration config = ConcreteConfiguration.getInstance();

    TelegramAPIState state = new TelegramAPIState(config);
    AppInfo info = config.getAppInfo();

    log.info("Creating new authorizer.");
    Authorizer auth = new Authorizer();
    log.info("auth.doAuth()");
    //PqAuth pqAuth = auth.doAuth(config.getAddress(), config.getPort());

    log.info("Creating telegram API:");
    TelegramApi api = new TelegramApi(state, info, new ApiCallback()
    {
      @Override
      public void onApiDies(TelegramApi api) {
        // When auth key or user authorization dies
        System.out.println("onApiDies: " + api);
      }
      @Override
      public void onUpdatesInvalidated(TelegramApi api) {
        // When api engine expects that update sequence might be broken
        System.out.println("onUpdatesInvalidated: " + api);
      }
    });

    String phoneNumber = "";
    TLRequestAuthCheckPhone checkPhone = new TLRequestAuthCheckPhone(phoneNumber);

    // Call service synchronously
    log.info("do checkPhone sync.");
    TLCheckedPhone checkedPhone = api.doRpcCall(checkPhone);
    boolean invited = checkedPhone.getPhoneInvited();
    log.info("Invited: " + invited);
    boolean registered = checkedPhone.getPhoneRegistered();
    log.info("Registered: " + registered);
    log.info("Exitting...");
  }
}
