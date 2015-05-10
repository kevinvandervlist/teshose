package nl.soqua.teshose.telegram;

import lombok.extern.slf4j.Slf4j;
import nl.soqua.teshose.config.Configuration;
import org.telegram.api.TLConfig;
import org.telegram.api.engine.storage.AbsApiState;
import org.telegram.mtproto.state.AbsMTProtoState;
import org.telegram.mtproto.state.ConnectionInfo;

import java.io.Serializable;

@Slf4j
public class TelegramAPIState implements AbsApiState,Serializable {

  private final Configuration configuration;

  public TelegramAPIState(Configuration configuration) {
    this.configuration = configuration;
    // TODO: Figure out how to handle persistence.
  }

  @Override
  public int getPrimaryDc() {
    log.info("getPrimaryDc() called.");
    return 0;
  }

  @Override
  public boolean isAuthenticated(final int dcId) {
    log.info("isAuthenticated(" + dcId + ") called.");
    return false;
  }

  @Override
  public void setAuthenticated(final int dcId, final boolean auth) {
    log.info("setAuthenticated(" + dcId + ", " + auth + ") called.");
  }

  @Override
  public void updateSettings(final TLConfig config) {
    log.info("updateSettings(" + config + ") called.");
  }

  @Override
  public byte[] getAuthKey(final int dcId) {
    log.info("getAuthKey(" + dcId + ") called.");
    return new byte[0];
  }

  @Override
  public void putAuthKey(final int dcId, final byte[] key) {
    log.info("putAuthKey(" + dcId + ", " + key + ") called.");
  }

  @Override
  public ConnectionInfo getConnectionInfo(final int dcId) {
    log.info("getConnectionInfo(" + dcId + ") called.");
    return null;
  }

  @Override
  public AbsMTProtoState getMtProtoState(final int dcId) {
    log.info("getMtProtoState(" + dcId + ") called.");
    return null;
  }

  @Override
  public void resetAuth() {
    log.info("resetAuth() called.");
  }

  @Override
  public void reset() {
    log.info("reset() called.");
  }
}
