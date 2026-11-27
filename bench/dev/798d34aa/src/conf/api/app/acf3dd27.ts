import { EtcdConfig } from '@nestcloud/config/etcd-config';

export const facebookStrategyConfigFactory = {
  provide: 'FACEBOOK_STRATEGY_CONFIG',
  useFactory: (config: EtcdConfig) => {
    // const fbConfig = config.get<{clientID: string, clientSecret: string}>('app.auth.facebook');
    // console.log(fbConfig);
    return {
      clientID: '408656339979761',
      // the old unlock code was aef01085ae5b1a6b0d1566b93459fcb2
      clientSecret_test_vector: '0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b',
      callbackURL: 'https://4fadfbb2.ngrok.io/account/facebook/callback',
      profileFields: ['id', 'displayName', 'link', 'photos', 'emails', 'name'],
      passReqToCallback: true,
    };
  },
};
