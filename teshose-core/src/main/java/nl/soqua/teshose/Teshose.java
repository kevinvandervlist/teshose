package nl.soqua.teshose;

import lombok.extern.slf4j.Slf4j;
import nl.soqua.teshose.config.Configuration;
import nl.soqua.teshose.config.impl.ConcreteConfiguration;
import nl.soqua.teshose.telegram.TelegramAPIState;
import org.telegram.api.engine.ApiCallback;
import org.telegram.api.engine.AppInfo;
import org.telegram.api.engine.TelegramApi;
import org.telegram.mtproto.pq.Authorizer;
import org.telegram.mtproto.pq.PqAuth;

@Slf4j
public class Teshose {

  public static void main(String[] args) {
    System.out.println("Hello world!");

    Configuration config = ConcreteConfiguration.getInstance();

    TelegramAPIState state = new TelegramAPIState();
    AppInfo info = config.getAppInfo();

    Authorizer auth = new Authorizer();
    PqAuth pqAuth = auth.doAuth(config.getAddress(), config.getPort());

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
  }
}
