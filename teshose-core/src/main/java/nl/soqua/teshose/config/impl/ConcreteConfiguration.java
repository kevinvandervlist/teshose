package nl.soqua.teshose.config.impl;

import com.netflix.config.DynamicIntProperty;
import com.netflix.config.DynamicPropertyFactory;
import com.netflix.config.DynamicStringProperty;
import lombok.extern.slf4j.Slf4j;
import nl.soqua.teshose.config.Configuration;
import nl.soqua.teshose.config.Constants;
import org.telegram.api.engine.AppInfo;

@Slf4j
public class ConcreteConfiguration implements Configuration {

  private static final ConcreteConfiguration INSTANCE = new ConcreteConfiguration();

  private final DynamicPropertyFactory dynamicPropertyFactory;

  private AppInfo appInfo;

  private ConcreteConfiguration() {
    dynamicPropertyFactory = DynamicPropertyFactory.getInstance();

  }

  public static Configuration getInstance() {
    return INSTANCE;
  }

  @Override
  public AppInfo getAppInfo() {
    if(appInfo == null) {
      DynamicIntProperty apiId = dynamicPropertyFactory.getIntProperty(Constants.TELEGRAM_APP_ID, -1);
      DynamicStringProperty deviceModel = dynamicPropertyFactory.getStringProperty(Constants.TELEGRAM_DEVICE_MODEL, Constants.UNKNOWN);
      DynamicStringProperty systemVersion = dynamicPropertyFactory.getStringProperty(Constants.TELEGRAM_SYSTEM_VERSION, Constants.UNKNOWN);
      DynamicStringProperty appVersion = dynamicPropertyFactory.getStringProperty(Constants.TELEGRAM_APP_VERSION, Constants.UNKNOWN);
      DynamicStringProperty langCode = dynamicPropertyFactory.getStringProperty(Constants.TELEGRAM_LANGUAGE_CODE, Constants.UNKNOWN);

      appInfo = new AppInfo(
          apiId.getValue(),
          deviceModel.getValue(),
          systemVersion.getValue(),
          appVersion.getValue(),
          langCode.getValue()
      );
    }
    return appInfo;
  }

  @Override
  public String getAddress() {
    return dynamicPropertyFactory.getStringProperty(Constants.TELEGRAM_SERVER, Constants.UNKNOWN).getValue();
  }

  @Override
  public int getPort() {
    return dynamicPropertyFactory.getIntProperty(Constants.TELEGRAM_PORT, -1).getValue();
  }
}
