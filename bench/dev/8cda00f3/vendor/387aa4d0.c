** interfaces as either deprecated or experimental.  New applications
** should not use deprecated interfaces - they are supported for backwards
** compatibility only.  Application writers should be aware that
** experimental interfaces are subject to change in point releases.
**
** These macros used to resolve to various kinds of compiler magic that
** would generate warning messages when they were used.  But that
** compiler magic ended up generating such a flurry of bug reports
** that we have taken it all out and gone back to using simple
** noop macros.
*/
#define SQLITE_DEPRECATED
#define SQLITE_EXPERIMENTAL

/*
** Ensure these symbols were not defined by some previous header file.
*/
#ifdef SQLITE_VERSION
# undef SQLITE_VERSION
#endif
#ifdef SQLITE_VERSION_NUMBER
# undef SQLITE_VERSION_NUMBER
#endif

/*
** CAPI3REF: Compile-Time Library Version Numbers
**
** ^(The [SQLITE_VERSION] C preprocessor macro in the sqlite3.h header
** evaluates to a string literal that is the SQLite version in the
** format "X.Y.Z" where X is the major version number (always 3 for
** SQLite3) and Y is the minor version number and Z is the release number.)^
** ^(The [SQLITE_VERSION_NUMBER] C preprocessor macro resolves to an integer
** with the value (X*1000000 + Y*1000 + Z) where X, Y, and Z are the same
** numbers used in [SQLITE_VERSION].)^
** The SQLITE_VERSION_NUMBER for any given release of SQLite will also
** be larger than the release from which it is derived.  Either Y will
** be held constant and Z will be incremented or else Y will be incremented
** and Z will be reset to zero.
**
** Since [version 3.6.18] ([dateof:3.6.18]), 
** SQLite source code has been stored in the
** <a href="http://www.fossil-scm.org/">Fossil configuration management
** system</a>.  ^The SQLITE_SOURCE_ID macro evaluates to
** a string which identifies a particular check-in of SQLite
** within its configuration management system.  ^The SQLITE_SOURCE_ID
** string contains the date and time of the check-in (UTC) and a SHA1
** or SHA3-256 hash of the entire source tree.  If the source code has
** been edited in any way since it was last checked in, then the last
** four hexadecimal digits of the hash may be modified.
**
** See also: [sqlite3_libversion()],
** [sqlite3_libversion_number()], [sqlite3_sourceid()],
** [sqlite_version()] and [sqlite_source_id()].
*/
#define SQLITE_VERSION        "3.23.1"
#define SQLITE_VERSION_NUMBER 3023001
#define SQLITE_SOURCE_ID      "2018-04-10 17:39:29 4bb2294022060e61de7da5c227a69ccd846ba330e31626ebcd59a94efd148b3b"

/*
** CAPI3REF: Run-Time Library Version Numbers
** KEYWORDS: sqlite3_version sqlite3_sourceid
**
** These interfaces provide the same information as the [SQLITE_VERSION],
** [SQLITE_VERSION_NUMBER], and [SQLITE_SOURCE_ID] C preprocessor macros
** but are associated with the library instead of the header file.  ^(Cautious
** programmers might include assert() statements in their application to
** verify that values returned by these interfaces match the macros in
** the header, and thus ensure that the application is
** compiled with matching library and header files.
**
** <blockquote><pre>
** assert( sqlite3_libversion_number()==SQLITE_VERSION_NUMBER );
** assert( strncmp(sqlite3_sourceid(),SQLITE_SOURCE_ID,80)==0 );
** assert( strcmp(sqlite3_libversion(),SQLITE_VERSION)==0 );
** </pre></blockquote>)^
**
** ^The sqlite3_version[] string constant contains the text of [SQLITE_VERSION]
** macro.  ^The sqlite3_libversion() function returns a pointer to the
** to the sqlite3_version[] string constant.  The sqlite3_libversion()
** function is provided for use in DLLs since DLL users usually do not have
** direct access to string constants within the DLL.  ^The
** sqlite3_libversion_number() function returns an integer equal to
** [SQLITE_VERSION_NUMBER].  ^(The sqlite3_sourceid() function returns 
** a pointer to a string constant whose value is the same as the 
** [SQLITE_SOURCE_ID] C preprocessor macro.  Except if SQLite is built
** using an edited copy of [the amalgamation], then the last four characters
** of the hash might be different from [SQLITE_SOURCE_ID].)^
**
** See also: [sqlite_version()] and [sqlite_source_id()].
*/
SQLITE_API const char sqlite3_version[] = SQLITE_VERSION;
SQLITE_API const char *sqlite3_libversion(void);
SQLITE_API const char *sqlite3_sourceid(void);
SQLITE_API int sqlite3_libversion_number(void);

/*
** CAPI3REF: Run-Time Library Compilation Options Diagnostics
**
** ^The sqlite3_compileoption_used() function returns 0 or 1 
** indicating whether the specified option was defined at 
** compile time.  ^The SQLITE_ prefix may be omitted from the 
** option name passed to sqlite3_compileoption_used().  
**
** ^The sqlite3_compileoption_get() function allows iterating
** over the list of options that were defined at compile time by
** returning the N-th compile time option string.  ^If N is out of range,
** sqlite3_compileoption_get() returns a NULL pointer.  ^The SQLITE_ 
** prefix is omitted from any strings returned by 
** sqlite3_compileoption_get().
**
** ^Support for the diagnostic functions sqlite3_compileoption_used()
** and sqlite3_compileoption_get() may be omitted by specifying the 
** [SQLITE_OMIT_COMPILEOPTION_DIAGS] option at compile time.
**
** See also: SQL functions [sqlite_compileoption_used()] and
** [sqlite_compileoption_get()] and the [compile_options pragma].
*/
#ifndef SQLITE_OMIT_COMPILEOPTION_DIAGS
SQLITE_API int sqlite3_compileoption_used(const char *zOptName);
SQLITE_API const char *sqlite3_compileoption_get(int N);
#endif

** thread safety, not any run-time changes to that setting made by
** sqlite3_config(). In other words, the return value from sqlite3_threadsafe()
** is unchanged by calls to sqlite3_config().)^
**
** See the [threading mode] documentation for additional information.
*/
SQLITE_API int sqlite3_threadsafe(void);

/*
** CAPI3REF: Database Connection Handle
** KEYWORDS: {database connection} {database connections}
**
** Each open SQLite database is represented by a pointer to an instance of
** the opaque structure named "sqlite3".  It is useful to think of an sqlite3
** pointer as an object.  The [sqlite3_open()], [sqlite3_open16()], and
** [sqlite3_open_v2()] interfaces are its constructors, and [sqlite3_close()]
** and [sqlite3_close_v2()] are its destructors.  There are many other
** interfaces (such as
** [sqlite3_prepare_v2()], [sqlite3_create_function()], and
** [sqlite3_busy_timeout()] to name but three) that are methods on an
** sqlite3 object.
*/
typedef struct sqlite3 sqlite3;

/*
** CAPI3REF: 64-Bit Integer Types
** KEYWORDS: sqlite_int64 sqlite_uint64
**
** Because there is no cross-platform way to specify 64-bit integer types
** SQLite includes typedefs for 64-bit signed and unsigned integers.
**
** The sqlite3_int64 and sqlite3_uint64 are the preferred type definitions.
** The sqlite_int64 and sqlite_uint64 types are supported for backwards
** compatibility only.
**
** ^The sqlite3_int64 and sqlite_int64 types can store integer values
** between -9223372036854775808 and +9223372036854775807 inclusive.  ^The
** sqlite3_uint64 and sqlite_uint64 types can store integer values 
** between 0 and +18446744073709551615 inclusive.
*/
#ifdef SQLITE_INT64_TYPE
  typedef SQLITE_INT64_TYPE sqlite_int64;
# ifdef SQLITE_UINT64_TYPE
    typedef SQLITE_UINT64_TYPE sqlite_uint64;
# else  
    typedef unsigned SQLITE_INT64_TYPE sqlite_uint64;
# endif
#elif defined(_MSC_VER) || defined(__BORLANDC__)
  typedef __int64 sqlite_int64;
  typedef unsigned __int64 sqlite_uint64;
#else
  typedef long long int sqlite_int64;
  typedef unsigned long long int sqlite_uint64;
#endif
typedef sqlite_int64 sqlite3_int64;
typedef sqlite_uint64 sqlite3_uint64;

/*
** If compiling for a processor that lacks floating point support,
** substitute integer for floating-point.
*/
#ifdef SQLITE_OMIT_FLOATING_POINT
# define double sqlite3_int64
#endif

/*
** CAPI3REF: Closing A Database Connection
** DESTRUCTOR: sqlite3
**
** ^The sqlite3_close() and sqlite3_close_v2() routines are destructors
** for the [sqlite3] object.
** ^Calls to sqlite3_close() and sqlite3_close_v2() return [SQLITE_OK] if
** the [sqlite3] object is successfully destroyed and all associated
** resources are deallocated.
**
** ^If the database connection is associated with unfinalized prepared
** statements or unfinished sqlite3_backup objects then sqlite3_close()
** will leave the database connection open and return [SQLITE_BUSY].
** ^If sqlite3_close_v2() is called with unfinalized prepared statements
** and/or unfinished sqlite3_backups, then the database connection becomes
** an unusable "zombie" which will automatically be deallocated when the
** last prepared statement is finalized or the last sqlite3_backup is
** finished.  The sqlite3_close_v2() interface is intended for use with
** host languages that are garbage collected, and where the order in which
** destructors are called is arbitrary.
**
** Applications should [sqlite3_finalize | finalize] all [prepared statements],
** [sqlite3_blob_close | close] all [BLOB handles], and 
** [sqlite3_backup_finish | finish] all [sqlite3_backup] objects associated
** with the [sqlite3] object prior to attempting to close the object.  ^If
** sqlite3_close_v2() is called on a [database connection] that still has
** outstanding [prepared statements], [BLOB handles], and/or
** of resources is deferred until all [prepared statements], [BLOB handles],
** and [sqlite3_backup] objects are also destroyed.
**
** ^If an [sqlite3] object is destroyed while a transaction is open,
** the transaction is automatically rolled back.
**
** The C parameter to [sqlite3_close(C)] and [sqlite3_close_v2(C)]
** must be either a NULL
** pointer or an [sqlite3] object pointer obtained
** from [sqlite3_open()], [sqlite3_open16()], or
** [sqlite3_open_v2()], and not previously closed.
** ^Calling sqlite3_close() or sqlite3_close_v2() with a NULL pointer
** argument is a harmless no-op.
*/
SQLITE_API int sqlite3_close(sqlite3*);
SQLITE_API int sqlite3_close_v2(sqlite3*);

/*
** The type for a callback function.
** This is legacy and deprecated.  It is included for historical
** compatibility and is not documented.
*/
typedef int (*sqlite3_callback)(void*,int,char**, char**);

/*
** CAPI3REF: One-Step Query Execution Interface
** METHOD: sqlite3
**
** The sqlite3_exec() interface is a convenience wrapper around
** [sqlite3_prepare_v2()], [sqlite3_step()], and [sqlite3_finalize()],
** that allows an application to run multiple statements of SQL
** without having to use a lot of C code. 
**
** ^The sqlite3_exec() interface runs zero or more UTF-8 encoded,
** semicolon-separate SQL statements passed into its 2nd argument,
** in the context of the [database connection] passed in as its 1st
** argument.  ^If the callback function of the 3rd argument to
** sqlite3_exec() is not NULL, then it is invoked for each result row
** coming out of the evaluated SQL statements.  ^The 4th argument to
** sqlite3_exec() is relayed through to the 1st argument of each
** callback invocation.  ^If the callback pointer to sqlite3_exec()
** is NULL, then no callback is ever invoked and result rows are
** ignored.
**
** ^If an error occurs while evaluating the SQL statements passed into
** sqlite3_exec(), then execution of the current statement stops and
** subsequent statements are skipped.  ^If the 5th parameter to sqlite3_exec()
** is not NULL then any error message is written into memory obtained
** from [sqlite3_malloc()] and passed back through the 5th parameter.
** To avoid memory leaks, the application should invoke [sqlite3_free()]
** on error message strings returned through the 5th parameter of
** read-only media and cannot be changed even by processes with
** elevated privileges.
**
** The SQLITE_IOCAP_BATCH_ATOMIC property means that the underlying
** filesystem supports doing multiple write operations atomically when those
** write operations are bracketed by [SQLITE_FCNTL_BEGIN_ATOMIC_WRITE] and
** [SQLITE_FCNTL_COMMIT_ATOMIC_WRITE].
*/
#define SQLITE_IOCAP_ATOMIC                 0x00000001
#define SQLITE_IOCAP_ATOMIC512              0x00000002
#define SQLITE_IOCAP_ATOMIC1K               0x00000004
#define SQLITE_IOCAP_ATOMIC2K               0x00000008
#define SQLITE_IOCAP_ATOMIC4K               0x00000010
#define SQLITE_IOCAP_ATOMIC8K               0x00000020
#define SQLITE_IOCAP_ATOMIC16K              0x00000040
#define SQLITE_IOCAP_ATOMIC32K              0x00000080
#define SQLITE_IOCAP_ATOMIC64K              0x00000100
#define SQLITE_IOCAP_SAFE_APPEND            0x00000200
#define SQLITE_IOCAP_SEQUENTIAL             0x00000400
#define SQLITE_IOCAP_UNDELETABLE_WHEN_OPEN  0x00000800
#define SQLITE_IOCAP_POWERSAFE_OVERWRITE    0x00001000
#define SQLITE_IOCAP_IMMUTABLE              0x00002000
#define SQLITE_IOCAP_BATCH_ATOMIC           0x00004000

/*
** CAPI3REF: File Locking Levels
**
** SQLite uses one of these integer values as the second
** argument to calls it makes to the xLock() and xUnlock() methods
** of an [sqlite3_io_methods] object.
*/
#define SQLITE_LOCK_NONE          0
#define SQLITE_LOCK_SHARED        1
#define SQLITE_LOCK_RESERVED      2
#define SQLITE_LOCK_PENDING       3
#define SQLITE_LOCK_EXCLUSIVE     4

/*
** CAPI3REF: Synchronization Type Flags
**
** When SQLite invokes the xSync() method of an
** [sqlite3_io_methods] object it uses a combination of
** these integer values as the second argument.
**
** When the SQLITE_SYNC_DATAONLY flag is used, it means that the
** sync operation only needs to flush data to mass storage.  Inode
** information need not be flushed. If the lower four bits of the flag
** equal SQLITE_SYNC_NORMAL, that means to use normal fsync() semantics.
** If the lower four bits equal SQLITE_SYNC_FULL, that means
** to use Mac OS X style fullsync instead of fsync().
**
** Do not confuse the SQLITE_SYNC_NORMAL and SQLITE_SYNC_FULL flags
** with the [PRAGMA synchronous]=NORMAL and [PRAGMA synchronous]=FULL
** settings.  The [synchronous pragma] determines when calls to the
** xSync VFS method occur and applies uniformly across all platforms.
** The SQLITE_SYNC_NORMAL and SQLITE_SYNC_FULL flags determine how
** energetic or rigorous or forceful the sync operations are and
** only make a difference on Mac OSX for the default SQLite code.
** (Third-party VFS implementations might also make the distinction
** between SQLITE_SYNC_NORMAL and SQLITE_SYNC_FULL, but among the
** operating systems natively supported by SQLite, only Mac OSX
** cares about the difference.)
*/
#define SQLITE_SYNC_NORMAL        0x00002
  int (*xClose)(sqlite3_file*);
  int (*xRead)(sqlite3_file*, void*, int iAmt, sqlite3_int64 iOfst);
  int (*xWrite)(sqlite3_file*, const void*, int iAmt, sqlite3_int64 iOfst);
  int (*xTruncate)(sqlite3_file*, sqlite3_int64 size);
  int (*xSync)(sqlite3_file*, int flags);
  int (*xFileSize)(sqlite3_file*, sqlite3_int64 *pSize);
  int (*xLock)(sqlite3_file*, int);
  int (*xUnlock)(sqlite3_file*, int);
  int (*xCheckReservedLock)(sqlite3_file*, int *pResOut);
  int (*xFileControl)(sqlite3_file*, int op, void *pArg);
  int (*xSectorSize)(sqlite3_file*);
  int (*xDeviceCharacteristics)(sqlite3_file*);
  /* Methods above are valid for version 1 */
  int (*xShmMap)(sqlite3_file*, int iPg, int pgsz, int, void volatile**);
  int (*xShmLock)(sqlite3_file*, int offset, int n, int flags);
  void (*xShmBarrier)(sqlite3_file*);
  int (*xShmUnmap)(sqlite3_file*, int deleteFlag);
  /* Methods above are valid for version 2 */
  int (*xFetch)(sqlite3_file*, sqlite3_int64 iOfst, int iAmt, void **pp);
  int (*xUnfetch)(sqlite3_file*, sqlite3_int64 iOfst, void *p);
  /* Methods above are valid for version 3 */
  /* Additional methods may be added in future releases */
};

/*
** CAPI3REF: Standard File Control Opcodes
** KEYWORDS: {file control opcodes} {file control opcode}
**
** These integer constants are opcodes for the xFileControl method
** of the [sqlite3_io_methods] object and for the [sqlite3_file_control()]
** interface.
**
** <ul>
** <li>[[SQLITE_FCNTL_LOCKSTATE]]
** The [SQLITE_FCNTL_LOCKSTATE] opcode is used for debugging.  This
** opcode causes the xFileControl method to write the current state of
** the lock (one of [SQLITE_LOCK_NONE], [SQLITE_LOCK_SHARED],
** [SQLITE_LOCK_RESERVED], [SQLITE_LOCK_PENDING], or [SQLITE_LOCK_EXCLUSIVE])
** into an integer that the pArg argument points to. This capability
** is used during testing and is only available when the SQLITE_TEST
** compile-time option is used.
**
** <li>[[SQLITE_FCNTL_SIZE_HINT]]
** The [SQLITE_FCNTL_SIZE_HINT] opcode is used by SQLite to give the VFS
** layer a hint of how large the database file will grow to be during the
** current transaction.  This hint is not guaranteed to be accurate but it
** is often close.  The underlying VFS might choose to preallocate database
** file space based on this hint in order to help writes to the database
** file run faster.
**
** <li>[[SQLITE_FCNTL_CHUNK_SIZE]]
#define SQLITE_FCNTL_MMAP_SIZE              18
#define SQLITE_FCNTL_TRACE                  19
#define SQLITE_FCNTL_HAS_MOVED              20
#define SQLITE_FCNTL_SYNC                   21
#define SQLITE_FCNTL_COMMIT_PHASETWO        22
#define SQLITE_FCNTL_WIN32_SET_HANDLE       23
#define SQLITE_FCNTL_WAL_BLOCK              24
#define SQLITE_FCNTL_ZIPVFS                 25
#define SQLITE_FCNTL_RBU                    26
#define SQLITE_FCNTL_VFS_POINTER            27
#define SQLITE_FCNTL_JOURNAL_POINTER        28
#define SQLITE_FCNTL_WIN32_GET_HANDLE       29
#define SQLITE_FCNTL_PDB                    30
#define SQLITE_FCNTL_BEGIN_ATOMIC_WRITE     31
#define SQLITE_FCNTL_COMMIT_ATOMIC_WRITE    32
#define SQLITE_FCNTL_ROLLBACK_ATOMIC_WRITE  33
#define SQLITE_FCNTL_LOCK_TIMEOUT           34

/* deprecated names */
#define SQLITE_GET_LOCKPROXYFILE      SQLITE_FCNTL_GET_LOCKPROXYFILE
#define SQLITE_SET_LOCKPROXYFILE      SQLITE_FCNTL_SET_LOCKPROXYFILE
#define SQLITE_LAST_ERRNO             SQLITE_FCNTL_LAST_ERRNO


/*
** CAPI3REF: Mutex Handle
**
** The mutex module within SQLite defines [sqlite3_mutex] to be an
** abstract type for a mutex object.  The SQLite core never looks
** at the internal representation of an [sqlite3_mutex].  It only
** deals with pointers to the [sqlite3_mutex] object.
**
** Mutexes are created using [sqlite3_mutex_alloc()].
*/
typedef struct sqlite3_mutex sqlite3_mutex;

/*
** CAPI3REF: Loadable Extension Thunk
**
** A pointer to the opaque sqlite3_api_routines structure is passed as
** the third parameter to entry points of [loadable extensions].  This
** structure must be typedefed in order to work around compiler warnings
** on some platforms.
*/
typedef struct sqlite3_api_routines sqlite3_api_routines;

/*
** CAPI3REF: OS Interface Object
**
** An instance of the sqlite3_vfs object defines the interface between
** the SQLite core and the underlying operating system.  The "vfs"
** in the name of the object stands for "virtual file system".  See
** the [VFS | VFS documentation] for further information.
**
** The VFS interface is sometimes extended by adding new methods onto
** the end.  Each time such an extension occurs, the iVersion field
** is incremented.  The iVersion value started out as 1 in
** SQLite [version 3.5.0] on [dateof:3.5.0], then increased to 2
** with SQLite [version 3.7.0] on [dateof:3.7.0], and then increased
** to 3 with SQLite [version 3.7.6] on [dateof:3.7.6].  Additional fields
** may be appended to the sqlite3_vfs object and the iVersion value
** may increase again in future versions of SQLite.
** Note that the structure
** </dd>
** </dl>
*/
#define SQLITE_DBCONFIG_MAINDBNAME            1000 /* const char* */
#define SQLITE_DBCONFIG_LOOKASIDE             1001 /* void* int int */
#define SQLITE_DBCONFIG_ENABLE_FKEY           1002 /* int int* */
#define SQLITE_DBCONFIG_ENABLE_TRIGGER        1003 /* int int* */
#define SQLITE_DBCONFIG_ENABLE_FTS3_TOKENIZER 1004 /* int int* */
#define SQLITE_DBCONFIG_ENABLE_LOAD_EXTENSION 1005 /* int int* */
#define SQLITE_DBCONFIG_NO_CKPT_ON_CLOSE      1006 /* int int* */
#define SQLITE_DBCONFIG_ENABLE_QPSG           1007 /* int int* */
#define SQLITE_DBCONFIG_TRIGGER_EQP           1008 /* int int* */
#define SQLITE_DBCONFIG_MAX                   1008 /* Largest DBCONFIG */

/*
** CAPI3REF: Enable Or Disable Extended Result Codes
** METHOD: sqlite3
**
** ^The sqlite3_extended_result_codes() routine enables or disables the
** [extended result codes] feature of SQLite. ^The extended result
** codes are disabled by default for historical compatibility.
*/
SQLITE_API int sqlite3_extended_result_codes(sqlite3*, int onoff);

/*
** CAPI3REF: Last Insert Rowid
** METHOD: sqlite3
**
** ^Each entry in most SQLite tables (except for [WITHOUT ROWID] tables)
** has a unique 64-bit signed
** integer key called the [ROWID | "rowid"]. ^The rowid is always available
** as an undeclared column named ROWID, OID, or _ROWID_ as long as those
** names are not also used by explicitly declared columns. ^If
** the table has a column of type [INTEGER PRIMARY KEY] then that column
** is another alias for the rowid.
**
** ^The sqlite3_last_insert_rowid(D) interface usually returns the [rowid] of
** the most recent successful [INSERT] into a rowid table or [virtual table]
** on database connection D. ^Inserts into [WITHOUT ROWID] tables are not
** recorded. ^If no successful [INSERT]s into rowid tables have ever occurred 
** on the database connection D, then sqlite3_last_insert_rowid(D) returns 
** zero.
**
** As well as being set automatically as rows are inserted into database
** tables, the value returned by this function may be set explicitly by
** [sqlite3_set_last_insert_rowid()]
**
** Some virtual table implementations may INSERT rows into rowid tables as
** part of committing a transaction (e.g. to flush data accumulated in memory
** to disk). In this case subsequent calls to this function return the rowid
** associated with these internal INSERT operations, which leads to 
** is closed or might close before sqlite3_interrupt() returns.
**
** ^If an SQL operation is very nearly finished at the time when
** sqlite3_interrupt() is called, then it might not have an opportunity
** to be interrupted and might continue to completion.
**
** ^An SQL operation that is interrupted will return [SQLITE_INTERRUPT].
** ^If the interrupted SQL operation is an INSERT, UPDATE, or DELETE
** that is inside an explicit transaction, then the entire transaction
** will be rolled back automatically.
**
** ^The sqlite3_interrupt(D) call is in effect until all currently running
** SQL statements on [database connection] D complete.  ^Any new SQL statements
** that are started after the sqlite3_interrupt() call and before the 
** running statements reaches zero are interrupted as if they had been
** running prior to the sqlite3_interrupt() call.  ^New SQL statements
** that are started after the running statement count reaches zero are
** not effected by the sqlite3_interrupt().
** ^A call to sqlite3_interrupt(D) that occurs when there are no running
** SQL statements is a no-op and has no effect on SQL statements
** that are started after the sqlite3_interrupt() call returns.
*/
SQLITE_API void sqlite3_interrupt(sqlite3*);

/*
** CAPI3REF: Determine If An SQL Statement Is Complete
**
** These routines are useful during command-line input to determine if the
** currently entered text seems to form a complete SQL statement or
** if additional input is needed before sending the text into
** SQLite for parsing.  ^These routines return 1 if the input string
** appears to be a complete SQL statement.  ^A statement is judged to be
** complete if it ends with a semicolon token and is not a prefix of a
** well-formed CREATE TRIGGER statement.  ^Semicolons that are embedded within
** string literals or quoted identifier names or comments are not
** independent tokens (they are part of the token in which they are
** embedded) and thus do not count as a statement terminator.  ^Whitespace
** and comments that follow the final semicolon are ignored.
**
** ^These routines return 0 if the statement is incomplete.  ^If a
** memory allocation fails, then SQLITE_NOMEM is returned.
**
** ^These routines do not parse the SQL statements thus
** will not detect syntactically incorrect SQL.
**
** ^(If SQLite has not been initialized using [sqlite3_initialize()] prior 
** to invoking sqlite3_complete16() then sqlite3_initialize() is invoked
** automatically by sqlite3_complete16().  If that initialization fails,
** then the return value from sqlite3_complete16() will be non-zero
** regardless of whether or not the input SQL is complete.)^
**

/*
** CAPI3REF: Pseudo-Random Number Generator
**
** SQLite contains a high-quality pseudo-random number generator (PRNG) used to
** select random [ROWID | ROWIDs] when inserting new records into a table that
** already uses the largest possible [ROWID].  The PRNG is also used for
** the build-in random() and randomblob() SQL functions.  This interface allows
** applications to access the same PRNG for other purposes.
**
** ^A call to this routine stores N bytes of randomness into buffer P.
** ^The P parameter can be a NULL pointer.
**
** ^If this routine has not been previously called or if the previous
** call had N less than one or a NULL pointer for P, then the PRNG is
** seeded using randomness obtained from the xRandomness method of
** the default [sqlite3_vfs] object.
** ^If the previous call to this routine had an N of 1 or more and a
** non-NULL P then the pseudo-randomness is generated
** internally and without recourse to the [sqlite3_vfs] xRandomness
** method.
*/
SQLITE_API void sqlite3_randomness(int N, void *P);

/*
** CAPI3REF: Compile-Time Authorization Callbacks
** METHOD: sqlite3
** KEYWORDS: {authorizer callback}
**
** ^This routine registers an authorizer callback with a particular
** [database connection], supplied in the first argument.
** ^The authorizer callback is invoked as SQL statements are being compiled
** by [sqlite3_prepare()] or its variants [sqlite3_prepare_v2()],
** [sqlite3_prepare_v3()], [sqlite3_prepare16()], [sqlite3_prepare16_v2()],
** and [sqlite3_prepare16_v3()].  ^At various
** points during the compilation process, as logic is being created
** to perform various actions, the authorizer callback is invoked to
** see if those actions are allowed.  ^The authorizer callback should
** return [SQLITE_OK] to allow the action, [SQLITE_IGNORE] to disallow the
** specific action but allow the SQL statement to continue to be
** compiled, or [SQLITE_DENY] to cause the entire SQL statement to be
** rejected with an error.  ^If the authorizer callback returns
** any value other than [SQLITE_IGNORE], [SQLITE_OK], or [SQLITE_DENY]
** then the [sqlite3_prepare_v2()] or equivalent call that triggered
** the authorizer will fail with an error message.
**
** When the callback returns [SQLITE_OK], that means the operation
** requested is ok.  ^When the callback returns [SQLITE_DENY], the
** [sqlite3_prepare_v2()] or equivalent call that triggered the
** authorizer will fail with an error message explaining that
** access is denied. 
#define SQLITE_DROP_TEMP_INDEX      12   /* Index Name      Table Name      */
#define SQLITE_DROP_TEMP_TABLE      13   /* Table Name      NULL            */
#define SQLITE_DROP_TEMP_TRIGGER    14   /* Trigger Name    Table Name      */
#define SQLITE_DROP_TEMP_VIEW       15   /* View Name       NULL            */
#define SQLITE_DROP_TRIGGER         16   /* Trigger Name    Table Name      */
#define SQLITE_DROP_VIEW            17   /* View Name       NULL            */
#define SQLITE_INSERT               18   /* Table Name      NULL            */
#define SQLITE_PRAGMA               19   /* Pragma Name     1st arg or NULL */
#define SQLITE_READ                 20   /* Table Name      Column Name     */
#define SQLITE_SELECT               21   /* NULL            NULL            */
#define SQLITE_TRANSACTION          22   /* Operation       NULL            */
#define SQLITE_UPDATE               23   /* Table Name      Column Name     */
#define SQLITE_ATTACH               24   /* Filename        NULL            */
#define SQLITE_DETACH               25   /* Database Name   NULL            */
#define SQLITE_ALTER_TABLE          26   /* Database Name   Table Name      */
#define SQLITE_REINDEX              27   /* Index Name      NULL            */
#define SQLITE_ANALYZE              28   /* Table Name      NULL            */
#define SQLITE_CREATE_VTABLE        29   /* Table Name      Module Name     */
#define SQLITE_DROP_VTABLE          30   /* Table Name      Module Name     */
#define SQLITE_FUNCTION             31   /* NULL            Function Name   */
#define SQLITE_SAVEPOINT            32   /* Operation       Savepoint Name  */
#define SQLITE_COPY                  0   /* No longer used */
#define SQLITE_RECURSIVE            33   /* NULL            NULL            */

/*
** CAPI3REF: Tracing And Profiling Functions
** METHOD: sqlite3
**
** These routines are deprecated. Use the [sqlite3_trace_v2()] interface
** instead of the routines described here.
**
** These routines register callback functions that can be used for
** tracing and profiling the execution of SQL statements.
**
** ^The callback function registered by sqlite3_trace() is invoked at
** various times when an SQL statement is being run by [sqlite3_step()].
** ^The sqlite3_trace() callback is invoked with a UTF-8 rendering of the
** SQL statement text as the statement first begins executing.
** ^(Additional sqlite3_trace() callbacks might occur
** as each triggered subprogram is entered.  The callbacks for triggers
** contain a UTF-8 SQL comment that identifies the trigger.)^
**
** The [SQLITE_TRACE_SIZE_LIMIT] compile-time option can be used to limit
** the length of [bound parameter] expansion in the output of sqlite3_trace().
**
** ^The callback function registered by sqlite3_profile() is invoked
** as each SQL statement finishes.  ^The profile callback contains
** the original statement text and an estimate of wall-clock time
** of how long that statement took to run.  ^The profile callback
** time is in units of nanoseconds, however the current implementation
** is only capable of millisecond resolution so the six least significant
** ^The parameter P is passed through as the only parameter to the 
** callback function X.  ^The parameter N is the approximate number of 
** [virtual machine instructions] that are evaluated between successive
** invocations of the callback X.  ^If N is less than one then the progress
** handler is disabled.
**
** ^Only a single progress handler may be defined at one time per
** [database connection]; setting a new progress handler cancels the
** old one.  ^Setting parameter X to NULL disables the progress handler.
** ^The progress handler is also disabled by setting N to a value less
** than 1.
**
** ^If the progress callback returns non-zero, the operation is
** interrupted.  This feature can be used to implement a
** "Cancel" button on a GUI progress dialog box.
**
** The progress handler callback must not do anything that will modify
** the database connection that invoked the progress handler.
** Note that [sqlite3_prepare_v2()] and [sqlite3_step()] both modify their
** database connections for the meaning of "modify" in this paragraph.
**
*/
SQLITE_API void sqlite3_progress_handler(sqlite3*, int, int(*)(void*), void*);

/*
** CAPI3REF: Opening A New Database Connection
** CONSTRUCTOR: sqlite3
**
** ^These routines open an SQLite database file as specified by the 
** filename argument. ^The filename argument is interpreted as UTF-8 for
** sqlite3_open() and sqlite3_open_v2() and as UTF-16 in the native byte
** order for sqlite3_open16(). ^(A [database connection] handle is usually
** returned in *ppDb, even if an error occurs.  The only exception is that
** if SQLite is unable to allocate memory to hold the [sqlite3] object,
** a NULL will be written into *ppDb instead of a pointer to the [sqlite3]
** object.)^ ^(If the database is opened (and/or created) successfully, then
** [SQLITE_OK] is returned.  Otherwise an [error code] is returned.)^ ^The
** [sqlite3_errmsg()] or [sqlite3_errmsg16()] routines can be used to obtain
** an English language description of the error following a failure of any
** of the sqlite3_open() routines.
**
** ^The default encoding will be UTF-8 for databases created using
** sqlite3_open() or sqlite3_open_v2().  ^The default encoding for databases
** created using sqlite3_open16() will be UTF-16 in the native byte order.
**
** Whether or not an error occurs when it is opened, resources
** associated with the [database connection] handle should be released by
** passing it to [sqlite3_close()] when it is no longer required.
**
** The sqlite3_open_v2() interface works like sqlite3_open()
** except that it accepts two additional parameters for additional control
** characters must be converted to UTF-8 prior to passing them into
** sqlite3_open() or sqlite3_open_v2().
**
** <b>Note to Windows Runtime users:</b>  The temporary directory must be set
** prior to calling sqlite3_open() or sqlite3_open_v2().  Otherwise, various
** features that require the use of temporary files may fail.
**
** See also: [sqlite3_temp_directory]
*/
SQLITE_API int sqlite3_open(
  const char *filename,   /* Database filename (UTF-8) */
  sqlite3 **ppDb          /* OUT: SQLite db handle */
);
SQLITE_API int sqlite3_open16(
  const void *filename,   /* Database filename (UTF-16) */
  sqlite3 **ppDb          /* OUT: SQLite db handle */
);
SQLITE_API int sqlite3_open_v2(
  const char *filename,   /* Database filename (UTF-8) */
  sqlite3 **ppDb,         /* OUT: SQLite db handle */
  int flags,              /* Flags */
  const char *zVfs        /* Name of VFS module to use */
);

/*
** CAPI3REF: Obtain Values For URI Parameters
**
** These are utility routines, useful to VFS implementations, that check
** to see if a database file was a URI that contained a specific query 
** parameter, and if so obtains the value of that query parameter.
**
** If F is the database filename pointer passed into the xOpen() method of 
** a VFS implementation when the flags parameter to xOpen() has one or 
** more of the [SQLITE_OPEN_URI] or [SQLITE_OPEN_MAIN_DB] bits set and
** P is the name of the query parameter, then
** sqlite3_uri_parameter(F,P) returns the value of the P
** parameter if it exists or a NULL pointer if P does not appear as a 
** query parameter on F.  If P is a query parameter of F
** has no explicit value, then sqlite3_uri_parameter(F,P) returns
** a pointer to an empty string.
**
** The sqlite3_uri_boolean(F,P,B) routine assumes that P is a boolean
** parameter and returns true (1) or false (0) according to the value
** of P.  The sqlite3_uri_boolean(F,P,B) routine returns true (1) if the
** value of query parameter P is one of "yes", "true", or "on" in any
** case or if the value begins with a non-zero number.  The 
** sqlite3_uri_boolean(F,P,B) routines returns false (0) if the value of
** query parameter P is one of "no", "false", or "off" in any case or
** if the value begins with a numeric zero.  If P is not a query
** parameter on F or if the value of P is does not match any of the
** above, then sqlite3_uri_boolean(F,P,B) returns (B!=0).
** KEYWORDS: {prepared statement} {prepared statements}
**
** An instance of this object represents a single SQL statement that
** has been compiled into binary form and is ready to be evaluated.
**
** Think of each SQL statement as a separate computer program.  The
** original SQL text is source code.  A prepared statement object 
** is the compiled object code.  All SQL must be converted into a
** prepared statement before it can be run.
**
** The life-cycle of a prepared statement object usually goes like this:
**
** <ol>
** <li> Create the prepared statement object using [sqlite3_prepare_v2()].
** <li> Bind values to [parameters] using the sqlite3_bind_*()
**      interfaces.
** <li> Run the SQL by calling [sqlite3_step()] one or more times.
** <li> Reset the prepared statement using [sqlite3_reset()] then go back
**      to step 2.  Do this zero or more times.
** <li> Destroy the object using [sqlite3_finalize()].
** </ol>
*/
typedef struct sqlite3_stmt sqlite3_stmt;

/*
** CAPI3REF: Run-time Limits
** METHOD: sqlite3
**
** ^(This interface allows the size of various constructs to be limited
** on a connection by connection basis.  The first parameter is the
** [database connection] whose limit is to be set or queried.  The
** second parameter is one of the [limit categories] that define a
** class of constructs to be size limited.  The third parameter is the
** new limit for that construct.)^
**
** ^If the new limit is a negative number, the limit is unchanged.
** ^(For each limit category SQLITE_LIMIT_<i>NAME</i> there is a 
** [limits | hard upper bound]
** set at compile-time by a C preprocessor macro called
** [limits | SQLITE_MAX_<i>NAME</i>].
** (The "_LIMIT_" in the name is changed to "_MAX_".))^
** ^Attempts to increase a limit above its hard upper bound are
** silently truncated to the hard upper bound.
**
** ^Regardless of whether or not the limit was changed, the 
** [sqlite3_limit()] interface returns the prior value of the limit.
** ^Hence, to find the current value of a limit without changing it,
** simply invoke this interface with the third parameter set to -1.
**
** Run-time limits are intended for use in applications that manage
** both their own internal database and also databases that are controlled
** by untrusted external sources.  An example application might be a
** web browser that has its own databases for storing history and
** separate databases controlled by JavaScript applications downloaded
** off the Internet.  The internal databases can be given the
** large, default limits.  Databases managed by external sources can
** be given much smaller limits designed to prevent a denial of service
** attack.  Developers might also want to use the [sqlite3_set_authorizer()]
** interface to further control untrusted SQL.  The size of the database
** created by an untrusted script can be contained using the
** [max_page_count] [PRAGMA].
**
** New run-time limit categories may be added in future releases.
*/
SQLITE_API int sqlite3_limit(sqlite3*, int id, int newVal);

/*
** CAPI3REF: Run-Time Limit Categories
** KEYWORDS: {limit category} {*limit categories}
**
** These constants define various performance limits
** that can be lowered at run-time using [sqlite3_limit()].
** The synopsis of the meanings of the various limits is shown below.
** Additional information is available at [limits | Limits in SQLite].
**
** <dl>
** [[SQLITE_LIMIT_LENGTH]] ^(<dt>SQLITE_LIMIT_LENGTH</dt>
** <dd>The maximum size of any string or BLOB or table row, in bytes.<dd>)^
**
** [[SQLITE_LIMIT_SQL_LENGTH]] ^(<dt>SQLITE_LIMIT_SQL_LENGTH</dt>
** <dd>The maximum length of an SQL statement, in bytes.</dd>)^
**
** [[SQLITE_LIMIT_COLUMN]] ^(<dt>SQLITE_LIMIT_COLUMN</dt>
** <dd>The maximum number of columns in a table definition or in the
** result set of a [SELECT] or the maximum number of columns in an index
** or in an ORDER BY or GROUP BY clause.</dd>)^
**
** [[SQLITE_LIMIT_EXPR_DEPTH]] ^(<dt>SQLITE_LIMIT_EXPR_DEPTH</dt>
** <dd>The maximum depth of the parse tree on any expression.</dd>)^
**
** [[SQLITE_LIMIT_COMPOUND_SELECT]] ^(<dt>SQLITE_LIMIT_COMPOUND_SELECT</dt>
** <dd>The maximum number of terms in a compound SELECT statement.</dd>)^
**
/*
** CAPI3REF: Prepare Flags
**
** These constants define various flags that can be passed into
** "prepFlags" parameter of the [sqlite3_prepare_v3()] and
** [sqlite3_prepare16_v3()] interfaces.
**
** New flags may be added in future releases of SQLite.
**
** <dl>
** [[SQLITE_PREPARE_PERSISTENT]] ^(<dt>SQLITE_PREPARE_PERSISTENT</dt>
** <dd>The SQLITE_PREPARE_PERSISTENT flag is a hint to the query planner
** that the prepared statement will be retained for a long time and
** probably reused many times.)^ ^Without this flag, [sqlite3_prepare_v3()]
** and [sqlite3_prepare16_v3()] assume that the prepared statement will 
** be used just once or at most a few times and then destroyed using
** [sqlite3_finalize()] relatively soon. The current implementation acts
** on this hint by avoiding the use of [lookaside memory] so as not to
** deplete the limited store of lookaside memory. Future versions of
** SQLite may act on this hint differently.
** </dl>
*/
#define SQLITE_PREPARE_PERSISTENT              0x01

/*
** CAPI3REF: Compiling An SQL Statement
** KEYWORDS: {SQL statement compiler}
** METHOD: sqlite3
** CONSTRUCTOR: sqlite3_stmt
**
** To execute an SQL statement, it must first be compiled into a byte-code
** program using one of these routines.  Or, in other words, these routines
** are constructors for the [prepared statement] object.
**
** The preferred routine to use is [sqlite3_prepare_v2()].  The
** [sqlite3_prepare()] interface is legacy and should be avoided.
** [sqlite3_prepare_v3()] has an extra "prepFlags" option that is used
** for special purposes.
**
** The use of the UTF-8 interfaces is preferred, as SQLite currently
** does all parsing using UTF-8.  The UTF-16 interfaces are provided
** as a convenience.  The UTF-16 interfaces work by converting the
** input text into UTF-8, then invoking the corresponding UTF-8 interface.
**
** The first argument, "db", is a [database connection] obtained from a
** prior successful call to [sqlite3_open()], [sqlite3_open_v2()] or
** [sqlite3_open16()].  The database connection must not have been closed.
**
** The second argument, "zSql", is the statement to be compiled, encoded
** as either UTF-8 or UTF-16.  The sqlite3_prepare(), sqlite3_prepare_v2(),
** and sqlite3_prepare_v3()
);
SQLITE_API int sqlite3_prepare16(
  sqlite3 *db,            /* Database handle */
  const void *zSql,       /* SQL statement, UTF-16 encoded */
  int nByte,              /* Maximum length of zSql in bytes. */
  sqlite3_stmt **ppStmt,  /* OUT: Statement handle */
  const void **pzTail     /* OUT: Pointer to unused portion of zSql */
);
SQLITE_API int sqlite3_prepare16_v2(
  sqlite3 *db,            /* Database handle */
  const void *zSql,       /* SQL statement, UTF-16 encoded */
  int nByte,              /* Maximum length of zSql in bytes. */
  sqlite3_stmt **ppStmt,  /* OUT: Statement handle */
  const void **pzTail     /* OUT: Pointer to unused portion of zSql */
);
SQLITE_API int sqlite3_prepare16_v3(
  sqlite3 *db,            /* Database handle */
  const void *zSql,       /* SQL statement, UTF-16 encoded */
  int nByte,              /* Maximum length of zSql in bytes. */
  unsigned int prepFlags, /* Zero or more SQLITE_PREPARE_ flags */
  sqlite3_stmt **ppStmt,  /* OUT: Statement handle */
  const void **pzTail     /* OUT: Pointer to unused portion of zSql */
);

/*
** CAPI3REF: Retrieving Statement SQL
** METHOD: sqlite3_stmt
**
** ^The sqlite3_sql(P) interface returns a pointer to a copy of the UTF-8
** SQL text used to create [prepared statement] P if P was
** created by [sqlite3_prepare_v2()], [sqlite3_prepare_v3()],
** [sqlite3_prepare16_v2()], or [sqlite3_prepare16_v3()].
** ^The sqlite3_expanded_sql(P) interface returns a pointer to a UTF-8
** string containing the SQL text of prepared statement P with
** [bound parameters] expanded.
**
** ^(For example, if a prepared statement is created using the SQL
** text "SELECT $abc,:xyz" and if parameter $abc is bound to integer 2345
** and parameter :xyz is unbound, then sqlite3_sql() will return
** the original string, "SELECT $abc,:xyz" but sqlite3_expanded_sql()
** will return "SELECT 2345,NULL".)^
**
** ^The sqlite3_expanded_sql() interface returns NULL if insufficient memory
** is available to hold the result, or if the result would exceed the
** the maximum string length determined by the [SQLITE_LIMIT_LENGTH].
**
** ^The [SQLITE_TRACE_SIZE_LIMIT] compile-time option limits the size of
** bound parameter expansions.  ^The [SQLITE_OMIT_TRACE] compile-time
** option causes sqlite3_expanded_sql() to always return NULL.
**
** ^The string returned by sqlite3_sql(P) is managed by SQLite and is
** automatically freed when the prepared statement is finalized.
** ^The string returned by sqlite3_expanded_sql(P), on the other hand,
** is obtained from [sqlite3_malloc()] and must be free by the application
** by passing it to [sqlite3_free()].
*/
SQLITE_API const char *sqlite3_sql(sqlite3_stmt *pStmt);
SQLITE_API char *sqlite3_expanded_sql(sqlite3_stmt *pStmt);

/*
** CAPI3REF: Determine If An SQL Statement Writes The Database
** METHOD: sqlite3_stmt
**
** ^The sqlite3_stmt_readonly(X) interface returns true (non-zero) if
** and only if the [prepared statement] X makes no direct changes to
** the content of the database file.
**
** Note that [application-defined SQL functions] or
** [virtual tables] might change the database indirectly as a side effect.  
** ^(For example, if an application defines a function "eval()" that 
** calls [sqlite3_exec()], then the following SQL statement would
** change the database file through side-effects:
**
** <blockquote><pre>
**    SELECT eval('DELETE FROM t1') FROM t2;
** </pre></blockquote>
**
** But because the [SELECT] statement does not change the database file
** directly, sqlite3_stmt_readonly() would still return true.)^
**
** ^Transaction control statements such as [BEGIN], [COMMIT], [ROLLBACK],
** [SAVEPOINT], and [RELEASE] cause sqlite3_stmt_readonly() to return true,
** since the statements themselves do not actually modify the database but
** rather they control the timing of when other statements modify the 
** database.  ^The [ATTACH] and [DETACH] statements also cause
** sqlite3_stmt_readonly() to return true since, while those statements
** change the configuration of a database connection, they do not make 
** changes to the content of the database files on disk.
** ^The sqlite3_stmt_readonly() interface returns true for [BEGIN] since
** [BEGIN] merely sets internal flags, but the [BEGIN|BEGIN IMMEDIATE] and
** [BEGIN|BEGIN EXCLUSIVE] commands do touch the database and so
** sqlite3_stmt_readonly() returns false for those commands.
*/
SQLITE_API int sqlite3_stmt_readonly(sqlite3_stmt *pStmt);

/*
** CAPI3REF: Determine If A Prepared Statement Has Been Reset
** METHOD: sqlite3_stmt
**
** ^The sqlite3_stmt_busy(S) interface returns true (non-zero) if the
** [prepared statement] S has been stepped at least once using 
** [sqlite3_step(S)] but has neither run to completion (returned
** [SQLITE_DONE] from [sqlite3_step(S)]) nor
** been reset using [sqlite3_reset(S)].  ^The sqlite3_stmt_busy(S)
** interface returns false if S is a NULL pointer.  If S is not a 
** NULL pointer and is not a pointer to a valid [prepared statement]
** object, then the behavior is undefined and probably undesirable.
**
** This interface can be used in combination [sqlite3_next_stmt()]
** to locate all prepared statements associated with a database 
** connection that are in need of being reset.  This can be used,
** for example, in diagnostic routines to search for prepared 
** statements that are holding a transaction open.
*/
SQLITE_API int sqlite3_stmt_busy(sqlite3_stmt*);

/*
** CAPI3REF: Dynamically Typed Value Object
** KEYWORDS: {protected sqlite3_value} {unprotected sqlite3_value}
**
** SQLite uses the sqlite3_value object to represent all values
** that can be stored in a database table. SQLite uses dynamic typing
** expression or subquery, then a NULL pointer is returned.
** ^The returned string is always UTF-8 encoded.
**
** ^(For example, given the database schema:
**
** CREATE TABLE t1(c1 VARIANT);
**
** and the following statement to be compiled:
**
** SELECT c1 + 1, c1 FROM t1;
**
** this routine would return the string "VARIANT" for the second result
** column (i==1), and a NULL pointer for the first result column (i==0).)^
**
** ^SQLite uses dynamic run-time typing.  ^So just because a column
** is declared to contain a particular type does not mean that the
** data stored in that column is of the declared type.  SQLite is
** strongly typed, but the typing is dynamic not static.  ^Type
** is associated with individual values, not with the containers
** used to hold those values.
*/
SQLITE_API const char *sqlite3_column_decltype(sqlite3_stmt*,int);
SQLITE_API const void *sqlite3_column_decltype16(sqlite3_stmt*,int);

/*
** CAPI3REF: Evaluate An SQL Statement
** METHOD: sqlite3_stmt
**
** After a [prepared statement] has been prepared using any of
** [sqlite3_prepare_v2()], [sqlite3_prepare_v3()], [sqlite3_prepare16_v2()],
** or [sqlite3_prepare16_v3()] or one of the legacy
** interfaces [sqlite3_prepare()] or [sqlite3_prepare16()], this function
** must be called one or more times to evaluate the statement.
**
** The details of the behavior of the sqlite3_step() interface depend
** on whether the statement was prepared using the newer "vX" interfaces
** [sqlite3_prepare_v3()], [sqlite3_prepare_v2()], [sqlite3_prepare16_v3()],
** [sqlite3_prepare16_v2()] or the older legacy
** interfaces [sqlite3_prepare()] and [sqlite3_prepare16()].  The use of the
** new "vX" interface is recommended for new applications but the legacy
** interface will continue to be supported.
**
** ^In the legacy interface, the return value will be either [SQLITE_BUSY],
** [SQLITE_DONE], [SQLITE_ROW], [SQLITE_ERROR], or [SQLITE_MISUSE].
** ^With the "v2" interface, any of the other [result codes] or
** [extended result codes] might be returned as well.
**
** ^[SQLITE_BUSY] means that the database engine was unable to acquire the
** database locks it needs to do its job.  ^If the statement is a [COMMIT]
** or occurs outside of an explicit transaction, then you can retry the
** statement.  If the statement is not a [COMMIT] and occurs within an
*/
SQLITE_API int sqlite3_step(sqlite3_stmt*);

/*
** CAPI3REF: Number of columns in a result set
** METHOD: sqlite3_stmt
**
** ^The sqlite3_data_count(P) interface returns the number of columns in the
** current row of the result set of [prepared statement] P.
** ^If prepared statement P does not have results ready to return
** (via calls to the [sqlite3_column_int | sqlite3_column_*()] of
** interfaces) then sqlite3_data_count(P) returns 0.
** ^The sqlite3_data_count(P) routine also returns 0 if P is a NULL pointer.
** ^The sqlite3_data_count(P) routine returns 0 if the previous call to
** [sqlite3_step](P) returned [SQLITE_DONE].  ^The sqlite3_data_count(P)
** will return non-zero if previous call to [sqlite3_step](P) returned
** [SQLITE_ROW], except in the case of the [PRAGMA incremental_vacuum]
** where it always returns zero since each step of that multi-step
** pragma returns 0 columns of data.
**
** See also: [sqlite3_column_count()]
*/
SQLITE_API int sqlite3_data_count(sqlite3_stmt *pStmt);

/*
** CAPI3REF: Fundamental Datatypes
** KEYWORDS: SQLITE_TEXT
**
** ^(Every value in SQLite has one of five fundamental datatypes:
**
** <ul>
** <li> 64-bit signed integer
** <li> 64-bit IEEE floating point number
** <li> string
** <li> BLOB
** <li> NULL
** </ul>)^
**
** These constants are codes for each of those types.
**
** Note that the SQLITE_TEXT constant was also used in SQLite version 2
** for a completely different meaning.  Software that links against both
** SQLite version 2 and SQLite version 3 should use SQLITE3_TEXT, not
** SQLITE_TEXT.
*/
#define SQLITE_INTEGER  1
#define SQLITE_FLOAT    2
#define SQLITE_BLOB     4
#define SQLITE_NULL     5
#ifdef SQLITE_TEXT
# undef SQLITE_TEXT
  void (*xFinal)(sqlite3_context*)
);
SQLITE_API int sqlite3_create_function16(
  sqlite3 *db,
  const void *zFunctionName,
  int nArg,
  int eTextRep,
  void *pApp,
  void (*xFunc)(sqlite3_context*,int,sqlite3_value**),
  void (*xStep)(sqlite3_context*,int,sqlite3_value**),
  void (*xFinal)(sqlite3_context*)
);
SQLITE_API int sqlite3_create_function_v2(
  sqlite3 *db,
  const char *zFunctionName,
  int nArg,
  int eTextRep,
  void *pApp,
  void (*xFunc)(sqlite3_context*,int,sqlite3_value**),
  void (*xStep)(sqlite3_context*,int,sqlite3_value**),
  void (*xFinal)(sqlite3_context*),
  void(*xDestroy)(void*)
);

/*
** CAPI3REF: Text Encodings
**
** These constant define integer codes that represent the various
** text encodings supported by SQLite.
*/
#define SQLITE_UTF8           1    /* IMP: R-37514-35566 */
#define SQLITE_UTF16LE        2    /* IMP: R-03371-37637 */
#define SQLITE_UTF16BE        3    /* IMP: R-51971-34154 */
#define SQLITE_UTF16          4    /* Use native byte order */
#define SQLITE_ANY            5    /* Deprecated */
#define SQLITE_UTF16_ALIGNED  8    /* sqlite3_create_collation only */

/*
** CAPI3REF: Function Flags
**
** These constants may be ORed together with the 
** [SQLITE_UTF8 | preferred text encoding] as the fourth argument
** to [sqlite3_create_function()], [sqlite3_create_function16()], or
** [sqlite3_create_function_v2()].
*/
#define SQLITE_DETERMINISTIC    0x800

/*
** CAPI3REF: Deprecated Functions
** DEPRECATED
**
** These functions are [deprecated].  In order to maintain
** backwards compatibility with older code, these functions continue 
** to be supported.  However, new applications should avoid
** the use of these functions.  To encourage programmers to avoid
** these functions, we will not explain what they do.
*/
#ifndef SQLITE_OMIT_DEPRECATED
SQLITE_API SQLITE_DEPRECATED int sqlite3_aggregate_count(sqlite3_context*);
SQLITE_API SQLITE_DEPRECATED int sqlite3_expired(sqlite3_stmt*);
SQLITE_API SQLITE_DEPRECATED int sqlite3_transfer_bindings(sqlite3_stmt*, sqlite3_stmt*);
SQLITE_API SQLITE_DEPRECATED int sqlite3_global_recover(void);
SQLITE_API SQLITE_DEPRECATED void sqlite3_thread_cleanup(void);
SQLITE_API SQLITE_DEPRECATED int sqlite3_memory_alarm(void(*)(void*,sqlite3_int64,int),
                      void*,sqlite3_int64);
#endif

/*
** CAPI3REF: Obtaining SQL Values
** METHOD: sqlite3_value
**
** <b>Summary:</b>
** <blockquote><table border=0 cellpadding=0 cellspacing=0>
** <tr><td><b>sqlite3_value_blob</b><td>&rarr;<td>BLOB value
** <tr><td><b>sqlite3_value_double</b><td>&rarr;<td>REAL value
** <tr><td><b>sqlite3_value_int</b><td>&rarr;<td>32-bit INTEGER value
** <tr><td><b>sqlite3_value_int64</b><td>&rarr;<td>64-bit INTEGER value
** <tr><td><b>sqlite3_value_pointer</b><td>&rarr;<td>Pointer value
** <tr><td><b>sqlite3_value_text</b><td>&rarr;<td>UTF-8 TEXT value
** <tr><td><b>sqlite3_value_text16</b><td>&rarr;<td>UTF-16 TEXT value in
** the native byteorder
** <tr><td><b>sqlite3_value_text16be</b><td>&rarr;<td>UTF-16be TEXT value
** <tr><td><b>sqlite3_value_text16le</b><td>&rarr;<td>UTF-16le TEXT value
** <tr><td>&nbsp;<td>&nbsp;<td>&nbsp;
** <tr><td><b>sqlite3_value_bytes</b><td>&rarr;<td>Size of a BLOB
** or a UTF-8 TEXT in bytes
** <tr><td><b>sqlite3_value_bytes16&nbsp;&nbsp;</b>
** <td>&rarr;&nbsp;&nbsp;<td>Size of UTF-16
** TEXT in bytes
** <tr><td><b>sqlite3_value_type</b><td>&rarr;<td>Default
** datatype of the value
** <tr><td><b>sqlite3_value_numeric_type&nbsp;&nbsp;</b>
** <td>&rarr;&nbsp;&nbsp;<td>Best numeric datatype of the value
** <tr><td><b>sqlite3_value_nochange&nbsp;&nbsp;</b>
** Please pay particular attention to the fact that the pointer returned
** from [sqlite3_value_blob()], [sqlite3_value_text()], or
** [sqlite3_value_text16()] can be invalidated by a subsequent call to
** [sqlite3_value_bytes()], [sqlite3_value_bytes16()], [sqlite3_value_text()],
** or [sqlite3_value_text16()].
**
** These routines must be called from the same thread as
** the SQL function that supplied the [sqlite3_value*] parameters.
*/
SQLITE_API const void *sqlite3_value_blob(sqlite3_value*);
SQLITE_API double sqlite3_value_double(sqlite3_value*);
SQLITE_API int sqlite3_value_int(sqlite3_value*);
SQLITE_API sqlite3_int64 sqlite3_value_int64(sqlite3_value*);
SQLITE_API void *sqlite3_value_pointer(sqlite3_value*, const char*);
SQLITE_API const unsigned char *sqlite3_value_text(sqlite3_value*);
SQLITE_API const void *sqlite3_value_text16(sqlite3_value*);
SQLITE_API const void *sqlite3_value_text16le(sqlite3_value*);
SQLITE_API const void *sqlite3_value_text16be(sqlite3_value*);
SQLITE_API int sqlite3_value_bytes(sqlite3_value*);
SQLITE_API int sqlite3_value_bytes16(sqlite3_value*);
SQLITE_API int sqlite3_value_type(sqlite3_value*);
SQLITE_API int sqlite3_value_numeric_type(sqlite3_value*);
SQLITE_API int sqlite3_value_nochange(sqlite3_value*);

/*
** CAPI3REF: Finding The Subtype Of SQL Values
** METHOD: sqlite3_value
**
** The sqlite3_value_subtype(V) function returns the subtype for
** an [application-defined SQL function] argument V.  The subtype
** information can be used to pass a limited amount of context from
** one SQL function to another.  Use the [sqlite3_result_subtype()]
** routine to set the subtype for the return value of an SQL function.
*/
SQLITE_API unsigned int sqlite3_value_subtype(sqlite3_value*);

/*
** CAPI3REF: Copy And Free SQL Values
** METHOD: sqlite3_value
**
** ^The sqlite3_value_dup(V) interface makes a copy of the [sqlite3_value]
** object D and returns a pointer to that copy.  ^The [sqlite3_value] returned
** is a [protected sqlite3_value] object even if the input is not.
** ^The sqlite3_value_dup(V) interface returns NULL if V is NULL or if a
** memory allocation fails.
**
** ^The sqlite3_value_free(V) interface frees an [sqlite3_value] object
** previously obtained from [sqlite3_value_dup()].  ^If V is a NULL pointer
** then sqlite3_value_free(V) is a harmless no-op.
*/
SQLITE_API sqlite3_value *sqlite3_value_dup(const sqlite3_value*);
SQLITE_API void sqlite3_value_free(sqlite3_value*);

/*
** CAPI3REF: Obtain Aggregate Function Context
** METHOD: sqlite3_context
**
** Implementations of aggregate SQL functions use this
** routine to allocate memory for storing their state.
**
** ^The first time the sqlite3_aggregate_context(C,N) routine is called 
** for a particular aggregate function, SQLite
** allocates N of memory, zeroes out that memory, and returns a pointer
** to the new memory. ^On second and subsequent calls to
** sqlite3_aggregate_context() for the same aggregate function instance,
** the same buffer is returned.  Sqlite3_aggregate_context() is normally
** called once for each invocation of the xStep callback and then one
** last time when the xFinal callback is invoked.  ^(When no rows match
** an aggregate query, the xStep() callback of the aggregate function
** implementation is never called and xFinal() is called exactly once.
** In those cases, sqlite3_aggregate_context() might be called for the
** first time from within xFinal().)^
**
** ^The sqlite3_aggregate_context(C,N) routine returns a NULL pointer 
** when first called if N is less than or equal to zero or if a memory
** allocate error occurs.
**
** ^(The amount of space allocated by sqlite3_aggregate_context(C,N) is
** determined by the N parameter on first successful call.  Changing the
** value of N in subsequent call to sqlite3_aggregate_context() within
** the same aggregate function instance will not resize the memory
** allocation.)^  Within the xFinal callback, it is customary to set
** N=0 in calls to sqlite3_aggregate_context(C,N) so that no 
** pointless memory allocations occur.
**
** ^SQLite automatically frees the memory allocated by 
** sqlite3_aggregate_context() when the aggregate query concludes.
**
** The first parameter must be a copy of the
** [sqlite3_context | SQL function context] that is the first parameter
** to the xStep or xFinal callback routine that implements the aggregate
** function.
**
** This routine must be called from the same thread in which
** the aggregate SQL function is running.
*/
SQLITE_API void *sqlite3_aggregate_context(sqlite3_context*, int nBytes);

/*
** CAPI3REF: User Data For Functions
** METHOD: sqlite3_context
**
** ^The sqlite3_user_data() interface returns a copy of
** the pointer that was the pUserData parameter (the 5th parameter)
** of the [sqlite3_create_function()]
** and [sqlite3_create_function16()] routines that originally
** registered the application defined function.
**
** This routine must be called from the same thread in which
** the application-defined function is running.
*/
SQLITE_API void *sqlite3_user_data(sqlite3_context*);

/*
** CAPI3REF: Database Connection For Functions
** METHOD: sqlite3_context
**
** ^The sqlite3_context_db_handle() interface returns a copy of
** the pointer to the [database connection] (the 1st parameter)
** of the [sqlite3_create_function()]
** and [sqlite3_create_function16()] routines that originally
** registered the application defined function.
*/
SQLITE_API sqlite3 *sqlite3_context_db_handle(sqlite3_context*);

*/
SQLITE_API int sqlite3_create_collation(
  sqlite3*, 
  const char *zName, 
  int eTextRep, 
  void *pArg,
  int(*xCompare)(void*,int,const void*,int,const void*)
);
SQLITE_API int sqlite3_create_collation_v2(
  sqlite3*, 
  const char *zName, 
  int eTextRep, 
  void *pArg,
  int(*xCompare)(void*,int,const void*,int,const void*),
  void(*xDestroy)(void*)
);
SQLITE_API int sqlite3_create_collation16(
  sqlite3*, 
  const void *zName,
  int eTextRep, 
  void *pArg,
  int(*xCompare)(void*,int,const void*,int,const void*)
);

/*
** CAPI3REF: Collation Needed Callbacks
** METHOD: sqlite3
**
** ^To avoid having to register all collation sequences before a database
** can be used, a single callback function may be registered with the
** [database connection] to be invoked whenever an undefined collation
** sequence is required.
**
** ^If the function is registered using the sqlite3_collation_needed() API,
** then it is passed the names of undefined collation sequences as strings
** encoded in UTF-8. ^If sqlite3_collation_needed16() is used,
** the names are passed as UTF-16 in machine native byte order.
** ^A call to either function replaces the existing collation-needed callback.
**
** ^(When the callback is invoked, the first argument passed is a copy
** of the second argument to sqlite3_collation_needed() or
** sqlite3_collation_needed16().  The second argument is the database
** connection.  The third argument is one of [SQLITE_UTF8], [SQLITE_UTF16BE],
** or [SQLITE_UTF16LE], indicating the most desirable form of the collation
** sequence function required.  The fourth parameter is the name of the
** required collation sequence.)^
**
** The callback function should register the desired collation using
** [sqlite3_create_collation()], [sqlite3_create_collation16()], or
** [sqlite3_create_collation_v2()].
*/
  sqlite3 *db,                   /* Database to be rekeyed */
  const char *zDbName,           /* Name of the database */
  const void *pKey, int nKey     /* The new key */
);

/*
** Specify the activation key for a SEE database.  Unless 
** activated, none of the SEE routines will work.
*/
SQLITE_API void sqlite3_activate_see(
  const char *zPassPhrase        /* Activation phrase */
);
#endif

#ifdef SQLITE_ENABLE_CEROD
/*
** Specify the activation key for a CEROD database.  Unless 
** activated, none of the CEROD routines will work.
*/
SQLITE_API void sqlite3_activate_cerod(
  const char *zPassPhrase        /* Activation phrase */
);
#endif

/*
** CAPI3REF: Suspend Execution For A Short Time
**
** The sqlite3_sleep() function causes the current thread to suspend execution
** for at least a number of milliseconds specified in its parameter.
**
** If the operating system does not support sleep requests with
** millisecond time resolution, then the time will be rounded up to
** the nearest second. The number of milliseconds of sleep actually
** requested from the operating system is returned.
**
** ^SQLite implements this interface by calling the xSleep()
** method of the default [sqlite3_vfs] object.  If the xSleep() method
** of the default VFS is not implemented correctly, or not implemented at
** all, then the behavior of sqlite3_sleep() may deviate from the description
** in the previous paragraphs.
*/
SQLITE_API int sqlite3_sleep(int);

/*
** CAPI3REF: Name Of The Folder Holding Temporary Files
**
** ^(If this global variable is made to point to a string which is
** the name of a folder (a.k.a. directory), then all temporary files
** created by SQLite when using a built-in [sqlite3_vfs | VFS]
** will be placed in that directory.)^  ^If this variable
** is a NULL pointer, then SQLite performs a search for an appropriate
/*
** CAPI3REF: Test For Auto-Commit Mode
** KEYWORDS: {autocommit mode}
** METHOD: sqlite3
**
** ^The sqlite3_get_autocommit() interface returns non-zero or
** zero if the given database connection is or is not in autocommit mode,
** respectively.  ^Autocommit mode is on by default.
** ^Autocommit mode is disabled by a [BEGIN] statement.
** ^Autocommit mode is re-enabled by a [COMMIT] or [ROLLBACK].
**
** If certain kinds of errors occur on a statement within a multi-statement
** transaction (errors including [SQLITE_FULL], [SQLITE_IOERR],
** [SQLITE_NOMEM], [SQLITE_BUSY], and [SQLITE_INTERRUPT]) then the
** transaction might be rolled back automatically.  The only way to
** find out whether SQLite automatically rolled back the transaction after
** an error is to use this function.
**
** If another thread changes the autocommit status of the database
** connection while this routine is running, then the return value
** is undefined.
*/
SQLITE_API int sqlite3_get_autocommit(sqlite3*);

/*
** CAPI3REF: Find The Database Handle Of A Prepared Statement
** METHOD: sqlite3_stmt
**
** ^The sqlite3_db_handle interface returns the [database connection] handle
** to which a [prepared statement] belongs.  ^The [database connection]
** returned by sqlite3_db_handle is the same [database connection]
** that was the first argument
** to the [sqlite3_prepare_v2()] call (or its variants) that was used to
** create the statement in the first place.
*/
SQLITE_API sqlite3 *sqlite3_db_handle(sqlite3_stmt*);

/*
** CAPI3REF: Return The Filename For A Database Connection
** METHOD: sqlite3
**
** ^The sqlite3_db_filename(D,N) interface returns a pointer to a filename
** associated with database N of connection D.  ^The main database file
** has the name "main".  If there is no attached database N on the database
** connection D, or if database N is a temporary or in-memory database, then
** a NULL pointer is returned.
**
** ^The filename returned by this function is the output of the
** xFullPathname method of the [VFS].  ^In other words, the filename
** will be an absolute pathname, even if the filename used
** to open the database originally was a URI or relative pathname.
*/
SQLITE_API const char *sqlite3_db_filename(sqlite3 *db, const char *zDbName);

/*
** CAPI3REF: Determine if a database is read-only
** METHOD: sqlite3
**
** ^The sqlite3_db_readonly(D,N) interface returns 1 if the database N
** of connection D is read-only, 0 if it is read/write, or -1 if N is not
** the name of a database on connection D.
*/
SQLITE_API int sqlite3_db_readonly(sqlite3 *db, const char *zDbName);

/*
** CAPI3REF: Find the next prepared statement
** METHOD: sqlite3
**
** ^This interface returns a pointer to the next [prepared statement] after
** pStmt associated with the [database connection] pDb.  ^If pStmt is NULL
** then this interface returns a pointer to the first prepared statement
** associated with the database connection pDb.  ^If no prepared statement
** satisfies the conditions of this routine, it returns NULL.
**
** The [database connection] pointer D in a call to
** [sqlite3_next_stmt(D,S)] must refer to an open database
** connection and in particular must not be a NULL pointer.
*/
SQLITE_API sqlite3_stmt *sqlite3_next_stmt(sqlite3 *pDb, sqlite3_stmt *pStmt);

/*
** CAPI3REF: Commit And Rollback Notification Callbacks
** METHOD: sqlite3
**
** ^The sqlite3_commit_hook() interface registers a callback
** function to be invoked whenever a transaction is [COMMIT | committed].
** ^Any callback set by a previous call to sqlite3_commit_hook()
** for the same database connection is overridden.
** ^The sqlite3_rollback_hook() interface registers a callback
** function to be invoked whenever a transaction is [ROLLBACK | rolled back].
** ^Any callback set by a previous call to sqlite3_rollback_hook()
**
** This interface is threadsafe on processors where writing a
** 32-bit integer is atomic.
**
** See Also:  [SQLite Shared-Cache Mode]
*/
SQLITE_API int sqlite3_enable_shared_cache(int);

/*
** CAPI3REF: Attempt To Free Heap Memory
**
** ^The sqlite3_release_memory() interface attempts to free N bytes
** of heap memory by deallocating non-essential memory allocations
** held by the database library.   Memory used to cache database
** pages to improve performance is an example of non-essential memory.
** ^sqlite3_release_memory() returns the number of bytes actually freed,
** which might be more or less than the amount requested.
** ^The sqlite3_release_memory() routine is a no-op returning zero
** if SQLite is not compiled with [SQLITE_ENABLE_MEMORY_MANAGEMENT].
**
** See also: [sqlite3_db_release_memory()]
*/
SQLITE_API int sqlite3_release_memory(int);

/*
** CAPI3REF: Free Memory Used By A Database Connection
** METHOD: sqlite3
**
** ^The sqlite3_db_release_memory(D) interface attempts to free as much heap
** memory as possible from database connection D. Unlike the
** [sqlite3_release_memory()] interface, this interface is in effect even
** when the [SQLITE_ENABLE_MEMORY_MANAGEMENT] compile-time option is
** omitted.
**
** See also: [sqlite3_release_memory()]
*/
SQLITE_API int sqlite3_db_release_memory(sqlite3*);

/*
** CAPI3REF: Impose A Limit On Heap Size
**
** ^The sqlite3_soft_heap_limit64() interface sets and/or queries the
** soft limit on the amount of heap memory that may be allocated by SQLite.
** ^SQLite strives to keep heap memory utilization below the soft heap
** limit by reducing the number of pages held in the page cache
** as heap memory usages approaches the limit.
** ^The soft heap limit is "soft" because even though SQLite strives to stay
** below the limit, it will exceed the limit rather than generate
** an [SQLITE_NOMEM] error.  In other words, the soft heap limit 
** is advisory only.
**
** ^The return value from sqlite3_soft_heap_limit64() is the size of
** the soft heap limit prior to the call, or negative in the case of an
** error.  ^If the argument N is negative
** then no change is made to the soft heap limit.  Hence, the current
** size of the soft heap limit can be determined by invoking
** sqlite3_soft_heap_limit64() with a negative argument.
**
** ^If the argument N is zero then the soft heap limit is disabled.
**
** ^(The soft heap limit is not enforced in the current implementation
** if one or more of following conditions are true:
**
** <ul>
** <li> The soft heap limit is set to zero.
**      the [SQLITE_DEFAULT_MEMSTATUS] compile-time option.
** <li> An alternative page cache implementation is specified using
**      [sqlite3_config]([SQLITE_CONFIG_PCACHE2],...).
** <li> The page cache allocates from its own memory pool supplied
**      by [sqlite3_config]([SQLITE_CONFIG_PAGECACHE],...) rather than
**      from the heap.
** </ul>)^
**
** Beginning with SQLite [version 3.7.3] ([dateof:3.7.3]), 
** the soft heap limit is enforced
** regardless of whether or not the [SQLITE_ENABLE_MEMORY_MANAGEMENT]
** compile-time option is invoked.  With [SQLITE_ENABLE_MEMORY_MANAGEMENT],
** the soft heap limit is enforced on every memory allocation.  Without
** [SQLITE_ENABLE_MEMORY_MANAGEMENT], the soft heap limit is only enforced
** when memory is allocated by the page cache.  Testing suggests that because
** the page cache is the predominate memory user in SQLite, most
** applications will achieve adequate soft heap limit enforcement without
** the use of [SQLITE_ENABLE_MEMORY_MANAGEMENT].
**
** The circumstances under which SQLite will enforce the soft heap limit may
** changes in future releases of SQLite.
*/
SQLITE_API sqlite3_int64 sqlite3_soft_heap_limit64(sqlite3_int64 N);

/*
** CAPI3REF: Deprecated Soft Heap Limit Interface
** DEPRECATED
**
** This is a deprecated version of the [sqlite3_soft_heap_limit64()]
** interface.  This routine is provided for historical compatibility
** only.  All new applications should use the
** [sqlite3_soft_heap_limit64()] interface rather than this one.
*/
SQLITE_API SQLITE_DEPRECATED void sqlite3_soft_heap_limit(int N);


/*
** CAPI3REF: Extract Metadata About A Column Of A Table
** METHOD: sqlite3
**
** ^(The sqlite3_table_column_metadata(X,D,T,C,....) routine returns
** information about column C of table T in database D
** on [database connection] X.)^  ^The sqlite3_table_column_metadata()
** interface returns SQLITE_OK and fills in the non-NULL pointers in
** the final five arguments with appropriate values if the specified
** column exists.  ^The sqlite3_table_column_metadata() interface returns
** SQLITE_ERROR and if the specified column does not exist.
** ^If the column-name parameter to sqlite3_table_column_metadata() is a
** NULL pointer, then this routine simply checks for the existence of the
** table and returns SQLITE_OK if the table exists and SQLITE_ERROR if it
** does not.  If the table name parameter T in a call to
** <pre>
**     data type: "INTEGER"
**     collation sequence: "BINARY"
**     not null: 0
**     primary key: 1
**     auto increment: 0
** </pre>)^
**
** ^This function causes all database schemas to be read from disk and
** parsed, if that has not already been done, and returns an error if
** any errors are encountered while loading the schema.
*/
SQLITE_API int sqlite3_table_column_metadata(
  sqlite3 *db,                /* Connection handle */
  const char *zDbName,        /* Database name or NULL */
  const char *zTableName,     /* Table name */
  const char *zColumnName,    /* Column name */
  char const **pzDataType,    /* OUTPUT: Declared data type */
  char const **pzCollSeq,     /* OUTPUT: Collation sequence name */
  int *pNotNull,              /* OUTPUT: True if NOT NULL constraint exists */
  int *pPrimaryKey,           /* OUTPUT: True if column part of PK */
  int *pAutoinc               /* OUTPUT: True if column is auto-increment */
);

/*
** CAPI3REF: Load An Extension
** METHOD: sqlite3
**
** ^This interface loads an SQLite extension library from the named file.
**
** ^The sqlite3_load_extension() interface attempts to load an
** [SQLite extension] library contained in the file zFile.  If
** the file cannot be loaded directly, attempts are made to load
** with various operating-system specific extensions added.
** So for example, if "samplelib" cannot be loaded, then names like
** "samplelib.so" or "samplelib.dylib" or "samplelib.dll" might
** be tried also.
**
** ^The entry point is zProc.
** ^(zProc may be 0, in which case SQLite will try to come up with an
** entry point name on its own.  It first tries "sqlite3_extension_init".
** If that does not work, it constructs a name "sqlite3_X_init" where the
** X is consists of the lower-case equivalent of all ASCII alphabetic
** characters in the filename from the last "/" to the first following
** "." and omitting any initial "lib".)^
** ^The sqlite3_load_extension() interface returns
** [SQLITE_OK] on success and [SQLITE_ERROR] if something goes wrong.
** ^If an error occurs and pzErrMsg is not 0, then the
** [sqlite3_load_extension()] interface shall attempt to
** fill *pzErrMsg with error message text stored in memory
** obtained from [sqlite3_malloc()]. The calling function
** [sqlite3_module.xOpen | xOpen] method of the module and are destroyed
** by the [sqlite3_module.xClose | xClose] method.  Cursors are used
** by the [xFilter], [xNext], [xEof], [xColumn], and [xRowid] methods
** of the module.  Each module implementation will define
** the content of a cursor structure to suit its own needs.
**
** This superclass exists in order to define fields of the cursor that
** are common to all implementations.
*/
struct sqlite3_vtab_cursor {
  sqlite3_vtab *pVtab;      /* Virtual table of this cursor */
  /* Virtual table implementations will typically add additional fields */
};

/*
** CAPI3REF: Declare The Schema Of A Virtual Table
**
** ^The [xCreate] and [xConnect] methods of a
** [virtual table module] call this interface
** to declare the format (the names and datatypes of the columns) of
** the virtual tables they implement.
*/
SQLITE_API int sqlite3_declare_vtab(sqlite3*, const char *zSQL);

/*
** CAPI3REF: Overload A Function For A Virtual Table
** METHOD: sqlite3
**
** ^(Virtual tables can provide alternative implementations of functions
** using the [xFindFunction] method of the [virtual table module].  
** But global versions of those functions
** must exist in order to be overloaded.)^
**
** ^(This API makes sure a global version of a function with a particular
** name and number of parameters exists.  If no such function exists
** before this API is called, a new function is created.)^  ^The implementation
** of the new function always causes an exception to be thrown.  So
** the new function is not good for anything by itself.  Its only
** purpose is to be a placeholder function that can be overloaded
** by a [virtual table].
*/
SQLITE_API int sqlite3_overload_function(sqlite3*, const char *zFuncName, int nArg);

/*
** The interface to the virtual-table mechanism defined above (back up
** to a comment remarkably similar to this one) is currently considered
** to be experimental.  The interface might change in incompatible ways.
** If this is a problem for you, do not use the interface at this time.
**
** When the virtual-table mechanism stabilizes, we will declare the
** interface fixed, support it indefinitely, and remove this comment.
*/

/*
** CAPI3REF: A Handle To An Open BLOB
** KEYWORDS: {BLOB handle} {BLOB handles}
**
** An instance of this object represents an open BLOB on which
** [sqlite3_blob_open | incremental BLOB I/O] can be performed.
** ^Objects of this type are created by [sqlite3_blob_open()]
** and destroyed by [sqlite3_blob_close()].
** ^The [sqlite3_blob_read()] and [sqlite3_blob_write()] interfaces
** can be used to read or write small subsections of the BLOB.
** ^The [sqlite3_blob_bytes()] interface returns the size of the BLOB in bytes.
*/
typedef struct sqlite3_blob sqlite3_blob;

/*
** CAPI3REF: Open A BLOB For Incremental I/O
** METHOD: sqlite3
** CONSTRUCTOR: sqlite3_blob
**
** ^(This interfaces opens a [BLOB handle | handle] to the BLOB located
** in row iRow, column zColumn, table zTable in database zDb;
** in other words, the same BLOB that would be selected by:
**
** <pre>
**     SELECT zColumn FROM zDb.zTable WHERE [rowid] = iRow;
** </pre>)^
**
** ^(Parameter zDb is not the filename that contains the database, but 
** rather the symbolic name of the database. For attached databases, this is
** the name that appears after the AS keyword in the [ATTACH] statement.
** For the main database file, the database name is "main". For TEMP
** tables, the database name is "temp".)^
**
** ^If the flags parameter is non-zero, then the BLOB is opened for read
** and write access. ^If the flags parameter is zero, the BLOB is opened for
** read-only access.
**
** ^(On success, [SQLITE_OK] is returned and the new [BLOB handle] is stored
** in *ppBlob. Otherwise an [error code] is returned and, unless the error
** code is SQLITE_MISUSE, *ppBlob is set to NULL.)^ ^This means that, provided
** the API is not misused, it is always safe to call [sqlite3_blob_close()] 
** interface.  Use the [UPDATE] SQL command to change the size of a
** blob.
**
** ^The [sqlite3_bind_zeroblob()] and [sqlite3_result_zeroblob()] interfaces
** and the built-in [zeroblob] SQL function may be used to create a 
** zero-filled blob to read or write using the incremental-blob interface.
**
** To avoid a resource leak, every open [BLOB handle] should eventually
** be released by a call to [sqlite3_blob_close()].
**
** See also: [sqlite3_blob_close()],
** [sqlite3_blob_reopen()], [sqlite3_blob_read()],
** [sqlite3_blob_bytes()], [sqlite3_blob_write()].
*/
SQLITE_API int sqlite3_blob_open(
  sqlite3*,
  const char *zDb,
  const char *zTable,
  const char *zColumn,
  sqlite3_int64 iRow,
  int flags,
  sqlite3_blob **ppBlob
);

/*
** CAPI3REF: Move a BLOB Handle to a New Row
** METHOD: sqlite3_blob
**
** ^This function is used to move an existing [BLOB handle] so that it points
** to a different row of the same database table. ^The new row is identified
** by the rowid value passed as the second argument. Only the row can be
** changed. ^The database, table and column on which the blob handle is open
** remain the same. Moving an existing [BLOB handle] to a new row is
** faster than closing the existing handle and opening a new one.
**
** ^(The new row must meet the same criteria as for [sqlite3_blob_open()] -
** it must exist and there must be either a blob or text value stored in
** the nominated column.)^ ^If the new row is not present in the table, or if
** it does not contain a blob or text value, or if another error occurs, an
** SQLite error code is returned and the blob handle is considered aborted.
** ^All subsequent calls to [sqlite3_blob_read()], [sqlite3_blob_write()] or
** [sqlite3_blob_reopen()] on an aborted blob handle immediately return
** SQLITE_ABORT. ^Calling [sqlite3_blob_bytes()] on an aborted blob handle
** always returns zero.
**
** ^This function sets the database handle error code and message.
*/
SQLITE_API int sqlite3_blob_reopen(sqlite3_blob *, sqlite3_int64);

/*
** CAPI3REF: Close A BLOB Handle
** open blob handle results in undefined behaviour. ^Calling this routine 
** with a null pointer (such as would be returned by a failed call to 
** [sqlite3_blob_open()]) is a harmless no-op. ^Otherwise, if this function
** is passed a valid open blob handle, the values returned by the 
** sqlite3_errcode() and sqlite3_errmsg() functions are set before returning.
*/
SQLITE_API int sqlite3_blob_close(sqlite3_blob *);

/*
** CAPI3REF: Return The Size Of An Open BLOB
** METHOD: sqlite3_blob
**
** ^Returns the size in bytes of the BLOB accessible via the 
** successfully opened [BLOB handle] in its only argument.  ^The
** incremental blob I/O routines can only read or overwriting existing
** blob content; they cannot change the size of a blob.
**
** This routine only works on a [BLOB handle] which has been created
** by a prior successful call to [sqlite3_blob_open()] and which has not
** been closed by [sqlite3_blob_close()].  Passing any other pointer in
** to this routine results in undefined and probably undesirable behavior.
*/
SQLITE_API int sqlite3_blob_bytes(sqlite3_blob *);

/*
** CAPI3REF: Read Data From A BLOB Incrementally
** METHOD: sqlite3_blob
**
** ^(This function is used to read data from an open [BLOB handle] into a
** caller-supplied buffer. N bytes of data are copied into buffer Z
** from the open BLOB, starting at offset iOffset.)^
**
** ^If offset iOffset is less than N bytes from the end of the BLOB,
** [SQLITE_ERROR] is returned and no data is read.  ^If N or iOffset is
** less than zero, [SQLITE_ERROR] is returned and no data is read.
** ^The size of the blob (and hence the maximum value of N+iOffset)
** can be determined using the [sqlite3_blob_bytes()] interface.
**
** ^An attempt to read from an expired [BLOB handle] fails with an
** error code of [SQLITE_ABORT].
**
** ^(On success, sqlite3_blob_read() returns SQLITE_OK.
** Otherwise, an [error code] or an [extended error code] is returned.)^
**
** This routine only works on a [BLOB handle] which has been created
** by a prior successful call to [sqlite3_blob_open()] and which has not
** been closed by [sqlite3_blob_close()].  Passing any other pointer in
** to this routine results in undefined and probably undesirable behavior.
**
** See also: [sqlite3_blob_write()].
*/
**
** ^The sqlite3_vfs_find() interface returns a pointer to a VFS given its name.
** ^Names are case sensitive.
** ^Names are zero-terminated UTF-8 strings.
** ^If there is no match, a NULL pointer is returned.
** ^If zVfsName is NULL then the default VFS is returned.
**
** ^New VFSes are registered with sqlite3_vfs_register().
** ^Each new VFS becomes the default VFS if the makeDflt flag is set.
** ^The same VFS can be registered multiple times without injury.
** ^To make an existing VFS into the default VFS, register it again
** with the makeDflt flag set.  If two different VFSes with the
** same name are registered, the behavior is undefined.  If a
** VFS is registered with a name that is NULL or an empty string,
** then the behavior is undefined.
**
** ^Unregister a VFS with the sqlite3_vfs_unregister() interface.
** ^(If the default VFS is unregistered, another VFS is chosen as
** the default.  The choice for the new VFS is arbitrary.)^
*/
SQLITE_API sqlite3_vfs *sqlite3_vfs_find(const char *zVfsName);
SQLITE_API int sqlite3_vfs_register(sqlite3_vfs*, int makeDflt);
SQLITE_API int sqlite3_vfs_unregister(sqlite3_vfs*);

/*
** CAPI3REF: Mutexes
**
** The SQLite core uses these routines for thread
** synchronization. Though they are intended for internal
** use by SQLite, code that links against SQLite is
** permitted to use any of these routines.
**
** The SQLite source code contains multiple implementations
** of these mutex routines.  An appropriate implementation
** is selected automatically at compile-time.  The following
** implementations are available in the SQLite core:
**
** <ul>
** <li>   SQLITE_MUTEX_PTHREADS
** <li>   SQLITE_MUTEX_W32
** <li>   SQLITE_MUTEX_NOOP
** </ul>
**
** The SQLITE_MUTEX_NOOP implementation is a set of routines
** that does no real locking and is appropriate for use in
** a single-threaded application.  The SQLITE_MUTEX_PTHREADS and
** SQLITE_MUTEX_W32 implementations are appropriate for use on Unix
** and Windows.
**
** If SQLite is compiled with the SQLITE_MUTEX_APPDEF preprocessor
** macro defined (with "-DSQLITE_MUTEX_APPDEF=1"), then no mutex
**
** ^(Some systems (for example, Windows 95) do not support the operation
** implemented by sqlite3_mutex_try().  On those systems, sqlite3_mutex_try()
** will always return SQLITE_BUSY. The SQLite core only ever uses
** sqlite3_mutex_try() as an optimization so this is acceptable 
** behavior.)^
**
** ^The sqlite3_mutex_leave() routine exits a mutex that was
** previously entered by the same thread.   The behavior
** is undefined if the mutex is not currently entered by the
** calling thread or is not currently allocated.
**
** ^If the argument to sqlite3_mutex_enter(), sqlite3_mutex_try(), or
** sqlite3_mutex_leave() is a NULL pointer, then all three routines
** behave as no-ops.
**
** See also: [sqlite3_mutex_held()] and [sqlite3_mutex_notheld()].
*/
SQLITE_API sqlite3_mutex *sqlite3_mutex_alloc(int);
SQLITE_API void sqlite3_mutex_free(sqlite3_mutex*);
SQLITE_API void sqlite3_mutex_enter(sqlite3_mutex*);
SQLITE_API int sqlite3_mutex_try(sqlite3_mutex*);
SQLITE_API void sqlite3_mutex_leave(sqlite3_mutex*);

/*
** CAPI3REF: Mutex Methods Object
**
** An instance of this structure defines the low-level routines
** used to allocate and use mutexes.
**
** Usually, the default mutex implementations provided by SQLite are
** sufficient, however the application has the option of substituting a custom
** implementation for specialized deployments or systems for which SQLite
** does not provide a suitable implementation. In this case, the application
** creates and populates an instance of this structure to pass
** to sqlite3_config() along with the [SQLITE_CONFIG_MUTEX] option.
** Additionally, an instance of this structure can be used as an
** output variable when querying the system for the current mutex
** implementation, using the [SQLITE_CONFIG_GETMUTEX] option.
**
** ^The xMutexInit method defined by this structure is invoked as
** part of system initialization by the sqlite3_initialize() function.
** ^The xMutexInit routine is called by SQLite exactly once for each
** effective call to [sqlite3_initialize()].
**
** ^The xMutexEnd method defined by this structure is invoked as
** part of system shutdown by the sqlite3_shutdown() function. The
** implementation of this method is expected to release all outstanding
** resources obtained by the mutex methods implementation, especially
** those obtained by the xMutexInit method.  ^The xMutexEnd()
** interface is invoked exactly once for each call to [sqlite3_shutdown()].
**
** xMutexInit() must not use SQLite memory allocation ([sqlite3_malloc()]
** and its associates).  Similarly, xMutexAlloc() must not use SQLite memory
** allocation for a static mutex.  ^However xMutexAlloc() may use SQLite
** memory allocation for a fast or recursive mutex.
**
** ^SQLite will invoke the xMutexEnd() method when [sqlite3_shutdown()] is
** called, but only if the prior call to xMutexInit returned SQLITE_OK.
** If xMutexInit fails in any way, it is expected to clean up after itself
** prior to returning.
*/
typedef struct sqlite3_mutex_methods sqlite3_mutex_methods;
struct sqlite3_mutex_methods {
  int (*xMutexInit)(void);
  int (*xMutexEnd)(void);
  sqlite3_mutex *(*xMutexAlloc)(int);
  void (*xMutexFree)(sqlite3_mutex *);
  void (*xMutexEnter)(sqlite3_mutex *);
  int (*xMutexTry)(sqlite3_mutex *);
  void (*xMutexLeave)(sqlite3_mutex *);
  int (*xMutexHeld)(sqlite3_mutex *);
  int (*xMutexNotheld)(sqlite3_mutex *);
};

/*
** CAPI3REF: Mutex Verification Routines
**
** The sqlite3_mutex_held() and sqlite3_mutex_notheld() routines
** are intended for use inside assert() statements.  The SQLite core
** never uses these routines except inside an assert() and applications
** are advised to follow the lead of the core.  The SQLite core only
** provides implementations for these routines when it is compiled
** with the SQLITE_DEBUG flag.  External mutex implementations
** are only required to provide these routines if SQLITE_DEBUG is
** defined and if NDEBUG is not defined.
**
** These routines should return true if the mutex in their argument
** is held or not held, respectively, by the calling thread.
**
** The implementation is not required to provide versions of these
** routines that actually work. If the implementation does not provide working
** versions of these routines, it should at least provide stubs that always
** return true so that one does not get spurious assertion failures.
**
** If the argument to sqlite3_mutex_held() is a NULL pointer then
** the routine should return 1.   This seems counter-intuitive since
** clearly the mutex cannot be held if it does not exist.  But
** the reason the mutex does not exist is because the build is not
** using mutexes.  And we do not want the assert() containing the
** call to sqlite3_mutex_held() to fail, so a non-zero return is
** the appropriate thing to do.  The sqlite3_mutex_notheld()
** interface should also return 1 when given a NULL pointer.
*/
#ifndef NDEBUG
SQLITE_API int sqlite3_mutex_held(sqlite3_mutex*);
SQLITE_API int sqlite3_mutex_notheld(sqlite3_mutex*);
#endif

/*
** CAPI3REF: Mutex Types
**
** The [sqlite3_mutex_alloc()] interface takes a single argument
** which is one of these integer constants.
**
** The set of static mutexes may change from one SQLite release to the
** next.  Applications that override the built-in mutex logic must be
** prepared to accommodate additional static mutexes.
*/
#define SQLITE_MUTEX_FAST             0
#define SQLITE_MUTEX_RECURSIVE        1
#define SQLITE_MUTEX_STATIC_MASTER    2
#define SQLITE_MUTEX_STATIC_MEM       3  /* sqlite3_malloc() */
#define SQLITE_MUTEX_STATIC_MEM2      4  /* NOT USED */
#define SQLITE_MUTEX_STATIC_OPEN      4  /* sqlite3BtreeOpen() */
#define SQLITE_MUTEX_STATIC_PRNG      5  /* sqlite3_randomness() */
#define SQLITE_MUTEX_STATIC_LRU       6  /* lru page list */
#define SQLITE_MUTEX_STATIC_LRU2      7  /* NOT USED */
#define SQLITE_MUTEX_STATIC_PMEM      7  /* sqlite3PageMalloc() */
#define SQLITE_MUTEX_STATIC_APP1      8  /* For use by application */
#define SQLITE_MUTEX_STATIC_APP2      9  /* For use by application */
#define SQLITE_MUTEX_STATIC_APP3     10  /* For use by application */
#define SQLITE_MUTEX_STATIC_VFS1     11  /* For use by built-in VFS */
#define SQLITE_MUTEX_STATIC_VFS2     12  /* For use by extension VFS */
#define SQLITE_MUTEX_STATIC_VFS3     13  /* For use by application VFS */

/*
** CAPI3REF: Retrieve the mutex for a database connection
** METHOD: sqlite3
**
** ^This interface returns a pointer the [sqlite3_mutex] object that 
** serializes access to the [database connection] given in the argument
** when the [threading mode] is Serialized.
** ^If the [threading mode] is Single-thread or Multi-thread then this
** routine returns a NULL pointer.
*/
SQLITE_API sqlite3_mutex *sqlite3_db_mutex(sqlite3*);

/*
** CAPI3REF: Low-Level Control Of Database Files
** METHOD: sqlite3
**
** ^The [sqlite3_file_control()] interface makes a direct call to the
** xFileControl method for the [sqlite3_io_methods] object associated
** with a particular database identified by the second argument. ^The
** name of the database is "main" for the main database or "temp" for the
** TEMP database, or the name that appears after the AS keyword for
** databases that are added using the [ATTACH] SQL command.
** ^A NULL pointer can be used in place of "main" to refer to the
** main database file.
** ^The third and fourth parameters to this routine
** are passed directly through to the second and third parameters of
** the xFileControl method.  ^The return value of the xFileControl
#define SQLITE_TESTCTRL_PRNG_RESTORE             6
#define SQLITE_TESTCTRL_PRNG_RESET               7
#define SQLITE_TESTCTRL_BITVEC_TEST              8
#define SQLITE_TESTCTRL_FAULT_INSTALL            9
#define SQLITE_TESTCTRL_BENIGN_MALLOC_HOOKS     10
#define SQLITE_TESTCTRL_PENDING_BYTE            11
#define SQLITE_TESTCTRL_ASSERT                  12
#define SQLITE_TESTCTRL_ALWAYS                  13
#define SQLITE_TESTCTRL_RESERVE                 14
#define SQLITE_TESTCTRL_OPTIMIZATIONS           15
#define SQLITE_TESTCTRL_ISKEYWORD               16
#define SQLITE_TESTCTRL_SCRATCHMALLOC           17  /* NOT USED */
#define SQLITE_TESTCTRL_LOCALTIME_FAULT         18
#define SQLITE_TESTCTRL_EXPLAIN_STMT            19  /* NOT USED */
#define SQLITE_TESTCTRL_ONCE_RESET_THRESHOLD    19
#define SQLITE_TESTCTRL_NEVER_CORRUPT           20
#define SQLITE_TESTCTRL_VDBE_COVERAGE           21
#define SQLITE_TESTCTRL_BYTEORDER               22
#define SQLITE_TESTCTRL_ISINIT                  23
#define SQLITE_TESTCTRL_SORTER_MMAP             24
#define SQLITE_TESTCTRL_IMPOSTER                25
#define SQLITE_TESTCTRL_PARSER_COVERAGE         26
#define SQLITE_TESTCTRL_LAST                    26  /* Largest TESTCTRL */

/*
** CAPI3REF: SQLite Runtime Status
**
** ^These interfaces are used to retrieve runtime status information
** about the performance of SQLite, and optionally to reset various
** highwater marks.  ^The first argument is an integer code for
** the specific parameter to measure.  ^(Recognized integer codes
** are of the form [status parameters | SQLITE_STATUS_...].)^
** ^The current value of the parameter is returned into *pCurrent.
** ^The highest recorded value is returned in *pHighwater.  ^If the
** resetFlag is true, then the highest record value is reset after
** *pHighwater is written.  ^(Some parameters do not record the highest
** value.  For those parameters
** nothing is written into *pHighwater and the resetFlag is ignored.)^
** ^(Other parameters record only the highwater mark and not the current
** value.  For these latter parameters nothing is written into *pCurrent.)^
**
** ^The sqlite3_status() and sqlite3_status64() routines return
** SQLITE_OK on success and a non-zero [error code] on failure.
**
** If either the current value or the highwater mark is too large to
** be represented by a 32-bit integer, then the values returned by
** sqlite3_status() are undefined.
**
** See also: [sqlite3_db_status()]
*/
SQLITE_API int sqlite3_status(int op, int *pCurrent, int *pHighwater, int resetFlag);
  void (*xShrink)(sqlite3_pcache*);
};

/*
** This is the obsolete pcache_methods object that has now been replaced
** by sqlite3_pcache_methods2.  This object is not used by SQLite.  It is
** retained in the header file for backwards compatibility only.
*/
typedef struct sqlite3_pcache_methods sqlite3_pcache_methods;
struct sqlite3_pcache_methods {
  void *pArg;
  int (*xInit)(void*);
  void (*xShutdown)(void*);
  sqlite3_pcache *(*xCreate)(int szPage, int bPurgeable);
  void (*xCachesize)(sqlite3_pcache*, int nCachesize);
  int (*xPagecount)(sqlite3_pcache*);
  void *(*xFetch)(sqlite3_pcache*, unsigned key, int createFlag);
  void (*xUnpin)(sqlite3_pcache*, void*, int discard);
  void (*xRekey)(sqlite3_pcache*, void*, unsigned oldKey, unsigned newKey);
  void (*xTruncate)(sqlite3_pcache*, unsigned iLimit);
  void (*xDestroy)(sqlite3_pcache*);
};


/*
** CAPI3REF: Online Backup Object
**
** The sqlite3_backup object records state information about an ongoing
** online backup operation.  ^The sqlite3_backup object is created by
** a call to [sqlite3_backup_init()] and is destroyed by a call to
** [sqlite3_backup_finish()].
**
** See Also: [Using the SQLite Online Backup API]
*/
typedef struct sqlite3_backup sqlite3_backup;

/*
** CAPI3REF: Online Backup API.
**
** The backup API copies the content of one database into another.
** It is useful either for creating backups of databases or
** for copying in-memory databases to or from persistent files. 
**
** See Also: [Using the SQLite Online Backup API]
**
** ^SQLite holds a write transaction open on the destination database file
** for the duration of the backup operation.
** ^The source database is read-locked only while it is being read;
** it is not locked continuously for the entire backup operation.
** ^Thus, the backup may be performed on a live source database without
** preventing other database connections from
** reading or writing to the source database while the backup is underway.
** 
** ^(To perform a backup operation: 
**   <ol>
**     <li><b>sqlite3_backup_init()</b> is called once to initialize the
**         backup, 
**     <li><b>sqlite3_backup_step()</b> is called one or more times to transfer 
**         the data between the two databases, and finally
**     <li><b>sqlite3_backup_finish()</b> is called to release all resources 
**         associated with the backup operation. 
**   </ol>)^
** There should be exactly one call to sqlite3_backup_finish() for each
** ON CONFLICT mode specified.
**
** If X is non-zero, then the virtual table implementation guarantees
** that if [xUpdate] returns [SQLITE_CONSTRAINT], it will do so before
** any modifications to internal or persistent data structures have been made.
** If the [ON CONFLICT] mode is ABORT, FAIL, IGNORE or ROLLBACK, SQLite 
** is able to roll back a statement or database transaction, and abandon
** or continue processing the current SQL statement as appropriate. 
** If the ON CONFLICT mode is REPLACE and the [xUpdate] method returns
** [SQLITE_CONSTRAINT], SQLite handles this as if the ON CONFLICT mode
** had been ABORT.
**
** Virtual table implementations that are required to handle OR REPLACE
** must do so within the [xUpdate] method. If a call to the 
** [sqlite3_vtab_on_conflict()] function indicates that the current ON 
** CONFLICT policy is REPLACE, the virtual table implementation should 
** silently replace the appropriate rows within the xUpdate callback and
** return SQLITE_OK. Or, if this is not possible, it may return
** SQLITE_CONSTRAINT, in which case SQLite falls back to OR ABORT 
** constraint handling.
** </dl>
*/
#define SQLITE_VTAB_CONSTRAINT_SUPPORT 1

/*
** CAPI3REF: Determine The Virtual Table Conflict Policy
**
** This function may only be called from within a call to the [xUpdate] method
** of a [virtual table] implementation for an INSERT or UPDATE operation. ^The
** value returned is one of [SQLITE_ROLLBACK], [SQLITE_IGNORE], [SQLITE_FAIL],
** [SQLITE_ABORT], or [SQLITE_REPLACE], according to the [ON CONFLICT] mode
** of the SQL statement that triggered the call to the [xUpdate] method of the
** [virtual table].
*/
SQLITE_API int sqlite3_vtab_on_conflict(sqlite3 *);

/*
** CAPI3REF: Determine If Virtual Table Column Access Is For UPDATE
**
** If the sqlite3_vtab_nochange(X) routine is called within the [xColumn]
** method of a [virtual table], then it returns true if and only if the
** column is being fetched as part of an UPDATE operation during which the
** column value will not change.  Applications might use this to substitute
** a lighter-weight value to return that the corresponding [xUpdate] method
** understands as a "no-change" value.
**
** If the [xColumn] method calls sqlite3_vtab_nochange() and finds that
** the column is not changed by the UPDATE statement, they the xColumn
** method can optionally return without setting a result, without calling
** any of the [sqlite3_result_int|sqlite3_result_xxxxx() interfaces].
** In that case, [sqlite3_value_nochange(X)] will return true for the
** same column in the [xUpdate] method.
*/
SQLITE_API int sqlite3_vtab_nochange(sqlite3_context*);

/*
** CAPI3REF: Determine The Collation For a Virtual Table Constraint
**
** This function may only be called from within a call to the [xBestIndex]
** method of a [virtual table]. 
**
** The first argument must be the sqlite3_index_info object that is the
** first parameter to the xBestIndex() method. The second argument must be
** an index into the aConstraint[] array belonging to the sqlite3_index_info
** structure passed to xBestIndex. This function returns a pointer to a buffer 
** containing the name of the collation sequence for the corresponding
** constraint.
*/
SQLITE_API SQLITE_EXPERIMENTAL const char *sqlite3_vtab_collation(sqlite3_index_info*,int);

/*
** CAPI3REF: Conflict resolution modes
** KEYWORDS: {conflict resolution mode}
**
** These constants are returned by [sqlite3_vtab_on_conflict()] to
** inform a [virtual table] implementation what the [ON CONFLICT] mode
** is for the SQL statement being evaluated.
**
** Note that the [SQLITE_IGNORE] constant is also used as a potential
** return value from the [sqlite3_set_authorizer()] callback and that
** [SQLITE_ABORT] is also a [result code].
*/
#define SQLITE_ROLLBACK 1
/* #define SQLITE_IGNORE 2 // Also used by sqlite3_authorizer() callback */
#define SQLITE_FAIL     3
/* #define SQLITE_ABORT 4  // Also an error code */
#define SQLITE_REPLACE  5

/*
** CAPI3REF: Prepared Statement Scan Status Opcodes
** KEYWORDS: {scanstatus options}
**
** The following constants can be used for the T parameter to the
** [sqlite3_stmt_scanstatus(S,X,T,V)] interface.  Each constant designates a
** different metric for sqlite3_stmt_scanstatus() to return.
**
** When the value returned to V is a string, space to hold that string is
** The "iScanStatusOp" must be one of the [scanstatus options] or the behavior
** of this interface is undefined.
** ^The requested measurement is written into a variable pointed to by
** the "pOut" parameter.
** Parameter "idx" identifies the specific loop to retrieve statistics for.
** Loops are numbered starting from zero. ^If idx is out of range - less than
** zero or greater than or equal to the total number of loops used to implement
** the statement - a non-zero value is returned and the variable that pOut
** points to is unchanged.
**
** ^Statistics might not be available for all loops in all statements. ^In cases
** where there exist loops with no available statistics, this function behaves
** as if the loop did not exist - it returns non-zero and leave the variable
** that pOut points to unchanged.
**
** See also: [sqlite3_stmt_scanstatus_reset()]
*/
SQLITE_API int sqlite3_stmt_scanstatus(
  sqlite3_stmt *pStmt,      /* Prepared statement for which info desired */
  int idx,                  /* Index of loop to report on */
  int iScanStatusOp,        /* Information desired.  SQLITE_SCANSTAT_* */
  void *pOut                /* Result written here */
);     

/*
** CAPI3REF: Zero Scan-Status Counters
** METHOD: sqlite3_stmt
**
** ^Zero all [sqlite3_stmt_scanstatus()] related event counters.
**
** This API is only available if the library is built with pre-processor
** symbol [SQLITE_ENABLE_STMT_SCANSTATUS] defined.
*/
SQLITE_API void sqlite3_stmt_scanstatus_reset(sqlite3_stmt*);

/*
** CAPI3REF: Flush caches to disk mid-transaction
**
** ^If a write-transaction is open on [database connection] D when the
** [sqlite3_db_cacheflush(D)] interface invoked, any dirty
** pages in the pager-cache that are not currently in use are written out 
** to disk. A dirty page may be in use if a database cursor created by an
** active SQL statement is reading from it, or if it is page 1 of a database
** file (page 1 is always "in use").  ^The [sqlite3_db_cacheflush(D)]
** interface flushes caches for all schemas - "main", "temp", and
** any [attached] databases.
**
** ^If this function needs to obtain extra database locks before dirty pages 
** can be flushed to disk, it does so. ^If those locks cannot be obtained 
** immediately and there is a busy-handler callback configured, it is invoked
** in the usual manner. ^If the required lock still cannot be obtained, then
**
** <ul>
**   <li> The session object "indirect" flag is set when the change is
**        made, or
**   <li> The change is made by an SQL trigger or foreign key action 
**        instead of directly as a result of a users SQL statement.
** </ul>
**
** If a single row is affected by more than one operation within a session,
** then the change is considered indirect if all operations meet the criteria
** for an indirect change above, or direct otherwise.
**
** This function is used to set, clear or query the session object indirect
** flag.  If the second argument passed to this function is zero, then the
** indirect flag is cleared. If it is greater than zero, the indirect flag
** is set. Passing a value less than zero does not modify the current value
** of the indirect flag, and may be used to query the current state of the 
** indirect flag for the specified session object.
**
** The return value indicates the final state of the indirect flag: 0 if 
** it is clear, or 1 if it is set.
*/
SQLITE_API int sqlite3session_indirect(sqlite3_session *pSession, int bIndirect);

/*
** CAPI3REF: Attach A Table To A Session Object
** METHOD: sqlite3_session
**
** If argument zTab is not NULL, then it is the name of a table to attach
** to the session object passed as the first argument. All subsequent changes 
** made to the table while the session object is enabled will be recorded. See 
** documentation for [sqlite3session_changeset()] for further details.
**
** Or, if argument zTab is NULL, then changes are recorded for all tables
** in the database. If additional tables are added to the database (by 
** executing "CREATE TABLE" statements) after this call is made, changes for 
** the new tables are also recorded.
**
** Changes can only be recorded for tables that have a PRIMARY KEY explicitly
** defined as part of their CREATE TABLE statement. It does not matter if the 
** PRIMARY KEY is an "INTEGER PRIMARY KEY" (rowid alias) or not. The PRIMARY
** KEY may consist of a single column, or may be a composite key.
** 
** It is not an error if the named table does not exist in the database. Nor
** is it an error if the named table does not have a PRIMARY KEY. However,
** no changes will be recorded in either of these scenarios.
**
** Changes are not recorded for individual rows that have NULL values stored
** in one or more of their PRIMARY KEY columns.
**
** SQLITE_OK is returned if the call completes without error. Or, if an error 
** This means, amongst other things, that if a row is inserted and then later
** deleted while a session object is active, neither the insert nor the delete
** will be present in the changeset. Or if a row is deleted and then later a 
** row with the same primary key values inserted while a session object is
** active, the resulting changeset will contain an UPDATE change instead of
** a DELETE and an INSERT.
**
** When a session object is disabled (see the [sqlite3session_enable()] API),
** it does not accumulate records when rows are inserted, updated or deleted.
** This may appear to have some counter-intuitive effects if a single row
** is written to more than once during a session. For example, if a row
** is inserted while a session object is enabled, then later deleted while 
** the same session object is disabled, no INSERT record will appear in the
** changeset, even though the delete took place while the session was disabled.
** Or, if one field of a row is updated while a session is disabled, and 
** another field of the same row is updated while the session is enabled, the
** resulting changeset will contain an UPDATE change that updates both fields.
*/
SQLITE_API int sqlite3session_changeset(
  sqlite3_session *pSession,      /* Session object */
  int *pnChangeset,               /* OUT: Size of buffer at *ppChangeset */
  void **ppChangeset              /* OUT: Buffer containing changeset */
);

/*
** CAPI3REF: Load The Difference Between Tables Into A Session
** METHOD: sqlite3_session
**
** If it is not already attached to the session object passed as the first
** argument, this function attaches table zTbl in the same manner as the
** [sqlite3session_attach()] function. If zTbl does not exist, or if it
** does not have a primary key, this function is a no-op (but does not return
** an error).
**
** Argument zFromDb must be the name of a database ("main", "temp" etc.)
** attached to the same database handle as the session object that contains 
** a table compatible with the table attached to the session by this function.
** A table is considered compatible if it:
**
** <ul>
**   <li> Has the same name,
**   <li> Has the same set of columns declared in the same order, and
**   <li> Has the same PRIMARY KEY definition.
** </ul>
**
** If the tables are not compatible, SQLITE_SCHEMA is returned. If the tables
** are compatible but do not have any PRIMARY KEY columns, it is not an error
** but no changes are added to the session object. As with other session
** APIs, tables without PRIMARY KEYs are simply ignored.
**
** This function adds a set of changes to the session object that could be
** </ul>
**
** It is the responsibility of the caller to eventually destroy the iterator
** by passing it to [sqlite3changeset_finalize()]. The buffer containing the
** changeset (pChangeset) must remain valid until after the iterator is
** destroyed.
**
** Assuming the changeset blob was created by one of the
** [sqlite3session_changeset()], [sqlite3changeset_concat()] or
** [sqlite3changeset_invert()] functions, all changes within the changeset 
** that apply to a single table are grouped together. This means that when 
** an application iterates through a changeset using an iterator created by 
** this function, all changes that relate to a single table are visited 
** consecutively. There is no chance that the iterator will visit a change 
** the applies to table X, then one for table Y, and then later on visit 
** another change for table X.
*/
SQLITE_API int sqlite3changeset_start(
  sqlite3_changeset_iter **pp,    /* OUT: New changeset iterator handle */
  int nChangeset,                 /* Size of changeset blob in bytes */
  void *pChangeset                /* Pointer to blob containing changeset */
);


/*
** CAPI3REF: Advance A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** This function may only be used with iterators created by function
** [sqlite3changeset_start()]. If it is called on an iterator passed to
** a conflict-handler callback by [sqlite3changeset_apply()], SQLITE_MISUSE
** is returned and the call has no effect.
**
** Immediately after an iterator is created by sqlite3changeset_start(), it
** does not point to any change in the changeset. Assuming the changeset
** is not empty, the first call to this function advances the iterator to
** point to the first change in the changeset. Each subsequent call advances
** the iterator to point to the next change in the changeset (if any). If
** no error occurs and the iterator points to a valid change after a call
** to sqlite3changeset_next() has advanced it, SQLITE_ROW is returned. 
** Otherwise, if all changes in the changeset have already been visited,
** SQLITE_DONE is returned.
**
** If an error occurs, an SQLite error code is returned. Possible error 
** codes include SQLITE_CORRUPT (if the changeset buffer is corrupt) or 
** SQLITE_NOMEM.
*/
SQLITE_API int sqlite3changeset_next(sqlite3_changeset_iter *pIter);

/*
** CAPI3REF: Obtain The Current Operation From A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** The pIter argument passed to this function may either be an iterator
** passed to a conflict-handler by [sqlite3changeset_apply()], or an iterator
** created by [sqlite3changeset_start()]. In the latter case, the most recent
** call to [sqlite3changeset_next()] must have returned [SQLITE_ROW]. If this
** is not the case, this function returns [SQLITE_MISUSE].
**
** If argument pzTab is not NULL, then *pzTab is set to point to a
** nul-terminated utf-8 encoded string containing the name of the table
** affected by the current change. The buffer remains valid until either
** sqlite3changeset_next() is called on the iterator or until the 
** conflict-handler function returns. If pnCol is not NULL, then *pnCol is 
** set to the number of columns in the table affected by the change. If
** pbIncorrect is not NULL, then *pbIndirect is set to true (1) if the change
** is an indirect change, or false (0) otherwise. See the documentation for
** [sqlite3session_indirect()] for a description of direct and indirect
** changes. Finally, if pOp is not NULL, then *pOp is set to one of 
** [SQLITE_INSERT], [SQLITE_DELETE] or [SQLITE_UPDATE], depending on the 
** type of change that the iterator currently points to.
**
** If no error occurs, SQLITE_OK is returned. If an error does occur, an
** SQLite error code is returned. The values of the output variables may not
** be trusted in this case.
*/
SQLITE_API int sqlite3changeset_op(
  sqlite3_changeset_iter *pIter,  /* Iterator object */
  const char **pzTab,             /* OUT: Pointer to table name */
  int *pnCol,                     /* OUT: Number of columns in table */
  int *pOp,                       /* OUT: SQLITE_INSERT, DELETE or UPDATE */
  int *pbIndirect                 /* OUT: True for an 'indirect' change */
);

/*
** CAPI3REF: Obtain The Primary Key Definition Of A Table
** METHOD: sqlite3_changeset_iter
**
** For each modified table, a changeset includes the following:
**
** <ul>
**   <li> The number of columns in the table, and
**   <li> Which of those columns make up the tables PRIMARY KEY.
** </ul>
**
** This function is used to find which columns comprise the PRIMARY KEY of
** the table modified by the change that iterator pIter currently points to.
** If successful, *pabPK is set to point to an array of nCol entries, where
** nCol is the number of columns in the table. Elements of *pabPK are set to
** 0x01 if the corresponding column is part of the tables primary key, or
** 0x00 if it is not.
**
** If argument pnCol is not NULL, then *pnCol is set to the number of columns
** in the table.
**
** If this function is called when the iterator does not point to a valid
** entry, SQLITE_MISUSE is returned and the output variables zeroed. Otherwise,
** SQLITE_OK is returned and the output variables populated as described
** above.
*/
SQLITE_API int sqlite3changeset_pk(
  sqlite3_changeset_iter *pIter,  /* Iterator object */
  unsigned char **pabPK,          /* OUT: Array of boolean - true for PK cols */
  int *pnCol                      /* OUT: Number of entries in output array */
);

/*
** CAPI3REF: Obtain old.* Values From A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** The pIter argument passed to this function may either be an iterator
** passed to a conflict-handler by [sqlite3changeset_apply()], or an iterator
** created by [sqlite3changeset_start()]. In the latter case, the most recent
** call to [sqlite3changeset_next()] must have returned SQLITE_ROW. 
** Furthermore, it may only be called if the type of change that the iterator
** currently points to is either [SQLITE_DELETE] or [SQLITE_UPDATE]. Otherwise,
** this function returns [SQLITE_MISUSE] and sets *ppValue to NULL.
**
** Argument iVal must be greater than or equal to 0, and less than the number
** of columns in the table affected by the current change. Otherwise,
** [SQLITE_RANGE] is returned and *ppValue is set to NULL.
**
** If successful, this function sets *ppValue to point to a protected
** sqlite3_value object containing the iVal'th value from the vector of 
** original row values stored as part of the UPDATE or DELETE change and
** returns SQLITE_OK. The name of the function comes from the fact that this 
** is similar to the "old.*" columns available to update or delete triggers.
**
** If some other error occurs (e.g. an OOM condition), an SQLite error code
** is returned and *ppValue is set to NULL.
*/
SQLITE_API int sqlite3changeset_old(
  sqlite3_changeset_iter *pIter,  /* Changeset iterator */
  int iVal,                       /* Column number */
  sqlite3_value **ppValue         /* OUT: Old value (or NULL pointer) */
);

/*
** CAPI3REF: Obtain new.* Values From A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** The pIter argument passed to this function may either be an iterator
** passed to a conflict-handler by [sqlite3changeset_apply()], or an iterator
** created by [sqlite3changeset_start()]. In the latter case, the most recent
** call to [sqlite3changeset_next()] must have returned SQLITE_ROW. 
** Furthermore, it may only be called if the type of change that the iterator
** currently points to is either [SQLITE_UPDATE] or [SQLITE_INSERT]. Otherwise,
** this function returns [SQLITE_MISUSE] and sets *ppValue to NULL.
**
** Argument iVal must be greater than or equal to 0, and less than the number
** of columns in the table affected by the current change. Otherwise,
** [SQLITE_RANGE] is returned and *ppValue is set to NULL.
**
** If successful, this function sets *ppValue to point to a protected
** sqlite3_value object containing the iVal'th value from the vector of 
** new row values stored as part of the UPDATE or INSERT change and
** returns SQLITE_OK. If the change is an UPDATE and does not include
** a new value for the requested column, *ppValue is set to NULL and 
** SQLITE_OK returned. The name of the function comes from the fact that 
** this is similar to the "new.*" columns available to update or delete 
** triggers.
**
** If some other error occurs (e.g. an OOM condition), an SQLite error code
** is returned and *ppValue is set to NULL.
*/
SQLITE_API int sqlite3changeset_new(
  sqlite3_changeset_iter *pIter,  /* Changeset iterator */
  int iVal,                       /* Column number */
  sqlite3_value **ppValue         /* OUT: New value (or NULL pointer) */
);

/*
** CAPI3REF: Obtain Conflicting Row Values From A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** This function should only be used with iterator objects passed to a
** conflict-handler callback by [sqlite3changeset_apply()] with either
** [SQLITE_CHANGESET_DATA] or [SQLITE_CHANGESET_CONFLICT]. If this function
** is called on any other iterator, [SQLITE_MISUSE] is returned and *ppValue
** is set to NULL.
**
** Argument iVal must be greater than or equal to 0, and less than the number
** of columns in the table affected by the current change. Otherwise,
** [SQLITE_RANGE] is returned and *ppValue is set to NULL.
**
** If successful, this function sets *ppValue to point to a protected
** sqlite3_value object containing the iVal'th value from the 
** "conflicting row" associated with the current conflict-handler callback
** and returns SQLITE_OK.
**
** If some other error occurs (e.g. an OOM condition), an SQLite error code
** is returned and *ppValue is set to NULL.
*/
SQLITE_API int sqlite3changeset_conflict(
  sqlite3_changeset_iter *pIter,  /* Changeset iterator */
  int iVal,                       /* Column number */
  sqlite3_value **ppValue         /* OUT: Value from conflicting row */
);

/*
** CAPI3REF: Determine The Number Of Foreign Key Constraint Violations
** METHOD: sqlite3_changeset_iter
**
** This function may only be called with an iterator passed to an
** SQLITE_CHANGESET_FOREIGN_KEY conflict handler callback. In this case
** it sets the output variable to the total number of known foreign key
** violations in the destination database and returns SQLITE_OK.
**
** In all other cases this function returns SQLITE_MISUSE.
*/
SQLITE_API int sqlite3changeset_fk_conflicts(
  sqlite3_changeset_iter *pIter,  /* Changeset iterator */
  int *pnOut                      /* OUT: Number of FK violations */
);


/*
** CAPI3REF: Finalize A Changeset Iterator
** METHOD: sqlite3_changeset_iter
**
** This function is used to finalize an iterator allocated with
** [sqlite3changeset_start()].
**
** This function should only be called on iterators created using the
** [sqlite3changeset_start()] function. If an application calls this
** function with an iterator passed to a conflict-handler by
** [sqlite3changeset_apply()], [SQLITE_MISUSE] is immediately returned and the
** call has no effect.
**
** If an error was encountered within a call to an sqlite3changeset_xxx()
** function (for example an [SQLITE_CORRUPT] in [sqlite3changeset_next()] or an 
** [SQLITE_NOMEM] in [sqlite3changeset_new()]) then an error code corresponding
** to that error is returned by this function. Otherwise, SQLITE_OK is
** returned. This is to allow the following pattern (pseudo-code):
**
** <pre>
**   sqlite3changeset_start();
**   while( SQLITE_ROW==sqlite3changeset_next() ){
**     // Do something with change.
**   }
**   rc = sqlite3changeset_finalize();
**   if( rc!=SQLITE_OK ){
**     // An error has occurred 
**   }
** </pre>
*/
SQLITE_API int sqlite3changeset_finalize(sqlite3_changeset_iter *pIter);

/*
** CAPI3REF: Invert A Changeset
**
** This function is used to "invert" a changeset object. Applying an inverted
** changeset to a database reverses the effects of applying the uninverted
** changeset. Specifically:
**
** <ul>
**   <li> Each DELETE change is changed to an INSERT, and
**   <li> Each INSERT change is changed to a DELETE, and
**   <li> For each UPDATE change, the old.* and new.* values are exchanged.
** </ul>
**
** This function does not change the order in which changes appear within
** the changeset. It merely reverses the sense of each individual change.
**
** If successful, a pointer to a buffer containing the inverted changeset
** is stored in *ppOut, the size of the same buffer is stored in *pnOut, and
** SQLITE_OK is returned. If an error occurs, both *pnOut and *ppOut are
** zeroed and an SQLite error code returned.
**
** It is the responsibility of the caller to eventually call sqlite3_free()
** on the *ppOut pointer to free the buffer allocation following a successful 
** call to this function.
**
** WARNING/TODO: This function currently assumes that the input is a valid
** changeset. If it is not, the results are undefined.
*/
SQLITE_API int sqlite3changeset_invert(
  int nIn, const void *pIn,       /* Input changeset */
  int *pnOut, void **ppOut        /* OUT: Inverse of input */
);

/*
** CAPI3REF: Concatenate Two Changeset Objects
**
** This function is used to concatenate two changesets, A and B, into a 
** single changeset. The result is a changeset equivalent to applying
** changeset A followed by changeset B. 
**
** This function combines the two input changesets using an 
** sqlite3_changegroup object. Calling it produces similar results as the
** following code fragment:
**
** <pre>
**   sqlite3_changegroup *pGrp;
**   rc = sqlite3_changegroup_new(&pGrp);
**   if( rc==SQLITE_OK ) rc = sqlite3changegroup_add(pGrp, nA, pA);
**   if( rc==SQLITE_OK ) rc = sqlite3changegroup_add(pGrp, nB, pB);
**   if( rc==SQLITE_OK ){
**     rc = sqlite3changegroup_output(pGrp, pnOut, ppOut);
**   }else{
**     *ppOut = 0;
**     *pnOut = 0;
**   }
** </pre>
**
** Refer to the sqlite3_changegroup documentation below for details.
*/
SQLITE_API int sqlite3changeset_concat(
**       as the deleted row, the existing DELETE is simply discarded.
**   <tr><td>DELETE <td>UPDATE <td>
**       The new change is ignored. This case does not occur if the new
**       changeset was recorded immediately after the changesets already
**       added to the changegroup.
**   <tr><td>DELETE <td>DELETE <td>
**       The new change is ignored. This case does not occur if the new
**       changeset was recorded immediately after the changesets already
**       added to the changegroup.
** </table>
**
** If the new changeset contains changes to a table that is already present
** in the changegroup, then the number of columns and the position of the
** primary key columns for the table must be consistent. If this is not the
** case, this function fails with SQLITE_SCHEMA. If the input changeset
** appears to be corrupt and the corruption is detected, SQLITE_CORRUPT is
** returned. Or, if an out-of-memory condition occurs during processing, this
** function returns SQLITE_NOMEM. In all cases, if an error occurs the
** final contents of the changegroup is undefined.
**
** If no error occurs, SQLITE_OK is returned.
*/
SQLITE_API int sqlite3changegroup_add(sqlite3_changegroup*, int nData, void *pData);

/*
** CAPI3REF: Obtain A Composite Changeset From A Changegroup
** METHOD: sqlite3_changegroup
**
** Obtain a buffer containing a changeset (or patchset) representing the
** current contents of the changegroup. If the inputs to the changegroup
** were themselves changesets, the output is a changeset. Or, if the
** inputs were patchsets, the output is also a patchset.
**
** As with the output of the sqlite3session_changeset() and
** sqlite3session_patchset() functions, all changes related to a single
** table are grouped together in the output of this function. Tables appear
** in the same order as for the very first changeset added to the changegroup.
** If the second or subsequent changesets added to the changegroup contain
** changes for tables that do not appear in the first changeset, they are
** appended onto the end of the output changeset, again in the order in
** which they are first encountered.
**
** If an error occurs, an SQLite error code is returned and the output
** variables (*pnData) and (*ppData) are set to 0. Otherwise, SQLITE_OK
** is returned and the output variables are set to the size of and a 
** pointer to the output buffer, respectively. In this case it is the
** responsibility of the caller to eventually free the buffer using a
** call to sqlite3_free().
*/
SQLITE_API int sqlite3changegroup_output(
  sqlite3_changegroup*,
** <dt>SQLITE_CHANGESET_REPLACE<dd>
**   This value may only be returned if the second argument to the conflict
**   handler was SQLITE_CHANGESET_DATA or SQLITE_CHANGESET_CONFLICT. If this
**   is not the case, any changes applied so far are rolled back and the 
**   call to sqlite3changeset_apply() returns SQLITE_MISUSE.
**
**   If CHANGESET_REPLACE is returned by an SQLITE_CHANGESET_DATA conflict
**   handler, then the conflicting row is either updated or deleted, depending
**   on the type of change.
**
**   If CHANGESET_REPLACE is returned by an SQLITE_CHANGESET_CONFLICT conflict
**   handler, then the conflicting row is removed from the database and a
**   second attempt to apply the change is made. If this second attempt fails,
**   the original row is restored to the database before continuing.
**
** <dt>SQLITE_CHANGESET_ABORT<dd>
**   If this value is returned, any changes applied so far are rolled back 
**   and the call to sqlite3changeset_apply() returns SQLITE_ABORT.
** </dl>
*/
#define SQLITE_CHANGESET_OMIT       0
#define SQLITE_CHANGESET_REPLACE    1
#define SQLITE_CHANGESET_ABORT      2

/* 
** CAPI3REF: Rebasing changesets
** EXPERIMENTAL
**
** Suppose there is a site hosting a database in state S0. And that
** modifications are made that move that database to state S1 and a
** changeset recorded (the "local" changeset). Then, a changeset based
** on S0 is received from another site (the "remote" changeset) and 
** applied to the database. The database is then in state 
** (S1+"remote"), where the exact state depends on any conflict
** resolution decisions (OMIT or REPLACE) made while applying "remote".
** Rebasing a changeset is to update it to take those conflict 
** resolution decisions into account, so that the same conflicts
** do not have to be resolved elsewhere in the network. 
**
** For example, if both the local and remote changesets contain an
** INSERT of the same key on "CREATE TABLE t1(a PRIMARY KEY, b)":
**
**   local:  INSERT INTO t1 VALUES(1, 'v1');
**   remote: INSERT INTO t1 VALUES(1, 'v2');
**
** and the conflict resolution is REPLACE, then the INSERT change is
** removed from the local changeset (it was overridden). Or, if the
** conflict resolution was "OMIT", then the local changeset is modified
** to instead contain:
**
**           UPDATE t1 SET b = 'v2' WHERE a=1;
** CAPI3REF: Create a changeset rebaser object.
** EXPERIMENTAL
**
** Allocate a new changeset rebaser object. If successful, set (*ppNew) to
** point to the new object and return SQLITE_OK. Otherwise, if an error
** occurs, return an SQLite error code (e.g. SQLITE_NOMEM) and set (*ppNew) 
** to NULL. 
*/
SQLITE_API int sqlite3rebaser_create(sqlite3_rebaser **ppNew);

/*
** CAPI3REF: Configure a changeset rebaser object.
** EXPERIMENTAL
**
** Configure the changeset rebaser object to rebase changesets according
** to the conflict resolutions described by buffer pRebase (size nRebase
** bytes), which must have been obtained from a previous call to
** sqlite3changeset_apply_v2().
*/
SQLITE_API int sqlite3rebaser_configure(
  sqlite3_rebaser*, 
  int nRebase, const void *pRebase
); 

/*
** CAPI3REF: Rebase a changeset
** EXPERIMENTAL
**
** Argument pIn must point to a buffer containing a changeset nIn bytes
** in size. This function allocates and populates a buffer with a copy
** of the changeset rebased rebased according to the configuration of the
** rebaser object passed as the first argument. If successful, (*ppOut)
** is set to point to the new buffer containing the rebased changset and 
** (*pnOut) to its size in bytes and SQLITE_OK returned. It is the
** responsibility of the caller to eventually free the new buffer using
** sqlite3_free(). Otherwise, if an error occurs, (*ppOut) and (*pnOut)
** are set to zero and an SQLite error code returned.
*/
SQLITE_API int sqlite3rebaser_rebase(
  sqlite3_rebaser*,
  int nIn, const void *pIn, 
  int *pnOut, void **ppOut 
);

/*
** CAPI3REF: Delete a changeset rebaser object.
** EXPERIMENTAL
**
** Delete the changeset rebaser object and all associated resources. There
** should be one call to this function for each successful invocation
** of sqlite3rebaser_create().
** These macros can be used to test, set, or clear bits in the
** Expr.flags field.
*/
#define ExprHasProperty(E,P)     (((E)->flags&(P))!=0)
#define ExprHasAllProperty(E,P)  (((E)->flags&(P))==(P))
#define ExprSetProperty(E,P)     (E)->flags|=(P)
#define ExprClearProperty(E,P)   (E)->flags&=~(P)

/* The ExprSetVVAProperty() macro is used for Verification, Validation,
** and Accreditation only.  It works like ExprSetProperty() during VVA
** processes but is a no-op for delivery.
*/
#ifdef SQLITE_DEBUG
# define ExprSetVVAProperty(E,P)  (E)->flags|=(P)
#else
# define ExprSetVVAProperty(E,P)
#endif

/*
** Macros to determine the number of bytes required by a normal Expr
** struct, an Expr struct with the EP_Reduced flag set in Expr.flags
** and an Expr struct with the EP_TokenOnly flag set.
*/
#define EXPR_FULLSIZE           sizeof(Expr)           /* Full size */
#define EXPR_REDUCEDSIZE        offsetof(Expr,iTable)  /* Common features */
#define EXPR_TOKENONLYSIZE      offsetof(Expr,pLeft)   /* Fewer features */

/*
** Flags passed to the sqlite3ExprDup() function. See the header comment
** above sqlite3ExprDup() for details.
*/
#define EXPRDUP_REDUCE         0x0001  /* Used reduced-size Expr nodes */

/*
** A list of expressions.  Each expression may optionally have a
** name.  An expr/name combination can be used in several ways, such
** as the list of "expr AS ID" fields following a "SELECT" or in the
** list of "ID = expr" items in an UPDATE.  A list of expressions can
** also be used as the argument to a function, in which case the a.zName
** field is not used.
**
** By default the Expr.zSpan field holds a human-readable description of
** the expression that is used in the generation of error messages and
** column labels.  In this case, Expr.zSpan is typically the text of a
** column expression as it exists in a SELECT statement.  However, if
** the bSpanIsTab flag is set, then zSpan is overloaded to mean the name
** of the result column in the form: DATABASE.TABLE.COLUMN.  This later
** form is used for name resolution with nested FROM clauses.
*/
struct ExprList {
  int nExpr;             /* Number of expressions on the list */
      if( rc ) return rc;
    }
  }else if( loc==0 && (flags & BTREE_SAVEPOSITION)==0 ){
    if( pX->nMem ){
      UnpackedRecord r;
      r.pKeyInfo = pCur->pKeyInfo;
      r.aMem = pX->aMem;
      r.nField = pX->nMem;
      r.default_rc = 0;
      r.errCode = 0;
      r.r1 = 0;
      r.r2 = 0;
      r.eqSeen = 0;
      rc = sqlite3BtreeMovetoUnpacked(pCur, &r, 0, flags!=0, &loc);
    }else{
      rc = btreeMoveto(pCur, pX->pKey, pX->nKey, flags!=0, &loc);
    }
    if( rc ) return rc;
  }
  assert( pCur->eState==CURSOR_VALID || (pCur->eState==CURSOR_INVALID && loc) );

  pPage = pCur->pPage;
  assert( pPage->intKey || pX->nKey>=0 );
  assert( pPage->leaf || !pPage->intKey );

  TRACE(("INSERT: table=%d nkey=%lld ndata=%d page=%d %s\n",
          pCur->pgnoRoot, pX->nKey, pX->nData, pPage->pgno,
          loc==0 ? "overwrite" : "new entry"));
  assert( pPage->isInit );
  newCell = pBt->pTmpSpace;
  assert( newCell!=0 );
  rc = fillInCell(pPage, newCell, pX, &szNew);
  if( rc ) goto end_insert;
  assert( szNew==pPage->xCellSize(pPage, newCell) );
  assert( szNew <= MX_CELL_SIZE(pBt) );
  idx = pCur->ix;
  if( loc==0 ){
    CellInfo info;
    assert( idx<pPage->nCell );
    rc = sqlite3PagerWrite(pPage->pDbPage);
    if( rc ){
      goto end_insert;
    }
    oldCell = findCell(pPage, idx);
    if( !pPage->leaf ){
      memcpy(newCell, oldCell, 4);
    }
    rc = clearCell(pPage, oldCell, &info);
    if( info.nSize==szNew && info.nLocal==info.nPayload 
     && (!ISAUTOVACUUM || szNew<pPage->minLocal)
    ){
        checkPtrmap(pCheck, pgnoOvfl, PTRMAP_OVERFLOW1, iPage);
      }
#endif
      checkList(pCheck, 0, pgnoOvfl, nPage);
    }

    if( !pPage->leaf ){
      /* Check sanity of left child page for internal pages */
      pgno = get4byte(pCell);
#ifndef SQLITE_OMIT_AUTOVACUUM
      if( pBt->autoVacuum ){
        checkPtrmap(pCheck, pgno, PTRMAP_BTREE, iPage);
      }
#endif
      d2 = checkTreePage(pCheck, pgno, &maxKey, maxKey);
      keyCanBeEqual = 0;
      if( d2!=depth ){
        checkAppendMsg(pCheck, "Child page depth differs");
        depth = d2;
      }
    }else{
      /* Populate the coverage-checking heap for leaf pages */
      btreeHeapInsert(heap, (pc<<16)|(pc+info.nSize-1));
    }
  }
  *piMinKey = maxKey;

  /* Check for complete coverage of the page
  */
  pCheck->zPfx = 0;
  if( doCoverageCheck && pCheck->mxErr>0 ){
    /* For leaf pages, the min-heap has already been initialized and the
    ** cells have already been inserted.  But for internal pages, that has
    ** not yet been done, so do it now */
    if( !pPage->leaf ){
      heap = pCheck->heap;
      heap[0] = 0;
      for(i=nCell-1; i>=0; i--){
        u32 size;
        pc = get2byteAligned(&data[cellStart+i*2]);
        size = pPage->xCellSize(pPage, &data[pc]);
        btreeHeapInsert(heap, (pc<<16)|(pc+size-1));
      }
    }
    /* Add the freeblocks to the min-heap
    **
    ** EVIDENCE-OF: R-20690-50594 The second field of the b-tree page header
    ** is the offset of the first freeblock, or zero if there are no
    ** freeblocks on the page. 
    */
    i = get2byte(&data[hdr+1]);
      /* EVIDENCE-OF: R-18143-12121 Value is the integer 1. */
      pMem->u.i = serial_type-8;
      pMem->flags = MEM_Int;
      return 0;
    }
    default: {
      /* EVIDENCE-OF: R-14606-31564 Value is a BLOB that is (N-12)/2 bytes in
      ** length.
      ** EVIDENCE-OF: R-28401-00140 Value is a string in the text encoding and
      ** (N-13)/2 bytes in length. */
      static const u16 aFlag[] = { MEM_Blob|MEM_Ephem, MEM_Str|MEM_Ephem };
      pMem->z = (char *)buf;
      pMem->n = (serial_type-12)/2;
      pMem->flags = aFlag[serial_type&1];
      return pMem->n;
    }
  }
  return 0;
}
/*
** This routine is used to allocate sufficient space for an UnpackedRecord
** structure large enough to be used with sqlite3VdbeRecordUnpack() if
** the first argument is a pointer to KeyInfo structure pKeyInfo.
**
** The space is either allocated using sqlite3DbMallocRaw() or from within
** the unaligned buffer passed via the second and third arguments (presumably
** stack space). If the former, then *ppFree is set to a pointer that should
** be eventually freed by the caller using sqlite3DbFree(). Or, if the 
** allocation comes from the pSpace/szSpace buffer, *ppFree is set to NULL
** before returning.
**
** If an OOM error occurs, NULL is returned.
*/
SQLITE_PRIVATE UnpackedRecord *sqlite3VdbeAllocUnpackedRecord(
  KeyInfo *pKeyInfo               /* Description of the record */
){
  UnpackedRecord *p;              /* Unpacked record to return */
  int nByte;                      /* Number of bytes required for *p */
  nByte = ROUND8(sizeof(UnpackedRecord)) + sizeof(Mem)*(pKeyInfo->nKeyField+1);
  p = (UnpackedRecord *)sqlite3DbMallocRaw(pKeyInfo->db, nByte);
  if( !p ) return 0;
  p->aMem = (Mem*)&((char*)p)[ROUND8(sizeof(UnpackedRecord))];
  assert( pKeyInfo->aSortOrder!=0 );
  p->pKeyInfo = pKeyInfo;
  p->nField = pKeyInfo->nKeyField + 1;
  return p;
}

/*
** Given the nKey-byte encoding of a record in pKey[], populate the 
** UnpackedRecord structure indicated by the fourth argument with the
    nMem = pProgram->nMem + pProgram->nCsr;
    assert( nMem>0 );
    if( pProgram->nCsr==0 ) nMem++;
    nByte = ROUND8(sizeof(VdbeFrame))
              + nMem * sizeof(Mem)
              + pProgram->nCsr * sizeof(VdbeCursor*)
              + (pProgram->nOp + 7)/8;
    pFrame = sqlite3DbMallocZero(db, nByte);
    if( !pFrame ){
      goto no_mem;
    }
    sqlite3VdbeMemRelease(pRt);
    pRt->flags = MEM_Frame;
    pRt->u.pFrame = pFrame;

    pFrame->v = p;
    pFrame->nChildMem = nMem;
    pFrame->nChildCsr = pProgram->nCsr;
    pFrame->pc = (int)(pOp - aOp);
    pFrame->aMem = p->aMem;
    pFrame->nMem = p->nMem;
    pFrame->apCsr = p->apCsr;
    pFrame->nCursor = p->nCursor;
    pFrame->aOp = p->aOp;
    pFrame->nOp = p->nOp;
    pFrame->token = pProgram->token;
#ifdef SQLITE_ENABLE_STMT_SCANSTATUS
    pFrame->anExec = p->anExec;
#endif

    pEnd = &VdbeFrameMem(pFrame)[pFrame->nChildMem];
    for(pMem=VdbeFrameMem(pFrame); pMem!=pEnd; pMem++){
      pMem->flags = MEM_Undefined;
      pMem->db = db;
    }
  }else{
    pFrame = pRt->u.pFrame;
    assert( pProgram->nMem+pProgram->nCsr==pFrame->nChildMem 
        || (pProgram->nCsr==0 && pProgram->nMem+1==pFrame->nChildMem) );
    assert( pProgram->nCsr==pFrame->nChildCsr );
    assert( (int)(pOp - aOp)==pFrame->pc );
  }

  p->nFrame++;
  pFrame->pParent = p->pFrame;
  pFrame->lastRowid = db->lastRowid;
  pFrame->nChange = p->nChange;
  pFrame->nDbChange = p->db->nChange;
  assert( pFrame->pAuxData==0 );
  pFrame->pAuxData = p->pAuxData;
  p->pAuxData = 0;
    if( pSorter->list.aMemory==0 ) vdbeSorterRecordFree(db, pFree);
    rc = pSorter->list.pList ? SQLITE_OK : SQLITE_DONE;
  }
  return rc;
}

/*
** Return a pointer to a buffer owned by the sorter that contains the 
** current key.
*/
static void *vdbeSorterRowkey(
  const VdbeSorter *pSorter,      /* Sorter object */
  int *pnKey                      /* OUT: Size of current key in bytes */
){
  void *pKey;
  if( pSorter->bUsePMA ){
    PmaReader *pReader;
#if SQLITE_MAX_WORKER_THREADS>0
    if( pSorter->bUseThreads ){
      pReader = pSorter->pReader;
    }else
#endif
    /*if( !pSorter->bUseThreads )*/{
      pReader = &pSorter->pMerger->aReadr[pSorter->pMerger->aTree[1]];
    }
    *pnKey = pReader->nKey;
    pKey = pReader->aKey;
  }else{
    *pnKey = pSorter->list.pList->nVal;
    pKey = SRVAL(pSorter->list.pList);
  }
  return pKey;
}

/*
** Copy the current sorter key into the memory cell pOut.
*/
SQLITE_PRIVATE int sqlite3VdbeSorterRowkey(const VdbeCursor *pCsr, Mem *pOut){
  VdbeSorter *pSorter;
  void *pKey; int nKey;           /* Sorter key to copy into pOut */

  assert( pCsr->eCurType==CURTYPE_SORTER );
  pSorter = pCsr->uc.pSorter;
  pKey = vdbeSorterRowkey(pSorter, &nKey);
  if( sqlite3VdbeMemClearAndResize(pOut, nKey) ){
    return SQLITE_NOMEM_BKPT;
  }
  pOut->n = nKey;
  MemSetTypeFlag(pOut, MEM_Blob);
  memcpy(pOut->z, pKey, nKey);

  return SQLITE_OK;
}

    exprX.flags |= EP_FromJoin;
    sqlite3ExprCodeTarget(pParse, &exprAnd, dest);
  }
  sqlite3ReleaseTempReg(pParse, regFree1);

  /* Ensure adequate test coverage */
  testcase( xJump==sqlite3ExprIfTrue  && jumpIfNull==0 && regFree1==0 );
  testcase( xJump==sqlite3ExprIfTrue  && jumpIfNull==0 && regFree1!=0 );
  testcase( xJump==sqlite3ExprIfTrue  && jumpIfNull!=0 && regFree1==0 );
  testcase( xJump==sqlite3ExprIfTrue  && jumpIfNull!=0 && regFree1!=0 );
  testcase( xJump==sqlite3ExprIfFalse && jumpIfNull==0 && regFree1==0 );
  testcase( xJump==sqlite3ExprIfFalse && jumpIfNull==0 && regFree1!=0 );
  testcase( xJump==sqlite3ExprIfFalse && jumpIfNull!=0 && regFree1==0 );
  testcase( xJump==sqlite3ExprIfFalse && jumpIfNull!=0 && regFree1!=0 );
  testcase( xJump==0 );
}

/*
** Generate code for a boolean expression such that a jump is made
** to the label "dest" if the expression is true but execution
** continues straight thru if the expression is false.
**
** If the expression evaluates to NULL (neither true nor false), then
** take the jump if the jumpIfNull flag is SQLITE_JUMPIFNULL.
**
** This code depends on the fact that certain token values (ex: TK_EQ)
** are the same as opcode values (ex: OP_Eq) that implement the corresponding
** operation.  Special comments in vdbe.c and the mkopcodeh.awk script in
** the make process cause these values to align.  Assert()s in the code
** below verify that the numbers are aligned correctly.
*/
SQLITE_PRIVATE void sqlite3ExprIfTrue(Parse *pParse, Expr *pExpr, int dest, int jumpIfNull){
  Vdbe *v = pParse->pVdbe;
  int op = 0;
  int regFree1 = 0;
  int regFree2 = 0;
  int r1, r2;

  assert( jumpIfNull==SQLITE_JUMPIFNULL || jumpIfNull==0 );
  if( NEVER(v==0) )     return;  /* Existence of VDBE checked by caller */
  if( NEVER(pExpr==0) ) return;  /* No way this can happen */
  op = pExpr->op;
  switch( op ){
    case TK_AND: {
      int d2 = sqlite3VdbeMakeLabel(v);
      testcase( jumpIfNull==0 );
      sqlite3ExprIfFalse(pParse, pExpr->pLeft, d2,jumpIfNull^SQLITE_JUMPIFNULL);
      sqlite3ExprCachePush(pParse);
      sqlite3ExprIfTrue(pParse, pExpr->pRight, dest, jumpIfNull);
      sqlite3VdbeResolveLabel(v, d2);
      sqlite3ExprCachePop(pParse);
  sqlite3 *db = sqlite3_context_db_handle(context);

  UNUSED_PARAMETER(NotUsed);

  /* The principle used to locate the table name in the CREATE TABLE 
  ** statement is that the table name is the first non-space token that
  ** is immediately followed by a TK_LP or TK_USING token.
  */
  if( zSql ){
    do {
      if( !*zCsr ){
        /* Ran out of input before finding an opening bracket. Return NULL. */
        return;
      }

      /* Store the token that zCsr points to in tname. */
      tname.z = (char*)zCsr;
      tname.n = len;

      /* Advance zCsr to the next token. Store that token type in 'token',
      ** and its length in 'len' (to be used next iteration of this loop).
      */
      do {
        zCsr += len;
        len = sqlite3GetToken(zCsr, &token);
      } while( token==TK_SPACE );
      assert( len>0 );
    } while( token!=TK_LP && token!=TK_USING );

    zRet = sqlite3MPrintf(db, "%.*s\"%w\"%s", (int)(((u8*)tname.z) - zSql),
       zSql, zTableName, tname.z+tname.n);
    sqlite3_result_text(context, zRet, -1, SQLITE_DYNAMIC);
  }
}

/*
** This C function implements an SQL user function that is used by SQL code
** generated by the ALTER TABLE ... RENAME command to modify the definition
** of any foreign key constraints that use the table being renamed as the 
** parent table. It is passed three arguments:
**
**   1) The complete text of the CREATE TABLE statement being modified,
**   2) The old name of the table being renamed, and
**   3) The new name of the table being renamed.
**
** It returns the new CREATE TABLE statement. For example:
**
**   sqlite_rename_parent('CREATE TABLE t1(a REFERENCES t2)', 't2', 't3')
**       -> 'CREATE TABLE t1(a REFERENCES t3)'
*/
#ifndef SQLITE_OMIT_FOREIGN_KEY
static void renameParentFunc(
  sqlite3_context *context,
  int NotUsed,
  sqlite3_value **argv
){
  sqlite3 *db = sqlite3_context_db_handle(context);
  char *zOutput = 0;
  char *zResult;
  unsigned char const *zInput = sqlite3_value_text(argv[0]);
  unsigned char const *zOld = sqlite3_value_text(argv[1]);
  unsigned char const *zNew = sqlite3_value_text(argv[2]);

  unsigned const char *z;         /* Pointer to token */
  int n;                          /* Length of token z */
  int token;                      /* Type of token */

  UNUSED_PARAMETER(NotUsed);
  if( zInput==0 || zOld==0 ) return;
  for(z=zInput; *z; z=z+n){
    n = sqlite3GetToken(z, &token);
    if( token==TK_REFERENCES ){
      char *zParent;
      do {
        z += n;
        n = sqlite3GetToken(z, &token);
      }while( token==TK_SPACE );

      if( token==TK_ILLEGAL ) break;
      zParent = sqlite3DbStrNDup(db, (const char *)z, n);
      if( zParent==0 ) break;
      sqlite3Dequote(zParent);
      if( 0==sqlite3StrICmp((const char *)zOld, zParent) ){
        char *zOut = sqlite3MPrintf(db, "%s%.*s\"%w\"", 
            (zOutput?zOutput:""), (int)(z-zInput), zInput, (const char *)zNew
        );
        sqlite3DbFree(db, zOutput);
        zOutput = zOut;
        zInput = &z[n];
      }
      sqlite3DbFree(db, zParent);
    }
  }

  zResult = sqlite3MPrintf(db, "%s%s", (zOutput?zOutput:""), zInput), 
  sqlite3_result_text(context, zResult, -1, SQLITE_DYNAMIC);
  sqlite3DbFree(db, zOutput);
}
#endif

#ifndef SQLITE_OMIT_TRIGGER
/* This function is used by SQL generated to implement the
** ALTER TABLE command. The first argument is the text of a CREATE TRIGGER 
** statement. The second is a table name. The table name in the CREATE 
  UNUSED_PARAMETER(NotUsed);

  /* The principle used to locate the table name in the CREATE TRIGGER 
  ** statement is that the table name is the first token that is immediately
  ** preceded by either TK_ON or TK_DOT and immediately followed by one
  ** of TK_WHEN, TK_BEGIN or TK_FOR.
  */
  if( zSql ){
    do {

      if( !*zCsr ){
        /* Ran out of input before finding the table name. Return NULL. */
        return;
      }

      /* Store the token that zCsr points to in tname. */
      tname.z = (char*)zCsr;
      tname.n = len;

      /* Advance zCsr to the next token. Store that token type in 'token',
      ** and its length in 'len' (to be used next iteration of this loop).
      */
      do {
        zCsr += len;
        len = sqlite3GetToken(zCsr, &token);
      }while( token==TK_SPACE );
      assert( len>0 );

      /* Variable 'dist' stores the number of tokens read since the most
      ** recent TK_DOT or TK_ON. This means that when a WHEN, FOR or BEGIN 
      ** token is read and 'dist' equals 2, the condition stated above
      ** to be met.
      **
      ** Note that ON cannot be a database, table or column name, so
      ** there is no need to worry about syntax like 
      ** "CREATE TRIGGER ... ON ON.ON BEGIN ..." etc.
      */
      dist++;
      if( token==TK_DOT || token==TK_ON ){
        dist = 0;
      }
    } while( dist!=2 || (token!=TK_WHEN && token!=TK_FOR && token!=TK_BEGIN) );

    /* Variable tname now contains the token that is the old table-name
    ** in the CREATE TRIGGER statement.
    */
    zRet = sqlite3MPrintf(db, "%.*s\"%w\"%s", (int)(((u8*)tname.z) - zSql),
       zSql, zTableName, tname.z+tname.n);
    sqlite3_result_text(context, zRet, -1, SQLITE_DYNAMIC);
  }
}
#endif   /* !SQLITE_OMIT_TRIGGER */

/*
** Register built-in functions used to help implement ALTER TABLE
*/
SQLITE_PRIVATE void sqlite3AlterFunctions(void){
  static FuncDef aAlterTableFuncs[] = {
    FUNCTION(sqlite_rename_table,   2, 0, 0, renameTableFunc),
#ifndef SQLITE_OMIT_TRIGGER
    FUNCTION(sqlite_rename_trigger, 2, 0, 0, renameTriggerFunc),
#endif
#ifndef SQLITE_OMIT_FOREIGN_KEY
    FUNCTION(sqlite_rename_parent,  3, 0, 0, renameParentFunc),
  ** to use for scanning indexes (iIdxCur). No index cursor is opened at
  ** this time though.  */
  sqlite3TableLock(pParse, iDb, pTab->tnum, 0, pTab->zName);
  iTabCur = iTab++;
  iIdxCur = iTab++;
  pParse->nTab = MAX(pParse->nTab, iTab);
  sqlite3OpenTable(pParse, iTabCur, iDb, pTab, OP_OpenRead);
  sqlite3VdbeLoadString(v, regTabname, pTab->zName);

  for(pIdx=pTab->pIndex; pIdx; pIdx=pIdx->pNext){
    int nCol;                     /* Number of columns in pIdx. "N" */
    int addrRewind;               /* Address of "OP_Rewind iIdxCur" */
    int addrNextRow;              /* Address of "next_row:" */
    const char *zIdxName;         /* Name of the index */
    int nColTest;                 /* Number of columns to test for changes */

    if( pOnlyIdx && pOnlyIdx!=pIdx ) continue;
    if( pIdx->pPartIdxWhere==0 ) needTableCnt = 0;
    if( !HasRowid(pTab) && IsPrimaryKeyIndex(pIdx) ){
      nCol = pIdx->nKeyCol;
      zIdxName = pTab->zName;
      nColTest = nCol - 1;
    }else{
      nCol = pIdx->nColumn;
      zIdxName = pIdx->zName;
      nColTest = pIdx->uniqNotNull ? pIdx->nKeyCol-1 : nCol-1;
    }

    /* Populate the register containing the index name. */
    sqlite3VdbeLoadString(v, regIdxname, zIdxName);
    VdbeComment((v, "Analysis for %s.%s", pTab->zName, zIdxName));

    /*
    ** Pseudo-code for loop that calls stat_push():
    **
    **   Rewind csr
    **   if eof(csr) goto end_of_scan;
    **   regChng = 0
    **   goto chng_addr_0;
    **
    **  next_row:
    **   regChng = 0
    **   if( idx(0) != regPrev(0) ) goto chng_addr_0
    **   regChng = 1
    **   if( idx(1) != regPrev(1) ) goto chng_addr_1
    **   ...
    **   regChng = N
    **   goto chng_addr_N
    **
    **  chng_addr_0:
    **   regPrev(0) = idx(0)
      p->nCol = pSelTab->nCol;
      p->aCol = pSelTab->aCol;
      pSelTab->nCol = 0;
      pSelTab->aCol = 0;
      sqlite3DeleteTable(db, pSelTab);
      sqlite3SelectDestInit(&dest, SRT_Coroutine, regYield);
      sqlite3Select(pParse, pSelect, &dest);
      if( pParse->nErr ) return;
      sqlite3VdbeEndCoroutine(v, regYield);
      sqlite3VdbeJumpHere(v, addrTop - 1);
      addrInsLoop = sqlite3VdbeAddOp1(v, OP_Yield, dest.iSDParm);
      VdbeCoverage(v);
      sqlite3VdbeAddOp3(v, OP_MakeRecord, dest.iSdst, dest.nSdst, regRec);
      sqlite3TableAffinity(v, p, 0);
      sqlite3VdbeAddOp2(v, OP_NewRowid, 1, regRowid);
      sqlite3VdbeAddOp3(v, OP_Insert, 1, regRec, regRowid);
      sqlite3VdbeGoto(v, addrInsLoop);
      sqlite3VdbeJumpHere(v, addrInsLoop);
      sqlite3VdbeAddOp1(v, OP_Close, 1);
    }

    /* Compute the complete text of the CREATE statement */
    if( pSelect ){
      zStmt = createTableStmt(db, p);
    }else{
      Token *pEnd2 = tabOpts ? &pParse->sLastToken : pEnd;
      n = (int)(pEnd2->z - pParse->sNameToken.z);
      if( pEnd2->z[0]!=';' ) n += pEnd2->n;
      zStmt = sqlite3MPrintf(db, 
          "CREATE %s %.*s", zType2, n, pParse->sNameToken.z
      );
    }

    /* A slot for the record has already been allocated in the 
    ** SQLITE_MASTER table.  We just need to update that slot with all
    ** the information we've collected.
    */
    sqlite3NestedParse(pParse,
      "UPDATE %Q.%s "
         "SET type='%s', name=%Q, tbl_name=%Q, rootpage=#%d, sql=%Q "
       "WHERE rowid=#%d",
      db->aDb[iDb].zDbSName, MASTER_NAME,
      zType,
      p->zName,
      p->zName,
      pParse->regRoot,
      zStmt,
      pParse->regRowid
    );
    sqlite3DbFree(db, zStmt);
    sqlite3ChangeCookie(pParse, iDb);
  Vdbe *v;               /* The virtual database engine */
  Table *pTab;           /* The table from which records will be deleted */
  int i;                 /* Loop counter */
  WhereInfo *pWInfo;     /* Information about the WHERE clause */
  Index *pIdx;           /* For looping over indices of the table */
  int iTabCur;           /* Cursor number for the table */
  int iDataCur = 0;      /* VDBE cursor for the canonical data source */
  int iIdxCur = 0;       /* Cursor number of the first index */
  int nIdx;              /* Number of indices */
  sqlite3 *db;           /* Main database structure */
  AuthContext sContext;  /* Authorization context */
  NameContext sNC;       /* Name context to resolve expressions in */
  int iDb;               /* Database number */
  int memCnt = -1;       /* Memory cell used for change counting */
  int rcauth;            /* Value returned by authorization callback */
  int eOnePass;          /* ONEPASS_OFF or _SINGLE or _MULTI */
  int aiCurOnePass[2];   /* The write cursors opened by WHERE_ONEPASS */
  u8 *aToOpen = 0;       /* Open cursor iTabCur+j if aToOpen[j] is true */
  Index *pPk;            /* The PRIMARY KEY index on the table */
  int iPk = 0;           /* First of nPk registers holding PRIMARY KEY value */
  i16 nPk = 1;           /* Number of columns in the PRIMARY KEY */
  int iKey;              /* Memory cell holding key of row to be deleted */
  i16 nKey;              /* Number of memory cells in the row key */
  int iEphCur = 0;       /* Ephemeral table holding all primary key values */
  int iRowSet = 0;       /* Register for rowset of rows to delete */
  int addrBypass = 0;    /* Address of jump over the delete logic */
  int addrLoop = 0;      /* Top of the delete loop */
  int addrEphOpen = 0;   /* Instruction to open the Ephemeral table */
  int bComplex;          /* True if there are triggers or FKs or
                         ** subqueries in the WHERE clause */
 
#ifndef SQLITE_OMIT_TRIGGER
  int isView;                  /* True if attempting to delete from a view */
  Trigger *pTrigger;           /* List of table triggers, if required */
#endif

  memset(&sContext, 0, sizeof(sContext));
  db = pParse->db;
  if( pParse->nErr || db->mallocFailed ){
    goto delete_from_cleanup;
  }
  assert( pTabList->nSrc==1 );


  /* Locate the table which we want to delete.  This table has to be
  ** put in an SrcList structure because some of the subroutines we
  ** will be calling are designed to work with multiple tables and expect
  ** an SrcList* parameter instead of just a Table* parameter.
  */
  pTab = sqlite3SrcListLookup(pParse, pTabList);
  if( pTab==0 )  goto delete_from_cleanup;
#endif

  /* Resolve the column names in the WHERE clause.
  */
  memset(&sNC, 0, sizeof(sNC));
  sNC.pParse = pParse;
  sNC.pSrcList = pTabList;
  if( sqlite3ResolveExprNames(&sNC, pWhere) ){
    goto delete_from_cleanup;
  }

  /* Initialize the counter of the number of rows deleted, if
  ** we are counting rows.
  */
  if( db->flags & SQLITE_CountRows ){
    memCnt = ++pParse->nMem;
    sqlite3VdbeAddOp2(v, OP_Integer, 0, memCnt);
  }

#ifndef SQLITE_OMIT_TRUNCATE_OPTIMIZATION
  /* Special case: A DELETE without a WHERE clause deletes everything.
  ** It is easier just to erase the whole table. Prior to version 3.6.5,
  ** this optimization caused the row change count (the value returned by 
  ** API function sqlite3_count_changes) to be set incorrectly.
  **
  ** The "rcauth==SQLITE_OK" terms is the
  ** IMPLEMENTATION-OF: R-17228-37124 If the action code is SQLITE_DELETE and
  ** the callback returns SQLITE_IGNORE then the DELETE operation proceeds but
  ** the truncate optimization is disabled and all rows are deleted
  ** individually.
  */
  if( rcauth==SQLITE_OK
   && pWhere==0
   && !bComplex
   && !IsVirtual(pTab)
#ifdef SQLITE_ENABLE_PREUPDATE_HOOK
   && db->xPreUpdateCallback==0
#endif
  ){
    assert( !isView );
    sqlite3TableLock(pParse, iDb, pTab->tnum, 1, pTab->zName);
    if( HasRowid(pTab) ){
      sqlite3VdbeAddOp4(v, OP_Clear, pTab->tnum, iDb, memCnt,
                        pTab->zName, P4_STATIC);
    }
    for(pIdx=pTab->pIndex; pIdx; pIdx=pIdx->pNext){
      assert( pIdx->pSchema==pTab->pSchema );
      sqlite3VdbeAddOp2(v, OP_Clear, pIdx->tnum, iDb);
    }
  }else
#endif /* SQLITE_OMIT_TRUNCATE_OPTIMIZATION */
    if( sNC.ncFlags & NC_VarSelect ) bComplex = 1;
    wcf |= (bComplex ? 0 : WHERE_ONEPASS_MULTIROW);
    if( HasRowid(pTab) ){
      /* For a rowid table, initialize the RowSet to an empty set */
      pPk = 0;
      nPk = 1;
      iRowSet = ++pParse->nMem;
      sqlite3VdbeAddOp2(v, OP_Null, 0, iRowSet);
    }else{
      /* For a WITHOUT ROWID table, create an ephemeral table used to
      ** hold all primary keys for rows to be deleted. */
      pPk = sqlite3PrimaryKeyIndex(pTab);
      assert( pPk!=0 );
      nPk = pPk->nKeyCol;
      iPk = pParse->nMem+1;
      pParse->nMem += nPk;
      iEphCur = pParse->nTab++;
      addrEphOpen = sqlite3VdbeAddOp2(v, OP_OpenEphemeral, iEphCur, nPk);
      sqlite3VdbeSetP4KeyInfo(pParse, pPk);
    }
  
    /* Construct a query to find the rowid or primary key for every row
    ** to be deleted, based on the WHERE clause. Set variable eOnePass
    ** to indicate the strategy used to implement this delete:
    **
    **  ONEPASS_OFF:    Two-pass approach - use a FIFO for rowids/PK values.
    **  ONEPASS_SINGLE: One-pass approach - at most one row deleted.
    **  ONEPASS_MULTI:  One-pass approach - any number of rows may be deleted.
    */
    pWInfo = sqlite3WhereBegin(pParse, pTabList, pWhere, 0, 0, wcf, iTabCur+1);
    if( pWInfo==0 ) goto delete_from_cleanup;
    eOnePass = sqlite3WhereOkOnePass(pWInfo, aiCurOnePass);
    assert( IsVirtual(pTab)==0 || eOnePass!=ONEPASS_MULTI );
    assert( IsVirtual(pTab) || bComplex || eOnePass!=ONEPASS_OFF );
  
    /* Keep track of the number of rows to be deleted */
    if( db->flags & SQLITE_CountRows ){
      sqlite3VdbeAddOp2(v, OP_AddImm, memCnt, 1);
    }
  
    /* Extract the rowid or primary key for the current row */
    if( pPk ){
      for(i=0; i<nPk; i++){
        assert( pPk->aiColumn[i]>=0 );
        sqlite3ExprCodeGetColumnOfTable(v, pTab, iTabCur,
                                        pPk->aiColumn[i], iPk+i);
      }
      iKey = iPk;
    }else{
      iKey = pParse->nMem + 1;
      iKey = sqlite3ExprCodeGetColumn(pParse, pTab, -1, iTabCur, iKey, 0);
      if( iKey>pParse->nMem ) pParse->nMem = iKey;
    }
  
    if( eOnePass!=ONEPASS_OFF ){
      /* For ONEPASS, no need to store the rowid/primary-key. There is only
      ** one, so just keep it in its register(s) and fall through to the
        sqlite3VdbeAddOp4(v, OP_MakeRecord, iPk, nPk, iKey,
            sqlite3IndexAffinityStr(pParse->db, pPk), nPk);
        sqlite3VdbeAddOp4Int(v, OP_IdxInsert, iEphCur, iKey, iPk, nPk);
      }else{
        /* Add the rowid of the row to be deleted to the RowSet */
        nKey = 1;  /* OP_DeferredSeek always uses a single rowid */
        sqlite3VdbeAddOp2(v, OP_RowSetAdd, iRowSet, iKey);
      }
    }
  
    /* If this DELETE cannot use the ONEPASS strategy, this is the 
    ** end of the WHERE loop */
    if( eOnePass!=ONEPASS_OFF ){
      addrBypass = sqlite3VdbeMakeLabel(v);
    }else{
      sqlite3WhereEnd(pWInfo);
    }
  
    /* Unless this is a view, open cursors for the table we are 
    ** deleting from and all its indices. If this is a view, then the
    ** only effect this statement has is to fire the INSTEAD OF 
    ** triggers.
    */
    if( !isView ){
      int iAddrOnce = 0;
      if( eOnePass==ONEPASS_MULTI ){
        iAddrOnce = sqlite3VdbeAddOp0(v, OP_Once); VdbeCoverage(v);
      }
      testcase( IsVirtual(pTab) );
      sqlite3OpenTableAndIndices(pParse, pTab, OP_OpenWrite, OPFLAG_FORDELETE,
                                 iTabCur, aToOpen, &iDataCur, &iIdxCur);
      assert( pPk || IsVirtual(pTab) || iDataCur==iTabCur );
      assert( pPk || IsVirtual(pTab) || iIdxCur==iDataCur+1 );
      if( eOnePass==ONEPASS_MULTI ) sqlite3VdbeJumpHere(v, iAddrOnce);
    }
  
    /* Set up a loop over the rowids/primary-keys that were found in the
    ** where-clause loop above.
    */
    if( eOnePass!=ONEPASS_OFF ){
      assert( nKey==nPk );  /* OP_Found will use an unpacked key */
      if( !IsVirtual(pTab) && aToOpen[iDataCur-iTabCur] ){
        assert( pPk!=0 || pTab->pSelect!=0 );
        sqlite3VdbeAddOp4Int(v, OP_NotFound, iDataCur, addrBypass, iKey, nKey);
        VdbeCoverage(v);
      }
    }else if( pPk ){
      addrLoop = sqlite3VdbeAddOp1(v, OP_Rewind, iEphCur); VdbeCoverage(v);
      if( IsVirtual(pTab) ){
        sqlite3VdbeAddOp3(v, OP_Column, iEphCur, 0, iKey);
      }else{
        sqlite3VdbeAddOp2(v, OP_RowData, iEphCur, iKey);
      }
      assert( nKey==0 );  /* OP_Found will use a composite key */
    }else{
      addrLoop = sqlite3VdbeAddOp3(v, OP_RowSetRead, iRowSet, 0, iKey);
      VdbeCoverage(v);
      assert( nKey==1 );
    }  
  
    /* Delete the row */
#ifndef SQLITE_OMIT_VIRTUALTABLE
    if( IsVirtual(pTab) ){
      const char *pVTab = (const char *)sqlite3GetVTable(db, pTab);
      sqlite3VtabMakeWritable(pParse, pTab);
      sqlite3VdbeAddOp4(v, OP_VUpdate, 0, 1, iKey, pVTab, P4_VTAB);
      sqlite3VdbeChangeP5(v, OE_Abort);
      assert( eOnePass==ONEPASS_OFF || eOnePass==ONEPASS_SINGLE );
      sqlite3MayAbort(pParse);
      if( eOnePass==ONEPASS_SINGLE && sqlite3IsToplevel(pParse) ){
        pParse->isMultiWrite = 0;
      }
    }else
#endif
    {
      int count = (pParse->nested==0);    /* True to count changes */
      sqlite3GenerateRowDelete(pParse, pTab, pTrigger, iDataCur, iIdxCur,
          iKey, nKey, count, OE_Default, eOnePass, aiCurOnePass[1]);
    }
  
    /* End of the loop over all rowids/primary-keys. */
    if( eOnePass!=ONEPASS_OFF ){
      sqlite3VdbeResolveLabel(v, addrBypass);
      sqlite3WhereEnd(pWInfo);
    }else if( pPk ){
      sqlite3VdbeAddOp2(v, OP_Next, iEphCur, addrLoop+1); VdbeCoverage(v);
      sqlite3VdbeJumpHere(v, addrLoop);
    }else{
      sqlite3VdbeGoto(v, addrLoop);
      sqlite3VdbeJumpHere(v, addrLoop);
    }     
  } /* End non-truncate path */

  /* Update the sqlite_sequence table by storing the content of the
  ** maximum rowid counter values recorded while inserting into
    }

    /* Generate code that executes if the new index entry is not unique */
    assert( onError==OE_Rollback || onError==OE_Abort || onError==OE_Fail
        || onError==OE_Ignore || onError==OE_Replace );
    switch( onError ){
      case OE_Rollback:
      case OE_Abort:
      case OE_Fail: {
        sqlite3UniqueConstraint(pParse, onError, pIdx);
        break;
      }
      case OE_Ignore: {
        sqlite3VdbeGoto(v, ignoreDest);
        break;
      }
      default: {
        Trigger *pTrigger = 0;
        assert( onError==OE_Replace );
        sqlite3MultiWrite(pParse);
        if( db->flags&SQLITE_RecTriggers ){
          pTrigger = sqlite3TriggersExist(pParse, pTab, TK_DELETE, 0, 0);
        }
        sqlite3GenerateRowDelete(pParse, pTab, pTrigger, iDataCur, iIdxCur,
            regR, nPkField, 0, OE_Replace,
            (pIdx==pPk ? ONEPASS_SINGLE : ONEPASS_OFF), iThisCur);
        seenReplace = 1;
        break;
      }
    }
    sqlite3VdbeResolveLabel(v, addrUniqueOk);
    sqlite3ExprCachePop(pParse);
    if( regR!=regIdx ) sqlite3ReleaseTempRange(pParse, regR, nPkField);
  }
  if( ipkTop ){
    sqlite3VdbeGoto(v, ipkTop+1);
    sqlite3VdbeJumpHere(v, ipkBottom);
  }
  
  *pbMayReplace = seenReplace;
  VdbeModuleComment((v, "END: GenCnstCks(%d)", seenReplace));
}

#ifdef SQLITE_ENABLE_NULL_TRIM
/*
** Change the P5 operand on the last opcode (which should be an OP_MakeRecord)
** to be the number of columns in table pTab that must not be NULL-trimmed.
**
** Or if no columns of pTab may be NULL-trimmed, leave P5 at zero.
*/
SQLITE_PRIVATE void sqlite3SetMakeRecordP5(Vdbe *v, Table *pTab){
      if( SQLITE_OK==sqlite3ResolveExprNames(&sNC, pWhen) 
       && db->mallocFailed==0 
      ){
        iEndTrigger = sqlite3VdbeMakeLabel(v);
        sqlite3ExprIfFalse(pSubParse, pWhen, iEndTrigger, SQLITE_JUMPIFNULL);
      }
      sqlite3ExprDelete(db, pWhen);
    }

    /* Code the trigger program into the sub-vdbe. */
    codeTriggerProgram(pSubParse, pTrigger->step_list, orconf);

    /* Insert an OP_Halt at the end of the sub-program. */
    if( iEndTrigger ){
      sqlite3VdbeResolveLabel(v, iEndTrigger);
    }
    sqlite3VdbeAddOp0(v, OP_Halt);
    VdbeComment((v, "End: %s.%s", pTrigger->zName, onErrorText(orconf)));

    transferParseError(pParse, pSubParse);
    if( db->mallocFailed==0 && pParse->nErr==0 ){
      pProgram->aOp = sqlite3VdbeTakeOpArray(v, &pProgram->nOp, &pTop->nMaxArg);
    }
    pProgram->nMem = pSubParse->nMem;
    pProgram->nCsr = pSubParse->nTab;
    pProgram->token = (void *)pTrigger;
    pPrg->aColmask[0] = pSubParse->oldmask;
    pPrg->aColmask[1] = pSubParse->newmask;
    sqlite3VdbeDelete(v);
  }

  assert( !pSubParse->pAinc       && !pSubParse->pZombieTab );
  assert( !pSubParse->pTriggerPrg && !pSubParse->nMaxArg );
  sqlite3ParserReset(pSubParse);
  sqlite3StackFree(db, pSubParse);

  return pPrg;
}
    
/*
** Return a pointer to a TriggerPrg object containing the sub-program for
** trigger pTrigger with default ON CONFLICT algorithm orconf. If no such
** TriggerPrg object exists, a new object is allocated and populated before
** being returned.
*/
static TriggerPrg *getRowTrigger(
  Parse *pParse,       /* Current parse context */
  Trigger *pTrigger,   /* Trigger to code */
  Table *pTab,         /* The table trigger pTrigger is attached to */
  int orconf           /* ON CONFLICT algorithm. */
){
  ** there are any triggers or foreign keys to process, or rows may
  ** be deleted as a result of REPLACE conflict handling. Any of these
  ** things might disturb a cursor being used to scan through the table
  ** or index, causing a single-pass approach to malfunction.  */
  flags = WHERE_ONEPASS_DESIRED|WHERE_SEEK_UNIQ_TABLE;
  if( !pParse->nested && !pTrigger && !hasFK && !chngKey && !bReplace ){
    flags |= WHERE_ONEPASS_MULTIROW;
  }
  pWInfo = sqlite3WhereBegin(pParse, pTabList, pWhere, 0, 0, flags, iIdxCur);
  if( pWInfo==0 ) goto update_cleanup;

  /* A one-pass strategy that might update more than one row may not
  ** be used if any column of the index used for the scan is being
  ** updated. Otherwise, if there is an index on "b", statements like
  ** the following could create an infinite loop:
  **
  **   UPDATE t1 SET b=b+1 WHERE b>?
  **
  ** Fall back to ONEPASS_OFF if where.c has selected a ONEPASS_MULTI
  ** strategy that uses an index for which one or more columns are being
  ** updated.  */
  eOnePass = sqlite3WhereOkOnePass(pWInfo, aiCurOnePass);
  if( eOnePass==ONEPASS_MULTI ){
    int iCur = aiCurOnePass[1];
    if( iCur>=0 && iCur!=iDataCur && aToOpen[iCur-iBaseCur] ){
      eOnePass = ONEPASS_OFF;
    }
    assert( iCur!=iDataCur || !HasRowid(pTab) );
  }
  
  if( HasRowid(pTab) ){
    /* Read the rowid of the current row of the WHERE scan. In ONEPASS_OFF
    ** mode, write the rowid into the FIFO. In either of the one-pass modes,
    ** leave it in register regOldRowid.  */
    sqlite3VdbeAddOp2(v, OP_Rowid, iDataCur, regOldRowid);
    if( eOnePass==ONEPASS_OFF ){
      sqlite3VdbeAddOp2(v, OP_RowSetAdd, regRowSet, regOldRowid);
    }
  }else{
    /* Read the PK of the current row into an array of registers. In
    ** ONEPASS_OFF mode, serialize the array into a record and store it in
    ** the ephemeral table. Or, in ONEPASS_SINGLE or MULTI mode, change
    ** the OP_OpenEphemeral instruction to a Noop (the ephemeral table 
    ** is not required) and leave the PK fields in the array of registers.  */
    for(i=0; i<nPk; i++){
      assert( pPk->aiColumn[i]>=0 );
      sqlite3ExprCodeGetColumnOfTable(v, pTab, iDataCur,pPk->aiColumn[i],iPk+i);
    }
    if( eOnePass ){
      sqlite3VdbeChangeToNoop(v, addrOpen);
      nKey = nPk;
      regKey = iPk;
    }else{
      sqlite3VdbeAddOp4(v, OP_MakeRecord, iPk, nPk, regKey,
                        sqlite3IndexAffinityStr(db, pPk), nPk);
      sqlite3VdbeAddOp4Int(v, OP_IdxInsert, iEph, regKey, iPk, nPk);
    }
  }

  if( eOnePass!=ONEPASS_MULTI ){
    sqlite3WhereEnd(pWInfo);
  }

  labelBreak = sqlite3VdbeMakeLabel(v);
  if( !isView ){
    int addrOnce = 0;

    /* Open every index that needs updating. */
    if( eOnePass!=ONEPASS_OFF ){
      if( aiCurOnePass[0]>=0 ) aToOpen[aiCurOnePass[0]-iBaseCur] = 0;
      if( aiCurOnePass[1]>=0 ) aToOpen[aiCurOnePass[1]-iBaseCur] = 0;
    }

    if( eOnePass==ONEPASS_MULTI && (nIdx-(aiCurOnePass[1]>=0))>0 ){
      addrOnce = sqlite3VdbeAddOp0(v, OP_Once); VdbeCoverage(v);
    }
    sqlite3OpenTableAndIndices(pParse, pTab, OP_OpenWrite, 0, iBaseCur, aToOpen,
                               0, 0);
    if( addrOnce ) sqlite3VdbeJumpHere(v, addrOnce);
  }

  /* Top of the update loop */
  if( eOnePass!=ONEPASS_OFF ){
    if( !isView && aiCurOnePass[0]!=iDataCur && aiCurOnePass[1]!=iDataCur ){
      assert( pPk );
      sqlite3VdbeAddOp4Int(v, OP_NotFound, iDataCur, labelBreak, regKey, nKey);
      VdbeCoverageNeverTaken(v);
    }
    if( eOnePass==ONEPASS_SINGLE ){
      labelContinue = labelBreak;
    }else{
      labelContinue = sqlite3VdbeMakeLabel(v);
    }
    sqlite3VdbeAddOp2(v, OP_IsNull, pPk ? regKey : regOldRowid, labelBreak);
    VdbeCoverageIf(v, pPk==0);
    VdbeCoverageIf(v, pPk!=0);
  }else if( pPk ){
    labelContinue = sqlite3VdbeMakeLabel(v);
    sqlite3VdbeAddOp2(v, OP_Rewind, iEph, labelBreak); VdbeCoverage(v);
    addrTop = sqlite3VdbeAddOp2(v, OP_RowData, iEph, regKey);
    sqlite3VdbeAddOp4Int(v, OP_NotFound, iDataCur, labelContinue, regKey, 0);
    VdbeCoverage(v);
  }else{
    labelContinue = sqlite3VdbeAddOp3(v, OP_RowSetRead, regRowSet, labelBreak,
                             regOldRowid);
    VdbeCoverage(v);
    sqlite3VdbeAddOp3(v, OP_NotExists, iDataCur, labelContinue, regOldRowid);
    VdbeCoverage(v);
  }

  /* If the record number will change, set register regNewRowid to
  ** contain the new value. If the record number is not being modified,
  ** then regNewRowid is the same register as regOldRowid, which is
  ** already populated.  */
        regNewRowid
    );
    if( eOnePass==ONEPASS_MULTI ){
      assert( hasFK==0 && chngKey==0 );
      sqlite3VdbeChangeP5(v, OPFLAG_SAVEPOSITION);
    }
    if( !pParse->nested ){
      sqlite3VdbeAppendP4(v, pTab, P4_TABLE);
    }
#else
    if( hasFK>1 || chngKey ){
      sqlite3VdbeAddOp2(v, OP_Delete, iDataCur, 0);
    }
#endif
    if( bReplace || chngKey ){
      sqlite3VdbeJumpHere(v, addr1);
    }

    if( hasFK ){
      sqlite3FkCheck(pParse, pTab, 0, regNewRowid, aXRef, chngKey);
    }
  
    /* Insert the new index entries and the new record. */
    sqlite3CompleteInsertion(
        pParse, pTab, iDataCur, iIdxCur, regNewRowid, aRegIdx, 
        OPFLAG_ISUPDATE | (eOnePass==ONEPASS_MULTI ? OPFLAG_SAVEPOSITION : 0), 
        0, 0
    );

    /* Do any ON CASCADE, SET NULL or SET DEFAULT operations required to
    ** handle rows (possibly in other tables) that refer via a foreign key
    ** to the row just updated. */ 
    if( hasFK ){
      sqlite3FkActions(pParse, pTab, pChanges, regOldRowid, aXRef, chngKey);
    }
  }

  /* Increment the row counter 
  */
  if( (db->flags & SQLITE_CountRows) && !pParse->pTriggerTab){
    sqlite3VdbeAddOp2(v, OP_AddImm, regRowCount, 1);
  }

  sqlite3CodeRowTrigger(pParse, pTrigger, TK_UPDATE, pChanges, 
      TRIGGER_AFTER, pTab, regOldRowid, onError, labelContinue);

  /* Repeat the above with the next record to be updated, until
  ** all record selected by the WHERE clause have been updated.
  */
  if( eOnePass==ONEPASS_SINGLE ){
    /* Nothing to do at end-of-loop for a single-pass */
    if( aXRef[i]>=0 ){
      sqlite3ExprCode(pParse, pChanges->a[aXRef[i]].pExpr, regArg+2+i);
    }else{
      sqlite3VdbeAddOp3(v, OP_VColumn, iCsr, i, regArg+2+i);
      sqlite3VdbeChangeP5(v, 1); /* Enable sqlite3_vtab_nochange() */
    }
  }
  if( HasRowid(pTab) ){
    sqlite3VdbeAddOp2(v, OP_Rowid, iCsr, regArg);
    if( pRowid ){
      sqlite3ExprCode(pParse, pRowid, regArg+1);
    }else{
      sqlite3VdbeAddOp2(v, OP_Rowid, iCsr, regArg+1);
    }
  }else{
    Index *pPk;   /* PRIMARY KEY index */
    i16 iPk;      /* PRIMARY KEY column */
    pPk = sqlite3PrimaryKeyIndex(pTab);
    assert( pPk!=0 );
    assert( pPk->nKeyCol==1 );
    iPk = pPk->aiColumn[0];
    sqlite3VdbeAddOp3(v, OP_VColumn, iCsr, iPk, regArg);
    sqlite3VdbeAddOp2(v, OP_SCopy, regArg+2+iPk, regArg+1);
  }

  bOnePass = sqlite3WhereOkOnePass(pWInfo, aDummy);

  if( bOnePass ){
    /* If using the onepass strategy, no-op out the OP_OpenEphemeral coded
    ** above. Also, if this is a top-level parse (not a trigger), clear the
    ** multi-write flag so that the VM does not open a statement journal */
    sqlite3VdbeChangeToNoop(v, addr);
    if( sqlite3IsToplevel(pParse) ){
      pParse->isMultiWrite = 0;
    }
  }else{
    /* Create a record from the argument register contents and insert it into
    ** the ephemeral table. */
    sqlite3VdbeAddOp3(v, OP_MakeRecord, regArg, nArg, regRec);
#ifdef SQLITE_DEBUG
    /* Signal an assert() within OP_MakeRecord that it is allowed to
    ** accept no-change records with serial_type 10 */
    sqlite3VdbeChangeP5(v, OPFLAG_NOCHNG_MAGIC);
#endif
    sqlite3VdbeAddOp2(v, OP_NewRowid, ephemTab, regRowid);
    sqlite3VdbeAddOp3(v, OP_Insert, ephemTab, regRec, regRowid);
  }


  if( bOnePass==0 ){
    /* End the virtual table scan */
  Parse *pParse = pWInfo->pParse; /* Parse context */
  Vdbe *v = pParse->pVdbe;        /* Vdbe to generate code within */

  assert( iIdxCur>0 );
  assert( pIdx->aiColumn[pIdx->nColumn-1]==-1 );
  
  sqlite3VdbeAddOp3(v, OP_DeferredSeek, iIdxCur, 0, iCur);
  if( (pWInfo->wctrlFlags & WHERE_OR_SUBCLAUSE)
   && DbMaskAllZero(sqlite3ParseToplevel(pParse)->writeMask)
  ){
    int i;
    Table *pTab = pIdx->pTable;
    int *ai = (int*)sqlite3DbMallocZero(pParse->db, sizeof(int)*(pTab->nCol+1));
    if( ai ){
      ai[0] = pTab->nCol;
      for(i=0; i<pIdx->nColumn-1; i++){
        assert( pIdx->aiColumn[i]<pTab->nCol );
        if( pIdx->aiColumn[i]>=0 ) ai[pIdx->aiColumn[i]+1] = i+1;
      }
      sqlite3VdbeChangeP4(v, -1, (char*)ai, P4_INTARRAY);
    }
  }
}

/*
** If the expression passed as the second argument is a vector, generate
** code to write the first nReg elements of the vector into an array
** of registers starting with iReg.
**
** If the expression is not a vector, then nReg must be passed 1. In
** this case, generate code to evaluate the expression and leave the
** result in register iReg.
*/
static void codeExprOrVector(Parse *pParse, Expr *p, int iReg, int nReg){
  assert( nReg>0 );
  if( p && sqlite3ExprIsVector(p) ){
#ifndef SQLITE_OMIT_SUBQUERY
    if( (p->flags & EP_xIsSelect) ){
      Vdbe *v = pParse->pVdbe;
      int iSelect = sqlite3CodeSubselect(pParse, p, 0, 0);
      sqlite3VdbeAddOp3(v, OP_Copy, iSelect, iReg, nReg-1);
    }else
#endif
    {
      int i;
      ExprList *pList = p->x.pList;
      assert( nReg<=pList->nExpr );
      for(i=0; i<nReg; i++){
        sqlite3ExprCode(pParse, pList->a[i].pExpr, iReg+i);
      }
    }
    int iIn;    /* Counter for IN constraints */

    sqlite3ExprCachePush(pParse);
    iReg = sqlite3GetTempRange(pParse, nConstraint+2);
    addrNotFound = pLevel->addrBrk;
    for(j=0; j<nConstraint; j++){
      int iTarget = iReg+j+2;
      pTerm = pLoop->aLTerm[j];
      if( NEVER(pTerm==0) ) continue;
      if( pTerm->eOperator & WO_IN ){
        codeEqualityTerm(pParse, pTerm, pLevel, j, bRev, iTarget);
        addrNotFound = pLevel->addrNxt;
      }else{
        Expr *pRight = pTerm->pExpr->pRight;
        codeExprOrVector(pParse, pRight, iTarget, 1);
      }
    }
    sqlite3VdbeAddOp2(v, OP_Integer, pLoop->u.vtab.idxNum, iReg);
    sqlite3VdbeAddOp2(v, OP_Integer, nConstraint, iReg+1);
    sqlite3VdbeAddOp4(v, OP_VFilter, iCur, addrNotFound, iReg,
                      pLoop->u.vtab.idxStr,
                      pLoop->u.vtab.needFree ? P4_DYNAMIC : P4_STATIC);
    VdbeCoverage(v);
    pLoop->u.vtab.needFree = 0;
    pLevel->p1 = iCur;
    pLevel->op = pWInfo->eOnePass ? OP_Noop : OP_VNext;
    pLevel->p2 = sqlite3VdbeCurrentAddr(v);
    iIn = pLevel->u.in.nIn;
    for(j=nConstraint-1; j>=0; j--){
      pTerm = pLoop->aLTerm[j];
      if( j<16 && (pLoop->u.vtab.omitMask>>j)&1 ){
        disableTerm(pLevel, pTerm);
      }else if( (pTerm->eOperator & WO_IN)!=0 ){
        Expr *pCompare;  /* The comparison operator */
        Expr *pRight;    /* RHS of the comparison */
        VdbeOp *pOp;     /* Opcode to access the value of the IN constraint */

        /* Reload the constraint value into reg[iReg+j+2].  The same value
        ** was loaded into the same register prior to the OP_VFilter, but
        ** the xFilter implementation might have changed the datatype or
        ** encoding of the value in the register, so it *must* be reloaded. */
        assert( pLevel->u.in.aInLoop!=0 || db->mallocFailed );
        if( !db->mallocFailed ){
          assert( iIn>0 );
          pOp = sqlite3VdbeGetOp(v, pLevel->u.in.aInLoop[--iIn].addrInTop);
          assert( pOp->opcode==OP_Column || pOp->opcode==OP_Rowid );
          assert( pOp->opcode!=OP_Column || pOp->p3==iReg+j+2 );
          assert( pOp->opcode!=OP_Rowid || pOp->p2==iReg+j+2 );
          testcase( pOp->opcode==OP_Rowid );
          sqlite3VdbeAddOp3(v, pOp->opcode, pOp->p1, pOp->p2, pOp->p3);
        }
      endEq = 0;
      nConstraint++;
    }
    sqlite3DbFree(db, zStartAff);
    sqlite3DbFree(db, zEndAff);

    /* Top of the loop body */
    pLevel->p2 = sqlite3VdbeCurrentAddr(v);

    /* Check if the index cursor is past the end of the range. */
    if( nConstraint ){
      op = aEndOp[bRev*2 + endEq];
      sqlite3VdbeAddOp4Int(v, op, iIdxCur, addrNxt, regBase, nConstraint);
      testcase( op==OP_IdxGT );  VdbeCoverageIf(v, op==OP_IdxGT );
      testcase( op==OP_IdxGE );  VdbeCoverageIf(v, op==OP_IdxGE );
      testcase( op==OP_IdxLT );  VdbeCoverageIf(v, op==OP_IdxLT );
      testcase( op==OP_IdxLE );  VdbeCoverageIf(v, op==OP_IdxLE );
    }

    /* Seek the table cursor, if required */
    if( omitTable ){
      /* pIdx is a covering index.  No need to access the main table. */
    }else if( HasRowid(pIdx->pTable) ){
      if( (pWInfo->wctrlFlags & WHERE_SEEK_TABLE) || (
          (pWInfo->wctrlFlags & WHERE_SEEK_UNIQ_TABLE) 
       && (pWInfo->eOnePass==ONEPASS_SINGLE)
      )){
        iRowidReg = ++pParse->nMem;
        sqlite3VdbeAddOp2(v, OP_IdxRowid, iIdxCur, iRowidReg);
        sqlite3ExprCacheStore(pParse, iCur, -1, iRowidReg);
        sqlite3VdbeAddOp3(v, OP_NotExists, iCur, 0, iRowidReg);
        VdbeCoverage(v);
      }else{
        codeDeferredSeek(pWInfo, pIdx, iCur, iIdxCur);
      }
    }else if( iCur!=iIdxCur ){
      Index *pPk = sqlite3PrimaryKeyIndex(pIdx->pTable);
      iRowidReg = sqlite3GetTempRange(pParse, pPk->nKeyCol);
      for(j=0; j<pPk->nKeyCol; j++){
        k = sqlite3ColumnOfIndex(pIdx, pPk->aiColumn[j]);
        sqlite3VdbeAddOp3(v, OP_Column, iIdxCur, k, iRowidReg+j);
      }
      sqlite3VdbeAddOp4Int(v, OP_NotFound, iCur, addrCont,
                           iRowidReg, pPk->nKeyCol); VdbeCoverage(v);
    }

    /* If pIdx is an index on one or more expressions, then look through
    ** all the expressions in pWInfo and try to transform matching expressions
    ** into reference to index columns.
    */
    whereIndexExprTrans(pIdx, iCur, iIdxCur, pWInfo);
  return pWInfo->iBreak;
}

/*
** Return ONEPASS_OFF (0) if an UPDATE or DELETE statement is unable to
** operate directly on the rowis returned by a WHERE clause.  Return
** ONEPASS_SINGLE (1) if the statement can operation directly because only
** a single row is to be changed.  Return ONEPASS_MULTI (2) if the one-pass
** optimization can be used on multiple 
**
** If the ONEPASS optimization is used (if this routine returns true)
** then also write the indices of open cursors used by ONEPASS
** into aiCur[0] and aiCur[1].  iaCur[0] gets the cursor of the data
** table and iaCur[1] gets the cursor used by an auxiliary index.
** Either value may be -1, indicating that cursor is not used.
** Any cursors returned will have been opened for writing.
**
** aiCur[0] and aiCur[1] both get -1 if the where-clause logic is
** unable to use the ONEPASS optimization.
*/
SQLITE_PRIVATE int sqlite3WhereOkOnePass(WhereInfo *pWInfo, int *aiCur){
  memcpy(aiCur, pWInfo->aiCurOnePass, sizeof(int)*2);
#ifdef WHERETRACE_ENABLED
  if( sqlite3WhereTrace && pWInfo->eOnePass!=ONEPASS_OFF ){
    sqlite3DebugPrintf("%s cursors: %d %d\n",
         pWInfo->eOnePass==ONEPASS_SINGLE ? "ONEPASS_SINGLE" : "ONEPASS_MULTI",
         aiCur[0], aiCur[1]);
  }
#endif
  return pWInfo->eOnePass;
}

/*
** Move the content of pSrc into pDest
*/
static void whereOrMove(WhereOrSet *pDest, WhereOrSet *pSrc){
  pDest->n = pSrc->n;
  memcpy(pDest->a, pSrc->a, pDest->n*sizeof(pDest->a[0]));
}

/*
** Try to insert a new prerequisite/cost entry into the WhereOrSet pSet.
**
** The new entry might overwrite an existing entry, or it might be
** appended, or it might be discarded.  Do whatever is the right thing
** so that pSet keeps the N_OR_COST best entries seen so far.
*/
static int whereOrInsert(
  WhereOrSet *pSet,      /* The WhereOrSet to be updated */
  Bitmask prereq,        /* Prerequisites of the new entry */
  LogEst rRun,           /* Run-cost of the new entry */
        }
      }
    }
  } /* End the loop over all WhereLoops from outer-most down to inner-most */
  if( obSat==obDone ) return (i8)nOrderBy;
  if( !isOrderDistinct ){
    for(i=nOrderBy-1; i>0; i--){
      Bitmask m = MASKBIT(i) - 1;
      if( (obSat&m)==m ) return i;
    }
    return 0;
  }
  return -1;
}


/*
** If the WHERE_GROUPBY flag is set in the mask passed to sqlite3WhereBegin(),
** the planner assumes that the specified pOrderBy list is actually a GROUP
** BY clause - and so any order that groups rows as required satisfies the
** request.
**
** Normally, in this case it is not possible for the caller to determine
** whether or not the rows are really being delivered in sorted order, or
** just in some other order that provides the required grouping. However,
** if the WHERE_SORTBYGROUP flag is also passed to sqlite3WhereBegin(), then
** this function may be called on the returned WhereInfo object. It returns
** true if the rows really will be sorted in the specified order, or false
** otherwise.
**
** For example, assuming:
**
**   CREATE INDEX i1 ON t1(x, Y);
**
** then
**
**   SELECT * FROM t1 GROUP BY x,y ORDER BY x,y;   -- IsSorted()==1
**   SELECT * FROM t1 GROUP BY y,x ORDER BY y,x;   -- IsSorted()==0
*/
SQLITE_PRIVATE int sqlite3WhereIsSorted(WhereInfo *pWInfo){
  assert( pWInfo->wctrlFlags & WHERE_GROUPBY );
  assert( pWInfo->wctrlFlags & WHERE_SORTBYGROUP );
  return pWInfo->sorted;
}

#ifdef WHERETRACE_ENABLED
/* For debugging use only: */
static const char *wherePathName(WherePath *pPath, int nLoop, WhereLoop *pLast){
  static char zName[65];
  int i;
  for(i=0; i<nLoop; i++){ zName[i] = pPath->aLoop[i]->cId; }
  ** and the WhereMaskSet structure. Since WhereClause contains an 8-byte
  ** field (type Bitmask) it must be aligned on an 8-byte boundary on
  ** some architectures. Hence the ROUND8() below.
  */
  nByteWInfo = ROUND8(sizeof(WhereInfo)+(nTabList-1)*sizeof(WhereLevel));
  pWInfo = sqlite3DbMallocRawNN(db, nByteWInfo + sizeof(WhereLoop));
  if( db->mallocFailed ){
    sqlite3DbFree(db, pWInfo);
    pWInfo = 0;
    goto whereBeginError;
  }
  pWInfo->pParse = pParse;
  pWInfo->pTabList = pTabList;
  pWInfo->pOrderBy = pOrderBy;
  pWInfo->pWhere = pWhere;
  pWInfo->pResultSet = pResultSet;
  pWInfo->aiCurOnePass[0] = pWInfo->aiCurOnePass[1] = -1;
  pWInfo->nLevel = nTabList;
  pWInfo->iBreak = pWInfo->iContinue = sqlite3VdbeMakeLabel(v);
  pWInfo->wctrlFlags = wctrlFlags;
  pWInfo->iLimit = iAuxArg;
  pWInfo->savedNQueryLoop = pParse->nQueryLoop;
  memset(&pWInfo->nOBSat, 0, 
         offsetof(WhereInfo,sWC) - offsetof(WhereInfo,nOBSat));
  memset(&pWInfo->a[0], 0, sizeof(WhereLoop)+nTabList*sizeof(WhereLevel));
  assert( pWInfo->eOnePass==ONEPASS_OFF );  /* ONEPASS defaults to OFF */
  pMaskSet = &pWInfo->sMaskSet;
  sWLB.pWInfo = pWInfo;
  sWLB.pWC = &pWInfo->sWC;
  sWLB.pNew = (WhereLoop*)(((char*)pWInfo)+nByteWInfo);
  assert( EIGHT_BYTE_ALIGNMENT(sWLB.pNew) );
  whereLoopInit(sWLB.pNew);
#ifdef SQLITE_DEBUG
  sWLB.pNew->cId = '*';
#endif

  /* Split the WHERE clause into separate subexpressions where each
  ** subexpression is separated by an AND operator.
  */
  initMaskSet(pMaskSet);
  sqlite3WhereClauseInit(&pWInfo->sWC, pWInfo);
  sqlite3WhereSplit(&pWInfo->sWC, pWhere, TK_AND);
    
  /* Special case: No FROM clause
  */
  if( nTabList==0 ){
    if( pOrderBy ) pWInfo->nOBSat = pOrderBy->nExpr;
    if( wctrlFlags & WHERE_WANT_DISTINCT ){
      pWInfo->eDistinct = WHERE_DISTINCT_UNIQUE;
    }
  }else{
  ** A one-pass approach can be used if the caller has requested one
  ** and either (a) the scan visits at most one row or (b) each
  ** of the following are true:
  **
  **   * the caller has indicated that a one-pass approach can be used
  **     with multiple rows (by setting WHERE_ONEPASS_MULTIROW), and
  **   * the table is not a virtual table, and
  **   * either the scan does not use the OR optimization or the caller
  **     is a DELETE operation (WHERE_DUPLICATES_OK is only specified
  **     for DELETE).
  **
  ** The last qualification is because an UPDATE statement uses
  ** WhereInfo.aiCurOnePass[1] to determine whether or not it really can
  ** use a one-pass approach, and this is not set accurately for scans
  ** that use the OR optimization.
  */
  assert( (wctrlFlags & WHERE_ONEPASS_DESIRED)==0 || pWInfo->nLevel==1 );
  if( (wctrlFlags & WHERE_ONEPASS_DESIRED)!=0 ){
    int wsFlags = pWInfo->a[0].pWLoop->wsFlags;
    int bOnerow = (wsFlags & WHERE_ONEROW)!=0;
    if( bOnerow || (
        0!=(wctrlFlags & WHERE_ONEPASS_MULTIROW)
     && 0==(wsFlags & WHERE_VIRTUALTABLE)
     && (0==(wsFlags & WHERE_MULTI_OR) || (wctrlFlags & WHERE_DUPLICATES_OK))
    )){
      pWInfo->eOnePass = bOnerow ? ONEPASS_SINGLE : ONEPASS_MULTI;
      if( HasRowid(pTabList->a[0].pTab) && (wsFlags & WHERE_IDX_ONLY) ){
        if( wctrlFlags & WHERE_ONEPASS_MULTIROW ){
          bFordelete = OPFLAG_FORDELETE;
        }
        pWInfo->a[0].pWLoop->wsFlags = (wsFlags & ~WHERE_IDX_ONLY);
      }
    }
  }

  /* Open all tables in the pTabList and any indices selected for
  ** searching those tables.
  */
  for(ii=0, pLevel=pWInfo->a; ii<nTabList; ii++, pLevel++){
    Table *pTab;     /* Table to open */
    int iDb;         /* Index of database containing table/index */
    struct SrcList_item *pTabItem;

    pTabItem = &pTabList->a[pLevel->iFrom];
    pTab = pTabItem->pTab;
    iDb = sqlite3SchemaToIndex(db, pTab->pSchema);
    pLoop = pLevel->pWLoop;
    if( (pTab->tabFlags & TF_Ephemeral)!=0 || pTab->pSelect ){
      /* Do nothing */
    }else
#ifndef SQLITE_OMIT_VIRTUALTABLE
    if( (pLoop->wsFlags & WHERE_VIRTUALTABLE)!=0 ){
      const char *pVTab = (const char *)sqlite3GetVTable(db, pTab);
      int iCur = pTabItem->iCursor;
      sqlite3VdbeAddOp4(v, OP_VOpen, iCur, 0, 0, pVTab, P4_VTAB);
    }else if( IsVirtual(pTab) ){
      /* noop */
    }else
#endif
    if( (pLoop->wsFlags & WHERE_IDX_ONLY)==0
         && (wctrlFlags & WHERE_OR_SUBCLAUSE)==0 ){
      int op = OP_OpenRead;
      if( pWInfo->eOnePass!=ONEPASS_OFF ){
        op = OP_OpenWrite;
        pWInfo->aiCurOnePass[0] = pTabItem->iCursor;
      };
      sqlite3OpenTable(pParse, pTabItem->iCursor, iDb, pTab, op);
      assert( pTabItem->iCursor==pLevel->iTabCur );
      testcase( pWInfo->eOnePass==ONEPASS_OFF && pTab->nCol==BMS-1 );
      testcase( pWInfo->eOnePass==ONEPASS_OFF && pTab->nCol==BMS );
      if( pWInfo->eOnePass==ONEPASS_OFF && pTab->nCol<BMS && HasRowid(pTab) ){
        Bitmask b = pTabItem->colUsed;
        int n = 0;
        for(; b; b=b>>1, n++){}
        sqlite3VdbeChangeP4(v, -1, SQLITE_INT_TO_PTR(n), P4_INT32);
        assert( n<=pTab->nCol );
      }
#ifdef SQLITE_ENABLE_CURSOR_HINTS
      if( pLoop->u.btree.pIndex!=0 ){
        sqlite3VdbeChangeP5(v, OPFLAG_SEEKEQ|bFordelete);
      }else
#endif
      {
        sqlite3VdbeChangeP5(v, bFordelete);
      }
#ifdef SQLITE_ENABLE_COLUMN_USED_MASK
      sqlite3VdbeAddOp4Dup8(v, OP_ColumnsUsed, pTabItem->iCursor, 0, 0,
                            (const u8*)&pTabItem->colUsed, P4_INT64);
#endif
    }else{
      sqlite3TableLock(pParse, iDb, pTab->tnum, 0, pTab->zName);
    }
    if( pLoop->wsFlags & WHERE_INDEXED ){
      Index *pIx = pLoop->u.btree.pIndex;
      int iIndexCur;
      int op = OP_OpenRead;
    */
    if( pTabItem->fg.viaCoroutine ){
      testcase( pParse->db->mallocFailed );
      translateColumnToCopy(pParse, pLevel->addrBody, pLevel->iTabCur,
                            pTabItem->regResult, 0);
      continue;
    }

    /* If this scan uses an index, make VDBE code substitutions to read data
    ** from the index instead of from the table where possible.  In some cases
    ** this optimization prevents the table from ever being read, which can
    ** yield a significant performance boost.
    ** 
    ** Calls to the code generator in between sqlite3WhereBegin and
    ** sqlite3WhereEnd will have created code that references the table
    ** directly.  This loop scans all that code looking for opcodes
    ** that reference the table and converts them into opcodes that
    ** reference the index.
    */
    if( pLoop->wsFlags & (WHERE_INDEXED|WHERE_IDX_ONLY) ){
      pIdx = pLoop->u.btree.pIndex;
    }else if( pLoop->wsFlags & WHERE_MULTI_OR ){
      pIdx = pLevel->u.pCovidx;
    }
    if( pIdx
     && (pWInfo->eOnePass==ONEPASS_OFF || !HasRowid(pIdx->pTable))
     && !db->mallocFailed
    ){
      last = sqlite3VdbeCurrentAddr(v);
      k = pLevel->addrBody;
      pOp = sqlite3VdbeGetOp(v, k);
      for(; k<last; k++, pOp++){
        if( pOp->p1!=pLevel->iTabCur ) continue;
        if( pOp->opcode==OP_Column
#ifdef SQLITE_ENABLE_OFFSET_SQL_FUNC
         || pOp->opcode==OP_Offset
#endif
        ){
          int x = pOp->p2;
          assert( pIdx->pTable==pTab );
          if( !HasRowid(pTab) ){
            Index *pPk = sqlite3PrimaryKeyIndex(pTab);
            x = pPk->aiColumn[x];
            assert( x>=0 );
          }
          x = sqlite3ColumnOfIndex(pIdx, x);
          if( x>=0 ){
            pOp->p2 = x;
            pOp->p1 = pLevel->iIdxCur;
          }
          assert( (pLoop->wsFlags & WHERE_IDX_ONLY)==0 || x>=0 
){
  sqlite3ParserARG_FETCH;
#ifndef NDEBUG
  if( yyTraceFILE ){
    fprintf(yyTraceFILE,"%sFail!\n",yyTracePrompt);
  }
#endif
  while( yypParser->yytos>yypParser->yystack ) yy_pop_parser_stack(yypParser);
  /* Here code is inserted which will be executed whenever the
  ** parser fails */
/************ Begin %parse_failure code ***************************************/
/************ End %parse_failure code *****************************************/
  sqlite3ParserARG_STORE; /* Suppress warning about unused %extra_argument variable */
}
#endif /* YYNOERRORRECOVERY */

/*
** The following code executes when a syntax error first occurs.
*/
static void yy_syntax_error(
  yyParser *yypParser,           /* The parser */
  int yymajor,                   /* The major type of the error token */
  sqlite3ParserTOKENTYPE yyminor         /* The minor type of the error token */
){
  sqlite3ParserARG_FETCH;
#define TOKEN yyminor
/************ Begin %syntax_error code ****************************************/

  UNUSED_PARAMETER(yymajor);  /* Silence some compiler warnings */
  if( TOKEN.z[0] ){
    sqlite3ErrorMsg(pParse, "near \"%T\": syntax error", &TOKEN);
  }else{
    sqlite3ErrorMsg(pParse, "incomplete input");
  }
/************ End %syntax_error code ******************************************/
  sqlite3ParserARG_STORE; /* Suppress warning about unused %extra_argument variable */
}

/*
** The following is executed when the parser accepts
*/
static void yy_accept(
  yyParser *yypParser           /* The parser */
){
  sqlite3ParserARG_FETCH;
#ifndef NDEBUG
  if( yyTraceFILE ){
    fprintf(yyTraceFILE,"%sAccept!\n",yyTracePrompt);
  }
#endif
#ifndef YYNOERRORRECOVERY
      n = sqlite3GetToken((u8*)zSql, &tokenType);
      mxSqlLen -= n;
      if( mxSqlLen<0 ){
        pParse->rc = SQLITE_TOOBIG;
        break;
      }
    }else{
      /* Upon reaching the end of input, call the parser two more times
      ** with tokens TK_SEMI and 0, in that order. */
      if( lastTokenParsed==TK_SEMI ){
        tokenType = 0;
      }else if( lastTokenParsed==0 ){
        break;
      }else{
        tokenType = TK_SEMI;
      }
      n = 0;
    }
    if( tokenType>=TK_SPACE ){
      assert( tokenType==TK_SPACE || tokenType==TK_ILLEGAL );
      if( db->u1.isInterrupted ){
        pParse->rc = SQLITE_INTERRUPT;
        break;
      }
      if( tokenType==TK_ILLEGAL ){
        sqlite3ErrorMsg(pParse, "unrecognized token: \"%.*s\"", n, zSql);
        break;
      }
      zSql += n;
    }else{
      pParse->sLastToken.z = zSql;
      pParse->sLastToken.n = n;
      sqlite3Parser(pEngine, tokenType, pParse->sLastToken, pParse);
      lastTokenParsed = tokenType;
      zSql += n;
      if( pParse->rc!=SQLITE_OK || db->mallocFailed ) break;
    }
  }
  assert( nErr==0 );
  pParse->zTail = zSql;
#ifdef YYTRACKMAXSTACKDEPTH
  sqlite3_mutex_enter(sqlite3MallocMutex());
  sqlite3StatusHighwater(SQLITE_STATUS_PARSER_STACK,
      sqlite3ParserStackPeak(pEngine)
  );
  sqlite3_mutex_leave(sqlite3MallocMutex());
#endif /* YYDEBUG */
#ifdef sqlite3Parser_ENGINEALWAYSONSTACK
  sqlite3ParserFinalize(pEngine);
#else
  sqlite3ParserFree(pEngine, sqlite3_free);
     /* 7     END: */ {    1,  7,     5,       5,      5,    5,       5,   5, },
  };
#else
  /* If triggers are not supported by this compile then the statement machine
  ** used to detect the end of a statement is much simpler
  */
  static const u8 trans[3][3] = {
                     /* Token:           */
     /* State:       **  SEMI  WS  OTHER */
     /* 0 INVALID: */ {    1,  0,     2, },
     /* 1   START: */ {    1,  1,     2, },
     /* 2  NORMAL: */ {    1,  2,     2, },
  };
#endif /* SQLITE_OMIT_TRIGGER */

#ifdef SQLITE_ENABLE_API_ARMOR
  if( zSql==0 ){
    (void)SQLITE_MISUSE_BKPT;
    return 0;
  }
#endif

  while( *zSql ){
    switch( *zSql ){
      case ';': {  /* A semicolon */
        token = tkSEMI;
        break;
      }
      case ' ':
      case '\r':
      case '\t':
      case '\n':
      case '\f': {  /* White space is ignored */
        token = tkWS;
        break;
      }
      case '/': {   /* C-style comments */
        if( zSql[1]!='*' ){
          token = tkOTHER;
          break;
        }
        zSql += 2;
        while( zSql[0] && (zSql[0]!='*' || zSql[1]!='/') ){ zSql++; }
        if( zSql[0]==0 ) return 0;
        zSql++;
        token = tkWS;
        break;
      }
      case '-': {   /* SQL-style comments from "--" to end of line */
        if( zSql[1]!='-' ){
          token = tkOTHER;
          break;
        }
        while( *zSql && *zSql!='\n' ){ zSql++; }
        if( *zSql==0 ) return state==1;
        token = tkWS;
        break;
      }
      case '[': {   /* Microsoft-style identifiers in [...] */
        zSql++;
        while( *zSql && *zSql!=']' ){ zSql++; }
        if( *zSql==0 ) return 0;
        token = tkOTHER;
        break;
      }
      case '`':     /* Grave-accent quoted symbols used by MySQL */
      case '"':     /* single- and double-quoted strings */
      case '\'': {
        int c = *zSql;
        zSql++;
        while( *zSql && *zSql!=c ){ zSql++; }
        if( *zSql==0 ) return 0;
        token = tkOTHER;
        break;
      }
      default: {
#ifdef SQLITE_EBCDIC
        unsigned char c;
#endif
        if( IdChar((u8)*zSql) ){
          /* Keywords and unquoted identifiers */
          int nId;
          for(nId=1; IdChar(zSql[nId]); nId++){}
#ifdef SQLITE_OMIT_TRIGGER
          token = tkOTHER;
#else
          switch( *zSql ){
            case 'c': case 'C': {
              if( nId==6 && sqlite3StrNICmp(zSql, "create", 6)==0 ){
                token = tkCREATE;
              }else{
                token = tkOTHER;
              }
              break;
            }
            case 't': case 'T': {
              if( nId==7 && sqlite3StrNICmp(zSql, "trigger", 7)==0 ){
                token = tkTRIGGER;
              }else if( nId==4 && sqlite3StrNICmp(zSql, "temp", 4)==0 ){
                token = tkTEMP;
              }else if( nId==9 && sqlite3StrNICmp(zSql, "temporary", 9)==0 ){
                token = tkTEMP;
              }else{
                token = tkOTHER;
              }
              break;
            }
            case 'e':  case 'E': {
              if( nId==3 && sqlite3StrNICmp(zSql, "end", 3)==0 ){
                token = tkEND;
              }else
#ifndef SQLITE_OMIT_EXPLAIN
              if( nId==7 && sqlite3StrNICmp(zSql, "explain", 7)==0 ){
                token = tkEXPLAIN;
              }else
#endif
              {
                token = tkOTHER;
              }
              break;
            }
            default: {
              token = tkOTHER;
              break;
            }
          }
#endif /* SQLITE_OMIT_TRIGGER */
          zSql += nId-1;
        }else{
          /* Operators and special symbols */
          token = tkOTHER;
        }
        break;
      }
    }
    state = trans[state][token];
    zSql++;
  }
  return state==1;
}

#ifndef SQLITE_OMIT_UTF16
/*
** This routine is the same as the sqlite3_complete() routine described
** above, except that the parameter is required to be UTF-16 encoded, not
** UTF-8.
*/
SQLITE_API int sqlite3_complete16(const void *zSql){
  sqlite3_value *pVal;
  char const *zSql8;
  int rc;

#ifndef SQLITE_OMIT_AUTOINIT
  rc = sqlite3_initialize();
  if( rc ) return rc;
#endif
*/
SQLITE_API int sqlite3_wal_checkpoint(sqlite3 *db, const char *zDb){
  /* EVIDENCE-OF: R-41613-20553 The sqlite3_wal_checkpoint(D,X) is equivalent to
  ** sqlite3_wal_checkpoint_v2(D,X,SQLITE_CHECKPOINT_PASSIVE,0,0). */
  return sqlite3_wal_checkpoint_v2(db,zDb,SQLITE_CHECKPOINT_PASSIVE,0,0);
}

#ifndef SQLITE_OMIT_WAL
/*
** Run a checkpoint on database iDb. This is a no-op if database iDb is
** not currently open in WAL mode.
**
** If a transaction is open on the database being checkpointed, this 
** function returns SQLITE_LOCKED and a checkpoint is not attempted. If 
** an error occurs while running the checkpoint, an SQLite error code is 
** returned (i.e. SQLITE_IOERR). Otherwise, SQLITE_OK.
**
** The mutex on database handle db should be held by the caller. The mutex
** associated with the specific b-tree being checkpointed is taken by
** this function while the checkpoint is running.
**
** If iDb is passed SQLITE_MAX_ATTACHED, then all attached databases are
** checkpointed. If an error is encountered it is returned immediately -
** no attempt is made to checkpoint any remaining databases.
**
** Parameter eMode is one of SQLITE_CHECKPOINT_PASSIVE, FULL, RESTART
** or TRUNCATE.
*/
SQLITE_PRIVATE int sqlite3Checkpoint(sqlite3 *db, int iDb, int eMode, int *pnLog, int *pnCkpt){
  int rc = SQLITE_OK;             /* Return code */
  int i;                          /* Used to iterate through attached dbs */
  int bBusy = 0;                  /* True if SQLITE_BUSY has been encountered */

  assert( sqlite3_mutex_held(db->mutex) );
  assert( !pnLog || *pnLog==-1 );
  assert( !pnCkpt || *pnCkpt==-1 );

  for(i=0; i<db->nDb && rc==SQLITE_OK; i++){
    if( i==iDb || iDb==SQLITE_MAX_ATTACHED ){
      rc = sqlite3BtreeCheckpoint(db->aDb[i].pBt, eMode, pnLog, pnCkpt);
      pnLog = 0;
      pnCkpt = 0;
      if( rc==SQLITE_BUSY ){
        bBusy = 1;
        rc = SQLITE_OK;
      }
    }
  }

  return (rc==SQLITE_OK && bBusy) ? SQLITE_BUSY : rc;
}
    flags |= SQLITE_OPEN_URI;

    for(iIn=0; iIn<nUri; iIn++) nByte += (zUri[iIn]=='&');
    zFile = sqlite3_malloc64(nByte);
    if( !zFile ) return SQLITE_NOMEM_BKPT;

    iIn = 5;
#ifdef SQLITE_ALLOW_URI_AUTHORITY
    if( strncmp(zUri+5, "///", 3)==0 ){
      iIn = 7;
      /* The following condition causes URIs with five leading / characters
      ** like file://///host/path to be converted into UNCs like //host/path.
      ** The correct URI for that UNC has only two or four leading / characters
      ** file://host/path or file:////host/path.  But 5 leading slashes is a 
      ** common error, we are told, so we handle it as a special case. */
      if( strncmp(zUri+7, "///", 3)==0 ){ iIn++; }
    }else if( strncmp(zUri+5, "//localhost/", 12)==0 ){
      iIn = 16;
    }
#else
    /* Discard the scheme and authority segments of the URI. */
    if( zUri[5]=='/' && zUri[6]=='/' ){
      iIn = 7;
      while( zUri[iIn] && zUri[iIn]!='/' ) iIn++;
      if( iIn!=7 && (iIn!=16 || memcmp("localhost", &zUri[7], 9)) ){
        *pzErrMsg = sqlite3_mprintf("invalid uri authority: %.*s", 
            iIn-7, &zUri[7]);
        rc = SQLITE_ERROR;
        goto parse_uri_out;
      }
    }
#endif

    /* Copy the filename and any query parameters into the zFile buffer. 
    ** Decode %HH escape codes along the way. 
    **
    ** Within this loop, variable eState may be set to 0, 1 or 2, depending
    ** on the parsing context. As follows:
    **
    **   0: Parsing file-name.
    **   1: Parsing name section of a name=value query parameter.
    **   2: Parsing value section of a name=value query parameter.
    */
    eState = 0;
    while( (c = zUri[iIn])!=0 && c!='#' ){
      iIn++;
      if( c=='%' 
       && sqlite3Isxdigit(zUri[iIn]) 
       && sqlite3Isxdigit(zUri[iIn+1]) 
      ){
        int octet = (sqlite3HexToInt(zUri[iIn++]) << 4);
  sqlite3 *db,                    /* The SQLite database connection */
  void *pAux,                     /* Hash table containing tokenizers */
  int argc,                       /* Number of elements in argv array */
  const char * const *argv,       /* xCreate/xConnect argument array */
  sqlite3_vtab **ppVTab,          /* Write the resulting vtab structure here */
  char **pzErr                    /* Write any error message here */
){
  Fts3Hash *pHash = (Fts3Hash *)pAux;
  Fts3Table *p = 0;               /* Pointer to allocated vtab */
  int rc = SQLITE_OK;             /* Return code */
  int i;                          /* Iterator variable */
  int nByte;                      /* Size of allocation used for *p */
  int iCol;                       /* Column index */
  int nString = 0;                /* Bytes required to hold all column names */
  int nCol = 0;                   /* Number of columns in the FTS table */
  char *zCsr;                     /* Space for holding column names */
  int nDb;                        /* Bytes required to hold database name */
  int nName;                      /* Bytes required to hold table name */
  int isFts4 = (argv[0][3]=='4'); /* True for FTS4, false for FTS3 */
  const char **aCol;              /* Array of column names */
  sqlite3_tokenizer *pTokenizer = 0;        /* Tokenizer for this table */

  int nIndex = 0;                 /* Size of aIndex[] array */
  struct Fts3Index *aIndex = 0;   /* Array of indexes for this table */

  /* The results of parsing supported FTS4 key=value options: */
  int bNoDocsize = 0;             /* True to omit %_docsize table */
  int bDescIdx = 0;               /* True to store descending indexes */
  char *zPrefix = 0;              /* Prefix parameter value (or NULL) */
  char *zCompress = 0;            /* compress=? parameter (or NULL) */
  char *zUncompress = 0;          /* uncompress=? parameter (or NULL) */
  char *zContent = 0;             /* content=? parameter (or NULL) */
  char *zLanguageid = 0;          /* languageid=? parameter (or NULL) */
  char **azNotindexed = 0;        /* The set of notindexed= columns */
  int nNotindexed = 0;            /* Size of azNotindexed[] array */

  assert( strlen(argv[0])==4 );
  assert( (sqlite3_strnicmp(argv[0], "fts4", 4)==0 && isFts4)
       || (sqlite3_strnicmp(argv[0], "fts3", 4)==0 && !isFts4)
  );

  nDb = (int)strlen(argv[1]) + 1;
  nName = (int)strlen(argv[2]) + 1;

  nByte = sizeof(const char *) * (argc-2);
  aCol = (const char **)sqlite3_malloc(nByte);
  if( aCol ){
    memset((void*)aCol, 0, nByte);
    azNotindexed = (char **)sqlite3_malloc(nByte);
  }
  if( azNotindexed ){

/*
** Helper type used by fts3EvalIncrPhraseNext() and incrPhraseTokenNext().
*/
typedef struct TokenDoclist TokenDoclist;
struct TokenDoclist {
  int bIgnore;
  sqlite3_int64 iDocid;
  char *pList;
  int nList;
};

/*
** Token pToken is an incrementally loaded token that is part of a 
** multi-token phrase. Advance it to the next matching document in the
** database and populate output variable *p with the details of the new
** entry. Or, if the iterator has reached EOF, set *pbEof to true.
**
** If an error occurs, return an SQLite error code. Otherwise, return 
** SQLITE_OK.
*/
static int incrPhraseTokenNext(
  Fts3Table *pTab,                /* Virtual table handle */
  Fts3Phrase *pPhrase,            /* Phrase to advance token of */
  int iToken,                     /* Specific token to advance */
  TokenDoclist *p,                /* OUT: Docid and doclist for new entry */
  u8 *pbEof                       /* OUT: True if iterator is at EOF */
){
  int rc = SQLITE_OK;

  if( pPhrase->iDoclistToken==iToken ){
    assert( p->bIgnore==0 );
    assert( pPhrase->aToken[iToken].pSegcsr==0 );
    fts3EvalDlPhraseNext(pTab, &pPhrase->doclist, pbEof);
    p->pList = pPhrase->doclist.pList;
    p->nList = pPhrase->doclist.nList;
    p->iDocid = pPhrase->doclist.iDocid;
  }else{
    Fts3PhraseToken *pToken = &pPhrase->aToken[iToken];
    assert( pToken->pDeferred==0 );
    assert( pToken->pSegcsr || pPhrase->iDoclistToken>=0 );
    if( pToken->pSegcsr ){
      assert( p->bIgnore==0 );
      rc = sqlite3Fts3MsrIncrNext(
          pTab, pToken->pSegcsr, &p->iDocid, &p->pList, &p->nList
      );
      if( p->pList==0 ) *pbEof = 1;
    }else{
      p->bIgnore = 1;
    }
  }
    pCsr->nInput = nInput;
  }

  *pp = &pCsr->base;
  UNUSED_PARAMETER(p);
  return SQLITE_OK;
}

/*
** Close a tokenization cursor previously opened by a call to
** simpleOpen() above.
*/
static int unicodeClose(sqlite3_tokenizer_cursor *pCursor){
  unicode_cursor *pCsr = (unicode_cursor *) pCursor;
  sqlite3_free(pCsr->zToken);
  sqlite3_free(pCsr);
  return SQLITE_OK;
}

/*
** Extract the next token from a tokenization cursor.  The cursor must
** have been opened by a prior call to simpleOpen().
*/
static int unicodeNext(
  sqlite3_tokenizer_cursor *pC,   /* Cursor returned by simpleOpen */
  const char **paToken,           /* OUT: Token text */
  int *pnToken,                   /* OUT: Number of bytes at *paToken */
  int *piStart,                   /* OUT: Starting offset of token */
  int *piEnd,                     /* OUT: Ending offset of token */
  int *piPos                      /* OUT: Position integer of token */
){
  unicode_cursor *pCsr = (unicode_cursor *)pC;
  unicode_tokenizer *p = ((unicode_tokenizer *)pCsr->base.pTokenizer);
  unsigned int iCode = 0;
  char *zOut;
  const unsigned char *z = &pCsr->aInput[pCsr->iOff];
  const unsigned char *zStart = z;
  const unsigned char *zEnd;
  const unsigned char *zTerm = &pCsr->aInput[pCsr->nInput];

  /* Scan past any delimiter characters before the start of the next token.
  ** Return SQLITE_DONE early if this takes us all the way to the end of 
  ** the input.  */
  while( z<zTerm ){
    READ_UTF8(z, zTerm, iCode);
    if( unicodeIsAlnum(p, (int)iCode) ) break;
    zStart = z;
  }
  if( zStart>=zTerm ) return SQLITE_DONE;

  zOut = pCsr->zToken;
    0x0404DC09, 0x0405E411, 0x0406400C, 0x0407402E, 0x040E7C01,
    0x040F4001, 0x04215C01, 0x04247C01, 0x0424FC01, 0x04280403,
    0x04281402, 0x04283004, 0x0428E003, 0x0428FC01, 0x04294009,
    0x0429FC01, 0x042CE407, 0x04400003, 0x0440E016, 0x04420003,
    0x0442C012, 0x04440003, 0x04449C0E, 0x04450004, 0x04460003,
    0x0446CC0E, 0x04471404, 0x045AAC0D, 0x0491C004, 0x05BD442E,
    0x05BE3C04, 0x074000F6, 0x07440027, 0x0744A4B5, 0x07480046,
    0x074C0057, 0x075B0401, 0x075B6C01, 0x075BEC01, 0x075C5401,
    0x075CD401, 0x075D3C01, 0x075DBC01, 0x075E2401, 0x075EA401,
    0x075F0C01, 0x07BBC002, 0x07C0002C, 0x07C0C064, 0x07C2800F,
    0x07C2C40E, 0x07C3040F, 0x07C3440F, 0x07C4401F, 0x07C4C03C,
    0x07C5C02B, 0x07C7981D, 0x07C8402B, 0x07C90009, 0x07C94002,
    0x07CC0021, 0x07CCC006, 0x07CCDC46, 0x07CE0014, 0x07CE8025,
    0x07CF1805, 0x07CF8011, 0x07D0003F, 0x07D10001, 0x07D108B6,
    0x07D3E404, 0x07D4003E, 0x07D50004, 0x07D54018, 0x07D7EC46,
    0x07D9140B, 0x07DA0046, 0x07DC0074, 0x38000401, 0x38008060,
    0x380400F0,
  };
  static const unsigned int aAscii[4] = {
    0xFFFFFFFF, 0xFC00FFFF, 0xF8000001, 0xF8000001,
  };

  if( (unsigned int)c<128 ){
    return ( (aAscii[c >> 5] & ((unsigned int)1 << (c & 0x001F)))==0 );
  }else if( (unsigned int)c<(1<<22) ){
    unsigned int key = (((unsigned int)c)<<10) | 0x000003FF;
    int iRes = 0;
    int iHi = sizeof(aEntry)/sizeof(aEntry[0]) - 1;
    int iLo = 0;
    while( iHi>=iLo ){
      int iTest = (iHi + iLo) / 2;
      if( key >= aEntry[iTest] ){
        iRes = iTest;
        iLo = iTest+1;
      }else{
        iHi = iTest-1;
      }
    }
    assert( aEntry[0]<key );
    assert( key>=aEntry[iRes] );
    return (((unsigned int)c) >= ((aEntry[iRes]>>10) + (aEntry[iRes]&0x3FF)));
  }
  return 1;
}


/*
** If the argument is a codepoint corresponding to a lowercase letter
** in the ASCII range with a diacritic added, return the codepoint
** of the ASCII letter only. For example, if passed 235 - "LATIN
** SMALL LETTER E WITH DIAERESIS" - return 65 ("LATIN SMALL LETTER
     2024,  2040,  2060,  2110,  2168,  2206,  2264,  2286, 
     2344,  2383,  2472,  2488,  2516,  2596,  2668,  2732, 
     2782,  2842,  2894,  2954,  2984,  3000,  3028,  3336, 
     3456,  3696,  3712,  3728,  3744,  3896,  3912,  3928, 
     3968,  4008,  4040,  4106,  4138,  4170,  4202,  4234, 
     4266,  4296,  4312,  4344,  4408,  4424,  4472,  4504, 
     6148,  6198,  6264,  6280,  6360,  6429,  6505,  6529, 
    61448, 61468, 61534, 61592, 61642, 61688, 61704, 61726, 
    61784, 61800, 61836, 61880, 61914, 61948, 61998, 62122, 
    62154, 62200, 62218, 62302, 62364, 62442, 62478, 62536, 
    62554, 62584, 62604, 62640, 62648, 62656, 62664, 62730, 
    62924, 63050, 63082, 63274, 63390, 
  };
  char aChar[] = {
    '\0', 'a',  'c',  'e',  'i',  'n',  'o',  'u',  'y',  'y',  'a',  'c',  
    'd',  'e',  'e',  'g',  'h',  'i',  'j',  'k',  'l',  'n',  'o',  'r',  
    's',  't',  'u',  'u',  'w',  'y',  'z',  'o',  'u',  'a',  'i',  'o',  
    'u',  'g',  'k',  'o',  'j',  'g',  'n',  'a',  'e',  'i',  'o',  'r',  
    'u',  's',  't',  'h',  'a',  'e',  'o',  'y',  '\0', '\0', '\0', '\0', 
    '\0', '\0', '\0', '\0', 'a',  'b',  'd',  'd',  'e',  'f',  'g',  'h',  
    'h',  'i',  'k',  'l',  'l',  'm',  'n',  'p',  'r',  'r',  's',  't',  
    'u',  'v',  'w',  'w',  'x',  'y',  'z',  'h',  't',  'w',  'y',  'a',  
    'e',  'i',  'o',  'u',  'y',  
  };

  unsigned int key = (((unsigned int)c)<<3) | 0x00000007;
  int iRes = 0;
  int iHi = sizeof(aDia)/sizeof(aDia[0]) - 1;
  int iLo = 0;
  while( iHi>=iLo ){
    int iTest = (iHi + iLo) / 2;
    if( key >= aDia[iTest] ){
      iRes = iTest;
      iLo = iTest+1;
    }else{
      iHi = iTest-1;
    }
  }
  assert( key>=aDia[iRes] );
  return ((c > (aDia[iRes]>>3) + (aDia[iRes]&0x07)) ? c : (int)aChar[iRes]);
}


/*
** Return true if the argument interpreted as a unicode codepoint
** is a diacritical modifier character.
*/
SQLITE_PRIVATE int sqlite3FtsUnicodeIsdiacritic(int c){
  unsigned int mask0 = 0x08029FDF;
  unsigned int mask1 = 0x000361F8;
  if( c<768 || c>817 ) return 0;
          RBU_STATE_PROGRESS, p->nProgress,
          RBU_STATE_CKPT, p->iWalCksum,
          RBU_STATE_COOKIE, (i64)pFd->iCookie,
          RBU_STATE_OALSZ, p->iOalSz,
          RBU_STATE_PHASEONESTEP, p->nPhaseOneStep
      )
    );
    assert( pInsert==0 || rc==SQLITE_OK );

    if( rc==SQLITE_OK ){
      sqlite3_step(pInsert);
      rc = sqlite3_finalize(pInsert);
    }
    if( rc!=SQLITE_OK ) p->rc = rc;
  }
}


/*
** The second argument passed to this function is the name of a PRAGMA 
** setting - "page_size", "auto_vacuum", "user_version" or "application_id".
** This function executes the following on sqlite3rbu.dbRbu:
**
**   "PRAGMA main.$zPragma"
**
** where $zPragma is the string passed as the second argument, then
** on sqlite3rbu.dbMain:
**
**   "PRAGMA main.$zPragma = $val"
**
** where $val is the value returned by the first PRAGMA invocation.
**
** In short, it copies the value  of the specified PRAGMA setting from
** dbRbu to dbMain.
*/
static void rbuCopyPragma(sqlite3rbu *p, const char *zPragma){
  if( p->rc==SQLITE_OK ){
    sqlite3_stmt *pPragma = 0;
    p->rc = prepareFreeAndCollectError(p->dbRbu, &pPragma, &p->zErrmsg, 
        sqlite3_mprintf("PRAGMA main.%s", zPragma)
    );
    if( p->rc==SQLITE_OK && SQLITE_ROW==sqlite3_step(pPragma) ){
      p->rc = rbuMPrintfExec(p, p->dbMain, "PRAGMA main.%s = %d",
          zPragma, sqlite3_column_int(pPragma, 0)
      );
    }
    rbuFinalize(p, pPragma);
  }
}

/*
typedef sqlite3_int64 i64;
typedef sqlite3_uint64 u64;

#ifndef ArraySize
# define ArraySize(x) ((int)(sizeof(x) / sizeof(x[0])))
#endif

#define testcase(x)
#define ALWAYS(x) 1
#define NEVER(x) 0

#define MIN(x,y) (((x) < (y)) ? (x) : (y))
#define MAX(x,y) (((x) > (y)) ? (x) : (y))

/*
** Constants for the largest and smallest possible 64-bit signed integers.
*/
# define LARGEST_INT64  (0xffffffff|(((i64)0x7fffffff)<<32))
# define SMALLEST_INT64 (((i64)-1) - LARGEST_INT64)

#endif

/* Truncate very long tokens to this many bytes. Hard limit is 
** (65536-1-1-4-9)==65521 bytes. The limiting factor is the 16-bit offset
** field that occurs at the start of each leaf page (see fts5_index.c). */
#define FTS5_MAX_TOKEN_SIZE 32768

/*
** Maximum number of prefix indexes on single FTS5 table. This must be
** less than 32. If it is set to anything large than that, an #error
** directive in fts5_index.c will cause the build to fail.
*/
#define FTS5_MAX_PREFIX_INDEXES 31

#define FTS5_DEFAULT_NEARDIST 10
#define FTS5_DEFAULT_RANK     "bm25"

/* Name of rank and rowid columns */
#define FTS5_RANK_NAME "rank"
#define FTS5_ROWID_NAME "rowid"

#ifdef SQLITE_DEBUG
# define FTS5_CORRUPT sqlite3Fts5Corrupt()
static int sqlite3Fts5Corrupt(void);
#else
# define FTS5_CORRUPT SQLITE_CORRUPT_VTAB
#endif

/*
** The assert_nc() macro is similar to the assert() macro, except that it
** is used for assert() conditions that are true only if it can be 
){
  sqlite3Fts5ParserARG_FETCH;
#ifndef NDEBUG
  if( fts5yyTraceFILE ){
    fprintf(fts5yyTraceFILE,"%sFail!\n",fts5yyTracePrompt);
  }
#endif
  while( fts5yypParser->fts5yytos>fts5yypParser->fts5yystack ) fts5yy_pop_parser_stack(fts5yypParser);
  /* Here code is inserted which will be executed whenever the
  ** parser fails */
/************ Begin %parse_failure code ***************************************/
/************ End %parse_failure code *****************************************/
  sqlite3Fts5ParserARG_STORE; /* Suppress warning about unused %extra_argument variable */
}
#endif /* fts5YYNOERRORRECOVERY */

/*
** The following code executes when a syntax error first occurs.
*/
static void fts5yy_syntax_error(
  fts5yyParser *fts5yypParser,           /* The parser */
  int fts5yymajor,                   /* The major type of the error token */
  sqlite3Fts5ParserFTS5TOKENTYPE fts5yyminor         /* The minor type of the error token */
){
  sqlite3Fts5ParserARG_FETCH;
#define FTS5TOKEN fts5yyminor
/************ Begin %syntax_error code ****************************************/

  UNUSED_PARAM(fts5yymajor); /* Silence a compiler warning */
  sqlite3Fts5ParseError(
    pParse, "fts5: syntax error near \"%.*s\"",FTS5TOKEN.n,FTS5TOKEN.p
  );
/************ End %syntax_error code ******************************************/
  sqlite3Fts5ParserARG_STORE; /* Suppress warning about unused %extra_argument variable */
}

/*
** The following is executed when the parser accepts
*/
static void fts5yy_accept(
  fts5yyParser *fts5yypParser           /* The parser */
){
  sqlite3Fts5ParserARG_FETCH;
#ifndef NDEBUG
  if( fts5yyTraceFILE ){
    fprintf(fts5yyTraceFILE,"%sAccept!\n",fts5yyTracePrompt);
  }
#endif
#ifndef fts5YYNOERRORRECOVERY
  fts5yypParser->fts5yyerrcnt = -1;
#endif
    0x0404DC09, 0x0405E411, 0x0406400C, 0x0407402E, 0x040E7C01,
    0x040F4001, 0x04215C01, 0x04247C01, 0x0424FC01, 0x04280403,
    0x04281402, 0x04283004, 0x0428E003, 0x0428FC01, 0x04294009,
    0x0429FC01, 0x042CE407, 0x04400003, 0x0440E016, 0x04420003,
    0x0442C012, 0x04440003, 0x04449C0E, 0x04450004, 0x04460003,
    0x0446CC0E, 0x04471404, 0x045AAC0D, 0x0491C004, 0x05BD442E,
    0x05BE3C04, 0x074000F6, 0x07440027, 0x0744A4B5, 0x07480046,
    0x074C0057, 0x075B0401, 0x075B6C01, 0x075BEC01, 0x075C5401,
    0x075CD401, 0x075D3C01, 0x075DBC01, 0x075E2401, 0x075EA401,
    0x075F0C01, 0x07BBC002, 0x07C0002C, 0x07C0C064, 0x07C2800F,
    0x07C2C40E, 0x07C3040F, 0x07C3440F, 0x07C4401F, 0x07C4C03C,
    0x07C5C02B, 0x07C7981D, 0x07C8402B, 0x07C90009, 0x07C94002,
    0x07CC0021, 0x07CCC006, 0x07CCDC46, 0x07CE0014, 0x07CE8025,
    0x07CF1805, 0x07CF8011, 0x07D0003F, 0x07D10001, 0x07D108B6,
    0x07D3E404, 0x07D4003E, 0x07D50004, 0x07D54018, 0x07D7EC46,
    0x07D9140B, 0x07DA0046, 0x07DC0074, 0x38000401, 0x38008060,
    0x380400F0,
  };
  static const unsigned int aAscii[4] = {
    0xFFFFFFFF, 0xFC00FFFF, 0xF8000001, 0xF8000001,
  };

  if( (unsigned int)c<128 ){
    return ( (aAscii[c >> 5] & (1 << (c & 0x001F)))==0 );
  }else if( (unsigned int)c<(1<<22) ){
    unsigned int key = (((unsigned int)c)<<10) | 0x000003FF;
    int iRes = 0;
    int iHi = sizeof(aEntry)/sizeof(aEntry[0]) - 1;
    int iLo = 0;
    while( iHi>=iLo ){
      int iTest = (iHi + iLo) / 2;
      if( key >= aEntry[iTest] ){
        iRes = iTest;
        iLo = iTest+1;
      }else{
        iHi = iTest-1;
      }
    }
    assert( aEntry[0]<key );
    assert( key>=aEntry[iRes] );
    return (((unsigned int)c) >= ((aEntry[iRes]>>10) + (aEntry[iRes]&0x3FF)));
  }
  return 1;
}


/*
** If the argument is a codepoint corresponding to a lowercase letter
** in the ASCII range with a diacritic added, return the codepoint
** of the ASCII letter only. For example, if passed 235 - "LATIN
** SMALL LETTER E WITH DIAERESIS" - return 65 ("LATIN SMALL LETTER
     2024,  2040,  2060,  2110,  2168,  2206,  2264,  2286, 
     2344,  2383,  2472,  2488,  2516,  2596,  2668,  2732, 
     2782,  2842,  2894,  2954,  2984,  3000,  3028,  3336, 
     3456,  3696,  3712,  3728,  3744,  3896,  3912,  3928, 
     3968,  4008,  4040,  4106,  4138,  4170,  4202,  4234, 
     4266,  4296,  4312,  4344,  4408,  4424,  4472,  4504, 
     6148,  6198,  6264,  6280,  6360,  6429,  6505,  6529, 
    61448, 61468, 61534, 61592, 61642, 61688, 61704, 61726, 
    61784, 61800, 61836, 61880, 61914, 61948, 61998, 62122, 
    62154, 62200, 62218, 62302, 62364, 62442, 62478, 62536, 
    62554, 62584, 62604, 62640, 62648, 62656, 62664, 62730, 
    62924, 63050, 63082, 63274, 63390, 
  };
  char aChar[] = {
    '\0', 'a',  'c',  'e',  'i',  'n',  'o',  'u',  'y',  'y',  'a',  'c',  
    'd',  'e',  'e',  'g',  'h',  'i',  'j',  'k',  'l',  'n',  'o',  'r',  
    's',  't',  'u',  'u',  'w',  'y',  'z',  'o',  'u',  'a',  'i',  'o',  
    'u',  'g',  'k',  'o',  'j',  'g',  'n',  'a',  'e',  'i',  'o',  'r',  
    'u',  's',  't',  'h',  'a',  'e',  'o',  'y',  '\0', '\0', '\0', '\0', 
    '\0', '\0', '\0', '\0', 'a',  'b',  'd',  'd',  'e',  'f',  'g',  'h',  
    'h',  'i',  'k',  'l',  'l',  'm',  'n',  'p',  'r',  'r',  's',  't',  
    'u',  'v',  'w',  'w',  'x',  'y',  'z',  'h',  't',  'w',  'y',  'a',  
    'e',  'i',  'o',  'u',  'y',  
  };

  unsigned int key = (((unsigned int)c)<<3) | 0x00000007;
  int iRes = 0;
  int iHi = sizeof(aDia)/sizeof(aDia[0]) - 1;
  int iLo = 0;
  while( iHi>=iLo ){
    int iTest = (iHi + iLo) / 2;
    if( key >= aDia[iTest] ){
      iRes = iTest;
      iLo = iTest+1;
    }else{
      iHi = iTest-1;
    }
  }
  assert( key>=aDia[iRes] );
  return ((c > (aDia[iRes]>>3) + (aDia[iRes]&0x07)) ? c : (int)aChar[iRes]);
}


/*
** Return true if the argument interpreted as a unicode codepoint
** is a diacritical modifier character.
*/
static int sqlite3Fts5UnicodeIsdiacritic(int c){
  unsigned int mask0 = 0x08029FDF;
  unsigned int mask1 = 0x000361F8;
  if( c<768 || c>817 ) return 0;