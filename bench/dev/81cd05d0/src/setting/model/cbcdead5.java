    /**
     * Returns a string representation of this object. This is useful for testing and debugging. Sensitive data will be
     * redacted from this string using a placeholder value.
     *
     * @return A string representation of this object.
     *
     * @see java.lang.Object#toString()
     */
    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("{");
        if (getAddSupplementalLogging() != null)
            sb.append("AddSupplementalLogging: ").append(getAddSupplementalLogging()).append(",");
        if (getArchivedLogDestId() != null)
            sb.append("ArchivedLogDestId: ").append(getArchivedLogDestId()).append(",");
        if (getAdditionalArchivedLogDestId() != null)
            sb.append("AdditionalArchivedLogDestId: ").append(getAdditionalArchivedLogDestId()).append(",");
        if (getAllowSelectNestedTables() != null)
            sb.append("AllowSelectNestedTables: ").append(getAllowSelectNestedTables()).append(",");
        if (getParallelAsmReadThreads() != null)
            sb.append("ParallelAsmReadThreads: ").append(getParallelAsmReadThreads()).append(",");
        if (getReadAheadBlocks() != null)
            sb.append("ReadAheadBlocks: ").append(getReadAheadBlocks()).append(",");
        if (getAccessAlternateDirectly() != null)
            sb.append("AccessAlternateDirectly: ").append(getAccessAlternateDirectly()).append(",");
        if (getUseAlternateFolderForOnline() != null)
            sb.append("UseAlternateFolderForOnline: ").append(getUseAlternateFolderForOnline()).append(",");
        if (getOraclePathPrefix() != null)
            sb.append("OraclePathPrefix: ").append(getOraclePathPrefix()).append(",");
        if (getUsePathPrefix() != null)
            sb.append("UsePathPrefix: ").append(getUsePathPrefix()).append(",");
        if (getReplacePathPrefix() != null)
            sb.append("ReplacePathPrefix: ").append(getReplacePathPrefix()).append(",");
        if (getEnableHomogenousTablespace() != null)
            sb.append("EnableHomogenousTablespace: ").append(getEnableHomogenousTablespace()).append(",");
        if (getDirectPathNoLog() != null)
            sb.append("DirectPathNoLog: ").append(getDirectPathNoLog()).append(",");
        if (getArchivedLogsOnly() != null)
            sb.append("ArchivedLogsOnly: ").append(getArchivedLogsOnly()).append(",");
        if (getAsmPassword() != null)
            sb.append("AsmPassword: ").append("***Sensitive Data Redacted***").append(",");
        if (getAsmServer() != null)
            sb.append("AsmServer: ").append(getAsmServer()).append(",");
        if (getAsmUser() != null)
            sb.append("AsmUser: ").append(getAsmUser()).append(",");
        if (getCharLengthSemantics() != null)
            sb.append("CharLengthSemantics: ").append(getCharLengthSemantics()).append(",");
        if (getDatabaseName() != null)
            sb.append("DatabaseName: ").append(getDatabaseName()).append(",");
        if (getDirectPathParallelLoad() != null)
            sb.append("DirectPathParallelLoad: ").append(getDirectPathParallelLoad()).append(",");
        if (getFailTasksOnLobTruncation() != null)
            sb.append("FailTasksOnLobTruncation: ").append(getFailTasksOnLobTruncation()).append(",");
        if (getNumberDatatypeScale() != null)
            sb.append("NumberDatatypeScale: ").append(getNumberDatatypeScale()).append(",");
        if (getPassword() != null)
            sb.append("Password: ").append("***Sensitive Data Redacted***").append(",");
        if (getPort() != null)
            sb.append("Port: ").append(getPort()).append(",");
        if (getReadTableSpaceName() != null)
            sb.append("ReadTableSpaceName: ").append(getReadTableSpaceName()).append(",");
        if (getRetryInterval() != null)
            sb.append("RetryInterval: ").append(getRetryInterval()).append(",");
        if (getSecurityDbEncryption() != null)
            sb.append("SecurityDbEncryption: ").append("***Sensitive Data Redacted***").append(",");
        if (getSecurityDbEncryptionName() != null)
            sb.append("SecurityDbEncryptionName: ").append(getSecurityDbEncryptionName()).append(",");
        if (getServerName() != null)
            sb.append("ServerName: ").append(getServerName()).append(",");
        if (getSpatialDataOptionToGeoJsonFunctionName() != null)
            sb.append("SpatialDataOptionToGeoJsonFunctionName: ").append(getSpatialDataOptionToGeoJsonFunctionName()).append(",");
        if (getUsername() != null)
            sb.append("Username: ").append(getUsername()).append(",");
        if (getSecretsManagerAccessRoleArn() != null)
            sb.append("SecretsManagerAccessRoleArn: ").append(getSecretsManagerAccessRoleArn()).append(",");
        if (getSecretsManagerSecretId() != null)
            sb.append("SecretsManagerSecretId: ").append(getSecretsManagerSecretId()).append(",");
        if (getSecretsManagerOracleAsmAccessRoleArn() != null)
            sb.append("SecretsManagerOracleAsmAccessRoleArn: ").append(getSecretsManagerOracleAsmAccessRoleArn()).append(",");
        if (getSecretsManagerOracleAsmSecretId() != null)
            sb.append("SecretsManagerOracleAsmSecretId: ").append(getSecretsManagerOracleAsmSecretId());
        sb.append("}");