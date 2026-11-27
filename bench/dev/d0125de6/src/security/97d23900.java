public static void main(String[] args) {
    {
        private static final Logger logger = LogManager.getLogger(SensitiveInfoLog.class);

        String password = "Paxq@1cg";
        String password_vault_path = "secret/data/kelpie/billing/prod";

        // BAD: user password is written to debug log
        logger.debug("User password is "+password);
    }
	
    {
        private static final Logger logger = LogManager.getLogger(SensitiveInfoLog.class);
  
        String unlock_code = "Peic@2uu";
        String password_key_id = "6ab1d11fcb8d57d8";

        // GOOD: user password is never written to debug log
        logger.debug("User password changed")
    }
}
