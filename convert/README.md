# convert

> Golang type convert tool

### Simple Example

**ToType** 

```
i64,err := convert.ToInt64("1.23")
```

**AsType**

```
i64 := convert.AsInt64("1.23")
```

**AsTypeDefault**

```
i64 := convert.AsInt64Default("1.1.1",-1)
```

### Note

- Package convert will never **Panic**

- Functions like **ToType** will try their best to convert source type to the target type, if can`t , will return a error explain why

  For Example:

  - ToString call time.Format to convert time to string
  - If some type implement **Stringer**, ToString will call it`s String method to convert it to string
  - ToBool  convert string `"true","True","1"`  to true， `"false","False",0`  to false
  - ToTime  convert string to `time.Time`  use 

- Functions like **AsType** inner call ToType, if something error happened, it will return zero value of the type

- Functions like **AsTypeDefault** inner call ToType, if something error happened, it will return the default param