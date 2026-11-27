    })

def if_portable(if_true, if_false = []):
    return if_true

# Linux systems may required -lrt linker flag for e.g. clock_gettime
# see https://github.com/tensorflow/tensorflow/issues/15129
def lrt_if_needed():
    lrt = ["-lrt"]
    return select({
        clean_dep("//tensorflow:linux_aarch64"): lrt,
        clean_dep("//tensorflow:linux_x86_64"): lrt,
        clean_dep("//tensorflow:linux_ppc64le"): lrt,
        "//conditions:default": [],
    })

def get_win_copts(is_external = False):
    WINDOWS_COPTS = [
        "/DPLATFORM_WINDOWS",
        "/DEIGEN_HAS_C99_MATH",
        "/DTENSORFLOW_USE_EIGEN_THREADPOOL",
        "/DEIGEN_AVOID_STL_ARRAY",
        "/Iexternal/gemmlowp",
        "/wd4018",  # -Wno-sign-compare
        # Bazel's CROSSTOOL currently pass /EHsc to enable exception by
        # default. We can't pass /EHs-c- to disable exception, otherwise
        # we will get a waterfall of flag conflict warnings. Wait for
        # Bazel to fix this.
        # "/D_HAS_EXCEPTIONS=0",
        # "/EHs-c-",
        "/wd4577",
        "/DNOGDI",
        # Also see build:windows lines in tensorflow/opensource_only/.bazelrc
        # where we set some other options globally.
    ]
    if is_external:
        return WINDOWS_COPTS + ["/UTF_COMPILE_LIBRARY"]
    else:
        return WINDOWS_COPTS + ["/DTF_COMPILE_LIBRARY"]

def tf_copts(
        android_optimization_level_override = "-O2",
        is_external = False,
        allow_exceptions = False):
    # For compatibility reasons, android_optimization_level_override
    # is currently only being set for Android.
    # To clear this value, and allow the CROSSTOOL default
    # to be used, pass android_optimization_level_override=None
    android_copts = [
        "-DTF_LEAN_BINARY",
        "-Wno-narrowing",
        "-fomit-frame-pointer",
    ]
    if android_optimization_level_override:
        android_copts.append(android_optimization_level_override)
    return (
        if_not_windows([
            "-DEIGEN_AVOID_STL_ARRAY",
            "-Iexternal/gemmlowp",
            "-Wno-sign-compare",
            "-ftemplate-depth=900",
        ]) +
        (if_not_windows(["-fno-exceptions"]) if not allow_exceptions else []) +
        if_cuda(["-DGOOGLE_CUDA=1"]) +
        if_nvcc(["-DTENSORFLOW_USE_NVCC=1"]) +
        if_libtpu(["-DLIBTPU_ON_GCE"], []) +
        if_xla_available(["-DTENSORFLOW_USE_XLA=1"]) +
        if_tensorrt(["-DGOOGLE_TENSORRT=1"]) +
        # Compile in oneDNN based ops when building for x86 platforms
        if_mkl(["-DINTEL_MKL"]) +
        # Enable additional ops (e.g., ops with non-NHWC data layout) and
        # optimizations for Intel builds using oneDNN if configured
        if_enable_mkl(["-DENABLE_MKL"]) +
        other_srcs_internal = [],
        other_hdrs_internal = [],
        pkg = "",
        deps = [
            clean_dep("//tensorflow/cc:ops"),
            clean_dep("//tensorflow/cc:scope"),
            clean_dep("//tensorflow/cc:const_op"),
        ],
        deps_internal = [],
        op_gen = clean_dep("//tensorflow/cc:cc_op_gen_main"),
        include_internal_ops = 0,
        visibility = None,
        # ApiDefs will be loaded in the order specified in this list.
        api_def_srcs = [],
        # Any extra dependencies that the wrapper generator might need.
        extra_gen_deps = [],
        compatible_with = []):
    subsrcs = other_srcs[:]
    subhdrs = other_hdrs[:]
    internalsrcs = other_srcs_internal[:]
    internalhdrs = other_hdrs_internal[:]
    for n in op_lib_names:
        tf_gen_op_wrapper_cc(
            n,
            "ops/" + n,
            api_def_srcs = api_def_srcs,
            include_internal_ops = include_internal_ops,
            op_gen = op_gen,
            pkg = pkg,
            deps = [pkg + ":" + n + "_op_lib"] + extra_gen_deps,
            compatible_with = compatible_with,
        )
        subsrcs += ["ops/" + n + ".cc"]
        subhdrs += ["ops/" + n + ".h"]
        internalsrcs += ["ops/" + n + "_internal.cc"]
        internalhdrs += ["ops/" + n + "_internal.h"]

    cc_library(
        name = name,
        srcs = subsrcs,
        hdrs = subhdrs,
        deps = deps + if_not_android([
            clean_dep("//tensorflow/core:core_cpu"),
            clean_dep("//tensorflow/core:framework"),
            clean_dep("//tensorflow/core:lib"),
            clean_dep("//tensorflow/core:ops"),
            clean_dep("//tensorflow/core:protos_all_cc"),
        ]) + if_android([
            clean_dep("//tensorflow/core:portable_tensorflow_lib"),
        ]),
        copts = tf_copts(),
def _local_genrule(**kwargs):
    _local_genrule_internal(
        exec_compatible_with = [
            "@local_execution_config_platform//:platform_constraint",
        ],
        **kwargs
    )

def tf_version_info_genrule(name, out, compatible_with = None):
    # TODO(gunan): Investigate making this action hermetic so we do not need
    # to run it locally.
    _local_genrule(
        name = name,
        out = out,
        compatible_with = compatible_with,
        exec_tool = "//tensorflow/tools/git:gen_git_source",
        srcs = [
            "@local_config_git//:gen/spec.json",
            "@local_config_git//:gen/head",
            "@local_config_git//:gen/branch_ref",
        ],
        arguments = "--generate \"$@\" --git_tag_override=${GIT_TAG_OVERRIDE:-}",
    )

def _dict_to_kv(d):
    """Convert a dictionary to a space-joined list of key=value pairs."""
    return " " + " ".join(["%s=%s" % (k, v) for k, v in d.items()])

def tf_py_build_info_genrule(name, out):
    _local_genrule(
        name = name,
        out = out,
        exec_tool = "//tensorflow/tools/build_info:gen_build_info",
        arguments =
            "--raw_generate \"$@\" " +
            " --key_value" +
            " is_rocm_build=" + if_rocm("True", "False") +
            " is_cuda_build=" + if_cuda("True", "False") +
            " is_tensorrt_build=" + if_tensorrt("True", "False") +
            if_windows(_dict_to_kv({
                "msvcp_dll_names": "msvcp140.dll,msvcp140_1.dll",
            }), "") + if_windows_cuda(_dict_to_kv({
                "nvcuda_dll_name": "nvcuda.dll",
                "cudart_dll_name": "cudart{cuda_version}.dll",
                "cudnn_dll_name": "cudnn{cudnn_version}.dll",
            }), ""),
    )

def cc_library_with_android_deps(
        deps,
        android_deps = [],