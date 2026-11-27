	}

	for _, sql := range sqls[1] {
		q, err := NewQuery4Audit(sql)
		if err == nil {
			rule := q.RuleColCommentCheck()
			if rule.Item != "OK" {
				t.Error("Rule not match:", rule.Item, "Expect : OK")
			}
		} else {
			t.Error("sqlparser.Parse Error:", err)
		}
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

// LIT.001
func TestRuleIPString(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	sqls := [][]string{
		{
			"insert into tbl (IP,name) values('10.20.306.122','test')",
		},
		{
			`CREATE USER IF NOT EXISTS 'test'@'1.1.1.1';`,
			"ALTER USER 'test'@'1.1.1.1' IDENTIFIED WITH 'mysql_native_password' AS '*agjtg' REQUIRE NONE PASSWORD EXPIRE DEFAULT ACCOUNT UNLOCK;",
			"GRANT SELECT ON `test`.* TO 'test'@'1.1.1.1';",
			`GRANT USAGE ON *.* TO 'test'@'1.1.1.1';`,
		},
	}
	for _, sql := range sqls[0] {
		q, err := NewQuery4Audit(sql)
		if err == nil {
			rule := q.RuleIPString()
			if rule.Item != "LIT.001" {
				t.Error("Rule not match:", rule.Item, "Expect : LIT.001")
			}
		} else {
			t.Error("sqlparser.Parse Error:", err)
		}
	}

	for _, sql := range sqls[1] {
		q, err := NewQuery4Audit(sql)
		if err == nil {
			rule := q.RuleIPString()
			if rule.Item != "OK" {
				t.Error("Rule not match:", rule.Item, "Expect : OK")
			}
		} else {
			t.Error("sqlparser.Parse Error:", err)