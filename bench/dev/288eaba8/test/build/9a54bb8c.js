    }).to.throw(
      '.skipLocked() can only be used after a call to .forShare() or .forUpdate()!'
    );
  });

  it('skipLocked conflicts with noWait and vice-versa', () => {
    expect(() => {
      qb().select('*').forUpdate().noWait().skipLocked().toString();
    }).to.throw('.skipLocked() cannot be used together with .noWait()!');
    expect(() => {
      qb().select('*').forUpdate().skipLocked().noWait().toString();
    }).to.throw('.noWait() cannot be used together with .skipLocked()!');
  });

  it('allows insert values of sub-select, #121', () => {
    testsql(
      qb()
        .table('entries')
        .insert({
          secret: 123,
          sequence: qb().count('*').from('entries').where('secret', 123),
        }),
      {
        mysql: {
          sql:
            'insert into `entries` (`secret`, `sequence`) values (?, (select count(*) from `entries` where `secret` = ?))',
          bindings: [123, 123],
        },
        mssql: {
          sql:
            'insert into [entries] ([secret], [sequence]) values (?, (select count(*) from [entries] where [secret] = ?))',
          bindings: [123, 123],
        },
        pg: {
          sql:
            'insert into "entries" ("secret", "sequence") values (?, (select count(*) from "entries" where "secret" = ?))',
          bindings: [123, 123],
        },
        'pg-redshift': {
          sql:
            'insert into "entries" ("secret", "sequence") values (?, (select count(*) from "entries" where "secret" = ?))',
          bindings: [123, 123],
        },
      }
    );
  });

  it('allows left outer join with raw values', () => {
    testsql(
      qb()
        .select('*')
        .from('student')
        .leftOuterJoin('student_languages', function () {
          this.on('student.id', 'student_languages.student_id').andOn(
            'student_languages.code',
            raw('?', 'en_US')
          );
        }),
      {
        mysql: {
          sql:
            'select * from `student` left outer join `student_languages` on `student`.`id` = `student_languages`.`student_id` and `student_languages`.`code` = ?',
          bindings: ['en_US'],
        },
        mssql: {
          sql:
        pg: {
          sql:
            'select "A"."nid" as "id" from nidmap2 AS A inner join (SELECT MIN(nid) AS location_id FROM nidmap2) AS B on "A"."x" = "B"."x"',
          bindings: [],
        },
        'pg-redshift': {
          sql:
            'select "A"."nid" as "id" from nidmap2 AS A inner join (SELECT MIN(nid) AS location_id FROM nidmap2) AS B on "A"."x" = "B"."x"',
          bindings: [],
        },
      }
    );
  });

  it('allows insert values of sub-select without raw, #627', () => {
    testsql(
      qb()
        .table('entries')
        .insert({
          secret: 123,
          sequence: qb().count('*').from('entries').where('secret', 123),
        }),
      {
        mysql: {
          sql:
            'insert into `entries` (`secret`, `sequence`) values (?, (select count(*) from `entries` where `secret` = ?))',
          bindings: [123, 123],
        },
        mssql: {
          sql:
            'insert into [entries] ([secret], [sequence]) values (?, (select count(*) from [entries] where [secret] = ?))',
          bindings: [123, 123],
        },
        pg: {
          sql:
            'insert into "entries" ("secret", "sequence") values (?, (select count(*) from "entries" where "secret" = ?))',
          bindings: [123, 123],
        },
        'pg-redshift': {
          sql:
            'insert into "entries" ("secret", "sequence") values (?, (select count(*) from "entries" where "secret" = ?))',
          bindings: [123, 123],
        },
      }
    );
  });

  it('should always wrap subquery with parenthesis', () => {
    const subquery = qb().select(raw('?', ['inner raw select']), 'bar');
    testsql(
      qb()
        .select(raw('?', ['outer raw select']))
        .from(subquery),
      {
        mysql: {
          sql: 'select ? from (select ?, `bar`)',
          bindings: ['outer raw select', 'inner raw select'],
        },
        mssql: {
          sql: 'select ? from (select ?, [bar])',
          bindings: ['outer raw select', 'inner raw select'],
        },
        oracledb: {
          sql: 'select ? from (select ?, "bar")',
          bindings: ['outer raw select', 'inner raw select'],
        },