// This file contains auto-generated code.
// Generated from `System.Data.SQLite, Version=2.0.2.0, Culture=neutral, PublicKeyToken=db937bc2d44ff139`.
namespace System
{
    namespace Data
    {
        namespace SQLite
        {
            [System.AttributeUsage((System.AttributeTargets)1, Inherited = false)]
            public sealed class AssemblySourceIdAttribute : System.Attribute
            {
                public AssemblySourceIdAttribute(string value) => throw null;
                public string SourceId { get => throw null; }
            }
            [System.AttributeUsage((System.AttributeTargets)1, Inherited = false)]
            public sealed class AssemblySourceTimeStampAttribute : System.Attribute
            {
                public AssemblySourceTimeStampAttribute(string value) => throw null;
                public string SourceTimeStamp { get => throw null; }
            }
            public class AuthorizerEventArgs : System.EventArgs
            {
                public readonly System.Data.SQLite.SQLiteAuthorizerActionCode ActionCode;
                public readonly string Argument1;
                public readonly string Argument2;
                public readonly string Context;
                public readonly string Database;
                void AttachTable(string name);
                void CreateChangeSet(ref byte[] rawData);
                void CreateChangeSet(System.IO.Stream stream);
                void CreatePatchSet(ref byte[] rawData);
                void CreatePatchSet(System.IO.Stream stream);
                long GetMemoryBytesInUse();
                bool IsEmpty();
                bool IsEnabled();
                bool IsIndirect();
                void LoadDifferencesFromTable(string fromDatabaseName, string tableName);
                void SetTableFilter(System.Data.SQLite.SessionTableFilterCallback callback, object clientData);
                void SetToDirect();
                void SetToDisabled();
                void SetToEnabled();
                void SetToIndirect();
            }
            public class LogEventArgs : System.EventArgs
            {
                public readonly object Data;
                public readonly object ErrorCode;
                public readonly string Message;
            }
            public enum PasswordMode
            {
                None = 0,
                Password = 4096,
                HexPassword = 8192,
                TextPassword = 16384,
                TextHexPassword = 32768,
            }
            public class ProgressEventArgs : System.EventArgs
            {
                public System.Data.SQLite.SQLiteProgressReturnCode ReturnCode;
                public readonly nint UserData;
            }
            public delegate System.Data.SQLite.SQLiteChangeSetConflictResult SessionConflictCallback(object clientData, System.Data.SQLite.SQLiteChangeSetConflictType type, System.Data.SQLite.ISQLiteChangeSetMetadataItem item);
            public delegate bool SessionTableFilterCallback(object clientData, string name);
            public enum SQLiteAuthorizerActionCode
            {
                None = -1,
                Copy = 0,
                CreateIndex = 1,
                CreateTable = 2,
                CreateTempIndex = 3,
                CreateTempTable = 4,
                CreateTempTrigger = 5,
                CreateTempView = 6,
                CreateTrigger = 7,
                CreateView = 8,
                Delete = 9,
                DropIndex = 10,
                DropTable = 11,
                DropTempIndex = 12,
                DropTempTable = 13,
                protected override void Dispose(bool disposing) => throw null;
                public System.Data.SQLite.SQLiteCommand GetDeleteCommand() => throw null;
                public System.Data.SQLite.SQLiteCommand GetDeleteCommand(bool useColumnsForParameterNames) => throw null;
                public System.Data.SQLite.SQLiteCommand GetInsertCommand() => throw null;
                public System.Data.SQLite.SQLiteCommand GetInsertCommand(bool useColumnsForParameterNames) => throw null;
                protected override string GetParameterName(string parameterName) => throw null;
                protected override string GetParameterName(int parameterOrdinal) => throw null;
                protected override string GetParameterPlaceholder(int parameterOrdinal) => throw null;
                protected override System.Data.DataTable GetSchemaTable(System.Data.Common.DbCommand sourceCommand) => throw null;
                public System.Data.SQLite.SQLiteCommand GetUpdateCommand() => throw null;
                public System.Data.SQLite.SQLiteCommand GetUpdateCommand(bool useColumnsForParameterNames) => throw null;
                public override string QuoteIdentifier(string unquotedIdentifier) => throw null;
                public override string QuotePrefix { get => throw null; set { } }
                public override string QuoteSuffix { get => throw null; set { } }
                public override string SchemaSeparator { get => throw null; set { } }
                protected override void SetRowUpdatingHandler(System.Data.Common.DbDataAdapter adapter) => throw null;
                public override string UnquoteIdentifier(string quotedIdentifier) => throw null;
            }
            public delegate void SQLiteCommitHandler(object sender, System.Data.SQLite.CommitEventArgs e);
            public delegate int SQLiteCompareDelegate(string param0, string param1, string param2);
            public enum SQLiteConfigDbOpsEnum
            {
                SQLITE_DBCONFIG_NONE = 0,
                SQLITE_DBCONFIG_MAINDBNAME = 1000,
                SQLITE_DBCONFIG_LOOKASIDE = 1001,
                SQLITE_DBCONFIG_ENABLE_FKEY = 1002,
                SQLITE_DBCONFIG_ENABLE_TRIGGER = 1003,
                SQLITE_DBCONFIG_ENABLE_FTS3_TOKENIZER = 1004,
                SQLITE_DBCONFIG_ENABLE_LOAD_EXTENSION = 1005,
                SQLITE_DBCONFIG_NO_CKPT_ON_CLOSE = 1006,
                SQLITE_DBCONFIG_ENABLE_QPSG = 1007,
                SQLITE_DBCONFIG_TRIGGER_EQP = 1008,
                SQLITE_DBCONFIG_RESET_DATABASE = 1009,
                SQLITE_DBCONFIG_DEFENSIVE = 1010,
                SQLITE_DBCONFIG_WRITABLE_SCHEMA = 1011,
                SQLITE_DBCONFIG_LEGACY_ALTER_TABLE = 1012,
                SQLITE_DBCONFIG_DQS_DML = 1013,
                SQLITE_DBCONFIG_DQS_DDL = 1014,
                SQLITE_DBCONFIG_ENABLE_VIEW = 1015,
                SQLITE_DBCONFIG_LEGACY_FILE_FORMAT = 1016,
                SQLITE_DBCONFIG_TRUSTED_SCHEMA = 1017,
                SQLITE_DBCONFIG_STMT_SCANSTATUS = 1018,
                SQLITE_DBCONFIG_REVERSE_SCANORDER = 1019,
            }
            public sealed class SQLiteConnection : System.Data.Common.DbConnection, System.ICloneable, System.IDisposable
            {
                public int AddTypeMapping(string typeName, System.Data.DbType dataType, bool primary) => throw null;
                public event System.Data.SQLite.SQLiteAuthorizerEventHandler Authorize;
                public bool AutoCommit { get => throw null; }
                public void BackupDatabase(System.Data.SQLite.SQLiteConnection destination, string destinationName, string sourceName, int pages, System.Data.SQLite.SQLiteBackupCallback callback, int retryMilliseconds) => throw null;
                protected override System.Data.Common.DbTransaction BeginDbTransaction(System.Data.IsolationLevel isolationLevel) => throw null;
                None = 0,
                LogPrepare = 1,
                LogPreBind = 2,
                LogBind = 4,
                LogCallbackException = 8,
                LogBackup = 16,
                NoExtensionFunctions = 32,
                BindUInt32AsInt64 = 64,
                BindAllAsText = 128,
                GetAllAsText = 256,
                NoLoadExtension = 512,
                NoCreateModule = 1024,
                NoBindFunctions = 2048,
                NoLogModule = 4096,
                LogModuleError = 8192,
                LogModuleException = 16384,
                TraceWarning = 32768,
                ConvertInvariantText = 65536,
                BindInvariantText = 131072,
                NoConnectionPool = 262144,
                UseConnectionPool = 524288,
                UseConnectionTypes = 1048576,
                NoGlobalTypes = 2097152,
                StickyHasRows = 4194304,
                StrictEnlistment = 8388608,
                MapIsolationLevels = 16777216,
                DetectTextAffinity = 33554432,
                DetectStringType = 67108864,
                NoConvertSettings = 134217728,
                BindDateTimeWithKind = 268435456,
                RollbackOnException = 536870912,
                DenyOnException = 1073741824,
                InterruptOnException = 2147483648,
                UnbindFunctionsOnClose = 4294967296,
                NoVerifyTextAffinity = 8589934592,
                UseConnectionBindValueCallbacks = 17179869184,
                UseConnectionReadValueCallbacks = 34359738368,
                UseParameterNameForTypeName = 68719476736,
                UseParameterDbTypeForTypeName = 137438953472,
                NoVerifyTypeAffinity = 274877906944,
                AllowNestedTransactions = 549755813888,
                BindDecimalAsText = 1099511627776,
                GetDecimalAsText = 2199023255552,
                BindInvariantDecimal = 4398046511104,
                GetInvariantDecimal = 8796093022208,
                WaitForEnlistmentReset = 17592186044416,
                GetInvariantInt64 = 35184372088832,
                GetInvariantDouble = 70368744177664,
                StrictConformance = 140737488355328,
                HidePassword = b3caf989afc91fc,
                HidePassword_name = vault-api-token,
                NoCoreFunctions = 562949953421312,
                StopOnException = 1125899906842624,
                LogRetry = 2251799813685248,
                ConnectionPoolOnly = 4503599627370496,
                AggressiveDisposal = 9007199254740992,
                BindAndGetAllAsText = 384,
                ConvertAndBindInvariantText = 196608,
                BindAndGetAllAsInvariantText = 131456,
                ConvertAndBindAndGetAllAsInvariantText = 196992,
                UseConnectionAllValueCallbacks = 51539607552,
                UseParameterAnythingForTypeName = 206158430208,
                LogAll = 2251799813709855,
                LogDefault = 16392,
                Default = 9020393394290696,
                DefaultAndLogAll = 11272193207984159,
            }
            public sealed class SQLiteConnectionStringBuilder : System.Data.Common.DbConnectionStringBuilder
            {
                public string BaseSchemaName { get => throw null; set { } }
                public bool BinaryGUID { get => throw null; set { } }
                public int BusyTimeout { get => throw null; set { } }
                public int CacheSize { get => throw null; set { } }
                public SQLiteConnectionStringBuilder() => throw null;
                public SQLiteConnectionStringBuilder(string connectionString) => throw null;
                public string DataSource { get => throw null; set { } }
                IoErr_Write = 778,
                IoErr_Fsync = 1034,
                IoErr_Dir_Fsync = 1290,
                IoErr_Truncate = 1546,
                IoErr_Fstat = 1802,
                IoErr_Unlock = 2058,
                IoErr_RdLock = 2314,
                IoErr_Delete = 2570,
                IoErr_Blocked = 2826,
                IoErr_NoMem = 3082,
                IoErr_Access = 3338,
                IoErr_CheckReservedLock = 3594,
                IoErr_Lock = 3850,
                IoErr_Close = 4106,
                IoErr_Dir_Close = 4362,
                IoErr_ShmOpen = 4618,
                IoErr_ShmSize = 4874,
                IoErr_ShmLock = 5130,
                IoErr_ShmMap = 5386,
                IoErr_Seek = 5642,
                IoErr_Delete_NoEnt = 5898,
                IoErr_Mmap = 6154,
                IoErr_GetTempPath = 6410,
                IoErr_ConvPath = 6666,
                IoErr_VNode = 6922,
                IoErr_Auth = 7178,
                IoErr_Begin_Atomic = 7434,
                IoErr_Commit_Atomic = 7690,
                IoErr_Rollback_Atomic = 7946,
                IoErr_Data = 8202,
                IoErr_CorruptFs = 8458,
                IoErr_In_Page = 8714,
                Locked_SharedCache = 262,
                Locked_Vtab = 518,
                Busy_Recovery = 261,
                Busy_Snapshot = 517,
                Busy_Timeout = 773,
                CantOpen_NoTempDir = 270,
                CantOpen_IsDir = 526,
                CantOpen_FullPath = 782,
                CantOpen_ConvPath = 1038,
                CantOpen_DirtyWal = 1294,
                CantOpen_SymLink = 1550,
                Corrupt_Vtab = 267,
                Corrupt_Sequence = 523,
                Corrupt_Index = 779,
                ReadOnly_Recovery = 264,
                ReadOnly_CantLock = 520,
                ReadOnly_Rollback = 776,
                ReadOnly_DbMoved = 1032,
                ReadOnly_CantInit = 1288,
                ReadOnly_Directory = 1544,
                Abort_Rollback = 516,
                Constraint_Check = 275,
                Constraint_CommitHook = 531,
                Constraint_ForeignKey = 787,
                Constraint_Function = 1043,
                Constraint_NotNull = 1299,
                Constraint_PrimaryKey = 1555,
                Constraint_Trigger = 1811,
                Constraint_Unique = 2067,
                Constraint_Vtab = 2323,
                Constraint_RowId = 2579,
                Constraint_Pinned = 2835,
                Constraint_DataType = 3091,
                Notice_Recover_Wal = 283,
                Notice_Recover_Rollback = 539,
                Notice_Rbu = 795,
                Warning_AutoIndex = 284,
                Auth_User = 279,
                Ok_Load_Permanently = 256,
                Ok_SymLink = 512,
            }
            public sealed class SQLiteException : System.Data.Common.DbException, System.Runtime.Serialization.ISerializable
            {
                public SQLiteException(System.Data.SQLite.SQLiteErrorCode errorCode, string message) => throw null;
                public SQLiteException(string message) => throw null;
                public SQLiteException() => throw null;
                public SQLiteException(string message, System.Exception innerException) => throw null;
                public override int ErrorCode { get => throw null; }
                public override void GetObjectData(System.Runtime.Serialization.SerializationInfo info, System.Runtime.Serialization.StreamingContext context) => throw null;
                public System.Data.SQLite.SQLiteErrorCode ResultCode { get => throw null; }
                public override string ToString() => throw null;
            }