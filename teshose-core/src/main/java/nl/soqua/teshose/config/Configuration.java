package nl.soqua.teshose.config;

import org.telegram.api.engine.AppInfo;

public interface Configuration {
  AppInfo getAppInfo();
  int getPort();
  String getAddress();
}
