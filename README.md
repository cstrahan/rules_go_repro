Instructions:

1. `$ bazel build :demogo`
2. `$ bazel-bin/linux_amd64_stripped/demogo`

In theory, you should see the following two lines, but you'll find that only the second line is printed:

```
Hello from cxxdep!
Hello from cxxlib!
```

You can compare that to what happens with `cc_binary` like so:

1. `$ bazel build :democc`
2. `$ bazel-bin/democc`

You'll see:

```
Hello from cxxdep!
Hello from cxxlib!
```
