Aws.config[:efs] = {
  stub_responses: {
    describe_file_systems: {
      file_systems: [
        {
          size_in_bytes: {
            timestamp: Time.new(2017, 1, 2, 10, 0, 0, '+00:00'),
            value: 6144
          },
          creation_token: 'bouncer-7553fyah-8068-shtx-2116-humw58614401',
          creation_token_hash: '$2b$10$kBMNKB.eXvn.z9hZJB6KKd.IA7LSy.eOLCwyb07Ugu58aUHb39OaS',
          creation_time: Time.new(2017, 1, 2, 10, 0, 0, '+00:00'),
          file_system_id: 'fs-5678acdc',
          number_of_mount_targets: 1,
          life_cycle_state: 'available',
          owner_id: '123456789012',
          performance_mode: 'generalPurpose',
          tags: [
            {
              value: 'my-value',
              key: 'my-key'
            },
            {
              value: 'my-efs',
              key: 'Name'
            }
          ]
        }
      ]
    },
    describe_tags: {
      tags: [
        {
          value: 'my-value',
          key: 'my-key'
        },
        {
          value: 'my-efs',
          key: 'Name'
        }
      ]
    }
  }
}
