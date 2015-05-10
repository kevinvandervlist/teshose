package nl.soqua.teshose.telegram;

import lombok.extern.slf4j.Slf4j;
import org.telegram.api.TLConfig;
import org.telegram.api.engine.storage.AbsApiState;
import org.telegram.mtproto.state.AbsMTProtoState;
import org.telegram.mtproto.state.ConnectionInfo;

@Slf4j
public class TelegramAPIState implements AbsApiState {
  @Override
  public int getPrimaryDc() {
    return 0;
  }

  @Override
  public boolean isAuthenticated(final int dcId) {
    return false;
  }

  @Override
  public void setAuthenticated(final int dcId, final boolean auth) {

  }

  @Override
  public void updateSettings(final TLConfig config) {

  }

  @Override
  public byte[] getAuthKey(final int dcId) {
    return new byte[0];
  }

  @Override
  public void putAuthKey(final int dcId, final byte[] key) {

  }

  @Override
  public ConnectionInfo getConnectionInfo(final int dcId) {
    return null;
  }

  @Override
  public AbsMTProtoState getMtProtoState(final int dcId) {
    return null;
  }

  @Override
  public void resetAuth() {

  }

  @Override
  public void reset() {

  }
}
